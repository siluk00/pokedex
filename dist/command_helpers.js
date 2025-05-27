export async function helperMapAndMapb(state, url) {
    try {
        if (url) {
            const locations = await state.pokeapi.fetchLocations(url.toString());
            const results = locations.results;
            for (const location of results) {
                console.log(location.name);
            }
        }
    }
    catch (err) {
        if (err instanceof Error) {
            console.log(err.message);
        }
    }
}
