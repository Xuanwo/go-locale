// +build !unit_test

package locale

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

var detectors = []detector{
	detectViaWin32OLE,
}

// osLanguageCode is a mapping from Microsoft Windows language code to language.Tag
//
// ref:
//   - https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-operatingsystem
//   - https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry
var osLanguageCode = map[uint32]string{
	0x4:   "zh-Hans-CN", // Chinese (Simplified)– China
	0x9:   "en",         // English
	0x404: "zh-Hant-TW", // Chinese (Traditional) – Taiwan
	0x409: "en-US",      // English – United States
	0x411: "ja",         // Japanese
	0x412: "ko",         // Korean
	0x804: "zh-Hans-CN", // Chinese (Simplified) – PRC
	0x809: "en-US",      // English – United Kingdom
}

// detectViaWin32OLE will detect system's language via w32 ole.
//
// code inspired from https://github.com/iamacarpet/go-win64api
func detectViaWin32OLE() (langs []string, err error) {
	defer func() {
		if err == nil {
			return
		}
		err = &Error{"detect via win32 ole", err}
	}()

	err = ole.CoInitialize(0)
	if err != nil {
		return
	}
	defer ole.CoUninitialize()

	unknown, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		return
	}
	defer unknown.Release()

	wmi, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return
	}
	defer wmi.Release()

	serviceRaw, err := oleutil.CallMethod(wmi, "ConnectServer")
	if err != nil {
		return
	}
	service := serviceRaw.ToIDispatch()
	defer service.Release()

	resultRaw, err := oleutil.CallMethod(service, "ExecQuery", "SELECT OSLanguage FROM Win32_OperatingSystem")
	if err != nil {
		return
	}
	result := resultRaw.ToIDispatch()
	defer result.Release()

	itemRaw, err := oleutil.CallMethod(result, "ItemIndex", 0)
	if err != nil {
		return
	}
	item := itemRaw.ToIDispatch()
	defer item.Release()

	languageCode, err := oleutil.GetProperty(item, "OSLanguage")
	if err != nil {
		return
	}

	lang, ok := osLanguageCode[uint32(languageCode.Val)]
	if !ok {
		err = ErrNotSupported
		return
	}
	return []string{lang}, nil
}
