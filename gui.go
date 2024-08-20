package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func DisplayMainMenu(scanner *bufio.Scanner) string {
	fmt.Println("1. Create New Budget")
	fmt.Println("2. View Budgets")
	fmt.Println("3. Save and Exit")

	scanner.Scan()
	return scanner.Text()
}

func DisplayBudgetMenu(scanner *bufio.Scanner) string {
	fmt.Println("1. Add Expense")
	fmt.Println("2. Delete Expense")
	fmt.Println("3. Finish")

	scanner.Scan()
	return scanner.Text()
}

func PromptForBudgetName(scanner *bufio.Scanner) string {
	fmt.Println("Enter budget name:")
	scanner.Scan()
	return scanner.Text()
}

func PromptForIncome(scanner *bufio.Scanner) (float64, error) {
	fmt.Println("Enter your total income:")
	scanner.Scan()
	incomeStr := scanner.Text()
	return strconv.ParseFloat(incomeStr, 64)
}

func PromptForExpenseName(scanner *bufio.Scanner) string {
	fmt.Println("Enter expense name:")
	scanner.Scan()
	return scanner.Text()
}

func PromptForExpenseAmount(scanner *bufio.Scanner) (float64, error) {
	fmt.Println("Enter expense amount:")
	scanner.Scan()
	expenseAmountStr := scanner.Text()
	return strconv.ParseFloat(expenseAmountStr, 64)
}

func DisplayBudgets(budgets []Budget) {
	if len(budgets) == 0 {
		fmt.Println("No budgets available.")
		return
	}

	for _, budget := range budgets {
		fmt.Printf("Budget Name: %s\n", budget.Name)
		fmt.Printf("Total Income: %.2f\n", budget.Income)
		fmt.Println("Expenses:")
		totalExpenses := 0.0
		for _, expense := range budget.Expenses {
			fmt.Printf("- %s: $%.2f\n", expense.Name, expense.Amount)
			totalExpenses += expense.Amount
		}
		fmt.Printf("Total Expenses: $%.2f\n", totalExpenses)
		fmt.Printf("Remaining Balance: $%.2f\n\n", budget.Income-totalExpenses)
	}
}

func DisplayExpenseNotFound() {
	fmt.Println("Expense not found.")
}

func DisplayExpenseAdded() {
	fmt.Println("Expense added successfully.")
}

func DisplayExpenseDeleted() {
	fmt.Println("Expense deleted successfully.")
}

func DisplayInvalidOption() {
	fmt.Println("Invalid option. Please choose again.")
}

func DisplayBudgetSaved() {
	fmt.Println("Budget saved successfully.")
}

func DisplayError(message string) {
	fmt.Println("Error:", message)
}
