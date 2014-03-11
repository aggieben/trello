package trello

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModelFieldsQueryParameterWithOne(t *testing.T) {
	var fields = [...]ModelField{FullName}
	mp := &ModelParams{Fields: fields[:]}

	qp := mp.FieldsQueryParameter()
	assert.Equal(t, fmt.Sprintf("%v", FullName), qp)
}

func TestModelFieldsQueryParameterWithTwo(t *testing.T) {
	var fields = [...]ModelField{FullName, Username}
	mp := &ModelParams{Fields: fields[:]}

	qp := mp.FieldsQueryParameter()
	assert.Equal(t, fmt.Sprintf("%v,%v", FullName, Username), qp)
}
