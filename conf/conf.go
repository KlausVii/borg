package conf

import (
	"io/ioutil"
	"os"
	"os/user"

	flag "github.com/juju/gnuflag"
	"gopkg.in/yaml.v2"
)

var (
	// F flag prints full request
	F = flag.Bool("f", false, "Print full results, ie. no more '...'")

	// L flag limit results to a number
	L = flag.Int("l", 5, "Result list limit. Defaults to 5")

	// H flag specifies the host to connect to
	S = flag.String("s", "ok-b.org", "Server to connect to")

	H = flag.Bool("h", false, "Display help")

	Help = flag.Bool("help", false, "Display help, same as -h")

	// P flag enables private search
	P = flag.Bool("p", false, "Private search. Your search won't leave a trace. Pinky promise. Don't use this all the time if you want to see the search result relevancy improved")

	// D flag enables debug mode
	D = flag.Bool("d", false, "Debug mode")
	// DontPipe
	DontPipe = flag.Bool("dontpipe", false, "Flag for internal use - ignore this")
	// Version flag displays current version
	Version = flag.Bool("version", false, "Print version number")
	// V flag displays current version
	V = flag.Bool("v", false, "Print version number")
)
var (
	// HomeDir of the config and other files
	HomeDir string
)

func init() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	HomeDir = usr.HomeDir
	os.Mkdir(HomeDir+"/.borg", os.ModePerm)
	os.Create(HomeDir + "/.borg/edit")
	if _, err := os.Stat(HomeDir + "/.borg/config.yml"); os.IsNotExist(err) {
		os.Create(HomeDir + "/.borg/config.yml")
	}
	if _, err := os.Stat(HomeDir + "/.borg/query"); os.IsNotExist(err) {
		os.Create(HomeDir + "/.borg/query")
	}
}

// Config file
type Config struct {
	Token       string
	DefaultTags []string
	Editor      string
	PipeTo      string
}

// Save config
func (c Config) Save() error {
	bs, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(HomeDir+"/.borg/config.yml", bs, os.ModePerm)
}

// Get config
func Get() (Config, error) {
	bs, err := ioutil.ReadFile(HomeDir + "/.borg/config.yml")
	if err != nil {
		panic(err)
	}
	c := &Config{}
	err = yaml.Unmarshal(bs, c)
	if err != nil {
		return *c, err
	}
	if len(c.Editor) == 0 {
		c.Editor = "vim"
	}
	return *c, nil
}
