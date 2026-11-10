package main

import ("fmt"
)


func main(){

	// For every price calculate the price  with the tax rate

var prices []float64 = []float64{10,20,30}
taxRates := []float64{0,0.07,0.1,0.15}

// map of float keys and float values
/* In Go, the built-in make function is used exclusively to initialize and allocate slices, maps, and channels. Unlike new, which returns a pointer to zeroed memory, make allocates space on the heap, initializes the underlying data structure, and returns an initialized value (not a pointer) that is ready for immediate use.*/
result := make(map[float64][]float64)

/* In Go, the for ... range loop uses positional assignment based strictly on the data structure you are iterating over. Go’s compiler determines what variables represent entirely by their order—the first variable is always the index (or key), and the second variable is always the value. */


for index, val := range prices {
	fmt.Println(index)
	fmt.Println(val)
}

fmt.Println(taxRates)
fmt.Println(result)

for _, taxRate := range taxRates {
	// make a slice of capacity 3
taxIncludedPrices := make([]float64, len(prices))
// In for loops the 1st position is always the index and the 2nd the value
for index, price := range prices {
	// add the results to the slice
	taxIncludedPrices[index] = price * (1 + taxRate)
}
// add the results to the map
result[taxRate] = taxIncludedPrices
}

fmt.Println(result)




} // End Main

