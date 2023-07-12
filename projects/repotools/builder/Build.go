package repoBuilder

import (
	"github.com/sam-caldwell/go/v2/projects/simpleLogger"
)

// Build - Run all builds for the repo or for a specific project path
func Build(logf simpleLogger.Logf, noop bool) (err error) {
	//const buildCommand = "go build %s"
	//
	//var projectList []string
	//if projectList, err = repotools.ListProjects("cmd", projectfilter.Enabled); err != nil {
	//	return err
	//}
	//
	//for _, project := range projectList {
	//
	//	logf(ansi.Blue(), "Build %s...", project)
	//
	//	if noop {
	//		logf(ansi.Yellow(), "NOOP: Skipping %s", project)
	//		continue
	//	}
	//
	//	if out, err := runcommand.ShellExecute(fmt.Sprintf(buildCommand, project)); err != nil {
	//		logf(ansi.Red(),
	//			"Build failed on %s\n"+
	//				"Error:%v\n"+
	//				"Details:%s\n",
	//			project, err, out)
	//		return err
	//	}
	//
	//	logf(ansi.Green(), "Build ok: %s", project)
	//}
	//
	//logf(ansi.Green(), "Build successful")

	return nil

}
