export function startREPL(state) {
    const rl = state.rl;
    console.log("Welcome to the Pokedex!");
    rl.prompt();
    rl.on('line', async (input) => {
        const cleanedInput = cleanInput(input);
        if (cleanedInput.length === 0) {
            rl.prompt();
        }
        const commands = state.commands;
        const command = commands[cleanedInput[0]];
        if (command) {
            await command.callback(state, ...cleanedInput.slice(1));
        }
        else {
            console.log("Unknown command");
        }
        rl.prompt();
    });
}
export function cleanInput(input) {
    return input.trimStart().trimEnd().split(/\s+/);
}
