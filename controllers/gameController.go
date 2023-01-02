package gameController

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

type Game struct {
	Autoexec Autoexec `json:"autoexec"`
	Bots     Bots     `json:"bots"`
	Map      string   `json:"map" binding:"required,alphanum,startswith=q3,max=10"`
	Server   Server   `json:"server"`
	Restart  string   `json:"restart" binding:"required,oneof=true false"`
}

type Autoexec struct {
	Vm_ui        uint8  `json:"vm_ui" binding:"gte=1,lte=5"`
	Vm_game      uint8  `json:"vm_game" binding:"gte=1,lte=5"`
	Vm_cgame     uint8  `json:"vm_cgame" binding:"gte=1,lte=5"`
	Net_port     uint16 `json:"net_port" binding:"gte=1,lte=65534"`
	Dedicated    uint8  `json:"dedicated" binding:"gte=0,lte=1"`
	Com_hunkmegs uint16 `json:"com_hunkmegs" binding:"gte=1,lte=4096"`
}

type Bots struct {
	Bot_enable     uint8 `json:"bot_enable" binding:"gte=0,lte=1"`
	Bot_minplayers uint8 `json:"bot_minplayers" binding:"gte=0,lte=16"`
	Bot_nochat     uint8 `json:"bot_nochat" binding:"gte=0,lte=1"`
	G_spskill      uint8 `json:"g_spskill" binding:"gte=1,lte=5"`
}

type Server struct {
	Capturelimit       uint8  `json:"capturelimit" binding:"gte=0,lte=100"`
	Cl_maxpackets      uint16 `json:"cl_maxpackets" binding:"gte=1,lte=65534"`
	Cl_packetdup       uint8  `json:"cl_packetdup" binding:"gte=0,lte=1"`
	Fraglimit          uint8  `json:"fraglimit" binding:"gte=0,lte=254"`
	G_forcerespawn     uint8  `json:"g_forcerespawn" binding:"gte=0,lte=1"`
	G_friendlyFire     uint8  `json:"g_friendlyFire" binding:"gte=0,lte=1"`
	G_gametype         string `json:"g_gametype" binding:"oneof=FFA TD CTF Tourney"`
	G_inactivity       uint8  `json:"g_inactivity" binding:"gte=5,lte=254"`
	G_log              string `json:"g_log" binding:"endswith=.log"`
	G_motd             string `json:"g_motd" binding:"max=255"`
	G_quadfactor       uint8  `json:"g_quadfactor" binding:"gte=1,lte=10"`
	G_teamAutoJoin     uint8  `json:"g_teamAutoJoin" binding:"gte=0,lte=1"`
	G_teamForceBalance uint8  `json:"g_teamForceBalance" binding:"gte=0,lte=1"`
	G_weaponrespawn    uint8  `json:"g_weaponrespawn" binding:"gte=1,lte=10"`
	Logfile            uint8  `json:"logfile" binding:"gte=1,lte=10"`
	Rate               uint32 `json:"rate" binding:"gte=1,lte=255999"`
	Rconpassword       string `json:"rconpassword" binding:"max=64"`
	Snaps              uint8  `json:"snaps" binding:"gte=1,lte=254"`
	Sv_hostname        string `json:"sv_hostname" binding:"max=64"`
	Sv_maxclients      uint8  `json:"sv_maxclients" binding:"gte=1,lte=16"`
	Sv_pure            uint8  `json:"sv_pure" binding:"gte=0,lte=1"`
	Timelimit          uint8  `json:"timelimit" binding:"gte=1,lte=180"`
}

// starts the server
func StartServer(c *gin.Context) {
	// async call
	go executeSH("./shellScripts/RunServer.sh")
	c.AsciiJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "server started", // cast it to string before showing
	})
}

