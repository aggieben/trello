// Copyright © Ben Collins <ben.collins@acm.org>
// See LICENSE for license.

// Package trello implements a Go-language SDK for the Trello
// (https://trello.com/docs/index.html)
package trello

import (
	"net/http"
)

func Version() string {
	return "0.1"
}

// Context is a type used to pass global, semi-const parameters to API
// requests, like access tokens, application secret, current user, etc.
type context struct {
	Version string
	Key     string
	Secret  string
	Token   string
	client  *http.Client
}

// Trello is the main type that users of the SDK will interact with.  It
// encapsulates an HTTP client and manages state.
type Trello struct {
	context *context

	Member
}

// TrelloParams provides initialization parameters to the trello client.
type TrelloParams struct {
	Version   string
	AppKey    string
	AppSecret string
	UserToken string
}

// NewTrello initializes a new Trello object using the provided parameters.
func NewTrello(params TrelloParams) *Trello {
	if params.Version == "" {
		params.Version = "1"
	}

	return &Trello{
		context: &context{params.Version, params.AppKey, params.AppSecret, params.UserToken, &http.Client{}},
	}
}
