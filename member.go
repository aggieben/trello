package trello

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
func (_ *Member) Me(context *context, params *ModelParams) <-chan *TrelloResponse {
	trc := make(chan *TrelloResponse)

	go func() {
		if context.token == "" {
			trc <- &TrelloResponse{error: errors.New("Cannot request members/me without user token.")}
		}

		req := MakeGetRequest(context, "members/me", fmt.Sprintf("fields=%v", params.FieldsQueryParameter()))

		resp, err := context.client.Do(req)
		if err != nil {
			trc <- &TrelloResponse{error: err}
			return
		}
		defer resp.Body.Close()

		log.Printf("got response: %v\n", resp)

		var m Member
		decoder := json.NewDecoder(resp.Body)
		log.Printf("decoder: %v\n", decoder)
		err = decoder.Decode(&m)
		if err != nil {
			log.Printf("error decoding json: %v\n", err)
			trc <- &TrelloResponse{error: err}
			return
		}
		log.Printf("model: %v\n", m)
		trc <- &TrelloResponse{model: m}
	}()

	return trc
}
