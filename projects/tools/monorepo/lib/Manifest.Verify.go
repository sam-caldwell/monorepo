package monorepo

import (
	"fmt"
	"strings"
)

func (m *Manifest) Verify() error {
	emptyString := func(s string) bool {
		return strings.TrimSpace(s) == ""
	}
	emptyStringList := func(s []string) bool {
		return len(s) == 0
	}

	stageTest := func(stage Stage, name string) error {
		if len(stage.Steps) < 0 {
			return fmt.Errorf("%s did not load properly", name)
		}
		for _, step := range stage.Steps {
			if emptyString(step.Command) {
				return fmt.Errorf("%s command must not be empty", name)
			}
			for _, env := range step.Environment {
				if emptyString(env.Key) {
					return fmt.Errorf("%s %s environment key cannot be empty", name, env.Key)
				}
				if len(env.Value) < 0 {
					return fmt.Errorf("%s %s environment value must exist", name, env.Key)
				}
			}
		}
		return nil
	}

	if emptyString(m.config.Header.Author) {
		return fmt.Errorf("author did not load properly")
	}
	if emptyString(m.config.Header.Email) {
		return fmt.Errorf("email did not load properly")
	}
	if emptyString(m.config.Header.Description) {
		return fmt.Errorf("description did not load properly")
	}

	if emptyStringList(m.config.Header.Arch) {
		return fmt.Errorf("arch did not load properly")
	}
	if emptyStringList(m.config.Header.Opsys) {
		return fmt.Errorf("opsys did not load properly")
	}
	if err := stageTest(m.config.Build, "build"); err != nil {
		return err
	}
	if err := stageTest(m.config.Clean, "clean"); err != nil {
		return err
	}
	if err := stageTest(m.config.Test, "test"); err != nil {
		return err
	}
	return nil
}
