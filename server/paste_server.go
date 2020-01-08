package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

const MaxListSize int = 50

func main() {
	textList := []string{}
	mutex := sync.RWMutex{}

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("../web", true)))
//        index := 0

	api := router.Group("/api")
	{
		api.GET("/gettext", func(c *gin.Context) {
			var text string = ""

			mutex.RLock()
			if len(textList) > 0 {
				text = textList[len(textList)-1]
			}
			fmt.Printf("len of text list: %v \n", len(textList))
			mutex.RUnlock()

			c.String(http.StatusOK, text)
		})
		
		api.GET("/back", func(c *gin.Context){
			c.String(http.StatusOK,"")
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
