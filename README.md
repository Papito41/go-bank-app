# Go Bank Backend

A secure, idiomatic Go implementation of a banking system backend. This project demonstrates core Go concepts including custom data structures, factory patterns, and memory management using pointers.

## 🚀 Features Built So Far

### Day 1: Foundation & Data Modeling
* **Nested Structs:** Implemented a `User` struct within an `Account` struct to mimic real-world data relationships.
* **Constructor Pattern:** Created a `createAccount` factory function to handle object initialization safely.

### Day 2: Financial Logic & Security
* **State Management:** Used **Pointer Receivers** (`*Account`) to allow functions to modify the actual account balance in memory.
* **Transaction Guard:** Implemented conditional logic (`if/else`) to prevent overdrafts and ensure "Insufficient Funds" handling.

### Day 3: Scaling to a Banking System
- **Slices:** Moved from single variables to a dynamic `[]Account` collection to manage multiple users.
- **Append:** Used the built-in `append` function to dynamically add new customers to the bank registry.
- **Loops:** Implemented `for range` loops to automate transactions across the entire registry.
- **Pointers in Slices:** Managed memory addresses within lists to ensure data persistence during bulk updates.

---

## 📅 Day 4: Data Structure Optimization (The O(1) Upgrade)

Today, I refactored the bank's core storage engine to move away from linear search patterns.

### Technical Achievements:
- **From O(n) to O(1):** Replaced slice-based storage with **Go Maps**. In a slice, finding a user requires "walking" through every element (Linear Time). With Maps, I implemented a Hash Table approach that allows for **instant** (Constant Time) lookups regardless of whether the bank has 10 users or 10 million.
- **The "Comma OK" Idiom:** Implemented robust search logic using Go’s `value, ok := map[key]` pattern to handle missing users gracefully without system panics.
- **Memory Management:** Practiced saving modified data back into the Map, acknowledging that Go Maps return copies of values.

### Why this matters:
Starting with memory-first optimization ensures that the "hot path" of the application is as fast as possible before introducing the latency of a physical database.

Bank App Evolution: Day 5 - The Database Handshake

Today, I moved from storing data in temporary RAM (Go Maps) to preparing for a real-world database (PostgreSQL). 

### 🚀 What's New
* **Database Driver:** Integrated the `pgx/v5` driver to allow the Go engine to talk to PostgreSQL.
* **Connection Logic:** Implemented a handshake script (`test_db.go`) to test database reachability using `conn.Ping()`.
* **Context Management:** Started using Go's `context` package to handle timeouts and prevent the app from hanging if the database is down.

### 🧠 Architectural Shift
We are moving away from **In-Memory Storage** because:
1.  **Persistence:** Data needs to survive a server restart.
2.  **Concurrency:** Maps aren't safe when multiple people withdraw at once (Atomic Transactions are coming next!).

### 🛠 How to Test the Handshake
1.  Run `go get github.com/jackc/pgx/v5` to install dependencies.
2.  Run the test script (currently set up as a standalone check).

## 🛠 Tech Stack
* **Language:** Go (Golang)
* **Concepts:** Pointers, Structs, Logic Flow, Formatting.

## 🏃 How to Run
1. Ensure you have Go installed on your machine.
2. Clone this repository.
3. Run the following command in your terminal:
   ```bash
   go run .