package main

import (
	"fmt"
	"strings"
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
