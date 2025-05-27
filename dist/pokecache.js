export class Cache {
    #cache;
    #reapIntervalId = undefined;
    #interval;
    constructor(interval) {
        this.#cache = new Map();
        this.#interval = interval;
        this.#startReapLoop();
    }
    add(key, val) {
        this.#cache.set(key, { createdAt: Date.now(), val: val });
        //console.log(this.#cache);
    }
    get(key) {
        return this.#cache.get(key);
    }
    #reap() {
        for (const [key, value] of this.#cache) {
            //console.log(Date.now()-value.createdAt);
            if (Date.now() - value.createdAt > this.#interval) {
                this.#cache.delete(key);
            }
        }
    }
    #startReapLoop() {
        this.#reapIntervalId = setInterval(this.#reap.bind(this), this.#interval);
    }
    stopReapLoop() {
        clearInterval(this.#reapIntervalId);
        this.#reapIntervalId = undefined;
    }
}
