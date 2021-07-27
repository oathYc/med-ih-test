package boot

import (
	"ihtest/library/global"
	"github.com/BurntSushi/toml"
)

func InitConfig() error {
	_, err := toml.DecodeFile("config/config.toml", &global.Config)
	if err != nil {
		return err
	}

	return nil
}
