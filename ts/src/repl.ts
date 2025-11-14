import { exit } from "process";
import { stringify } from "querystring";
import { Readline } from "readline/promises";
import readline from "node:readline";
import Stream from "stream";
import { getCommands } from "./registry.js";
import { State } from "./state.js";


export function startREPL(state: State) {
  const rl = state.rl;
  console.log("Welcome to the Pokedex!");
  rl.prompt();
  rl.on('line', async (input: string) => {
    
    const cleanedInput = cleanInput(input);
    if (cleanedInput.length === 0) {
      rl.prompt();
    }
  
    const commands = state.commands;
    const command = commands[cleanedInput[0]];
  
    if (command) {
      await command.callback(state, ...cleanedInput.slice(1));
    } else {
      console.log("Unknown command");
    }
    
    rl.prompt();
  });
}

export function cleanInput(input: string): string[] {
  return input.trimStart().trimEnd().split(/\s+/);
}