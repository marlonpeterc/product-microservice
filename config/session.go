package config

import (
	"os"
	"github.com/kataras/iris/sessions"
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
)

var manager *sessions.Sessions

func init() {
	var (
		hashKey  = []byte(os.Getenv("SESSION_HASH_SECRET"))
		blockKey = []byte(os.Getenv("SESSION_BLOCK_SECRET"))
	)
	var secureCookie = securecookie.New(hashKey, blockKey)
	manager = sessions.New(sessions.Config{
		Cookie: "session_id",
		Encode: secureCookie.Encode,
		Decode: secureCookie.Decode,
	})
}

func GetSession(ctx iris.Context) *sessions.Session {
	return manager.Start(ctx)
}

func AllSessions(ctx iris.Context) (string, string) {
	session := GetSession(ctx)
	id := session.GetString("id")
	username := session.GetString("username")
	return id, username
}
