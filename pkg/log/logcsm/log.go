package logcsm

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

// Log ...
type Log struct {
	logChannel chan string
	running    bool

	path            string
	nameWithoutType string

	bDebug bool
	bInfo  bool
	bError bool
}

var myLog *Log

// GetLog ...
func GetLog() *Log {
	if nil == myLog {
		myLog = new(Log)
	}
	return myLog
}

// OpenLogFile ...
func (thi *Log) OpenLogFile(year int, month time.Month, day int) (file *os.File, err error) {
	//filePre := fmt.Sprintf("%04d%02d%02d%02d%02d%02d", year, month, day, hour, min, sec)
	filePre := fmt.Sprintf("%04d%02d%02d", year, month, day)
	filePath := thi.path + thi.nameWithoutType + "_" + filePre + ".log"

	err = os.MkdirAll(thi.path, 0777)
	if err != nil {
		return nil, err
	}

	// f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModeAppend)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	return f, err
}

// Start ...
func (thi *Log) Start() error {

	thi.logChannel = make(chan string, 10000)
	thi.running = true

	thi.path = "./log/"
	thi.nameWithoutType = "something"

	thi.bDebug = true
	thi.bInfo = true
	thi.bError = true

	errChannel := make(chan error, 1)

	go func() {

		now := time.Now()
		bgnYear, bgnMon, bgnDay := now.Date()
		//hour, min, sec := now.Clock()

		logfile, err := thi.OpenLogFile(bgnYear, bgnMon, bgnDay)
		if nil != err {
			errChannel <- err
			return
		}

		defer logfile.Close()

		logger := log.New(logfile, "", log.Ldate|log.Ltime|log.Lmicroseconds)
		errChannel <- nil

		for {
			select {
			case text := <-thi.logChannel:
				{
					curYear, curMon, curDay := time.Now().Date()
					if curDay != bgnDay || curMon != bgnMon || curYear != bgnYear {

						logfile.Close()
						logfile = nil

						//hour, min, sec = time.Now().Clock()

						logfile, err = thi.OpenLogFile(curYear, curMon, curDay)
						if err == nil {
							logger = log.New(logfile, "", log.Ldate|log.Ltime|log.Lmicroseconds)
						}

						bgnYear = curYear
						bgnMon = curMon
						bgnDay = curDay
					}

					if logfile != nil {
						logger.Println(text)
					}

				}
			default:
				{
					time.Sleep(300 * time.Millisecond)
					if !thi.running {
						break
					}
				}
			}
		}
	}()

	err := <-errChannel
	if err != nil {
		close(errChannel)
		close(thi.logChannel)
		thi.running = false
	}
	return err
}

// Stop ...
func (thi *Log) Stop() {
	thi.running = false
}

func (thi *Log) Write(text string) {
	select {
	case thi.logChannel <- text:
		{
		}
	case <-time.After(1 * time.Second):
		{
			//fmt.Println("log write, input log channel time out")
		}
	}

}

// LogDebug ...
func LogDebug(text string) {
	if GetLog().bDebug == true {
		_, file, line, ok := runtime.Caller(1)
		// arrStr := strings.Split(file, "/")
		if ok {
			GetLog().Write(filepath.Base(file) + " " + strconv.FormatInt(int64(line), 10) + ": " + "[Debug] " + text)
		} else {
			GetLog().Write("[Debug] " + text)
		}
	}
}

// LogInfo ...
func LogInfo(text string) {
	if GetLog().bInfo == true {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			GetLog().Write(filepath.Base(file) + " " + strconv.FormatInt(int64(line), 10) + ": " + "[Info] " + text)
		} else {
			GetLog().Write("[Info] " + text)
		}
	}
}

// LogErr ...
func LogErr(text string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		GetLog().Write(filepath.Base(file) + " " + strconv.FormatInt(int64(line), 10) + ": " + "[Error] " + text)
	} else {
		GetLog().Write("[Error] " + text)
	}
}