// fetch gamestatus
func GetStatus(c *gin.Context) {
	// Reads all the config files and parses the parameters in json and returns to user

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

	// check barer token, return if not authorized
	// bearerToken := c.Request.Header.Get("Authorization")
	// bearerToken = strings.TrimPrefix(bearerToken, "Bearer ")
	// if bearerToken != "70B9VW8igFT1lZSxVd22w9HOPz6DQu7Y" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid token"})
	// 	return
	// }
	isValid := validateToken(c)
	if !isValid {
		return
	}
	// create game object
	game := Game{}
	// initialize game object with default values, to differentiate prestine state
	initializeGameObject(&game) // this initializes all the variables in the obhect with a non zero default value (uint8 <  0xFF, uint16 <- 0xFFFF, string <- ""). As zero can be a valid inputy

	// input validation and load post body in game object
	err := c.BindJSON(&game)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// update bots1.cfg
	updateConfigFiles(&game)
	// if game.Bots  {
	// 	fmt.Println("here")
	// }
	// searchAndReplaceTextFile("./config/bots1.cfg", "g_spskill 4", "g_spskill 5")

	c.JSON(http.StatusAccepted, &game)

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

func updateConfigFiles(game *Game) {
	// game is the structure that has all the data received from API

	// update bots.cfg file----------------------------------------------------------------------------------------------
	// read bots.cfg file to know the current values. The current values are required for search and replace operation
	lines := readFileLines("./config/bots1.cfg") // lines is a slice of string, each element is one line the text file
	botsMap := line2map(lines)                   // keys and values are string variable

	// convert game.bot object into gameBotMap
	var gameBotMap map[string]interface{}
	data, _ := json.Marshal(game.Bots)
	json.Unmarshal(data, &gameBotMap)

	// iterate through all settings in the bots.cfg file and look up if new value of that field is received in POST body
	for fieldName, currentValue := range botsMap {
		newValue := fmt.Sprint(gameBotMap[fieldName]) // convert float to string for comparision in line 210

		//  [  received value  ] 		   [current value]  	[ received value ]
		if gameBotMap[fieldName] != 0xFF && currentValue != newValue { // if the field value is received in POST body and if old value (in .cfg file) is different than the new value received

			// 											   [                   bot_enable   1          ]  [                   bot_enable   0                   ]
			searchAndReplaceTextFile("./config/bots1.cfg", fmt.Sprintf("%s %s", fieldName, currentValue), fmt.Sprintf("%s %v", fieldName, gameBotMap[fieldName])) // This actually updates the file
		}
	}

	// update autoexec.efg file ---------------------------------------------------------------------------------------------
	// read autoexec.cfg file to know the current values. The current values are required for search and replace operation
	lines = readFileLines("./config/autoexec1.cfg") // lines is a slice of string, each element is one line the text file
	autoexecMap := line2map(lines)                  // keys and values are string variable

	// convert game.autoexec object into a map
	var gameAutoexectMap map[string]interface{}
	data, _ = json.Marshal(game.Autoexec)
	json.Unmarshal(data, &gameAutoexectMap)

	// iterate through all settings in the bots.cfg file and look up if new value of that field is received in POST body
	for fieldName, currentValue := range autoexecMap {
		newValue := fmt.Sprint(gameAutoexectMap[fieldName]) // convert float to string for comparision in line 237

		//  [  received value  ] 		   [current value]  	[ received value ]
		if gameAutoexectMap[fieldName] != 0xFF && currentValue != newValue { // if the field value is received in POST body and if old value (in .cfg file) is different than the new value received

			// 											   [                   bot_enable   1          ]  [                   bot_enable   0                   ]
			searchAndReplaceTextFile("./config/autoexec1.cfg", fmt.Sprintf("%s %s", fieldName, currentValue), fmt.Sprintf("%s %v", fieldName, gameAutoexectMap[fieldName]))
		}
	}

	// update server.cfg file ---------------------------------------------------------------------------------------------
	// read server.cfg file to know the current values. The current values are required for search and replace operation
	lines = readFileLines("./config/server1.cfg") // lines is a slice of string, each element is one line the text file
	serverMap := line2map(lines)                  // keys and values are string variable

	// convert game.server object into a map
	var gameServerMap map[string]interface{}
	data, _ = json.Marshal(game.Server)
	json.Unmarshal(data, &gameServerMap)

	// iterate through all settings in the server.cfg file and look up if new value of that field is received in POST body
	for fieldName, currentValue := range serverMap {
		newValue := fmt.Sprint(gameServerMap[fieldName]) // convert float to string for comparision in line 237
		//  [  received value  ] 		   [current value]  	[ received value ]

		if gameServerMap[fieldName] != 0xFF && currentValue != newValue { // if the field value is received in POST body and if old value (in .cfg file) is different than the new value received

			searchStr := fmt.Sprintf("%s %s", fieldName, currentValue)
			replaceStr := fmt.Sprintf("%s %v", fieldName, gameServerMap[fieldName])

			// special case where i will have to add quote '"' around the values
			if fieldName == "sv_hostname" || fieldName == "g_motd" || fieldName == "rconpassword" || fieldName == "rate" || fieldName == "snaps" || fieldName == "cl_maxpackets" || fieldName == "cl_packetdup" { // special case where the values are quoted ("") in the cfg file
				searchStr = fmt.Sprintf("%s \"%s\"", fieldName, currentValue)              // add the extra quote around the value
				replaceStr = fmt.Sprintf("%s \"%v\"", fieldName, gameServerMap[fieldName]) // add the extra quote around the value
			}

			if fieldName == "g_gametype" {
				fmt.Println("---------------------------")
				fmt.Println(fieldName)
				fmt.Println(gameServerMap[fieldName])
				fmt.Println(currentValue)
				fmt.Println(newValue)

				// comment the current active block
				switch currentValue { // Currently active is CTF
				case "4":
					if newValue == "CTF" {
						break
					}

					for lineIndex := 10; lineIndex <= 15; lineIndex++ {
						commentLine("./config/server1.cfg", uint16(lineIndex))
					}
				case "0":
					if newValue == "FFA" {
						break
					}
					for lineIndex := 25; lineIndex <= 27; lineIndex++ {
						commentLine("./config/server1.cfg", uint16(lineIndex))
					}
				case "3":
					if newValue == "TD" {
						break
					}
					for lineIndex := 18; lineIndex <= 22; lineIndex++ {
						commentLine("./config/server1.cfg", uint16(lineIndex))
					}
				}

				// uncomment the new block based on gametype
				switch gameServerMap[fieldName] {
				case "FFA":
					if currentValue == "0" {
						break
					}
					for i := 25; i <= 27; i++ {
						uncommentLine("./config/server1.cfg", uint16(i))
					}
				case "CTF":
					if currentValue == "4" {
						break
					}
					for i := 10; i <= 15; i++ {
						uncommentLine("./config/server1.cfg", uint16(i))
					}
				case "TD":
					if currentValue == "3" {
						break
					}
					for i := 18; i <= 22; i++ {
						uncommentLine("./config/server1.cfg", uint16(i))
					}
				}
				continue
			}
			// 		                                      bot_enable 1  bot_enable   0
			searchAndReplaceTextFile("./config/server1.cfg", searchStr, replaceStr)
		}
	}

	// update levels1.cfg file ---------------------------------------------------------------------------------------------
	// raed levels.cfg
	lines = readFileLines("./config/levels1.cfg") // lines is a slice of string, each element is one line the text file1
	currentMapName := strings.TrimRight(strings.SplitN(lines[0], " ", int(5))[3], ";")
	searchAndReplaceTextFile("./config/levels1.cfg", fmt.Sprintf("set dm1 \"map %s", currentMapName), fmt.Sprintf("set dm1 \"map %s", game.Map))
}

func searchAndReplaceTextFile(filename string, search string, replace string) bool {
	// this function performs a search-and-replace operation on a text file.

	// read text file
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return false
	}

	// replace text
	output := bytes.Replace(input, []byte(search), []byte(replace), -1)

	// write file
	err = ioutil.WriteFile(filename, output, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return false
	}

	return true
}

