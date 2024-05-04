package cli

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	AtlassianTypes "github.com/sam-caldwell/monorepo/go/atlassian/types"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/tools/atlassian/commands"
	"net/http"
)

// Execute Given all context, execute the API action
func (client *JiraClient[T]) Execute(action commands.Commands) (err error) {
	var (
		issueOrKey   string
		resp         *http.Response
		responseBody []byte
		web          http.Client
		request      *http.Request
		jql          *AtlassianTypes.JqlQuery
	)

	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()

	if client.debug {
		ansi.Blue().Printf("action: %s", action.String()).LF().Reset()
	}

	switch action {
	case commands.Create:
		if request, err = client.descriptor.Create(&client.domain); err != nil {
			return err
		}
	case commands.Read:
		if request, err = client.descriptor.Read(&client.domain, issueOrKey); err != nil {
			return err
		}
	case commands.Update:
		if request, err = client.descriptor.Update(&client.domain, issueOrKey); err != nil {
			return err
		}
	case commands.Delete:
		if request, err = client.descriptor.Delete(&client.domain, issueOrKey); err != nil {
			return err
		}
	case commands.List:
		if request, err = client.descriptor.List(&client.domain, jql); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown/unexpected command")
	}

	request.Header.Set(words.ContentType, words.ApplicationJson)
	if err := client.client.SetAuthHeader(request); err != nil {
		return err
	}

	if resp, err = web.Do(request); err != nil {
		return fmt.Errorf("error sending request (%v)", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error in response: %d (%s)", resp.StatusCode, resp.Status)
	}

	if _, err = resp.Body.Read(responseBody); err != nil {
		return fmt.Errorf("error reading response body (%v)", err)
	}

	if responseBody != nil {
		//ToDo: better output formatting...
		ansi.Reset().Printf("%v", responseBody).LF().Reset()
	}

	return err

}
