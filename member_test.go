package trello

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestMemberGetMeWithMinimalFields(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.URL.Path, "/1/members/me")
		assert.Equal(t, r.URL.Query(), url.Values{"fields": []string{"fullName,username"}})

		fmt.Fprintln(w, `{"id":"4f11b44baf3eab192c009ff7","fullName":"Benjamin Collins","username":"aggieben"}"`)
	}))
	defer svr.Close()

	trello := NewTrello(TrelloParams{Version: "1", AppKey: "key", UserToken: "token", baseUrl: svr.URL})
	assert.NotNil(t, trello)

	rx := trello.Members.Me(trello.context, &ModelParams{Fields: trello.Members.MinimalFields()})

	resp := <-rx
	assert.Nil(t, resp.error, "err: %v", resp.error)
	assert.NotNil(t, resp.model)

	fmt.Printf("got model of type %T: %v\n", resp.model, resp.model)
}

func TestMemberGetMeWithoutUserToken(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Fail(t, "error: request should never have been made.")
	}))
	defer svr.Close()

	trello := NewTrello(TrelloParams{Version: "1", AppKey: "key", baseUrl: svr.URL})
	assert.NotNil(t, trello)

	rx := trello.Members.Me(trello.context, &ModelParams{Fields: trello.Members.MinimalFields()})
	resp := <-rx

	err, ok := resp.error.(error)
	assert.True(t, ok, "error was not an error after all")
	assert.Error(t, err, "without a token, Members.Me should return an error.")

	fmt.Printf("successfully received error: %v\n", err)
}
