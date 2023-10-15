package main

import (
	"github.com/zii/go-log"
)

func main() {
	// set relative depth of project base dir
	log.SetRoot(1) // 1=..
	log.SetLevel(log.LvMax)

	log.Trace("ttttt")
	log.Tracef("ttttt: %d", 1)
	log.Debug("ddddd")
	log.Debugf("ddddd: %d", 2)
	log.Info("iiiii")
	log.Infof("iiiii: %d", 3)
	log.Warn("wwwww")
	log.Warnf("wwwww: %d", 4)
	log.Error("eeeee")
	log.Errorf("eeeee: %d", 1)
	log.Fatal("fffff")
	log.Fatalf("fffff: %d", 4)
}
