package sessionService

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"scoresystem/app/models"
	"scoresystem/app/services/userService"
)

func ClearUserSession(c *gin.Context) {
	webSession := sessions.Default(c)
	webSession.Delete("id")
	webSession.Save()
	return
}

func CheckUserSession(c *gin.Context) bool {
	webSession := sessions.Default(c)
	id := webSession.Get("id")
	if id == nil {
		return false
	}
	return true
}

func GetUserSession(c *gin.Context) (*models.User, error) {
	webSession := sessions.Default(c)
	id := webSession.Get("id")
	if id == nil {
		return nil, errors.New("")
	}
	user, _ := userService.GetUserById(id.(int))
	if user == nil {
		ClearUserSession(c)
		return nil, errors.New("")
	}
	return user, nil
}

func SetStudentSession(c *gin.Context, user *models.User) error {
	webSession := sessions.Default(c)
	webSession.Options(sessions.Options{MaxAge: 3600 * 24 * 7})
	webSession.Set("id", user.ID)
	return webSession.Save()
}
