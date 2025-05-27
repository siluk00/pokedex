import { commandExit } from "./command_exit.js";
import { commandExplore } from "./command_explore.js";
import { commandCatch } from "./command_catch.js";
import { commandHelp } from "./command_help.js";
import { commandMap } from "./command_map.js";
import { commandMapb } from "./command_mapb.js";
import { commandInspect } from "./command_inspect.js";
import { commandPokedex } from "./command_pokedex.js";
export function getCommands() {
    return {
        exit: {
            name: "exit",
            description: "exits pokedex",
            callback: commandExit,
        },
        help: {
            name: "help",
            description: "exhibits help",
            callback: commandHelp,
        },
        map: {
            name: "map",
            description: "shows next 20 locations on map",
            callback: commandMap,
        },
        mapb: {
            name: "mapb",
            description: "shows the previous 20 locations on map",
            callback: commandMapb,
        },
        explore: {
            name: "explore",
            description: "allows to explore a location by name or id",
            callback: commandExplore,
        },
        catch: {
            name: "catch",
            description: "produces an attempt to catch pokemon",
            callback: commandCatch,
        },
        inspect: {
            name: "inspect",
            description: "inspects a pokemon in the pokedex",
            callback: commandInspect,
        },
        pokedex: {
            name: "pokedex",
            description: "allows you to see the pokemon you have caught",
            callback: commandPokedex,
        },
    };
}
