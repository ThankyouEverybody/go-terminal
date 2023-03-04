package setup

import (
	"fmt"
	"github.com/go-terminal-server/common/config"
	"github.com/go-terminal-server/service"
)

var s *setup = new(setup)

type setup struct {
}

func (s *setup) loadConfig() error {
	loadConfig := config.GlobalConfig
	fmt.Printf("start\n%v\n", loadConfig)
	return nil
}

func (s *setup) initDB() error {
	err := service.UserService.InitUser()
	if err != nil {
		return err
	}
	return nil
}

func Start() error {

	var err error

	//加载配置
	err = s.loadConfig()
	if err != nil {
		goto errLabel
	}
	err = s.initDB()
	if err != nil {
		goto errLabel
	}

	return nil
errLabel:
	return err
}
