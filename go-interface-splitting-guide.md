# Splitting Interfaces & Decoupling I/O in Go

This guide explains how splitting a combined `IOManager` interface into single-method `LineReader` and `ResultWriter` interfaces decouples `CMDManager` from `FileManager` and simplifies `TaxIncludedPriceJob.Process()`.

---

## 1. Why Split the Interface?

In Go, small, single-method interfaces (like `io.Reader` and `io.Writer`) create flexible abstractions:

| Aspect | Single Combined Interface (`IOManager`) | Split Interfaces (`LineReader` & `ResultWriter`) |
| :--- | :--- | :--- |
| **Coupling** | `CMDManager` must import or embed `FileManager` just to satisfy `WriteResult()`. | `CMDManager` has zero knowledge of `FileManager` or files. |
| **Mix & Match** | Hard to read from CLI and write to JSON without wrapper structs. | Mix any reader (`cmd`, `fm`) with any writer (`fm`, mock, console). |
| **Unit Testing** | Mocks must implement both methods even when testing one path. | Test `ReadLines()` and `WriteResult()` with tiny single-method mocks. |

---

## 2. Defining the Split Interfaces in `prices`

Interfaces in Go belong in the **consumer package** (`prices`) that calls them. Store interface fields by **value**—never as pointers to interfaces (`*LineReader`).

```go
package prices

type LineReader interface {
	ReadLines() ([]string, error)
}

type ResultWriter interface {
	WriteResult(data any) error
}

type TaxIncludedPriceJob struct {
	TaxRate float64
	Reader  LineReader
	Writer  ResultWriter
}

func NewTaxIncludedPriceJob(taxRate float64, r LineReader, w ResultWriter) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		Reader:  r,
		Writer:  w,
	}
}
```

---

## 3. Connecting `CMDManager` (Implicit Satisfaction)

Go interfaces are satisfied **implicitly**. `CMDManager` does not need an `implements` keyword, does not import `prices`, and does not embed `*FileManager`. Defining `ReadLines() ([]string, error)` is enough for `*CMDManager` to satisfy `prices.LineReader`:

```go
package cmdmanager

import "fmt"

type CMDManager struct{}

func New() *CMDManager {
	return &CMDManager{}
}

// Implicitly satisfies prices.LineReader
func (cmd *CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Enter prices (confirm each with ENTER, type 0 to finish):")
	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scanln(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}
```

---

## 4. Working Inside `Process()` & Wiring in `main.go`

### Inside `TaxIncludedPriceJob.Process()`

`Process()` calls `job.Reader.ReadLines()` and `job.Writer.WriteResult(...)` without knowing whether the underlying implementations use disk files, terminal input, or test mocks:

```go
func (job *TaxIncludedPriceJob) Process() (TaxIncludedPriceRecord, error) {
	lines, err := job.Reader.ReadLines()
	if err != nil {
		return TaxIncludedPriceRecord{}, err
	}

	// Parse lines & calculate tax-included prices...
	result := make(map[string]string)
	// ...

	record := TaxIncludedPriceRecord{
		TaxRate:           job.TaxRate,
		TaxIncludedPrices: result,
	}

	if err := job.Writer.WriteResult(record); err != nil {
		return TaxIncludedPriceRecord{}, err
	}

	return record, nil
}
```

### Wiring Dependencies in `main.go`

```go
cmd := cmdmanager.New()

for _, taxRate := range taxRates.TaxRates {
	fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

	// Read from CLI (cmd), write JSON to disk (fm)
	priceJob := prices.NewTaxIncludedPriceJob(taxRate, cmd, fm)

	record, err := priceJob.Process()
	if err != nil {
		fmt.Println(err)
		continue
	}
	records = append(records, record)
}
```
