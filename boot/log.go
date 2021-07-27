package boot

import (
	"git.medlinker.com/golang/xerror"
	"git.medlinker.com/golang/xlog"
	"ihtest/library/global"
	"git.medlinker.com/golang/xlog/support/logrus"
)

func InitLogger() error {
	path := global.Config.Log.Dir

	if "" == path {
		return xerror.New("log.dir invalid")
	}

	global.Log = logrus.NewLoggerWith(path, global.Config.Log.Category, xlog.DailyStack)
	global.Log.SetStdPrint(global.Config.Log.StdPrint)
	if len(global.Config.Log.Level) > 0 {
		global.Log.SetLevelByString(global.Config.Log.Level)
	}

	global.Log.Debug("configs: %+v", global.Config)
	return nil
}