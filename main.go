package main

import (
	"fmt"
	"taxcalculator/cmdmanager"
	"taxcalculator/filemanager"
	"taxcalculator/prices"
	"taxcalculator/printtable"
	"taxcalculator/taxrates"
)

func main() {

	// sampletaxRates := []float64{0, 0.07, 0.1, 0.2}

	var records []map[string]float64

	var record map[string]float64

	var err error

	cmd := cmdmanager.New(nil)
	taxRates := taxrates.NewTaxRate()
	taxRates, err = taxRates.ReadTaxRates("taxrates.txt")

	if err != nil {
		fmt.Println(err)
		return
	}


	for _, taxRate := range taxRates.TaxRates {
		// Filemanager.New() returns a pointer to a filemanager struct therefore in the prices constructor I must dereference fm as prices is expecting a copy not a pointer. I can change the prices constructor to accept a pointer, which I have done.
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmd.SetFileManager(fm)

		// Creates a new TaxIncludedPriceJob Struct that accepts a taxrate value and a filemanager struct. This allows you to access the filemanager functions from the prices struct.
		prices := prices.NewTaxIncludedPriceJob(taxRate, cmd)
		// Calls the Process function on prices struct
		err, record = prices.Process()

		if err != nil {
			fmt.Println(err)
		}
		records = append(records, record)
	}

	fmt.Println(records)

	printtable.PrintTable(records)

	/*
	for p := range records[0] {
		fmt.Println(p)
	} */

	for key := range records {
		fmt.Println("Key Name", records[key])
	}
	/*
	   	fmt.Println(records[1]["12.00"]) // The map value at row 2, key = "12.00"

	   	fmt.Println(records[1]) // The map at row 2

	   	fm := filemanager.New("", "result.json") // creates a new filemanager struct instance
	   	cmd := cmdmanager.New(fm)                // creates a new cmdmanager struct instance. This allows cmd to access the fm.WriteResult(data) function to write to a json file as specified by the fm output path.

	   	data, e := cmd.ReadLines()

	   	if e != nil {
	   		fmt.Println(e)
	   	}

	   	fmt.Println(data)
	   //	cmd.WriteResult(data)
	*/

} // end main
