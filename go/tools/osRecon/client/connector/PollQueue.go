package connector

import (
	"encoding/json"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/server"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/threatQL"
	"net/http"
)

// PollQueue - Poll server for queries in queue
func (svr *Connector) PollQueue(queue chan<- threatQL.Query) *Connector {
	response, err := http.Get(fmt.Sprintf("https://%s:%d%s", svr.host, svr.port, server.ApiV1CheckIn))
	if err == nil {
		switch response.StatusCode {
		case http.StatusOK:
			//Todo: update stats
			decoder := json.NewDecoder(response.Body)
			var query threatQL.Query
			err := decoder.Decode(&query)
			if err != nil {
				//Todo: update stats
				ansi.Red().Printf("Invalid Server Query: %v", err).Reset()
			} else {
				//Todo: update stats
				queue <- query
			}
		case http.StatusUnauthorized:
			//Todo: update stats
			ansi.Red().Println("Client is unauthorized").Fatal(exit.GeneralError).Reset()
		default:
			//Todo: update stats
			ansi.Red().Println("Unexpected server response").Reset()
		}
	}
	return svr
}
