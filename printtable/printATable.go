package printtable

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
)


func PrintTable(records []map[string]float64) {

	if len(records) == 0 {
		fmt.Println("No records to print.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	columns := []string{"Price", "T1", "T2", "T3", "T4"}
	fmt.Fprintln(w, strings.Join(columns, "\t"))

	// 1. Collect all base price keys from the first record as numbers
	var pricesList []float64
	for priceStr := range records[0] {
		// Convert string keys (e.g. "10.00") to float64 for numeric sorting
		if p, err := strconv.ParseFloat(priceStr, 64); err == nil {
			pricesList = append(pricesList, p)
		}
	}

	// 2. Sort the base prices in ascending numeric order
	sort.Float64s(pricesList)

	// 3. Construct and print a row for each base price
	for _, p := range pricesList {
		priceStr := fmt.Sprintf("%.2f", p)
		row := []string{priceStr}

		// Pull the calculated price for this base price across all tax rate records
		for _, r := range records {
			row = append(row, fmt.Sprintf("%.2f", r[priceStr]))
		}

		// Write the tab-delimited row to the tabwriter
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}

	w.Flush()

}