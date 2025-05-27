import { CLICommand, State } from "./state.js";

export async function commandHelp(state: State){
    for (const [name, command] of Object.entries(state.commands)) {{
        console.log(`${name}: ${command.description}`);
    }}    
}