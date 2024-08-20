package main

import (
	"bufio"
)

func main() {
	createGUI()
}

func createBudget(scanner *bufio.Scanner) {
	budgetName := PromptForBudgetName(scanner) // Prompt for budget name

	income, err := PromptForIncome(scanner) // Prompt for income amount
	if err != nil {
		DisplayError("Invalid income amount, please enter a number.") // Display error message
		return
	}

	budget := Budget{ // Create new budget
		Name:   budgetName,
		Income: income,
	}

	addExpenses(scanner, &budget)
	budgets = append(budgets, budget) // Add budget to the list of budgets
	DisplayBudgetSaved()              // Display success message
}

func addExpenses(scanner *bufio.Scanner, budget *Budget) {
	for {
		choice := DisplayBudgetMenu(scanner) // Display expense menu and get user choice

		switch choice {
		case "1":
			expenseName := PromptForExpenseName(scanner) // Prompt for expense name

			expenseAmount, err := PromptForExpenseAmount(scanner) // Prompt for expense amount
			if err != nil {
				DisplayError("Invalid amount, please enter a number.") // Display error message
				continue
			}

			expense := Expense{ // Create new expense
				Name:   expenseName,
				Amount: expenseAmount,
			}

			budget.Expenses = append(budget.Expenses, expense) // Add expense to the budget
			DisplayExpenseAdded()                              // Display success message

		case "2":
			expenseName := PromptForExpenseName(scanner) // Prompt for expense name to delete

			index := -1
			for i, expense := range budget.Expenses {
				if expense.Name == expenseName {
					index = i
					break
				}
			}

			if index == -1 {
				DisplayExpenseNotFound() // Display message if expense not found
			} else {
				budget.Expenses = append(budget.Expenses[:index], budget.Expenses[index+1:]...) // Remove expense
				DisplayExpenseDeleted()                                                         // Display success message
			}

		case "3":
			return

		default:
			DisplayInvalidOption() // Display invalid option message
		}
	}
}

func viewBudgets() {
	DisplayBudgets(budgets) // Display all budgets
}

func loadBudgets() {
	// Function to load budgets from file
	// This can be implemented using JSON or other storage methods
}

func saveBudgets() {
	// Function to save budgets to file
	// This can be implemented using JSON or other storage methods
}
