package trello

import "net/http"
import "net/url"
import "runtime"
import "fmt"

func MakeRequest(context *context, path string, query string) *http.Request {
	req := &http.Request{}
	req.Header.Add("User-Agent", fmt.Sprintf("go %s / github.com/aggieben/trello %s", runtime.Version(), Version()))

	var err error
	if req.URL, err = url.Parse(fmt.Sprintf("https://api.trello.com/1/%s?%s", path, query)); err != nil {
		return nil
	}

	return req
}

type ResponseChannel chan interface{}

func SendRequest(client *http.Client, req *http.Request, rx ResponseChannel) {
	var msg interface{}
	if resp, err := client.Do(req); err == nil {
		msg = resp
	} else {
		msg = &err
	}

	rx <- msg
}
