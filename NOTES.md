# Pokedex CLI — Project Map

## Files
- `main.go` — entry point
- `repl.go` — REPL loop, command registry
- `command_*.go` — one file per command
- `internal/pokeapi/` — API client + types
- `files under pokeapi` - client    
                      - location_get
                      - location_list 
                      - pokeapi 
                      - types_locations

- `internal/pokecache/` — cache

## Commands
| Command | File | Status |
|---------|------|--------|
| help    | command_help.go | done |
| map     | command_map.go  | done |
| explore | command_explore.go | done |
| catch   | command_catch.go | done |
| inspect | command_inspect.go | WIP|

## TODO
- add inspect command 
- [ ] adding feature to see name, height, stats and types 


## Roles Details
- `client` - the messenger, which get info from http.
        - `owns` - the client struct, cache , timeouts. 

- `types_*` - defines whats pokeapi responses looks like in Go.
          -  structs that match JSON fields.

- `_get/ _list` - function per endpoint. 
              - `owns` - fetch one pokemon name and fetch many location

