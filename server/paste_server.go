package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

const MaxListSize int = 5

func main() {
	textList := []string{}
	mutex := sync.Mutex{}

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("../web", true)))

	api := router.Group("/api")
	{
		api.GET("/gettext", func(c *gin.Context) {
			var text string = ""

			mutex.Lock()
			if len(textList) > 0 {
				text = textList[len(textList)-1]
			}
			fmt.Printf("len of text list: %v \n", len(textList))
			mutex.Unlock()

			c.String(http.StatusOK, text)
		})

		api.POST("/pushtext", func(c *gin.Context) {
			body := c.Request.Body
			x, _ := ioutil.ReadAll(body)
			fmt.Printf("%s \n", string(x))

			//save
			mutex.Lock()
			textList = append(textList, string(x))
			if len(textList) > MaxListSize {
				textList = textList[1:]
			}

			mutex.Unlock()

			c.String(http.StatusOK, "")
		})
	}

	router.Run(":8008")
}
