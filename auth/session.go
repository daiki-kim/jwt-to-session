package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var (
	Store = sessions.NewCookieStore([]byte("secret"))
)

const (
	SessionKey  = "user"
	SessionName = "alamoa"
)

func SetSession(c *gin.Context, email string) error {
	session, _ := Store.Get(c.Request, SessionName)
	session.Values[SessionKey] = email
	return session.Save(c.Request, c.Writer)
}

func GetSession(c *gin.Context) (string, error) {
	session, _ := Store.Get(c.Request, SessionName)
	user, ok := session.Values[SessionKey].(string)
	if !ok {
		return "", errors.New("session not found")
	}
	return user, nil
}

func ClearSession(c *gin.Context) error {
	session, _ := Store.Get(c.Request, SessionName)
	delete(session.Values, SessionKey)
	return session.Save(c.Request, c.Writer)
}
