package main

import (
	"fmt"
	"net/http"

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

	// wg := &sync.WaitGroup{}
	r := gin.Default()
	r.GET("/start", gameController.StartServer)
	r.GET("/status", gameController.GetStatus) // curl --location --request GET '43.204.235.64:5000/status'

	r.Run(":5000")
	fmt.Println("main ends")
}

// func getPID(wg *sync.WaitGroup, command string) (string){
// 	defer wg.Done()
// 	cmd := exec.Command(command)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatalf("cmd.Run() failed with %s\n", err)
// 	}

// 	return cmd.Stdout.String()
// }
