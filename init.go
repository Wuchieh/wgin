package wgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuchieh/wtype"
)

func init() {
	SetErrorHandler(func(c *gin.Context, err error) {
		c.String(http.StatusOK, err.Error())
	})

	SetResponseHandler(func(c *gin.Context, a any) {
		if a == nil {
			c.String(http.StatusOK, "")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "success",
				"data":    a,
			})
		}
	})

	SetBindHandler(func(c *gin.Context, a any) error {
		err := c.ShouldBind(a)
		if err != nil {
			return err
		}
		wtype.StructStringTrim(a)
		return nil
	})
}
