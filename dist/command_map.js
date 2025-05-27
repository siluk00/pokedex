import { helperMapAndMapb } from "./command_helpers.js";
export async function commandMap(state) {
    await helperMapAndMapb(state, state.nextLocationsURL);
    const offset = state.nextLocationsURL.searchParams.get("offset");
    if (offset) {
        const offsetInt = parseInt(offset);
        state.nextLocationsURL = new URL(`https://pokeapi.co/api/v2/location-area?offset=${offsetInt + 20}`);
        state.prevLocationsURL = new URL(`https://pokeapi.co/api/v2/location-area?offset=${offsetInt}`);
    }
}
