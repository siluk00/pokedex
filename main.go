package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	pokecache "github.com/siluk00/pokedex/internal"
)

type CliCommand struct {
	name        string
	description string
	callback    func(string, *Config)
}

type Config struct {
	page  int
	cache pokecache.Cache
}

func main() {
	var cfg Config
	cfg.page = 1
	duration, err := time.ParseDuration("60s")
	if err != nil {
		fmt.Printf("Error parsing duration: %s\n", err)
	}
	cfg.cache = *pokecache.NewCache(duration)

	for {
		fmt.Print("Pokédex >")
		cmdArgs, err := bufio.NewReader(os.Stdin).ReadString('\n')
		cmdArgs = strings.TrimSpace(cmdArgs)
		cmdArgsSlice := strings.Split(cmdArgs, " ")
		cmd := cmdArgsSlice[0]
		args := strings.Join(cmdArgsSlice[1:], " ")

		if err != nil {
			log.Fatalf("Couldn't read: %s", err)
		}

		Act(cmd, args, &cfg)
	}
}

func Act(cmd string, args string, cfg *Config) {
	commandList := getCommandList()
	cmdStruct, ok := commandList[cmd]
	if !ok {
		fmt.Printf("Command not recognized. Type 'help' to see options\n")
		return
	}
	cmdStruct.callback(args, cfg)
}

func commandHelp(args string, cfg *Config) {
	fmt.Printf("Welcome to pokédex! Usage: \n\n")

	cmdList := getCommandList()

	for i := range cmdList {
		fmt.Printf("%s:\t%s\n", cmdList[i].name, cmdList[i].description)
	}
}

func commandExit(args string, cfg *Config) {
	os.Exit(0)
}

func getCommandList() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the name of 20 locations in Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the name of last 20 locations in Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore region descripted in args. To see explorable regions use map.",
			callback:    commandExplore,
		},
	}
}

func getUrl(res string) string {
	return fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", res)
}

func commandMapb(args string, cfg *Config) {
	mapHelper(cfg, true)
}

func commandMap(args string, cfg *Config) {
	mapHelper(cfg, false)
}

func mapHelper(cfg *Config, reverse bool) {

	if reverse {
		if cfg.page > 20 {
			cfg.page -= 20
		} else {
			cfg.page = 0
		}
	}

	i := cfg.page

	for i < cfg.page+20 {
		url := getUrl(strconv.Itoa(i))
		resource, ok := cfg.cache.Get(url)
		var nextMap pokemap

		if !ok {
			err := fetchResource(url, &nextMap, cfg)
			if err != nil {
				fmt.Printf("Error fetching resource: %s\n", err)
				break
			}
		} else { //The value exists on the cache
			err := json.Unmarshal(resource, &nextMap)
			if err != nil {
				fmt.Printf("Error unmarshalling response: %s\n", err)
				break
			}
		}

		fmt.Println(nextMap.Name)
		i++
	}

	cfg.page = i
}

func commandExplore(args string, cfg *Config) {
	if len(strings.Split(args, " ")) != 1 {
		fmt.Printf("Invalid argument. Use 'explore region' or 'explore id' instead\n")
		return
	}

	var mapToExplore pokemap
	url := getUrl(args)
	fetchResource(url, &mapToExplore, cfg)

}

func fetchResource(url string, mapToExplore *pokemap, cfg *Config) error {
	res, err := http.Get(url)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code %d and body %s", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	cfg.cache.Add(url, body)
	err = json.Unmarshal(body, mapToExplore)
	if err != nil {
		return err
	}

	return nil

}

type pokemap struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
