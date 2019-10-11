package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"fmt"

	"github.com/urfave/cli"
)

const ServerAddr string = "http://localhost:8008"

func FetchText() string {
	fetchurl := ServerAddr + "/api/gettext"
	resp, err := http.Get(fetchurl)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}

func PushText(text string) {
	pushurl := ServerAddr + "/api/pushtext"
	_, err := http.Post(pushurl, "", strings.NewReader(text))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	myApp := cli.NewApp()
	myApp.Name = "ncp client"
	//myApp.Usage = "server(with SMUX)"
	//myApp.Version = VERSION
	myApp.Commands = []cli.Command{
		{
			Name:    "fetch",
			Aliases: []string{"f", "g"},
			Usage:   "fetch text",
			Action: func(c *cli.Context) error {
				text := FetchText()
				fmt.Println(text)
				return nil

			},
		},
		{
			Name:    "push",
			Aliases: []string{"p", "c"},
			Usage:   "copy text",
			Action: func(c *cli.Context) error {
				text := c.Args().First()
				fmt.Println("text: ", text)
				PushText(text)
				return nil

			},
		},
	}

	err := myApp.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
