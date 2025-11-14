import { State } from "./state";

export async function commandPokedex(state: State) {
    console.log("Your Pokedex:");
    for (const [_, pokemon] of Object.entries(state.pokedex)) {
        console.log(`\t- ${pokemon.name}`);
    }
}