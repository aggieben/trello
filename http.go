package trello

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func MakeGetRequest(context *context, path string, query string) *http.Request {
	baseUrl := fmt.Sprintf("https://api.trello.com")
	if context.baseUrl != "" {
		baseUrl = context.baseUrl
	}

	queryMap := map[string]string{
		"token": context.token,
		"key":   context.key,
	}

	var internalQueryStrings []string
	for k := range queryMap {
		if queryMap[k] != "" {
			internalQueryStrings = append(internalQueryStrings, fmt.Sprintf("%s=%s", k, queryMap[k]))
		}
	}

	if query != "" {
		query = fmt.Sprintf("%s&%s", query, strings.Join(internalQueryStrings, "&"))
	}

	var req *http.Request = nil
	var err error
	if req, err = http.NewRequest("GET", fmt.Sprintf("%s/1/%s?%s", baseUrl, path, query), nil); err != nil {
		log.Fatalln("Unable to create HTTP request: %v", err)
	}

	req.Header.Add("User-Agent",
		fmt.Sprintf("%s/github.com/aggieben/trello %s", runtime.Version(), Version()))

	if context.token != "" {
		req.URL.Query().Add("token", context.token)
		log.Println("added token parameter")
	}

	log.Printf("created request: %v", req)
	return req
}

type TrelloResponse struct {
	Model Model
	Error interface{}
}
