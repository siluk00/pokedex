import { Cache } from "./pokecache.js";
import { expect, test } from "vitest";
test.concurrent.each([
    {
        key: "https://example.com",
        val: "testdata",
        interval: 500, // 1/2 second
    },
    {
        key: "https://example.com/path",
        val: "moretestdata",
        interval: 1000, // 1 second
    },
])("Test Caching $interval ms", async ({ key, val, interval }) => {
    const cache = new Cache(interval);
    cache.add(key, val);
    const cached = cache.get(key);
    expect(cached?.val).toBe(val);
    await new Promise((resolve) => setTimeout(resolve, 2 * interval));
    const reaped = cache.get(key);
    expect(reaped?.val).toBe(undefined);
    cache.stopReapLoop();
});
