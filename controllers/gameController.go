package gameController

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

// starts the server
func StartServer(c *gin.Context) {
	// async call
	go executeSH("./shellScripts/RunServer.sh")
	c.AsciiJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "server started", // cast it to string before showing
	})
}

// returns game status. Returns if server is running and all the config settings
func GetStatus(c *gin.Context) {
	// read autoexec.cfg
	lines := readFileLines("./config/autoexec.cfg") // lines is a slice of string, each element is one line the text file
	autoexecMap := line2map(lines)                  // converts lines into key value pair

	// raed server.cfg
	lines = readFileLines("./config/server.cfg") // lines is a slice of string, each element is one line the text file
	serverMap := line2map(lines)                 // converts lines into key value pair
	// convert gametype from code to string
	switch serverMap["g_gametype"] {
	case "4":
		serverMap["g_gametype"] = "CTF"
	case "3":
		serverMap["g_gametype"] = "TD"
	case "1":
		serverMap["g_gametype"] = "Tourney"
	case "0":
		serverMap["g_gametype"] = "FFA"
	}

	// raed bots.cfg
	lines = readFileLines("./config/bots.cfg") // lines is a slice of string, each element is one line the text file
	botsMap := line2map(lines)                 // converts lines into key value pair

	// raed levels.cfg
	lines = readFileLines("./config/levels.cfg") // lines is a slice of string, each element is one line the text file1
	mapName := strings.TrimRight(strings.SplitN(lines[0], " ", int(5))[3], ";")

	c.JSON(http.StatusOK, gin.H{"autoexec": autoexecMap, "server": serverMap, "bots": botsMap, "map": mapName})
}

// updates the config files and restart / start the game
func UpdateGame(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// Handle error
	// }
	fmt.Println(string(jsonData))
}

// --------------------------- internal functions ---------------------------------
func executeSH(script string) {
	cmd := exec.Command(script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("ERROR: script %s failed with err: %s", script, err)
	}
}

// reads a text file line by line and returns a slice of string
func readFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("[ERROR] | file open failed | %s | error | %s", filename, err)
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
	return lines
}

func line2map(lines []string) map[string]interface{} {
	/*
		input: ["set vm_game 2", "set vm_cgame 2", "set vm_ui 2", ...]
		output: { "vm_game": 2, "vm_cgame": 2, "et vm_ui": 2 ...}
	*/
	obj := map[string]interface{}{}
	var NUM_OF_LINE_SEGMENTS uint8 = 4
	for _, line := range lines {
		if string(line[0]) == "#" || string(line[0]) == "/" || string(line[0]) == " " { // exclude any line that start with '/ 'or '#', or ' ' as they are comments or line breaks
			continue
		}
		obj[strings.SplitN(line, " ", int(NUM_OF_LINE_SEGMENTS))[1]] = strings.Trim(strings.SplitN(line, " ", int(NUM_OF_LINE_SEGMENTS))[2], "\"") // key = 2nd word; value = 3rd word

	}
	return obj
}
