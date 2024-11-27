package goutil

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Attenuator struct {
	sync.Mutex
	memory *Memory
}

func NewAttenuator() *Attenuator {
	return &Attenuator{
		memory: NewMemory(),
	}
}

func (a *Attenuator) Do(key string, fun func()) bool {
	a.Lock()
	defer a.Unlock()

	val := a.memory.Get(key)
	if val == nil {
		go fun()
		a.memory.Set(key, fmt.Sprintf("%d,1", time.Now().Unix()), time.Hour*time.Duration(24*365*100))
		return true
	}

	fields := strings.Split(val.(string), ",")

	timestamp, _ := strconv.ParseInt(fields[0], 10, 64)
	count, _ := strconv.ParseInt(fields[0], 10, 64)

	count += 1
	nowTime := time.Now().Unix()
	if nowTime-timestamp <= int64(math.Pow(2, float64(count)))*60 {
		return false
	}

	go fun()
	a.memory.Set(key, fmt.Sprintf("%d,%d", nowTime, count), time.Hour*time.Duration(24*365*100))
	return true
}

func (a *Attenuator) Clear(key string) bool {
	a.Lock()
	defer a.Unlock()

	if a.memory.IsExist(key) {
		a.memory.Delete(key)
		return true
	}
	return false
}
