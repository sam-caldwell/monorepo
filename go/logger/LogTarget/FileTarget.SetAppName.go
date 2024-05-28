package LogTarget

// SetAppName - Set the Application Name
//
//	(c) 2023 Sam Caldwell.  MIT License
func (out *FileTarget) SetAppName(appName *string) {
	out.appName = *appName
}
