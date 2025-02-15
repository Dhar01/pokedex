# PokeDex CLI

A command-line tool to look up Pokémon information!

# Commands

- **`catch`**: Catch a Pokémon by specifying its name.
    <br>*Example*: `catch <pokemon-name>`
- **`exit`**: Exit the application gracefully.
- **`explore`**: Explore Pokémon available in a specific location.
    <br>*Example*: `explore <location-name>`
- **`help`**: Looking for help?
- **`mapb`**: Displays the names of the previous 20 location areas.
- **`map`**: Displays the names of 20 location areas.
- **`pokedex`**: Show a list of all the Pokémon you have caught so far!

# Execution

- **Requirements**: Go 1.20 or higher
- **Run**: Execute the `run.sh` script from the command line.

Enjoy exploring the Pokémon world from your terminal!

# To-Do

- [ ] Update the CLI to support the “up” arrow to cycle through previous commands
- [ ] Simulate battles between pokemon
- [ ] Add more unit tests
- [ ] Refactor your code to organize it better and make it more testable
- [ ] Keep pokemon in a “party” and allow them to level up
- [ ] Allow for pokemon that are caught to evolve after a set amount of time
- [ ] Persist a user’s Pokedex to disk so they can save progress between sessions
- [ ] Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type “left” or “right”
- [ ] Random encounters with wild pokemon
- [ ] Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon


# Progress

Visit the [Project](https://github.com/users/Dhar01/projects/1/views/1) section on GitHub to track development progress!


![code coverage badge](https://github.com/Dhar01/pokedex/actions/workflows/ci.yml/badge.svg)