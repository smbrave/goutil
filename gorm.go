package goutil

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"reflect"
	"time"
)

type GORMLogger struct {
	Threshold int64
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
		if d.Threshold > 0 && time.Since(begin).Milliseconds() > d.Threshold {
			log.Errorf("[SQL]sql=%s affect=%d cost=%dms", sql, affects, time.Since(begin).Milliseconds())
		} else {
			log.Debugf("[SQL]sql=%s affect=%d cost=%dms", sql, affects, time.Since(begin).Milliseconds())
		}
	}
}

func ModelInitColumn(md interface{}, col interface{}) {
	modelValue := reflect.TypeOf(md).Elem()
	modelCol := reflect.ValueOf(col).Elem()
	for i := 0; i < modelValue.NumField(); i++ {
		field := modelValue.Field(i)
		fValue := modelCol.FieldByName(field.Name)
		if fValue.IsValid() && fValue.CanSet() {
			fValue.Set(reflect.ValueOf(SnakeString(field.Name)))
		}
	}
}

func ModelFieldValue(md interface{}, fields ...string) map[string]interface{} {
	data := make(map[string]interface{})
	obj := reflect.ValueOf(md).Elem()

	for _, f := range fields {
		val := obj.FieldByName(CamelString(f))
		if val.IsValid() {
			data[f] = val.Interface()
		}
	}
	return data
}
