package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	if userType != role {
		err = errors.New("Unauthorized access")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uId := c.GetString("user_id")
	err = nil

	if userType == "USER" && uId != userId {
		err = errors.New("Unauthorized access to this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err

}
