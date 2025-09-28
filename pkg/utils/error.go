package utils

import (
	utils "json-pipeline/pkg"
	"os"
	"runtime"
)

var log = utils.GetLogger()

func ExitOnError(err error) {
	if err != nil {
		_, file, no, ok := runtime.Caller(1)
		if ok {
			log.Errorf("Fatal error in %s#%d: %s", file, no, err.Error())
		} else {
			log.Errorf("Fatal error: %s", err.Error())
		}

		os.Exit(1)
	}
}
