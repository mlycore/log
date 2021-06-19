package log

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func Test_Log(t *testing.T) {
	NewDefaultLogger()
	SetFormatter(&TextFormatter{})
	logger.SetLevelByName("TRACE")
	printall(logger.Level)

	logger.SetLevelByName("DEBUG")
	printall(logger.Level)

	logger.SetLevelByName("INFO")
	printall(logger.Level)

	logger.SetLevelByName("WARN")
	printall(logger.Level)

	logger.SetLevelByName("ERROR")
	printall(logger.Level)

	logger.SetLevelByName("FATAL")
	printall(logger.Level)
}

func printall(level int) {
	l := LogLevelMap[level]
	str := fmt.Sprintf("this is level %s", l)
	Traceln("traceln: " + str)
	Tracef("tracef: %s", str)
	Debugln("debugln: " + str)
	Debugf("debugf: %s", str)
	Infoln("infoln: " + str)
	Infof("infof: %s", str)
	Warnln("warnln: " + str)
	Warnf("warnf: %s", str)
	Errorln("errorln: " + str)
	Errorf("errorf: %s", str)
	//Fatalln("fatalln: " + str)
	//Fatalf("fatalf: %s", str)
}

func Test_FuncInfo(t *testing.T) {
	Traceln("this is traceln")
	data := make([]byte, 10240)
	runtime.Stack(data, true)
	fmt.Printf("%s\n", string(data))
}

func Test_GetShortFileName(t *testing.T) {
	name := "github.com/xuyun-io/kubestar/pkg/controllers/crd/application.(*ApplicationController).List"
	NewLogger(os.Stdout, 0, 3, true)
	SetLevel("ERROR")
	Errorf("%s", getShortFileName(name))
	// t.Logf("%s\n", getShortFileName(name))
}

func Test_FormatterLogger(t *testing.T) {
	NewDefaultLogger()
	SetFormatter(&JSONFormatter{})
	// SetFormatter(&TextFormatter{})
	/*
		SetContext(Context{
			"namespace": "default",
			"deployment": "kubestar",
		})
	*/

	Infoln("deployment create success")
	Errorf("service create error")

	SetContext(Context{
		"namespace":  "starpay",
		"deployment": "uuid",
	})
	Infoln("deployment list success")
}

func Test_DefaultLogger(t *testing.T) {
	NewDefaultLogger()
	SetFormatter(&TextFormatter{})
	logger.SetLevelByName("WaRN")
	printall(logger.Level)
}
