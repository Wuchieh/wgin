package wgin

import "github.com/gin-gonic/gin"

type BaseContext = *gin.Context

type Context[T any] struct {
	BaseContext
	body           *T
	getBodyHandler GetBodyHandler[T]
}

func (c *Context[T]) GetBody() (*T, error) {
	if c.body == nil {
		var err error

		if c.getBodyHandler != nil {
			handler, err := c.getBodyHandler(c)
			if err != nil {
				return nil, err
			}
			c.body = handler
		} else if bindHandler != nil {
			err = bindHandler(c.BaseContext, &c.body)
		} else {
			err = c.ShouldBind(&c.body)
		}

		if err != nil {
			return nil, err
		}
	}

	return c.body, nil
}
