package confluence

import (
	"bytes"
	"github.com/choshsh/wrike-go"
	"html/template"
	"os"
	"strings"
)

func NewTemplate() string {
	tmpl, err := template.ParseFiles("table.html")
	errHandler(err)

	wrikeAPI := wrike.NewWrike(os.Getenv("BEARER"), nil)
	sprints := wrikeAPI.Sprints("2022.03.SP1")

	var tmplString bytes.Buffer
	err = tmpl.Execute(&tmplString, sprints)
	errHandler(err)

	// 태그 사이 공백 제거
	return strings.ReplaceAll(tmplString.String(), `/\>\s+\</m`, `><`)
}
