# Go Bank Backend

A secure, idiomatic Go implementation of a banking system backend. This project demonstrates core Go concepts including custom data structures, factory patterns, and memory management using pointers.

## 🚀 Features Built So Far

### Day 1: Foundation & Data Modeling
* **Nested Structs:** Implemented a `User` struct within an `Account` struct to mimic real-world data relationships.
* **Constructor Pattern:** Created a `createAccount` factory function to handle object initialization safely.

### Day 2: Financial Logic & Security
* **State Management:** Used **Pointer Receivers** (`*Account`) to allow functions to modify the actual account balance in memory.
* **Transaction Guard:** Implemented conditional logic (`if/else`) to prevent overdrafts and ensure "Insufficient Funds" handling.

## 🛠 Tech Stack
* **Language:** Go (Golang)
* **Concepts:** Pointers, Structs, Logic Flow, Formatting.

## 🏃 How to Run
1. Ensure you have Go installed on your machine.
2. Clone this repository.
3. Run the following command in your terminal:
   ```bash
   go run .