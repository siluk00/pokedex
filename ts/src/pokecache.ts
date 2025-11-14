export type CacheEntry<T> = {
    createdAt: number;
    val: T;
};

export class Cache {
    #cache: Map<string, CacheEntry<any>>;
    #reapIntervalId: NodeJS.Timeout | undefined = undefined;
    #interval: number;

    constructor(interval: number) {
        this.#cache =  new Map<string, CacheEntry<any>>();
        this.#interval = interval;
        this.#startReapLoop();
    }

    add<T>(key: string, val: T) {
        this.#cache.set(key, {createdAt: Date.now(), val: val});
        //console.log(this.#cache);
    }

    get<T>(key:string) {
        return this.#cache.get(key);
    }

    #reap() {
        for (const [key, value] of this.#cache) {
            //console.log(Date.now()-value.createdAt);
            if (Date.now()-value.createdAt > this.#interval) {
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