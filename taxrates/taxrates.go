package taxrates

import (
	"os"
	"bufio"
	"strconv"

)

type TaxRate struct {
	TaxRates []float64 `json:"tax_rates"`
}


func NewTaxRate() *TaxRate {
	return &TaxRate{}
}


func (tr *TaxRate) ReadTaxRates(path string) (*TaxRate, error) {

file, err := os.Open(path)

if err != nil {
	return nil, err
}

defer file.Close()

var rates []float64

scanner := bufio.NewScanner(file)

for scanner.Scan() {
	rate, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil { // deals with the parse error
		return nil, err
	}
	rates = append(rates, rate)
}

if err := scanner.Err(); err != nil { // deals with the scanner error
	return nil, err
}

tr.TaxRates = rates
	return tr, nil

}
