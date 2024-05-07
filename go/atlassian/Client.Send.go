package Atlassian

import (
	"bytes"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"io"
	"log"
	"net/http"
)

// Send - Send HTTP request for client and return an error and []byte serialized output.
func (client *Client) Send(method string, path string, body []byte) (output []byte, err error) {

	var (
		request *http.Request
		resp    *http.Response
		web     http.Client
	)

	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()

	if client.debug {
		ansi.Blue().Printf("path: %s", path).LF().Reset()
	}

	request, err = http.NewRequest(
		method,
		JiraUrlFactory(JiraUrlPattern, client.domain.Get(), path),
		bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	request.Header.Set(words.ContentType, ContentType)

	if err = client.SetAuthHeader(request); err != nil {
		return nil, err
	}

	if client.debug {
		ansi.Blue().
			Line("-", 40).
			Println("request headers")
		for key, value := range request.Header {
			ansi.Blue().Printf("  %s  : %s", key, value).LF().Reset()
		}
		ansi.Line("-", 40).
			LF().Reset()
	}

	if client.noop {
		ansi.Blue().Printf("noop: web request would be sent").LF().Reset()
		return []byte("noop-1"), nil
	}

	if resp, err = web.Do(request); err != nil {
		return nil, fmt.Errorf("error sending request (%v)", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("error in response: %d (%s)", resp.StatusCode, resp.Status)
	}

	if output, err = io.ReadAll(resp.Body); err != nil {
		log.Fatalln(err)
	}

	ansi.Reset().
		Printf("Result:OK").LF().
		Printf("%s", output).LF().
		Reset()

	return output, err
}
