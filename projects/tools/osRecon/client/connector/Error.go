package connector

// Error - Return the error state
func (svr *Connector) Error() error {
	return svr.err
}
