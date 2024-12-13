package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
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
	argsUse     string
}

type Config struct {
	page    int
	cache   pokecache.Cache
	pokedex map[string]Pokemon
}

func main() {
	var cfg Config
	cfg.page = 1
	duration, err := time.ParseDuration("60s")
	if err != nil {
		fmt.Printf("Error parsing duration: %s\n", err)
	}
	cfg.cache = *pokecache.NewCache(duration)
	cfg.pokedex = make(map[string]Pokemon)

	for {
		fmt.Print("Pokédex >")
		cmdArgs, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			fmt.Printf("Couldn't read: %s", err)
		}

		cmdArgs = strings.TrimSpace(cmdArgs)
		cmdArgsSlice := strings.Split(cmdArgs, " ")
		cmd := cmdArgsSlice[0]
		args := strings.Join(cmdArgsSlice[1:], " ")
		Run(cmd, args, &cfg)
	}
}

func parseArgs(cmd CliCommand, args string) error {
	switch cmd.name {
	case "explore":
	case "catch":
	case "inspect":
		if len(strings.Split(args, " ")) != 1 {
			return fmt.Errorf(cmd.argsUse)
		}

	default:
		if args != "" {
			return fmt.Errorf(cmd.argsUse)
		}
	}

	return nil
}

func Run(cmd string, args string, cfg *Config) {
	commandList := getCommandList()
	cmdStruct, ok := commandList[cmd]
	if !ok {
		fmt.Printf("Command not recognized. Type 'help' to see options\n")
		return
	}
	err := parseArgs(cmdStruct, args)
	if err != nil {
		fmt.Printf("Invalid format: %s", err)
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
			argsUse:     "help takes no arguments",
		},
		"exit": {
			name:        "exit",
			description: "Exit pokedex",
			callback:    commandExit,
			argsUse:     "exit takes no arguments",
		},
		"map": {
			name:        "map",
			description: "Displays the name of 20 locations in Pokemon world",
			callback:    commandMap,
			argsUse:     "map takes no arguments",
		},
		"mapb": {
			name:        "mapb",
			description: "Display the name of last 20 locations in Pokemon world",
			callback:    commandMapb,
			argsUse:     "mapb takes no arguments",
		},
		"explore": {
			name:        "explore",
			description: "Explore region descripted in args. To see explorable regions use map.",
			callback:    commandExplore,
			argsUse:     "invalid argument. Use 'explore region' or 'explore id' instead",
		},
		"catch": {
			name:        "catch",
			description: "A trial to catch the pokemon descripted in args. If you're succesful the pokemon is going to appear in your pokedex",
			callback:    commandCatch,
			argsUse:     "invalid argument. Use 'catch pokemon_name' or 'catch pokemon_id' instead",
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect pokemons on pokedex",
			callback:    commandInspect,
			argsUse:     "invalid argument. Use 'inspect pokemon_name' or 'inspect pokemon_id' instead",
		},
	}
}

func getMapUrl(res string) string {
	return fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", res)
}

func getPokemonUrl(res string) string {
	return fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", res)
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
		url := getMapUrl(strconv.Itoa(i))
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
	url := getMapUrl(args)
	fetchResource(url, &mapToExplore, cfg)
	for _, encounter := range mapToExplore.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
}

func fetchResource(url string, resourceStruct interface{}, cfg *Config) error {
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
	switch v := resourceStruct.(type) {
	case *pokemap:
		err = json.Unmarshal(body, v)
	case *Pokemon:
		err = json.Unmarshal(body, v)
	default:
		return fmt.Errorf("unsupported type")
	}

	if err != nil {
		return err
	}

	return nil
}

func commandCatch(args string, cfg *Config) {
	url := getPokemonUrl(args)
	var resourcePokemon Pokemon
	err := fetchResource(url, &resourcePokemon, cfg)

	if err != nil {
		fmt.Printf("Error fetching pokemon %s\n", err)
	}

	baseExp := resourcePokemon.BaseExperience
	baseExp = int(max(baseExp, 200.0))
	name := resourcePokemon.Name
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	randomNumber := rand.Intn(250)
	if randomNumber >= baseExp {
		cfg.pokedex[name] = resourcePokemon
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
}

func commandInspect(args string, cfg *Config) {
	if pokemon_struct, ok := cfg.pokedex[args]; ok {
		fmt.Printf("Name: %s\n", pokemon_struct.Name)
		fmt.Printf("Height: %d\n", pokemon_struct.Height)
		fmt.Printf("Weight: %d\n", pokemon_struct.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon_struct.Stats {
			fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, type_pokemon := range pokemon_struct.Types {
			fmt.Printf("\t-%s\n", type_pokemon.Type.Name)
		}
	} else {
		fmt.Println("You have not caught that pokemon")
	}
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

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
