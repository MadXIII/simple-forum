package front

import (
	"forum/errorhandle"
	"html/template"
	"io"
)

var tpl *template.Template

var err error

//SetUp - ...
func SetUp() {
	tpl, err = template.ParseGlob("./static/templates/*")
	errorhandle.CheckErr(err)
}

// ExecTemplate - ...
func ExecTemplate(wr io.Writer, name string, data interface{}) error {
	return tpl.ExecuteTemplate(wr, name, data)
}
