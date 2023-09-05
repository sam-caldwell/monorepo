package gitignore

import (
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	. "gopkg.in/check.v1"
)

type MatcherSuite struct {
	GFS  billy.Filesystem // git repository root
	RFS  billy.Filesystem // root that contains user home
	RFSR billy.Filesystem // root that contains user home, but with relative ~/.gitignore_global
	RFSU billy.Filesystem // root that contains user home, but with relative ~user/.gitignore_global
	MCFS billy.Filesystem // root that contains user home, but missing ~/.gitconfig
	MEFS billy.Filesystem // root that contains user home, but missing excludesfile entry
	MIFS billy.Filesystem // root that contains user home, but missing .gitignore

	SFS billy.Filesystem // root that contains /etc/gitconfig
}

var _ = Suite(&MatcherSuite{})

func (s *MatcherSuite) SetUpTest(c *C) {
	// setup generic git repository root
	fs := memfs.New()

	err := fs.MkdirAll(".git/info", os.ModePerm)
	c.Assert(err, IsNil)
	f, err := fs.Create(".git/info/exclude")
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("exclude.crlf\r\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	f, err = fs.Create(".gitignore")
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("vendor/g*/\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("ignore.crlf\r\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	err = fs.MkdirAll("vendor", os.ModePerm)
	c.Assert(err, IsNil)
	f, err = fs.Create("vendor/.gitignore")
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("!github.com/\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	err = fs.MkdirAll("another", os.ModePerm)
	c.Assert(err, IsNil)
	err = fs.MkdirAll("exclude.crlf", os.ModePerm)
	c.Assert(err, IsNil)
	err = fs.MkdirAll("ignore.crlf", os.ModePerm)
	c.Assert(err, IsNil)
	err = fs.MkdirAll("vendor/github.com", os.ModePerm)
	c.Assert(err, IsNil)
	err = fs.MkdirAll("vendor/gopkg.in", os.ModePerm)
	c.Assert(err, IsNil)

	err = fs.MkdirAll("multiple/sub/ignores/first", os.ModePerm)
	c.Assert(err, IsNil)
	err = fs.MkdirAll("multiple/sub/ignores/second", os.ModePerm)
	c.Assert(err, IsNil)
	f, err = fs.Create("multiple/sub/ignores/first/.gitignore")
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("ignore_dir\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)
	f, err = fs.Create("multiple/sub/ignores/second/.gitignore")
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("ignore_dir\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)
	err = fs.MkdirAll("multiple/sub/ignores/first/ignore_dir", os.ModePerm)
	c.Assert(err, IsNil)
	err = fs.MkdirAll("multiple/sub/ignores/second/ignore_dir", os.ModePerm)
	c.Assert(err, IsNil)

	s.GFS = fs

	// setup root that contains user home
	home, err := os.UserHomeDir()
	c.Assert(err, IsNil)

	fs = memfs.New()
	err = fs.MkdirAll(home, os.ModePerm)
	c.Assert(err, IsNil)

	f, err = fs.Create(fs.Join(home, gitconfigFile))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("[core]\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("	excludesfile = " + strconv.Quote(fs.Join(home, ".gitignore_global")) + "\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	f, err = fs.Create(fs.Join(home, ".gitignore_global"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("# IntelliJ\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte(".idea/\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("*.iml\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	s.RFS = fs

	// root that contains user home, but with relative ~/.gitignore_global
	fs = memfs.New()
	err = fs.MkdirAll(home, os.ModePerm)
	c.Assert(err, IsNil)

	f, err = fs.Create(fs.Join(home, gitconfigFile))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("[core]\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("	excludesfile = ~/.gitignore_global" + "\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	f, err = fs.Create(fs.Join(home, ".gitignore_global"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("# IntelliJ\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte(".idea/\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("*.iml\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	s.RFSR = fs

	// root that contains user home, but with relative ~user/.gitignore_global
	fs = memfs.New()
	err = fs.MkdirAll(home, os.ModePerm)
	c.Assert(err, IsNil)

	f, err = fs.Create(fs.Join(home, gitconfigFile))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("[core]\n"))
	c.Assert(err, IsNil)
	currentUser, err := user.Current()
	c.Assert(err, IsNil)
	// remove domain for windows
	username := currentUser.Username[strings.Index(currentUser.Username, "\\")+1:]
	_, err = f.Write([]byte("	excludesfile = ~" + username + "/.gitignore_global" + "\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	f, err = fs.Create(fs.Join(home, ".gitignore_global"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("# IntelliJ\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte(".idea/\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("*.iml\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	s.RFSU = fs

	// root that contains user home, but missing ~/.gitconfig
	fs = memfs.New()
	err = fs.MkdirAll(home, os.ModePerm)
	c.Assert(err, IsNil)

	f, err = fs.Create(fs.Join(home, ".gitignore_global"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("# IntelliJ\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte(".idea/\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("*.iml\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	s.MCFS = fs

	// setup root that contains user home, but missing excludesfile entry
	fs = memfs.New()
	err = fs.MkdirAll(home, os.ModePerm)
	c.Assert(err, IsNil)

	f, err = fs.Create(fs.Join(home, gitconfigFile))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("[core]\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	f, err = fs.Create(fs.Join(home, ".gitignore_global"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("# IntelliJ\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte(".idea/\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("*.iml\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	s.MEFS = fs

	// setup root that contains user home, but missing .gitnignore
	fs = memfs.New()
	err = fs.MkdirAll(home, os.ModePerm)
	c.Assert(err, IsNil)

	f, err = fs.Create(fs.Join(home, gitconfigFile))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("[core]\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("	excludesfile = " + strconv.Quote(fs.Join(home, ".gitignore_global")) + "\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	s.MIFS = fs

	// setup root that contains user home
	fs = memfs.New()
	err = fs.MkdirAll("etc", os.ModePerm)
	c.Assert(err, IsNil)

	f, err = fs.Create(systemFile)
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("[core]\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("	excludesfile = /etc/gitignore_global\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	f, err = fs.Create("/etc/gitignore_global")
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("# IntelliJ\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte(".idea/\n"))
	c.Assert(err, IsNil)
	_, err = f.Write([]byte("*.iml\n"))
	c.Assert(err, IsNil)
	err = f.Close()
	c.Assert(err, IsNil)

	s.SFS = fs
}

func (s *MatcherSuite) TestDir_ReadPatterns(c *C) {
	checkPatterns := func(ps []Pattern) {
		c.Assert(ps, HasLen, 6)
		m := NewMatcher(ps)

		c.Assert(m.Match([]string{"exclude.crlf"}, true), Equals, true)
		c.Assert(m.Match([]string{"ignore.crlf"}, true), Equals, true)
		c.Assert(m.Match([]string{"vendor", "gopkg.in"}, true), Equals, true)
		c.Assert(m.Match([]string{"vendor", "github.com"}, true), Equals, false)
		c.Assert(m.Match([]string{"multiple", "sub", "ignores", "first", "ignore_dir"}, true), Equals, true)
		c.Assert(m.Match([]string{"multiple", "sub", "ignores", "second", "ignore_dir"}, true), Equals, true)
	}

	ps, err := ReadPatterns(s.GFS, nil)
	c.Assert(err, IsNil)
	checkPatterns(ps)

	// passing an empty slice with capacity to check we don't hit a bug where the extra capacity is reused incorrectly
	ps, err = ReadPatterns(s.GFS, make([]string, 0, 6))
	c.Assert(err, IsNil)
	checkPatterns(ps)
}

func (s *MatcherSuite) TestDir_ReadRelativeGlobalGitIgnore(c *C) {
	for _, fs := range []billy.Filesystem{s.RFSR, s.RFSU} {
		ps, err := LoadGlobalPatterns(fs)
		c.Assert(err, IsNil)
		c.Assert(ps, HasLen, 2)

		m := NewMatcher(ps)
		c.Assert(m.Match([]string{".idea/"}, true), Equals, false)
		c.Assert(m.Match([]string{"*.iml"}, true), Equals, true)
		c.Assert(m.Match([]string{"IntelliJ"}, true), Equals, false)
	}
}

func (s *MatcherSuite) TestDir_LoadGlobalPatterns(c *C) {
	ps, err := LoadGlobalPatterns(s.RFS)
	c.Assert(err, IsNil)
	c.Assert(ps, HasLen, 2)

	m := NewMatcher(ps)
	c.Assert(m.Match([]string{"go-git.v4.iml"}, true), Equals, true)
	c.Assert(m.Match([]string{".idea"}, true), Equals, true)
}

func (s *MatcherSuite) TestDir_LoadGlobalPatternsMissingGitconfig(c *C) {
	ps, err := LoadGlobalPatterns(s.MCFS)
	c.Assert(err, IsNil)
	c.Assert(ps, HasLen, 0)
}

func (s *MatcherSuite) TestDir_LoadGlobalPatternsMissingExcludesfile(c *C) {
	ps, err := LoadGlobalPatterns(s.MEFS)
	c.Assert(err, IsNil)
	c.Assert(ps, HasLen, 0)
}

func (s *MatcherSuite) TestDir_LoadGlobalPatternsMissingGitignore(c *C) {
	ps, err := LoadGlobalPatterns(s.MIFS)
	c.Assert(err, IsNil)
	c.Assert(ps, HasLen, 0)
}

func (s *MatcherSuite) TestDir_LoadSystemPatterns(c *C) {
	ps, err := LoadSystemPatterns(s.SFS)
	c.Assert(err, IsNil)
	c.Assert(ps, HasLen, 2)

	m := NewMatcher(ps)
	c.Assert(m.Match([]string{"go-git.v4.iml"}, true), Equals, true)
	c.Assert(m.Match([]string{".idea"}, true), Equals, true)
}
