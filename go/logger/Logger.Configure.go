package logger

import (
	configuration "github.com/sam-caldwell/monorepo/go/configuration/Map"
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/ratelimiter/tokensPerSecond"
	"strings"
)

// Configure - Configure the logger
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Configure(cfg *configuration.Map[string, string]) *Logger {
	if cfg == nil {
		//Just accept default state (syslog logging)
		return log.DefaultConfiguration()
	}
	switch target := cfg.ExpectOrIgnore(words.Target); strings.TrimSpace(strings.ToLower(target)) {
	case words.File:
		log.ConfigureFile(cfg)
	case words.Syslog:
		log.ConfigureSyslog(cfg)
	case words.Http:
		log.ConfigureHttp(cfg)
	case words.Stdout:
		log.ConfigureStdout(cfg)
	default:
		log.ConfigureStdout(cfg)
	}
	log.ratelimit = tokensPerSecond.NewRateLimiter(int(convert.StringToUint(cfg.ExpectOrIgnore(words.RateLimit))))
	log.appName = strings.TrimSpace(strings.ToLower(cfg.ExpectOrIgnore(words.AppName)))
	log.msgId = cfg.ExpectOrIgnore(words.MsgId)
	return log
}
