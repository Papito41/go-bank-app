package main

import "fmt"

// 1. Defining the "Boxes" (Data Structures)
type User struct {
	Name string
	ID   int
}

type Account struct {
	Owner   User // This connects the User to the Account
	Balance float64
}

// 2. The Factory (Constructor)
// This fixes the error where you tried to use "InitialBalance" as a type.
func createAccount(newName string, newID int, startMoney float64) Account {
	// Create the User first
	newUser := User{Name: newName, ID: newID}

	// Put that User into the Account
	newAcc := Account{
		Owner:   newUser,
		Balance: startMoney,
	}

	return newAcc // Hand back the finished product
}

// 3. The Inspector (The View)
func checkBalance(c Account) {
	// Reach inside: Account -> Owner -> Name
	fmt.Printf("Owner: %s | ID: %d | Balance: ₦%.2f\n", c.Owner.Name, c.Owner.ID, c.Balance)
}

func deposit(a *Account, amount float64) {
	a.Balance = a.Balance + amount
}

func withdraw(a *Account, amount float64) {
	if amount > a.Balance {
		fmt.Println("Insufficient funds!")
	} else {
		a.Balance = a.Balance - amount
	}
}

func main() {
	// Calling the factory to make your account
	myAcc := createAccount("Paul", 1, 5000.50)
	fmt.Println("Initial State:")
	checkBalance(myAcc)

	deposit(&myAcc, 2000.00)
	fmt.Println("\n After Deposit of 2000.00:")
	checkBalance(myAcc)

	withdraw(&myAcc, 1500.00)
	fmt.Println("\n After Withdrawal of 1500.00")
	checkBalance(myAcc)

	fmt.Println("\n Testing Insufficient (Attempting 10,000.00):")
	withdraw(&myAcc, 10000)

	// Using the inspector to see it
	checkBalance(myAcc)
}
