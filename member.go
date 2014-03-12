package trello

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Member struct {
	Id       string
	FullName string
	Username string
}

var minimalFields = [...]ModelField{FullName, Username}

func (m Member) MinimalFields() []ModelField {
	return minimalFields[:]
}

// Get Member model of authenticated user
func (_ *Member) Me(_context interface{}, params *ModelParams) <-chan *TrelloResponse {
	trc := make(chan *TrelloResponse)
	context := _context.(*context)
	go func() {
		if context.token == "" {
			trc <- &TrelloResponse{Error: errors.New("Cannot request members/me without user token.")}
		}

		req := MakeGetRequest(context, "members/me", fmt.Sprintf("fields=%v", params.FieldsQueryParameter()))

		resp, err := context.client.Do(req)
		if err != nil {
			trc <- &TrelloResponse{Error: err}
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			trc <- &TrelloResponse{Error: errors.New(resp.Status)}
			return
		}

		var m Member
		decoder := json.NewDecoder(resp.Body)
		if err = decoder.Decode(&m); err == io.EOF {
			trc <- &TrelloResponse{}
		} else if err != nil {
			trc <- &TrelloResponse{Error: err}
			return
		}

		trc <- &TrelloResponse{Model: m}
	}()

	return trc
}
