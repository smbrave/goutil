package goutil

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type GORMLogger struct {
}

func (d *GORMLogger) LogMode(level logger.LogLevel) logger.Interface {
	return d
}

func (d *GORMLogger) Info(context.Context, string, ...interface{}) {

}

func (d *GORMLogger) Warn(context.Context, string, ...interface{}) {

}

func (d *GORMLogger) Error(context.Context, string, ...interface{}) {

}

func (d *GORMLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, affects := fc()

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Errorf("[SQL]sql=%s affect=%d cost=%dms error=%v", sql, affects, time.Since(begin).Milliseconds(), err)
	} else {
		log.Debugf("[SQL]sql=%s affect=%d cost=%dms", sql, affects, time.Since(begin).Milliseconds())
	}
}
