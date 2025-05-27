import { helperMapAndMapb } from "./command_helpers.js";
export async function commandMapb(state) {
    await helperMapAndMapb(state, state.prevLocationsURL);
    const url = state.prevLocationsURL;
    if (url) {
        const offset = url.searchParams.get("offset");
        if (offset) {
            const offsetInt = parseInt(offset);
            state.nextLocationsURL = new URL(`https://pokeapi.co/api/v2/location-area?offset=${offsetInt}`);
            if (offsetInt >= 20) {
                state.prevLocationsURL = new URL(`https://pokeapi.co/api/v2/location-area?offset=${offsetInt - 20}`);
            }
            else {
                state.prevLocationsURL = undefined;
            }
        }
    }
    else {
        state.nextLocationsURL = new URL(`https://pokeapi.co/api/v2/location-area?offset=0`);
        state.prevLocationsURL = undefined;
    }
}
