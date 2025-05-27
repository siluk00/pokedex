export async function commandHelp(state) {
    for (const [name, command] of Object.entries(state.commands)) {
        {
            console.log(`${name}: ${command.description}`);
        }
    }
}
