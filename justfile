# Run tests for a specific chapter and level
# Usage: just test 1 4
test chapter level:
    cd chapter-{{chapter}}/level-{{level}} && go test -v

# Show all available commands
help:
    @just --list
