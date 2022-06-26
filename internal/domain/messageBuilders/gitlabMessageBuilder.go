package messagebuilders

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
)

type GitlabMessageBuilder struct {
}

func (gmb *GitlabMessageBuilder) Build(event []byte, templates map[string]string) (string, error) {
	m := make(map[string]interface{})
	if err := json.Unmarshal(event, &m); err != nil {
		return "", err
	}
	objKind, ok := m["object_kind"]
	if !ok {
		return "", errors.New(`failed to get "object_kind" value`)
	}
	mTempl, ok := templates[objKind.(string)]
	if !ok {
		return "", fmt.Errorf("failed to get template of type: %v from map", objKind)
	}
	t, err := template.New("").Parse(mTempl)
	if err != nil {
		return "", err
	}
	buf := bytes.Buffer{}
	if err := t.Execute(&buf, &m); err != nil {
		return "", err
	}
	return buf.String(), nil
}
