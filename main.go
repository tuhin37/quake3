package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	gameController "github.com/Binayaka/Q3AServer/controllers"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// func StartGame(c *gin.Context) {

// 	go payload(wg, "./RunServer.sh")
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Server started",
// 	})

// }

func main() {
	fmt.Println("server starts...")

	Init() // this function initialize the environment variables

	// wg := &sync.WaitGroup{}
	r := gin.Default()
	r.GET("/start", gameController.StartServer)
	r.GET("/status", gameController.GetStatus) // curl --location --request GET '43.204.235.64:5000/status'
	r.PUT("/update", gameController.UpdateGame)
	r.PUT("/restore", gameController.RestoreDefault)
	r.GET("/stop", gameController.StopServer)

	r.Run(":5000")
	fmt.Println("main ends")
}

func Init() {
	loadENV()
	// set network port
	gameController.RestoreDefaultConfigs()

	// network port
	networkPort := os.Getenv("PORT")
	if networkPort == "" {
		networkPort = "27960"
	}
	gameController.SearchAndReplaceTextFile(gameController.AUTOEXEC_CFG, "set net_port 27960", fmt.Sprintf("set net_port %s", networkPort))

	// ram
	ram := os.Getenv("RAM")
	if ram == "" {
		ram = "128"
	}
	gameController.SearchAndReplaceTextFile(gameController.AUTOEXEC_CFG, "set com_hunkmegs 128", fmt.Sprintf("set com_hunkmegs %s", ram))

	// password
	password := os.Getenv("PASSWORD")
	if password == "" {
		password = "adminadmin" // if env is not found then use this as defauul value
	}
	gameController.SearchAndReplaceTextFile(gameController.SERVER_CFG, "seta rconpassword \"adminadmin\"", fmt.Sprintf("seta rconpassword \"%s\"", password))

	// TOKEN
	token := os.Getenv("TOKEN")
	if token == "" {
		token = "70B9VW8igFT1lZSxVd22w9HOPz6DQu7Y" // if env is not found then use this as defauul value
		os.Setenv("TOKEN", token)
	}

}

func loadENV() {
	// reads .env files, has to be at the root directory along with main.go file
	file, err := os.Open("./.env")
	if err != nil {
		fmt.Println(".env file not found")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		lines = append(lines, scanner.Text())
	}

	// now load the variable
	for _, line := range lines {
		if string(line[0]) == "#" || string(line[0]) == "/" || string(line[0]) == " " { // exclude any line that start with '/ 'or '#', or ' ' as they are comments or line breaks
			continue
		}
		os.Setenv(strings.SplitN(line, "=", 2)[0], strings.SplitN(line, "=", 2)[1])
	}

}
