package logrustool

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	// "github.com/onrik/logrus"
	// "github.com/keepeye/logrus-filename"
	// "github.com/rifflock/lfshook"
)

// LogConfig ...
type LogConfig struct {
	FilePath          string `json:"filePath"`
	FileName          string `json:"fileName"`
	Level             string `json:"level"`
	Formatter         string `json:"formatter"`
	LineLocateEnabled bool   `json:"lineLocateEnabled"`
}

func InitLog(cfg *LogConfig) {
	// func InitLog() {

	setFormatter(cfg.Formatter)
	setOutput(cfg)
	setLevel(cfg.Level)

	if cfg.LineLocateEnabled == true {
		// log.SetReportCaller(true)
		addLineFieldHook()
	}
}

func setOutput(cfg *LogConfig) {

	er := os.MkdirAll(cfg.FilePath, 0777)
	if er != nil {
		log.WithFields(log.Fields{
			"error": er,
		}).Fatal("os.MkdirAll fail!")
	}

	filePath := cfg.FilePath + cfg.FileName + ".log"

	rWriter, err := rotatelogs.New(
		filePath+".%Y%m%d",
		rotatelogs.WithLinkName(filePath),
		rotatelogs.WithRotationTime(time.Hour*24),
		// rotatelogs.WithMaxAge(time.Hour*24),
		// rotatelogs.WithRotationCount(18),
	)
	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}

	log.SetOutput(rWriter)
}

func setLevel(cfgLevel string) {
	level, err := log.ParseLevel(cfgLevel)
	if err != nil {
		level = log.ErrorLevel
	}
	log.SetLevel(level)
}

func setFormatter(cfgFormatter string) {
	log.SetFormatter(&log.TextFormatter{
		// DisableColors: true,
		// FullTimestamp: true,
	})
	if cfgFormatter == "JSON" {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func addLineFieldHook() {
	lineFieldHook := NewHook()
	lineFieldHook.Field = "line"
	log.AddHook(lineFieldHook)
}
