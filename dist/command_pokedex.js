export async function commandPokedex(state) {
    console.log("Your Pokedex:");
    for (const [_, pokemon] of Object.entries(state.pokedex)) {
        console.log(`\t- ${pokemon.name}`);
    }
}
