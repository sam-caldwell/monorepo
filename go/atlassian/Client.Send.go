package Atlassian

import (
	"bytes"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/version"
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

	url := JiraUrlFactory(JiraUrlPattern, client.domain.Get(), path)
	if client.debug {
		ansi.Blue().Printf("domain: %s", client.domain.Get()).LF().Reset()
		ansi.Blue().Printf("path: %s", path).LF().Reset()
		ansi.Blue().Printf("url: %s", url).LF().Reset()
	}

	request, err = http.NewRequest(
		method,
		url,
		bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	request.Header.Set(words.ContentType, ContentType)
	request.Header.Set("User-Agent", fmt.Sprintf("jira-cli/%s", version.Version))

	if err = client.SetAuthHeader(request); err != nil {
		return nil, err
	}

	if client.noop {
		ansi.Blue().Printf("noop: web request would be sent").LF().Reset()
		return []byte("noop-1"), nil
	}

	if resp, err = web.Do(request); err != nil {
		return nil, fmt.Errorf("error sending request (%v): %v", err, resp.Body)
	}

	if output, err = io.ReadAll(resp.Body); err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("error in response: %d (%s) output:%s",
			resp.StatusCode, resp.Status, output)
	}

	return output, err
}
