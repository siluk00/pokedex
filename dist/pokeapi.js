import { Cache } from "./pokecache.js";
export class PokeAPI {
    static baseURL = "https://pokeapi.co/api/v2";
    #cache = new Cache(10000);
    constructor() { }
    async fetchLocations(pageURL) {
        const fullURL = pageURL ? pageURL : `${PokeAPI.baseURL}/location-area?offset=0`;
        return await this.#fetchCache(fullURL);
    }
    async fetchPokemon(pokemon) {
        const fullURL = `${PokeAPI.baseURL}/pokemon/${pokemon}`;
        return await this.#fetchCache(fullURL);
    }
    async #fetchCache(fullURL) {
        const cacheData = this.#cache.get(fullURL);
        //console.log(cacheData);
        if (cacheData) {
            return cacheData.val;
        }
        const response = await fetch(fullURL, { method: 'GET' });
        const responseJson = await response.json();
        this.#cache.add(fullURL, responseJson);
        return responseJson;
    }
    async fetchLocation(locationName) {
        const fullURL = `${PokeAPI.baseURL}/location-area/${locationName}`;
        return await this.#fetchCache(fullURL);
    }
}
