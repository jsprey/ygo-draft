package query

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"
	"ygodraft/backend/customerrors"
)

// ErrorTemplateDoesNotExist this error is returned when a template with a given name does not exist.
var ErrorTemplateDoesNotExist = customerrors.WithCode{
	Code:        "", // TODO add code
	InternalMsg: "template %s does not exist",
}

type sqlQueryTemplater struct {
	Templates map[string]*template.Template
}

// NewSqlQueryTemplater create a new object responsible to parse and template sql queries.
func NewSqlQueryTemplater() (*sqlQueryTemplater, error) {
	templater := sqlQueryTemplater{Templates: map[string]*template.Template{}}

	templateStringMap := &map[string]string{}
	templater.AddCardTemplates(templateStringMap)
	templater.AddUserTemplates(templateStringMap)
	templater.AddFriendsTemplates(templateStringMap)

	for templateName, templateString := range *templateStringMap {
		parsedTemplate, err := template.New(templateName).Funcs(customFunctions).Parse(templateString)
		if err != nil {
			return nil, fmt.Errorf("failed to parse template [%s]: %w", templateName, err)
		}

		templater.Templates[templateName] = parsedTemplate
	}

	return &templater, nil
}

// Template receives a template name and an objects and templates the correct template with the values from the object.
func (sqt *sqlQueryTemplater) Template(templateName string, templateObject any) (string, error) {
	sqlTemplate, ok := sqt.Templates[templateName]
	if !ok {
		return "", ErrorTemplateDoesNotExist.WithParam(templateName)
	}

	buf := new(bytes.Buffer)
	err := sqlTemplate.Execute(buf, templateObject)
	if err != nil {
		return "", fmt.Errorf("failed to execute sqlTemplate [%s]: %w", templateName, err)
	}

	return buf.String(), nil
}

var customFunctions = template.FuncMap{
	"notLast": func(x int, a interface{}) bool {
		return x < reflect.ValueOf(a).Len()-1
	},
}

func escape(input string) string {
	return fmt.Sprintf("'%s'", strings.Replace(input, "'", "''", -1))
}
