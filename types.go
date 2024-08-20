package main

type Expense struct {
	Name   string
	Amount float64
}

type Budget struct {
	Name     string
	Income   float64
	Expenses []Expense
}

var budgets []Budget
