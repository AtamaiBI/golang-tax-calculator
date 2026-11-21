package filemanager

import (
	"bufio"
	"os"
	"strconv"
	"errors"
	"encoding/json"

)

type FileManager struct {
InputFilePath string
OutputFilePath string
}



// Receives a string filepath and returns a slice of floats or and error
func (fm *FileManager) ReadLines() ([]float64, error) {

file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("Failed to open file")
	}
	defer file.Close()

	var prices []float64

	scanner := bufio.NewScanner(file)
	// scanner.scan() returns a bool and the for loop iterates through the list of values until the end of the file which then returns false to exit the for loop
	for scanner.Scan() {
		price, err := strconv.ParseFloat(scanner.Text(), 64) // 64 is the bitSize
		if err != nil {
			return nil, err
		}
		prices = append(prices, price)
	} // end loop

	// Loop exited because scanner.Scan() == false
	// Now check if it stopped due to an error or a normal EOF
	if err := scanner.Err(); err != nil {

		return nil, err
	}

	return prices, nil
}


/*
Why interface{} Works Here

• Zero Method Requirements: In Go, an interface defines a set of methods a type must implement. Because interface{} specifies zero methods, every Go type automatically satisfies it.
• Single Reusable Function: When building APIs or serialization helpers, you want one WriteJSON function to handle a User struct, a []Product slice, or a map[string]string error message. interface{} makes that flexibility possible.
• Runtime Reflection: Under the hood, encoding/json (like json.Marshal(data)) uses Go's reflect package to inspect the underlying type and fields of data at runtime and convert them into JSON text.
*/

// 
func (fm *FileManager) WriteResult(data any) error {

	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	/* Yes, this is Go's "if with a short statement" syntax, which lets you run a quick setup statement (like declaring a variable) right before evaluating the if condition.

How It Works

The syntax has two parts separated by a semicolon (;):
go
if [initialization statement]; [boolean condition] {// runs if condition is true
}
*/
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil

}

func New(inputFilePath string, outputFilePath string) *FileManager {
	return &FileManager{
		InputFilePath: inputFilePath,
		OutputFilePath: outputFilePath,
	}
}