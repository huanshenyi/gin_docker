package log_source

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	log_conf := LoadLogConfig()

	// log output
	file, err := os.OpenFile(log_conf.LogDir, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	Log.Out = file
	// log leve
	level_mapping := map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}
	Log.SetLevel(level_mapping[log_conf.LogLevel])

	// log のフォーマット
	Log.SetFormatter(&logrus.JSONFormatter{})

}
