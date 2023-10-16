# go-log

log for golang

## Usage

```go
package main

import (
	"github.com/zii/go-log"
	"time"
)

func main() {
	// how many up levels to project base dir, 1 means ..
	log.SetRoot(1)
	log.SetLevel(log.LvMax)            // default level is INFO
	log.SetTimeFormat(time.StampMilli) // leave empty to hide time

	log.Trace("ttttt")
	log.Tracef("ttttt: %d", 1)
	log.Debug("ddddd")
	log.Debugf("ddddd: %d", 2)
	log.Info("iiiii")
	log.Infof("iiiii: %d", 3)
	log.Println("ppppp")
	log.Printf("ppppp: %d", 3)
	log.Warn("wwwww")
	log.Warnf("wwwww: %d", 4)
	log.Error("eeeee")
	log.Errorf("eeeee: %d", 1)
	log.Fatal("fffff")
	log.Fatalf("fffff: %d", 4)
}
```

## Installation

```
$ go get github.com/zii/go-log
```

## License

MIT

## Thanks

* https://github.com/fatih/color
* https://github.com/mattn/go-isatty
