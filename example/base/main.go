package main

import (
	"fmt"

	"github.com/duke-git/lancet/v2/pointer"
	"github.com/gin-gonic/gin"
	"github.com/wuchieh/wgin"
	_ "github.com/wuchieh/wgin/autosetup"
)

var (
	helloWorldHandler = wgin.NewBaseHandler[string](helloWorld).
				AddMiddleware(helloWorldMiddleware1).
				AddMiddleware(helloWorldMiddleware2)

	helloUserHandler = wgin.NewHandler[helloUserRequest, string](helloUser)
)

func main() {
	r := gin.Default()

	wgin.UseGet(r, "/hello", helloWorldHandler)
	wgin.UsePOST(r, "/hello", helloUserHandler)

	r.Run(":8080")
}

func helloWorldMiddleware1(c *gin.Context) {
	fmt.Println(1)
	c.Next()
	fmt.Println(5)
}

func helloWorldMiddleware2(c *gin.Context) {
	fmt.Println(2)
	c.Next()
	fmt.Println(4)
}

func helloWorld(c *gin.Context) (*string, error) {
	fmt.Println(3)
	return pointer.Of("hello world"), nil
}

type helloUserRequest struct {
	Name string `json:"name"`
}

func helloUser(c *wgin.Context[helloUserRequest]) (*string, error) {
	body, err := c.GetBody()
	if err != nil {
		return nil, err
	}
	return pointer.Of(fmt.Sprintf("hello %s", body.Name)), nil
}
