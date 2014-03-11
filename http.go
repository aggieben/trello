package trello

import "net/http"

import "runtime"
import "fmt"
import "log"

const apiHost = "api.trello.com"

func MakeGetRequest(context *context, path string, query string) *http.Request {
	baseUrl := fmt.Sprintf("https://%s/1", apiHost)
	if context.baseUrl != "" {
		baseUrl = context.baseUrl
	}

	var req *http.Request = nil
	var err error
	if req, err = http.NewRequest("GET", fmt.Sprintf("%s/1/%s?%s", baseUrl, path, query), nil); err != nil {
		log.Fatalln("Unable to create HTTP request: %v", err)
	}

	req.Header.Add("User-Agent", fmt.Sprintf("go %s / github.com/aggieben/trello %s", runtime.Version(), Version()))

	return req
}

type TrelloResponse struct {
	model Model
	error interface{}
}
