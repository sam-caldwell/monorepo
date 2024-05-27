package LogTarget

import configuration "github.com/sam-caldwell/monorepo/go/configuration/Map"

func (out *StdoutTarget) Configure(cfg configuration.Map[string, string]) {
	//ToDo: Do nothing because stdout should not really need configuration
}