func initializeGameObject(game *Game) {
	// This function initializes an object with a default value (uint8 <- 255, uint16 <- 65535 string = "" ...)
	game.Map = ""
	game.Restart = ""
	// autoexec
	game.Autoexec.Vm_ui = 0xFF
	game.Autoexec.Vm_game = 0xFF
	game.Autoexec.Vm_cgame = 0xFF
	game.Autoexec.Net_port = 0xFFFF
	game.Autoexec.Dedicated = 0xFF
	game.Autoexec.Com_hunkmegs = 0xFFF
	// Bots
	game.Bots.Bot_enable = 0xFF
	game.Bots.Bot_minplayers = 0xFF
	game.Bots.Bot_nochat = 0xFF
	game.Bots.G_spskill = 0xFF
	// server
	game.Server.Capturelimit = 0xFF
	game.Server.Cl_maxpackets = 0xFFFF
	game.Server.Cl_packetdup = 0xFF
	game.Server.Fraglimit = 0xFF
	game.Server.G_forcerespawn = 0xFF
	game.Server.G_friendlyFire = 0xFF
	game.Server.G_gametype = ""
	game.Server.G_inactivity = 0xFF
	game.Server.G_log = ""
	game.Server.G_motd = ""
	game.Server.G_quadfactor = 0xFF
	game.Server.G_teamAutoJoin = 0xFF
	game.Server.G_teamForceBalance = 0xFF
	game.Server.G_weaponrespawn = 0xFF
	game.Server.Logfile = 0xFF
	game.Server.Rate = 0xFFFFFFFF
	game.Server.Rconpassword = ""
	game.Server.Snaps = 0xFF
	game.Server.Sv_hostname = ""
	game.Server.Sv_maxclients = 0xFF
	game.Server.Sv_pure = 0xFF
	game.Server.Timelimit = 0xFF
}

func validateToken(c *gin.Context) bool {
	// This function validates the baerer token
	bearerToken := c.Request.Header.Get("Authorization")
	bearerToken = strings.TrimPrefix(bearerToken, "Bearer ")
	if bearerToken != "70B9VW8igFT1lZSxVd22w9HOPz6DQu7Y" { // TODO read this value from environment variale
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid token"})
		return false
	}
	return true
}

func commentLine(fileName string, lineNumber uint16) {
	// create file handler
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	// read lines
	lines := strings.Split(string(input), "\n")

	// edit
	selectedLine := lines[lineNumber-1]
	runes := []rune(selectedLine)
	if runes[0] != 47 && runes[1] != 47 {
		lines[lineNumber-1] = fmt.Sprintf("//%s", selectedLine)
	}

	// write back
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
func uncommentLine(fileName string, lineNumber uint16) {
	// create file handler
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	// read lines
	lines := strings.Split(string(input), "\n")

	// edit
	selectedLine := lines[lineNumber-1]
	runes := []rune(selectedLine)
	if runes[0] == '/' && runes[1] == '/' {
		lines[lineNumber-1] = string(selectedLine[2:])
	}

	// write back
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
