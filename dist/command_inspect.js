export async function commandInspect(state, pokemonName) {
    const pokedex = state.pokedex;
    if (pokedex[pokemonName]) {
        const pokemon = pokedex[pokemonName];
        console.log(`Name: ${pokemon.name}`);
        console.log(`Height: ${pokemon.height}`);
        console.log(`Weight: ${pokemon.weight}`);
        console.log(`Stats:`);
        for (const stat of pokemon.stats) {
            console.log(`\t- ${stat.stat.name}: ${stat.base_stat}`);
        }
        console.log("Types: ");
        for (const pokType of pokemon.types) {
            console.log(`\t- ${pokType.type}`);
        }
    }
    else {
        console.log("you have not caught that pokemon");
    }
}
