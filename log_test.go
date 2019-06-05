package log

import (
	"fmt"
	"testing"
)

func Test_Log(t *testing.T) {
	logger.SetLevelByName("TRACE")
	printall("TRACE")

	logger.SetLevelByName("DEBUG")
	printall("DEBUG")

	logger.SetLevelByName("INFO")
	printall("INFO")

	logger.SetLevelByName("WARN")
	printall("WARN")

	logger.SetLevelByName("ERROR")
	printall("ERROR")

	logger.SetLevelByName("FATAL")
	printall("FATAL")
}

func printall(level string) {
	str := fmt.Sprintf("this is level %s", level)
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
