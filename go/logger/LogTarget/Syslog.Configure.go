package LogTarget

import configuration "github.com/sam-caldwell/monorepo/go/configuration/Map"

func (out *SyslogTarget) Configure(cfg configuration.Map[string, string]) {
	//ToDo: parse cfg for host, port, etc. and configure the Syslog Target object.
}
