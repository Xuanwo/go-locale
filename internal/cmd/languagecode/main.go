// language code is used to parse [MS-LCID] tables to generate a map in go.
package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"
)

const msLCID = "https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-lcid/a9eac961-e77d-41a6-90a5-ce1a8b0cdb9c"

type lcid struct {
	Language         string
	Location         string
	LanguageID       string
	LanguageTag      string
	SupportedVersion string
}

func main() {
	// Get content from microsoft openspecs
	resp, err := http.Get(msLCID)
	if err != nil {
		log.Fatalf("Get MS-LCID: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Get MS-LCID with status code: %d %s", resp.StatusCode, resp.Status)
	}

	// Load document.
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Create document: %v", err)
	}

	var data []lcid
	idx := 0
	doc.Find("tbody").Each(func(i int, s *goquery.Selection) {
		if i != 1 {
			return
		}
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			value := strings.Trim(s.Text(), "\n ")
			// Remove all \n in value.
			value = strings.ReplaceAll(value, "\n", "")
			switch i % 5 {
			case 0:
				data = append(data, lcid{})
				data[idx].Language = value
			case 1:
				data[idx].Location = value
			case 2:
				data[idx].LanguageID = value
			case 3:
				data[idx].LanguageTag = value
			case 4:
				data[idx].SupportedVersion = value
				idx++
			}
		})
		// Keep this for later debug
		// fmt.Printf("%s", data)
	})

	f, err := os.OpenFile("locale_windows_generated.go", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Open file: %v", err)
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatalf("Template generate: %v", err)
	}
}

var tmpl = template.Must(template.New("language code").Parse(`// +build !unit_test

package locale

// osLanguageCode is a mapping from Microsoft Windows language code to language.Tag
// which genereated via internal/cmd/languagecode, data is from microsoft openspecs.
//
// Microsoft will assign 0x1000 to languages that doesn't have LCID, application should
// handle this, and we will return Und instead.
//
// ref:
//   - https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-operatingsystem
//   - https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry
//   - https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-lcid/a9eac961-e77d-41a6-90a5-ce1a8b0cdb9c
var osLanguageCode = map[uint32]string{
{{- range $_, $v := . }}
{{- if ne .LanguageID "0x1000" }}
	{{ .LanguageID }}: "{{ .LanguageTag }}", // {{ .Language }} - {{ .Location }}, supported from {{ .SupportedVersion }}
{{- end }}
{{- end }}

	0x1000: "Und",
}
`))
