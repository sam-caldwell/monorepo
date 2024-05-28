package LogTarget

// SetAppName - Set the Application Name
//
//	(c) 2023 Sam Caldwell.  MIT License
func (out *HttpTarget) SetAppName(appName *string) {
	out.appName = *appName
}
