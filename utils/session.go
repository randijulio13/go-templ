package utils

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func SetSession(c echo.Context, key string, value interface{}) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values[key] = value
	sess.Save(c.Request(), c.Response())
}

func GetSession(c echo.Context, key string) interface{} {
	sess, _ := session.Get("session", c)
	return sess.Values[key]
}

func DestroySession(c echo.Context, key string) {
	sess, _ := session.Get("session", c)
	// sess.Options = &sessions.Options{
	// 	Path:     "/",
	// 	MaxAge:   86400 * 7,
	// 	HttpOnly: true,
	// }
	delete(sess.Values, key)
	// sess.Values[key] = nil
	sess.Save(c.Request(), c.Response())
}
