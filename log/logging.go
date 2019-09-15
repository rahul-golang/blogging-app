package log

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"runtime"

	log "github.com/Sirupsen/logrus"
)

var logger *log.Logger

const REQUESTID = "requestID"

func init() {
	logger = log.New()
	logger.SetLevel(log.TraceLevel)
	logger.Formatter = &log.TextFormatter{}
}

//Logger with fields
func Logger(ctx context.Context) *log.Entry {
	var depth = 1
	var requestid string
	if ctxRqID, ok := ctx.Value(REQUESTID).(string); ok {
		requestid = ctxRqID
	}
	function, file, line, _ := runtime.Caller(depth)
	functionObject := runtime.FuncForPC(function)
	entry := logger.WithFields(log.Fields{
		"requestid": requestid,
		"file":      file,
		"function":  functionObject.Name(),
		"line":      line,
	})
	var filename string = "logfile.log"
	//Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		logger.SetOutput(f)
	}
	return entry

}

// WithRqID returns a context which knows its request ID
func WithRqID(ctx context.Context) context.Context {
	return context.WithValue(ctx, REQUESTID, randNumberRunes(10))
}

var numberRunes = []rune("0123456789")

func randNumberRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numberRunes[rand.Intn(len(numberRunes))]
	}
	return string(b)
}
