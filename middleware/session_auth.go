package middleware

import (
	"errors"
	"github.com/James2333/go_gateway/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if adminInfo,ok:=session.Get(public.AdminSessionInfoKey).(string);!ok||adminInfo==""{
			ResponseError(c, InternalErrorCode, errors.New("admin not login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
