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

	if context.token == "" {
		trc <- &TrelloResponse{error: errors.New("Cannot request members/me without user token.")}
	}

	req := MakeGetRequest(context, "members/me", fmt.Sprintf("fields=%v", params.FieldsQueryParameter()))

	go func() {
		resp, err := context.client.Do(req)
		if err != nil {
			trc <- &TrelloResponse{error: err}
			return
		}
		defer resp.Body.Close()

		log.Printf("got response: %v", resp)

		var m Member
		decoder := json.NewDecoder(resp.Body)
		log.Println("decoder: %v", decoder)
		err = decoder.Decode(&m)
		if err != nil {
			log.Printf("error decoding json: %v", err)
			trc <- &TrelloResponse{error: err}
			return
		}
		log.Println("model: %v", m)
		trc <- &TrelloResponse{model: m}
	}()

	return trc
}
