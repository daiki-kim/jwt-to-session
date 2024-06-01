package auth

import (
	"errors"
	"log"

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
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60,
		Secure:   false,
		HttpOnly: true,
	}
	log.Printf("SetSession: %v", session.Values)
	session.Values[SessionKey] = email
	return session.Save(c.Request, c.Writer)
}

func GetSession(c *gin.Context) (string, error) {
	session, _ := Store.Get(c.Request, SessionName)
	log.Printf("GetSession: %v", session.Values)
	user, ok := session.Values[SessionKey].(string)
	if !ok {
		return "", errors.New("session not found")
	}
	return user, nil
}

func ClearSession(c *gin.Context) error {
	session, _ := Store.Get(c.Request, SessionName)
	log.Printf("ClearSession before: %v", session.Values)
	session.Options.MaxAge = -1
	delete(session.Values, SessionKey)
	err := session.Save(c.Request, c.Writer)
	log.Printf("ClearSession after: %v", session.Values)
	return err
}
