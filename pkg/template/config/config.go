package config

var ETC_TPL = `debug: true
address: "0.0.0.0:{{.ProjectPort}}"
logfile: "/tmp/{{.ProjectName}}.log"
secret: ""

mysql:
  master: "root:@tcp(127.0.0.1:3306)/mysql?charset=utf8"
  slave1: "root:@tcp(127.0.0.1:3306)/mysql?charset=utf8"
`

var CONFIG_TPL = `package config

import (
	"os"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

type Settings struct {
	Debug   bool      ` + "`" + `yaml:"debug"` + "`" + `
	Address string    ` + "`" + `yaml:"address"` + "`" + `
	LogFile string    ` + "`" + `yaml:"logfile"` + "`" + `
	Secret  string    ` + "`" + `yaml:"secret"` + "`" + `
	Mysql   MysqlInfo ` + "`" + `yaml:"mysql"` + "`" + `
}

type MysqlInfo struct {
	Master string ` + "`" + `yaml:"master"` + "`" + `
	Slave  string ` + "`" + `yaml:"slave"` + "`" + `
}

var (
	setting Settings
	lock    = new(sync.RWMutex)
)

func ParseConfig(cfgFile string) {
	buf, err := os.ReadFile(cfgFile)
	if err != nil {
		panic(err)
	}

	lock.Lock()
	defer lock.Unlock()

	if err := yaml.Unmarshal(buf, &setting); err != nil {
		panic(err)
	}
}

func Config() Settings {
	lock.RLock()
	defer lock.RUnlock()
	return setting
}
`
