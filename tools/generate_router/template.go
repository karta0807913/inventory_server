package main

const CreateOrUpdateTemplate = `
package {{ .Package }}

import (
{{ if ne (len .OptionalFields) 0}}"errors"{{ end }}

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// data will put into struct
func (insert *{{ .StructName }}){{ .FuncName }}(c *gin.Context, db *gorm.DB) error {
    type Body struct {
{{ with .IndexField }}{{ .Name }} {{ .Type }} {{ .Tag }} {{ end }}
{{ range .RequiredFields }}{{ .Name }}  {{ .Type}} {{ .Tag }}
{{ end }}
{{ range .OptionalFields }}{{ .Name }} *{{ .Type}} {{ .Tag }}
{{ end }}
    }
    var body Body;
    err := c.{{ .Decoder }}(&body)
    if err != nil {
        return err
    }

{{/* select array */}}
{{ if ne (len .RequiredFields) 0}}
  selectField := []string {
    {{ range .RequiredFields }}"{{ .Name }}",
    {{ end }}
  }
{{ else }}
  selectField := make([]string, 0)
{{ end }}

{{/* put options */}}
{{ range .OptionalFields }}
if body.{{ .Name }} == nil {
  body.{{ .Name }} = new({{ .Type }})
} else {
  selectField = append(selectField, "{{ .Name }}")
}
{{ end }}

{{/* check options */}}
{{ if ne (len .OptionalFields) 0}}
  if len(selectField) == {{ len .RequiredFields }} {
    return errors.New("rqeuire at least one option")
  }
{{ end }}

{{ with .IndexField }}insert.{{ .Name }} = body.{{ .Name }}
{{ end }}{{ range .RequiredFields }}insert.{{ .Name }} = body.{{ .Name }}
{{ end }}{{ range .OptionalFields }}insert.{{ .Name }} = *body.{{ .Name }}
{{ end }}

{{/* create or update */}}
return db.Select(
selectField[0], selectField[1:],
){{ with .IndexField }}.Where("{{.Name}}=?", body.{{.Name}}){{ end }}.{{ .Mode }}(&insert).Error
}
`

const SearchTemplate = `
package {{ .Package }}

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file generate by go generate, please don't edit it
// search options will put into struct
func (item *{{ .StructName }}) {{ .FuncName }}(c *gin.Context, db *gorm.DB) (*gorm.DB, error) {
    type Body struct {
{{ range .RequiredFields }}{{ .Name }}  {{ .Type }} {{ .Tag }}
{{ end }}
{{ range .OptionalFields }}{{ .Name }} *{{ .Type }} {{ .Tag }}
{{ end }}
    }
    var body Body;
    err := c.{{ .Decoder }}(&body)
    if err != nil {
        return nil, err
    }

{{/* select array */}}
{{ if ne (len .RequiredFields) 0}}
  whereField := []string {
    {{ range .RequiredFields }}"{{ .Name }}=?",
    {{ end }}
  }
  valueField := []interface{}{
    {{ range .RequiredFields }}body.{{ .Name }},
    {{ end }}
  }
  {{ range .RequiredFields }}
  item.{{ .Name }} = body.{{ .Name }}{{ end }}
{{ else }}
  whereField := make([]string, 0)
  valueField := make([]interface{}, 0)
{{ end }}

{{/* put options */}}
{{ range .OptionalFields }}
  if body.{{ .Name }} != nil {
    whereField = append(whereField, "{{ .Name }}=?")
    valueField = append(valueField, body.{{ .Name }})
    item.{{ .Name }} = *body.{{ .Name }}
  }
{{ end }}

{{/* return build db */}}
return db.Where(
whereField,{{ range .RequiredFields }}
 body.{{.Name}},{{ end }}
), nil
}`

type TemplateRoot struct {
	Package        string
	FuncName       string
	StructName     string
	Decoder        string
	Mode           string
	IndexField     *TemplateField
	RequiredFields []TemplateField
	OptionalFields []TemplateField
}

type TemplateField struct {
	Name string
	Type string
	Tag  string
}
