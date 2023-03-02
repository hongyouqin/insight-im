package utils

import (
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

// AsyncIgnored 静默异步执行
func AsyncIgnored(log *logrus.Entry, handler func() error) {
	go func() {
		defer func() {
			if v := recover(); v != nil {
				log.Error("execute async handler panic: %v", v)
			}
		}()
		if err := handler(); err != nil {
			log.Errorf("execute async handler error: %+v", err)
		}
	}()
}

func NullToDefault[T any](v *T, def T) T {
	if v == nil {
		return def
	} else {
		return *v
	}
}

func MapList[F any, T any](rows []F, mapper func(item F) T) []T {
	return lo.Map(rows, func(item F, _ int) T {
		return mapper(item)
	})
}

/*##########################以下是 msg-new 的代码#####################################*/

func GetSelfFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return cleanUpFuncName(runtime.FuncForPC(pc).Name())
}

func cleanUpFuncName(funcName string) string {
	end := strings.LastIndex(funcName, ".")
	if end == -1 {
		return ""
	}
	return funcName[end+1:]
}

func OperationIDGenerator() string {
	return strconv.FormatInt(time.Now().UnixNano()+int64(rand.Uint32()), 10)
}

func StringToInt(i string) int {
	j, _ := strconv.Atoi(i)
	return j
}

//Get the current timestamp by Second

func GetCurrentTimestampBySecond() int64 {
	return time.Now().Unix()
}

// Get the current timestamp by Mill
func GetCurrentTimestampByMill() int64 {
	return time.Now().UnixNano() / 1e6
}

// Convert nano timestamp to time.Time type
func UnixNanoSecondToTime(nanoSecond int64) time.Time {
	return time.Unix(0, nanoSecond)
}

// Get the current timestamp by Nano
func GetCurrentTimestampByNano() int64 {
	return time.Now().UnixNano()
}
