package config

import (
	"time"
	"log"

	"github.com/joho/godotenv"
	"github.com/kataras/iris"
)

func init() {
	godotenv.Load()
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func Err(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}

func MeOrNot(ctx iris.Context, user string) bool {
	id, _ := AllSessions(ctx)
	if id != user {
		return false
	}
	return true
}
