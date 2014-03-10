package trello

import (
	"errors"
	"fmt"
)

type Member struct {
	Id       string
	FullName string
	Username string
}

var minimalFields = [...]ModelField{FullName, Username}

func (m *Member) MinimalFields() []ModelField {
	return minimalFields[:]
}

// Get Member model of authenticated user
func (m *Member) Me(context *context, params *ModelParams) (ResponseChannel, error) {
	if context.Token == "" {
		return nil, errors.New("Cannot request members/me without user token.")
	}

	rc := make(ResponseChannel)
	req := MakeRequest(context, "members/me", "")

	go SendRequest(context.client, req, rc)

	return rc, nil
}

// Get Member by username
func (m *Member) Get(context *context, params *ModelParams, username string) (ResponseChannel, *error) {
	rc := make(ResponseChannel)
	req := MakeRequest(context, fmt.Sprintf("members/%s", username), "")

	go SendRequest(context.client, req, rc)

	return rc, nil
}
