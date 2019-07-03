package main

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func findByID(model string) (string, error) {
	const body = `func {{ .funcName }}(db *gorm.DB, id int) (*{{ .model }}, error) {
	var obj {{ .model }}
	if err := db.First(&obj, id); err != nil {
		return nil, xerrors.Errorf("failed to find {{ .model }} with id %v: %w", id, err)
	}
	return &obj, nil
}`
	modelCmp := strings.Split(model, ".")
	structName := modelCmp[len(modelCmp)-1]
	return renderTpl(body, map[string]interface{}{
		"funcName": fmt.Sprintf("find%sByID", structName),
		"model":    model,
	})
}

func findByParam(model string, paramName string, paramType string) (string, error) {
	const body = `func {{ .funcName }}(db *gorm.DB, arg {{ .paramType }}) (*{{ .model }}, error) {
	var obj {{ .model }}
	if err := db.Find(&obj, "{{ .condition }}", arg); err != nil {
		return nil, xerrors.Errorf("failed to find {{ .model }} by {{ .paramName }}: %w", err)
	}
	return &obj, nil
}`
	modelCmp := strings.Split(model, ".")
	structName := modelCmp[len(modelCmp)-1]
	colName := strcase.ToSnake(paramName)
	return renderTpl(body, map[string]interface{}{
		"funcName":  fmt.Sprintf("find%sBy%s", structName, paramName),
		"model":     model,
		"condition": fmt.Sprintf("`%s` = ?", colName),
		"paramName": paramName,
		"paramType": paramType,
	})
}
