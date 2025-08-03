# Paradigms for interview problems
## About
This repository offers examples of how to solve common programming challenges using different **programming paradigms** and **problem-solving strategies**.  
The primary language used is **Go**, with a focus on writing clear and idiomatic code.

Each problem is solved in multiple ways to highlight:

- The differences between **imperative** (_procedural_) and **declarative** (_functional_) programming styles
- The use of strategies such as **dynamic programming**, **concurrency**, and the more traditional _iterative_ approach.

While many of the problems are similar to those found in coding interviews, the goal of this project is broader:  
to teach effective programming techniques, paradigms, and structured problem solving.

### Goals

- Serve as a practical resource for learning **Go**
- Introduce and compare **programming paradigms**
- Explore common **problem-solving strategies**
- Help developers prepare for **technical interviews** through deeper understanding
- Encourage writing clean, maintainable, and well-structured code

---

This repository is ideal for:

- Students and self-learners looking to improve their Go skills
- Developers interested in exploring different paradigms
- Anyone preparing for interviews

## How It's Organized
Each code file is organized following a consistent structure:
-  **Folder/module**: `problem` Each go mod represents a specific exercise (e.g. _primeNumber_). You can test all solution at this level using `go test ./...`
   - **folder**:  `paradigm` Inside each module, a folder defines the programming paradigm used (e.g., imperative or declarative).
        - **Classical**: `package` with a the verbose and highly commented solution that illustrates the basic, straightforward way to solve the problem. You can test each solution using `go test`at this level
        - **Dynamic**: `package` with a version using dynamic programming techniques, memoization or caching. You can test each solution using `go test`at this level
        - **Concurrent / Parallel**: `package` with a version that introduces concurrency or parallelism, often using threads or parallel streams.You can test each solution using `go test`at this level
