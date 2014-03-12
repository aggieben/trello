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
		assert.Equal(t, r.URL.Query(), url.Values{
			"fields": []string{"fullName,username"},
			"key":    []string{"key"},
			"token":  []string{"token"}})

		fmt.Fprintln(w, `{"id":"4f11b44baf3eab192c009ff7","fullName":"Benjamin Collins","username":"aggieben"}"`)
	}))
	defer svr.Close()

	trello := NewTrello(TrelloParams{Version: "1", AppKey: "key", UserToken: "token", baseUrl: svr.URL})
	assert.NotNil(t, trello)

	rx := trello.Members.Me(trello.Context, &ModelParams{Fields: trello.Members.MinimalFields()})

	resp := <-rx
	assert.Nil(t, resp.Error, "err: %v", resp.Error)
	assert.NotNil(t, resp.Model)

	fmt.Printf("got model of type %T: %v\n", resp.Model, resp.Model)
}

func TestMemberGetMeWithoutUserToken(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Fail(t, "error: request should never have been made.")
	}))
	defer svr.Close()

	trello := NewTrello(TrelloParams{Version: "1", AppKey: "key", baseUrl: svr.URL})
	assert.NotNil(t, trello)

	rx := trello.Members.Me(trello.Context, &ModelParams{Fields: trello.Members.MinimalFields()})
	resp := <-rx

	err, ok := resp.Error.(error)
	assert.True(t, ok, "error was not an error after all")
	assert.Error(t, err, "without a token, Members.Me should return an error.")

	fmt.Printf("successfully received error: %v\n", err)
}

func TestMemberGetMeWithHttp404(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer svr.Close()

	trello := NewTrello(TrelloParams{Version: "1", AppKey: "key", UserToken: "token", baseUrl: svr.URL})
	rx := trello.Members.Me(trello.Context, &ModelParams{Fields: trello.Members.MinimalFields()})
	resp := <-rx

	assert.Nil(t, resp.Model, "expecting no model in response")
	err, ok := resp.Error.(error)
	assert.True(t, ok, "error was not an error after all")
	assert.Error(t, err, "expecting http error response")

	fmt.Printf("received error in response: %v\n", resp.Error)
}

func TestMemberGetMeWithJsonErrors(t *testing.T) {
	responseIndex := 0
	jsonResponses := []string{`[1, "string",`, `[1, "string",,:]`, `[1, 2, 3]`}
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, jsonResponses[responseIndex])
		responseIndex += 1
	}))
	defer svr.Close()

	trello := NewTrello(TrelloParams{Version: "1", AppKey: "key", UserToken: "token", baseUrl: svr.URL})
	for i := 0; i < len(jsonResponses); i++ {
		resp := <-trello.Members.Me(trello.Context, &ModelParams{Fields: trello.Members.MinimalFields()})
		assert.Nil(t, resp.Model, "expecting no model in response")
		err, ok := resp.Error.(error)
		assert.True(t, ok, "error was not an error after all")
		assert.Error(t, err, "expecting a json error response")

		fmt.Printf("received error in response: %v\n", resp.Error)
	}
}
