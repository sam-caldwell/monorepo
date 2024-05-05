package Atlassian

import (
	"bytes"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"net/http"
)

func (client *Client) Send(method string, path string, body []byte) (output []byte, err error) {

	var (
		request      *http.Request
		resp         *http.Response
		responseBody []byte
		web          http.Client
	)

	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()

	request, err = http.NewRequest(
		method,
		JiraUrlFactory(JiraUrlPattern, client.domain.Get(), path),
		bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	request.Header.Set(words.ContentType, words.ApplicationJson)
	if err = client.SetAuthHeader(request); err != nil {
		return nil, err
	}

	if client.noop {
		ansi.Blue().Printf("noop: web request would be sent").LF().Reset()
		return []byte("noop-1"), nil
	}

	if resp, err = web.Do(request); err != nil {
		return nil, fmt.Errorf("error sending request (%v)", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error in response: %d (%s)", resp.StatusCode, resp.Status)
	}

	if _, err = resp.Body.Read(responseBody); err != nil {
		return nil, fmt.Errorf("error reading response body (%v)", err)
	}

	if responseBody != nil {
		//ToDo: better output formatting...
		ansi.Reset().Printf("%v", responseBody).LF().Reset()
	}
	return output, err
}
