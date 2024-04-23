package connector

// Connect - Connect to the server and authenticate the connection.
func (svr *Connector) Connect() *Connector {
	//ToDo: create a persistent connection between client and server
	//      note that if the connection already exists and is okay,
	//      we will just reuse the existing connection.
	return svr
}
