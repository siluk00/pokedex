import { createInterface, type Interface } from "node:readline";
import { getCommands } from "./registry.js";
import { PokeAPI, Pokemon } from "./pokeapi.js";

export type CLICommand = {
    name: string;
    description: string;
    callback: (state: State, ...args: string[]) => Promise<void>;
};

export type State = {
    rl: Interface;
    commands: Record<string, CLICommand>;
    pokeapi: PokeAPI;
    nextLocationsURL: URL;
    prevLocationsURL: URL | undefined;
    pokedex: Record<string, Pokemon>;
};

export function initState(): State {
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