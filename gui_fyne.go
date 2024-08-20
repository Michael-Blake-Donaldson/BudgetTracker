package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createGUI() {
	// Create the application
	myApp := app.New()
	myWindow := myApp.NewWindow("Budget App")

	// Create UI elements
	budgetNameEntry := widget.NewEntry()
	budgetNameEntry.SetPlaceHolder("Enter Budget Name")

	incomeEntry := widget.NewEntry()
	incomeEntry.SetPlaceHolder("Enter Income")

	expenseNameEntry := widget.NewEntry()
	expenseNameEntry.SetPlaceHolder("Enter Expense Name")

	expenseAmountEntry := widget.NewEntry()
	expenseAmountEntry.SetPlaceHolder("Enter Expense Amount")

	expenseList := widget.NewList(
		func() int {
			return len(budgets)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			if len(budgets) > 0 && len(budgets[0].Expenses) > i {
				o.(*widget.Label).SetText(fmt.Sprintf("%s: $%.2f", budgets[0].Expenses[i].Name, budgets[0].Expenses[i].Amount))
			}
		})

	addExpenseButton := widget.NewButton("Add Expense", func() {
		name := expenseNameEntry.Text
		amount, err := strconv.ParseFloat(expenseAmountEntry.Text, 64)
		if err != nil {
			fmt.Println("Invalid amount")
			return
		}
		budget := &budgets[0] // Assuming working with the first budget for simplicity
		budget.Expenses = append(budget.Expenses, Expense{
			Name:   name,
			Amount: amount,
		})
		expenseList.Refresh()
	})

	saveBudgetButton := widget.NewButton("Save Budget", func() {
		budgetName := budgetNameEntry.Text
		income, err := strconv.ParseFloat(incomeEntry.Text, 64)
		if err != nil {
			fmt.Println("Invalid income")
			return
		}
		budget := Budget{
			Name:   budgetName,
			Income: income,
		}
		budgets = append(budgets, budget)
		fmt.Println("Budget saved:", budgetName)
		expenseList.Refresh()
	})

	// Layout the UI elements
	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Create Budget"),
		budgetNameEntry,
		incomeEntry,
		saveBudgetButton,
		widget.NewLabel("Add Expense"),
		expenseNameEntry,
		expenseAmountEntry,
		addExpenseButton,
		expenseList,
	))

	// Show the window
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}
