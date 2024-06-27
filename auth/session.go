package auth

import (
	"errors"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	Store sessions.Store
)

const (
	SessionKey  = "user"
	SessionName = "alamoa"
)

func SetStore(store sessions.Store) {
	Store = store
}

func SetSession(c *gin.Context, email string) error {
	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60,
		Secure:   false,
		HttpOnly: true,
	})
	session.Set(SessionKey, email)
	log.Printf("SetSession: %v", email)
	return session.Save()
}

func GetSession(c *gin.Context) (string, error) {
	session := sessions.Default(c)
	email, ok := session.Get(SessionKey).(string)
	if !ok {
		return "", errors.New("session not found")
	}
	log.Printf("GetSession: %v", email)
	return email, nil
}

func ClearSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	session.Delete(SessionKey)
	err := session.Save()
	log.Println("Cleared session")
	return err
}
