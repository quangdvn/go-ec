package utils

import (
	"bytes"
	"fmt"
	"text/template"
)

func GetMailTemplate(name string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)

	t := template.Must(template.New(name).ParseFiles(fmt.Sprintf("mailTemplates/%s", name)))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}
