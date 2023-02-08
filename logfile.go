package goutil

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

type LogFile struct {
}

func (f *LogFile) Format(entry *log.Entry) ([]byte, error) {
	funcs := strings.SplitN(entry.Caller.Function, ".", 2)
	msg := fmt.Sprintf("[%s] [%s] [%s:%d] %s\n",
		entry.Time.Format("2006-01-02 15:04:05,000"),
		strings.ToUpper(entry.Level.String()),
		funcs[0]+"/"+filepath.Base(entry.Caller.File),
		entry.Caller.Line,
		entry.Message)
	return []byte(msg), nil
}
