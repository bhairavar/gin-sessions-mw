package gin_sessions_mw

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("reached session middleware.....")
		session := sessions.Default(c)
		if session == nil {
			fmt.Println("No Session Found.")
			session.Set("is_authenticated", "false")
			session.Clear()
			session.Save()
			c.AbortWithStatus(401)
			return
		} else {
			userID :=  session.Get("user_id")
			fmt.Println(userID)
			if userID == nil {
				session.Clear()
				session.Set("is_authenticated", "false")
				session.Save()
				c.AbortWithStatus(401)
				return
			} else {
				fmt.Println("Authenticated.")
				c.Next()
				return
			}
		}

	}
}


