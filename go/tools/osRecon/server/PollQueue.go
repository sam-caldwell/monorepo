package server

import "github.com/sam-caldwell/monorepo/go/types"

// PollQueue - Poll server for queries in queue
func (svr *Server) PollQueue(queue chan<- types.ThreatQlQuery) *Server {
	return svr
}
