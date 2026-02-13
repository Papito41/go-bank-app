package main

import "fmt"

// --- YESTERDAY'S BLUEPRINTS ---
type User struct {
	Name string
	ID   int
}

type Account struct {
	Customer User
	Balance  float64
}

// --- YESTERDAY'S TOOLS ---
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

func withdraw(acc *Account, amount float64) {
	if amount > acc.Balance {
		fmt.Println("Insufficient funds!")
		return
	}
	acc.Balance -= amount
	fmt.Printf("Withdrew %.2f. Remaining Balance: %.2f\n", amount, acc.Balance)
}

// --- TODAY'S SYSTEM (DAY 3) ---
func main() {
	// 1. We create the "Bank" (The Slice/List)
	var bank []Account

	// 2. We create individual accounts (Yesterday's work)
	paul := createAccount("Paul", 1, 5000.0)
	jane := createAccount("Jane", 2, 3000.0)

	// 3. We "Park" them in the Bank list (Today's work)
	bank = append(bank, paul)
	bank = append(bank, jane)

	// This loop "walks" through the bank and gives everyone 1000.0
	for i := range bank {
		// 1. &bank[i] gets the address of the person at position 'i'
		// 2. 1000.0 is the gift amount
		deposit(&bank[i], 1000.0)
	}

	// Final check to see the updated balances for everyone
	fmt.Println("Final Bank State:", bank)

	fmt.Println("--- Welcome to the Bank ---")

	// 4. How to use yesterday's functions on today's list:
	// We use &bank[0] to give the deposit function Paul's address in the list
	deposit(&bank[0], 2000)

	// We use &bank[1] to give the withdraw function Jane's address
	withdraw(&bank[1], 500)

	// 5. Printing the entire bank to see everyone's updated status
	fmt.Println("\nFinal Bank Registry:", bank)
}
