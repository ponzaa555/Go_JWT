package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ponzaa555/Go_JWT/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Have to login in"})
			c.Abort()
			return
		}
		claim, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			// ใช้ในกรณี middle ware เช่น จะเช็ค token ก่อนถึงให้ใช้ ดึงข้อมูล user ได้ Flow : authFunc -> getUser ถ้าใช้ Abort() ในกรณี
			// authFunc ไม่ผ่านจะการันตีไม่ให้ไปเรียก getUser
			c.Abort()
			return
		}
		// set return
		c.Set("email", claim.Email)
		c.Set("first_name", claim.First_Name)
		c.Set("last_name", claim.Last_Name)
		c.Set("uid", claim.Uid)
		c.Set("user_type", claim.User_type)
		// Next() คือทำงาน function ต่อไป
		c.Next()
	}
}
