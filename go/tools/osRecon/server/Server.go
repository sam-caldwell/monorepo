package server

import (
	"github.com/sam-caldwell/monorepo/go/net"
	"github.com/sam-caldwell/monorepo/go/types"
	"golang.org/x/net/http2"
)

type Server struct {
	host   net.Fqdn       //the server's hostname
	port   net.PortNumber // the tcp port number used for the server connect
	apiKey types.ApiKey   // The auth key used to connect ot the server
	err    error          //The server connector's error state
	//Todo: add some stats

	server http2.ClientConnPool
}
