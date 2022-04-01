package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func checkUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorised to access this resource")
		return err
	}
	return err
}

func MatchUserTypeUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorised to access this resource")

		return err
	}
	err = checkUserType(c, userType)
	return err
}
