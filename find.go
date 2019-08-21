package snippet

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

// FindByID find object by id
func FindByID(model string) (string, error) {
	const body = `func {{ .funcName }}(db *gorm.DB, id int) (*{{ .model }}, error) {
	var obj {{ .model }}
	if err := db.Find(&obj, id).Error; err != nil {
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

// FindByParam find object by param value
func FindByParam(model string, paramName string, paramType string) (string, error) {
	const body = `func {{ .funcName }}(db *gorm.DB, arg {{ .paramType }}) (*{{ .model }}, error) {
{{ .argValidation }}
	var obj {{ .model }}
	if err := db.Find(&obj, "{{ .condition }}", arg).Error; err != nil {
		return nil, xerrors.Errorf("failed to find {{ .model }} by {{ .paramName }}: %w", err)
	}
	return &obj, nil
}`
	modelCmp := strings.Split(model, ".")
	structName := modelCmp[len(modelCmp)-1]
	colName := strcase.ToSnake(paramName)
	argValidation := ""
	if paramType == "string" {
		argValidation = fmt.Sprintf(`	if len(arg) == 0 {
		return nil, xerrors.Errorf("%s must be non-nil.")
	}`, paramName)
	}
	return renderTpl(body, map[string]interface{}{
		"funcName":      fmt.Sprintf("find%sBy%s", structName, paramName),
		"model":         model,
		"condition":     fmt.Sprintf("`%s` = ?", colName),
		"paramName":     paramName,
		"paramType":     paramType,
		"argValidation": argValidation,
	})
}

// BatchFindByID find objects for every id at once
func BatchFindByID(model string) (string, error) {
	const body = `func {{ .funcName }}(db *gorm.DB, idSlice []int) ([]*{{ .model }}, error) {
	if len(idSlice) == 0 {
		return nil, nil
	}
	var objects []*{{ .model }}
	if err := db.Find(&objects, "id in (?)", idSlice).Error; err != nil {
		return nil, xerrors.Errorf("failed to find {{ .model }} by id slice %v: %w", idSlice, err)
	}
	return objects, nil
}`
	modelCmp := strings.Split(model, ".")
	structName := modelCmp[len(modelCmp)-1]
	return renderTpl(body, map[string]interface{}{
		"funcName": fmt.Sprintf("batchFind%sByID", structName),
		"model":    model,
	})
}

// BatchFindByParam find objects for every id at once
func BatchFindByParam(model string, paramName string, paramType string) (string, error) {
	const body = `func {{ .funcName }}(db *gorm.DB, a []{{ .paramType }}) ([]*{{ .model }}, error) {
	if len(a) == 0 {
		return nil, nil
	}
	var objects []*{{ .model }}
	if err := db.Find(&objects, "{{ .condition }}", a).Error; err != nil {
		return nil, xerrors.Errorf("failed to find {{ .model }} by {{ .paramName }} slice %v: %w", a, err)
	}
	return objects, nil
}`
	modelCmp := strings.Split(model, ".")
	structName := modelCmp[len(modelCmp)-1]
	colName := strcase.ToSnake(paramName)
	return renderTpl(body, map[string]interface{}{
		"funcName":  fmt.Sprintf("batchFind%sBy%s", structName, paramName),
		"model":     model,
		"condition": fmt.Sprintf("`%s` in (?)", colName),
		"paramName": paramName,
		"paramType": paramType,
	})
}
