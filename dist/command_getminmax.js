export async function commandGetMinMax(state) {
    let min = 9999999, max = -999999;
    for (let i = 1; i < 1400; i++) {
        try {
            const pokemon = await state.pokeapi.fetchPokemon(i.toString());
            if (pokemon.base_experience < min) {
                min = pokemon.base_experience;
            }
            if (pokemon.base_experience > max) {
                max = pokemon.base_experience;
            }
            console.log(`min:${min}, max:${max}`);
        }
        catch (err) {
            continue;
        }
    }
}
