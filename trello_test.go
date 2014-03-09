package trello

import "testing"

func TestNewTrelloHasContext(t *testing.T) {
	trello := NewTrello("blah", "dee", "dah")
	if trello.context == nil {
		t.Error("Expected allocated and initialized context, got ", trello.context, " instead.")
	}
}
