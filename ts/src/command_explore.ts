import { State } from "./state";

export async function commandExplore(state: State, ...args: string[]) {
    try {
        const response = await state.pokeapi.fetchLocation(args[0]);
        const pokemonsInLocation = response.pokemon_encounters;
        console.log(`Exploring ${args[0]}...`);
        console.log("Found Pokemon:")
        for (const pokemon of pokemonsInLocation) {
            console.log(`- ${pokemon.pokemon.name}`);
        }
    } catch (err: unknown) {
        if (err instanceof Error) {
            console.log(err.message);
        }
    }

}