export async function commandCatch(state, ...args) {
    try {
        const pokemon = await state.pokeapi.fetchPokemon(args[0]);
        const pokemonName = pokemon.name;
        console.log(`Throwing a Pokeball at ${pokemonName}...`);
        const baseExp = pokemon.base_experience; //min:36, max:635
        const trial = Math.random() * 700;
        if (trial > baseExp) {
            state.pokedex[pokemonName] = pokemon;
            console.log(`${pokemonName} was caught!`);
        }
        else {
            console.log(`${pokemonName} escaped!`);
        }
    }
    catch (err) {
        if (err instanceof Error) {
            console.log(err.message);
        }
    }
}
