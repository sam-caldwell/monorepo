package LogTarget

// TargetSet - A type of all LogTargetTypes
//
//	(c) 2023 Sam Caldwell.  MIT License
type TargetSet interface {
	StdoutTarget | FileTarget | HttpTarget | SyslogTarget
}
