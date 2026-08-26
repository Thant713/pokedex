# Pokedex

A CLI REPL that lets users explore the Pokemon world via the PokeAPI, catch Pokemon, and inspect their stats.

## Language

**Location Area**:
A named region in the Pokemon world where Pokemon can be found. Fetched from the PokeAPI location-area endpoint in paginated batches of 20.
_Avoid_: location, region, zone

**Catch**:
A probabilistic attempt to capture a Pokemon. Success depends on the Pokemon's base experience — higher experience means lower catch rate.
_Avoid_: capture, grab, collect

**Pokedex**:
The player's in-memory collection of successfully caught Pokemon. Lost when the session ends.
_Avoid_: inventory, party, team

**Inspect**:
Display a caught Pokemon's name, height, weight, and stats.
_Avoid_: view, show, details
