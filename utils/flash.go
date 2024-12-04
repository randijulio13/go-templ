package utils

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Set adds a new message to the cookie store
func SetFlashmessages(c echo.Context, kind, value string) {
	// sess, _ := getCookieStore().Get(c.Request(), session_name)
	sess, _ := session.Get("session", c)

	sess.AddFlash(value, kind)

	sess.Save(c.Request(), c.Response())
}

// Get receives flash messages from cookie store
func GetFlashmessages(c echo.Context, kind string) []string {
	// sess, _ := getCookieStore().Get(c.Request(), session_name)
	sess, _ := session.Get("session", c)

	fm := sess.Flashes(kind)

	// if there are some messages…
	if len(fm) > 0 {
		sess.Save(c.Request(), c.Response())

		// we start an empty strings slice that we
		// then return with messages
		var flashes []string
		for _, fl := range fm {
			// we add the messages to the slice
			flashes = append(flashes, fl.(string))
		}

		return flashes
	}

	return nil
}
