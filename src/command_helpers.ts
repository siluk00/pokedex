import { State } from "./state";

export async function helperMapAndMapb(state: State, url: URL| undefined) {
    try {
        if (url) {
            const locations = await state.pokeapi.fetchLocations(url.toString());
            const results = locations.results;
            for (const location of results) {
                console.log(location.name);
            }
        }
    } catch(err: unknown) {
        if (err instanceof Error) {
            console.log(err.message);
        }
    }
}