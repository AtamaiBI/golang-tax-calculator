package prices

import (
	//"bufio"
	"fmt"
	"taxcalculator/conversion"
	//"taxcalculator/filemanager"
	"taxcalculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64 `json:"tax_rate"`
	InputPrices       []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]float64 `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"` // "-" means ignore this key and value in the output

}

// Add a constructor function for the struct. Returns a pointer to a struct.
func NewTaxIncludedPriceJob(taxRate float64, io iomanager.IOManager) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		//InputPrices:       []float64{10, 20, 30}, // will read values in InputPrices later
		InputPrices:       []float64{},
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]float64),
		IOManager:          io,
	}
}




// Receiver functon iterates over the struct and prints the result map of prices with tax. Must have a pointer parameter as the constructor returns a pointer and we want to work with that rather than a copy
func (job *TaxIncludedPriceJob) Process() (error, map[string]float64) {

	err := job.LoadData()

	if err != nil {
		fmt.Println(err)
		return err, nil
	}

	result := make(map[string]float64)

	//	taxIncludedPrices := make([]float64, len(job.InputPrices))
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = conversion.RoundToTwo(taxIncludedPrice)
	}

	// Special if statement. The first part is the expression and the second part after ; is the test
	if err := job.IOManager.WriteResult(result); err != nil {
		return err, nil
	}

	// fmt.Println(result) // prints a map of prices with tax for a single tax rates
	return nil, result

} // ------------ END FUNC ---------------------------------------

// receiver function on a pointer to add a list of prices to the struct. If this receiver func wasn't a pointer then the update to its fields would be on a copy of the struct
func (job *TaxIncludedPriceJob) LoadData() error {

var err error
	job.InputPrices, err =  job.IOManager.ReadLines()// adds a slice of prices to the TaxIncludedPriceJob struct which is a pointer. Because IOManager (a filemanager type)
	// If successful return nil as not error occurred
if err != nil {

	fmt.Println("Error in file reader")
}

	return nil
}

type Record struct {
	ID int
	Name string

}


