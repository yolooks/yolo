package model

var TPL = `package model

import (
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

const (
	ROLE_ADMIN = "admin"
	ROLE_RD    = "rd"
	ROLE_OP    = "op"
)

var (
	DefaultDriver = "mysql"
	NotFound      = fmt.Errorf("query data not found")
)

var (
	once    sync.Once
	MEngine *xorm.Engine
	SEngine *xorm.Engine
)

func Connect(driver, master string, slaves ...string) {
	m, err := xorm.NewEngine(driver, master)
	if err != nil {
		log.Panicf("Connect master: %s database error: %s", master, err)
	}

	sList := make([]*xorm.Engine, 0)
	for _, slave := range slaves {
		s, err := xorm.NewEngine(driver, slave)
		if err != nil {
			log.Panicf("Connect slave: %s database error: %s", slave, err)
		}
		sList = append(sList, s)
	}

	eg, err := xorm.NewEngineGroup(m, sList)
	if err != nil {
		log.Panicf("Create engine group failed: %s", err)
	}

	eg.ShowSQL(false)
	eg.SetMapper(names.GonicMapper{})

	once.Do(func() {
		MEngine = eg.Master()
		SEngine = eg.Slave()
	})
} `
