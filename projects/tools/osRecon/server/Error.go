package server

// Error - return the server's current error state
func (svr *Server) Error() error {
	return svr.err
}
