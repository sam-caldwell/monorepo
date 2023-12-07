package monorepo

import "github.com/sam-caldwell/monorepo/go/ansi"

// Clean - Run project-specific clean-up operations (as defined in the project manifest).
func (m *Monorepo) Clean() (err error) {
	pass := 0
	fail := 0
	for _, manifest := range m.manifestList {
		if err = manifest.Run("clean", &m.Root, m.Debug); err != nil {
			fail++
			return err
		} else {
			pass++
		}
	}
	ansi.Cyan().Printf("Statistics\n"+
		"\tpass: %d\n"+
		"\tfail: %d\n", pass, fail).
		LF().Reset()
	return err
}
