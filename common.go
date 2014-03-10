package trello

type ModelField string

const (
	FullName ModelField = "fullName"
	Username ModelField = "username"
)

type ModelParams interface {
	Fields() *[]ModelField
}

type Model interface {
	MinimalFieldSet() *[]ModelField
}
