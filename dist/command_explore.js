export async function commandExplore(state, ...args) {
    try {
        const response = await state.pokeapi.fetchLocation(args[0]);
        const pokemonsInLocation = response.pokemon_encounters;
        console.log(`Exploring ${args[0]}...`);
        console.log("Found Pokemon:");
        for (const pokemon of pokemonsInLocation) {
            console.log(`- ${pokemon.pokemon.name}`);
        }
    }
    catch (err) {
        if (err instanceof Error) {
            console.log(err.message);
        }
    }
}
