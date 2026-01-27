package gotool

import (
	"fmt"
	"log"
)

var logLevel = 3

/**
0 : fatal
1 : error
2 : ok
3 : debug
*/

func SetLogLevel(level int) {
	logLevel = level
}

func LogFatal(msg ...any) {
	if logLevel >= 0 {
		log.Fatal("\033[31m", fmt.Sprint(msg...), "\033[0m")
	}
}

func LogError(msg ...any) {
	if logLevel >= 1 {
		log.Println("\033[35m", fmt.Sprint(msg...), "\033[0m")
	}
}

func LogOk(msg ...any) {
	if logLevel >= 2 {
		log.Println("\033[32m", fmt.Sprint(msg...), "\033[0m")
	}
}

func LogDebug(msg ...any) {
	if logLevel >= 3 {
		log.Println("\033[30m", fmt.Sprint(msg...), "\033[0m")
	}
}
