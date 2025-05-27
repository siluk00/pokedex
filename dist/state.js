import { createInterface } from "node:readline";
import { getCommands } from "./registry.js";
import { PokeAPI } from "./pokeapi.js";
export function initState() {
    const rl = createInterface({
        input: process.stdin,
        output: process.stdout,
        prompt: 'Pokedex >',
    });
    return {
        rl: rl,
        commands: getCommands(),
        pokeapi: new PokeAPI(),
        nextLocationsURL: new URL("https://pokeapi.co/api/v2/location-area?offset=0"),
        prevLocationsURL: undefined,
        pokedex: {},
    };
}
