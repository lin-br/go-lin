This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Repository Overview

This is a Go learning repository (`go-lin`) structured as a progressive tutorial series covering fundamental Go concepts. The codebase follows the "Learn Go with Tests" approach, with each lesson contained in its own numbered directory under `learn-go/`.

### Module Structure
- **Module**: `github.com/lin-br/go-lin`
- **Go Version**: 1.24 (using modern Go features like iterators)
- **Architecture**: Educational codebase with 17 structured lessons covering Go fundamentals to advanced concepts

## Common Development Commands

### Testing
```bash
# Run all tests across the entire codebase
go test ./...

# Run tests for a specific lesson with verbose output
go test -v ./learn-go/01-hello-world

# Run a specific test case
go test -run TestHello ./learn-go/01-hello-world

# Run tests with coverage
go test -cover ./...

# Run tests in a specific directory
cd learn-go/05-structs && go test -v
```

### Building and Running
```bash
# Run a specific lesson's main function (where applicable)
go run ./learn-go/09-mocking

# Build a specific lesson
go build ./learn-go/09-mocking

# Format all Go code
go fmt ./...

# Vet code for common issues
go vet ./...
```

## Code Architecture and Patterns

### Learning Path Structure
The repository follows a progressive learning structure with lessons covering:

1. **Basic Concepts** (01-04): Hello World, integers, iteration, arrays
2. **Core Go Features** (05-08): Structs, pointers, maps, dependency injection
3. **Advanced Patterns** (09-14): Mocking, concurrency, select, reflection, sync, context
4. **Specialized Topics** (15-17): Property-based testing, mathematics, file I/O

### Key Architectural Patterns

#### Interface-Driven Design
- Clean separation of concerns using interfaces (e.g., `Sleeper` interface in lesson 09)
- Dependency injection patterns for testability
- Mock implementations for unit testing

#### Test-Driven Development
- Each lesson follows TDD principles with comprehensive test coverage
- Subtests using `t.Run()` for organized test scenarios
- Test helpers with `t.Helper()` for cleaner error messages
- Assertion patterns with clear error messages

#### Modern Go Features (Go 1.24)
- **Iterators**: Uses `iter.Seq[T]` for custom iteration patterns (lesson 09)
- **Range over functions**: Modern `for i := range countDownFrom(start)` syntax
- **Context handling**: Proper context propagation in HTTP handlers (lesson 14)

#### Concurrency Patterns
- Goroutines with channels for concurrent operations (lesson 10)
- Select statements for channel orchestration (lesson 11)
- Proper synchronization using sync package (lesson 13)
- Context-based cancellation and timeouts (lesson 14)

### Testing Patterns
- **Table-driven tests**: Not extensively used, focus on subtests instead
- **Test doubles**: Mock objects and spy patterns for behavior verification
- **Property-based testing**: Introduction in lesson 15
- **Acceptance testing**: Used for complex scenarios like clockface rendering

## Development Workflow

### Working with Individual Lessons
Each lesson is self-contained and can be worked on independently:
```bash
# Navigate to a specific lesson
cd learn-go/12-reflection

# Run tests for just that lesson
go test -v

# Run the main function if it exists
go run .
```

### Code Quality
- All code follows standard Go formatting (`go fmt`)
- No external dependencies beyond the standard library
- Clean, readable code with meaningful variable names
- Comprehensive test coverage for all functionality

### Git Workflow
The repository uses standard Git practices. When making changes:
- Tests should pass before committing (`go test ./...`)
- Code should be formatted (`go fmt ./...`)
- Consider running `go vet ./...` to catch common issues
