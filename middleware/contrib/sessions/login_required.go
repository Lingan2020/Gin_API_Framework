package sessions

import (
	"github.com/gin-gonic/gin"
	//"fmt"
	"net/http"
)




// Login Require Decorator
func LoginRequired(handle gin.HandlerFunc) gin.HandlerFunc {

	return func(c *gin.Context) {
		userToken, cookie_err := c.GetSecureCookie("user_token",1)

		var is_login  bool = true

		if cookie_err != nil{
			is_login = false
		}

		//Tudo 添加查数据库逻辑

		if is_login == false{
			c.JSON(http.StatusUnauthorized,
				gin.H{
					"status":  "failed",
					"desc": "login requierd",
				})
		}else {
			handle(c)
			c.Set("currentUserId",userToken)
			c.Set("currentUser", userToken)
		}
	}
}