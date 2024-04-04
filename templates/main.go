package templates

import (
	"project/variables"
	"text/template"
)

func InitTemplates() {
	variables.Tmpl, _ = template.ParseGlob("./templates/*.html")
}
