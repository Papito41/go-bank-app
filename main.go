package main

import "fmt"

type User struct {
	Name string
	ID   int
}

type Account struct {
	Customer User
	Balance  float64
}

func createAccount(name string, id int, initialBalance float64) Account {
	return Account{
		Customer: User{Name: name, ID: id},
		Balance:  initialBalance,
	}
}

func deposit(acc *Account, amount float64) {
	acc.Balance += amount
	fmt.Printf("Deposited %.2f. New Balance: %.2f\n", amount, acc.Balance)
}

func main() {
	// Initialize Bank Map (O(1) Access)
	bank := make(map[string]Account)

	// Adding Users
	bank["Paul"] = createAccount("Paul", 1, 5000.0)
	bank["Jane"] = createAccount("Jane", 2, 3000.0)
	bank["Bose"] = createAccount("Bose", 3, 10000.0)

	fmt.Println("--- Bank System Online (In-Memory) ---")

	// The Search Logic
	searchName := "Chinedu"
	acc, found := bank[searchName]

	if !found {
		fmt.Printf("Alert: Account for '%s' not found. Search complexity: O(1)\n", searchName)
	} else {
		fmt.Printf("Success: Found %s. Balance: %.2f\n", acc.Customer.Name, acc.Balance)
		deposit(&acc, 500)
		bank[searchName] = acc // Save back to map
	}

	fmt.Println("\nCurrent Bank Registry:", bank)
}
