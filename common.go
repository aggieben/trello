package trello

import (
	"bytes"
	"fmt"
)

type ModelField string

const (
	FullName ModelField = "fullName"
	Username ModelField = "username"
)

type ModelParams struct {
	Fields []ModelField
}

// FieldString returns a URL-friendly query parameter value, like
// "field1,field2,..."
func (mp *ModelParams) FieldsQueryParameter() string {
	var buf bytes.Buffer
	for i, f := range mp.Fields {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("%s", f))
	}
	return buf.String()
}

type Model interface {
	MinimalFields() []ModelField
}
