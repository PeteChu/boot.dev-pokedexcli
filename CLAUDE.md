# CLAUDE.md - Pokedex CLI

## Build/Test Commands
- Build: `go build -o pokedexcli`
- Run: `./pokedexcli`
- Test all: `go test ./...`
- Test a specific file: `go test -v ./path/to/file_test.go`
- Test a specific function: `go test -v -run=TestFunctionName`
- Test with coverage: `go test -cover ./...`

## Code Style Guidelines
- **Imports**: Group stdlib imports first, then external imports, then local imports
- **Formatting**: Use `gofmt` standard formatting
- **Types**: Define structs with proper JSON tags (`json:"name,omitempty"`)
- **Naming**: 
  - Use camelCase for private functions/variables
  - Use PascalCase for exported functions/variables/types
  - Use descriptive names for variables and functions
- **Error Handling**: Always check errors and return them with context
- **Documentation**: Document packages and exported functions with godoc-style comments
- **Code Organization**: Use the internal/ directory for packages not meant for external use

## Application Structure
- Main CLI logic in main.go
- Cache implementation in internal/pokecache/pokecache.go
- Tests in *_test.go files alongside the code they test