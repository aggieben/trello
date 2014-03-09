// Copyright © Ben Collins <ben.collins@acm.org>
// See LICENSE for license.

// Package trello implements a Go-language SDK for the Trello
// (https://trello.com/docs/index.html)
package trello

import ()

// Context is a type used to pass global, semi-const parameters to API
// requests, like access tokens, application secret, current user, etc.
type context struct {
	Key    string
	Secret string
	Token  string
}

// Trello is the main type that users of the SDK will interact with.  It
// encapsulates an HTTP client and manages state.
type Trello struct {
	context *context
}

// NewTrello initializes a new Trello object with the provided appKey,
// appSecret, and userToken
func NewTrello(appKey string, appSecret string, userToken string) *Trello {
	return &Trello{&context{appKey, appSecret, userToken}}
}
