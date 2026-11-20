package main

import (

	"taxcalculator/prices"
	"fmt"
	"taxcalculator/printtable"
)

func main() {

sampletaxRates := []float64{0, 0.07, 0.1, 0.15}

var records []map[string]float64

var record map[string]float64

var err error



for _, taxRate := range sampletaxRates {
	// Creates a new TaxIncludedPriceJob Struct
	prices := prices.NewTaxIncludedPriceJob(taxRate)
	// Calls the Process function on prices struct
	err, record = prices.Process()


if err != nil {
	fmt.Println(err)
}
records = append(records, record)
}

fmt.Println(records)

printtable.PrintTable(records)

for p := range records[0] {
	fmt.Println(p)
}

for key := range records {
	fmt.Println("Key Name" , records[key])
}

fmt.Println(records[1]["12.00"]) // The map value at row 2, key = "12.00"

fmt.Println(records[1]) // The map at row 2


} // end main