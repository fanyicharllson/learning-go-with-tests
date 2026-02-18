# Learning Go with TDD 🚀

This repository tracks my journey of mastering **Go (Golang)** using **Test-Driven Development (TDD)**. Following the "Red-Green-Refactor" cycle, I am building a foundation of high-performance backend engineering.

## 🧠 Why TDD?
In Go, testing isn't an afterthought—it's a core feature. By writing tests first, I ensure:
- **Zero Regressions:** I can refactor code with confidence.
- **Self-Documenting Code:** The tests explain exactly how the logic should behave.
- **Edge Case Coverage:** Handling empty strings, nil pointers, and overflows from the start.

## 🛠 Project Structure
Each folder represents a core concept from the [Learn Go with Tests](https://quii.gitbook.io) curriculum:
- `hello-world/`: Subtests and basic syntax.
- `integers/`: Working with types and documentation.
- `iteration/`: Loops and benchmarking.
- *(Coming Soon)*: Slices, Maps, and Pointers.

## 🚦 How to run tests
To run all tests in this repository:
```bash
go test ./... -v
