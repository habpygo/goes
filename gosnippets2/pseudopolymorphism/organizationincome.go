package main

import "fmt"

// The Income interface defined above contains two methods calculate()
// which calculates and returns the income from the source and source()
// which returns the name of the source.
type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

// Next step would be to define methods on these struct types which
// calculate and return the actual income and source of income.

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

/*
In the case of FixedBilling projects, the income is the just the amount bid for the project.
Hence we return this from the calculate() method of FixedBilling type.

In the case of TimeAndMaterial projects, the income is the product of the noOfHours and hourlyRate.
This value is returned from the calculate() method with receiver type TimeAndMaterial.

We return the name of the project as source of income from the source() method.

Since both  FixedBilling and TimeAndMaterial structs provide definitions for the calculate()
and source() methods of the Income interface, both structs implement the Income interface.
*/

func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income from %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netincome)
}

/*
The calculateNetIncome function above accepts a slice of Income interfaces as argument.
It calculates the total income by iterating over the slice and calling calculate() method on each of its items.
It also displays the income source by calling source() method.
Depending on the concrete type of the Income interface, different calculate() and source() methods will be called.
Thus we have achieved polymorphism in the  calculateNetIncome function.
*/

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
}
