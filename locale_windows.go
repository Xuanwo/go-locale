package locale

import (
	"errors"
	"fmt"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"golang.org/x/text/language"
)

var detect = func() (tag language.Tag, err error) {
	errorMessage := "detect: %w"

	tag, err = detectViaWin32OLE()
	if err == nil {
		return
	}

	err = fmt.Errorf(errorMessage, ErrNotDetected)
	return
}

// osLanguageCode is a mapping from Microsoft Windows language code to language.Tag
//
// ref: https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-operatingsystem
var osLanguageCode = map[uint32]language.Tag{
	0x4:   language.SimplifiedChinese,  // Chinese (Simplified)– China
	0x9:   language.English,            // English
	0x404: language.TraditionalChinese, // Chinese (Traditional) – Taiwan
	0x409: language.AmericanEnglish,    // English – United States
	0x411: language.Japanese,           // Japanese
	0x412: language.Korean,             // Korean
	0x804: language.SimplifiedChinese,  // Chinese (Simplified) – PRC
	0x809: language.BritishEnglish,     // English – United Kingdom
}

// detectViaWin32OLE will detect system's language via w32 ole.
//
// code inspired from https://github.com/iamacarpet/go-win64api
func detectViaWin32OLE() (tag language.Tag, err error) {
	errorMessage := "detect via win32 OLE: %w"

	err = ole.CoInitialize(0)
	if err != nil {
		err = fmt.Errorf(errorMessage, err)
		return
	}
	defer ole.CoUninitialize()

	unknown, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		err = fmt.Errorf(errorMessage, err)
		return
	}
	defer unknown.Release()

	wmi, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		err = fmt.Errorf(errorMessage, err)
		return
	}
	defer wmi.Release()

	serviceRaw, err := oleutil.CallMethod(wmi, "ConnectServer")
	if err != nil {
		err = fmt.Errorf(errorMessage, err)
		return
	}
	service := serviceRaw.ToIDispatch()
	defer service.Release()

	resultRaw, err := oleutil.CallMethod(service, "ExecQuery", "SELECT OSLanguage FROM Win32_OperatingSystem")
	if err != nil {
		err = fmt.Errorf(errorMessage, err)
		return
	}
	result := resultRaw.ToIDispatch()
	defer result.Release()

	itemRaw, err := oleutil.CallMethod(result, "ItemIndex", 0)
	if err != nil {
		err = fmt.Errorf(errorMessage, err)
		return
	}
	item := itemRaw.ToIDispatch()
	defer item.Release()

	languageCode, err := oleutil.GetProperty(item, "OSLanguage")
	if err != nil {
		err = fmt.Errorf(errorMessage, err)
		return
	}

	tag, ok := osLanguageCode[uint32(languageCode.Val)]
	if !ok {
		err = fmt.Errorf(errorMessage, errors.New("language code not exist"))
		return
	}
	return
}
