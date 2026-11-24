package cmdmanager

import (
	"fmt"
	"taxcalculator/filemanager"
)

type CmdManager struct {
	// Originally an empty struct as we're just using this to attach methods but added a filemanager as I want access to the filemanager. WriteLines() function
	fmWrite *filemanager.FileManager
	Prices  []float64
}

func New(fm *filemanager.FileManager) *CmdManager {

	return &CmdManager{
		fmWrite: fm,
	}
}

func (cmd *CmdManager) SetFileManager(fm *filemanager.FileManager) {
	cmd.fmWrite = fm
}

func (cmd *CmdManager) ReadLines() ([]float64, error) {

	fmt.Println("Enter Prices, 0 to stop Reading lines")
	var prices []float64

	if cmd.Prices != nil {
		return cmd.Prices, nil
	}


	for {
		var price float64
		_, err := fmt.Scanln(&price) // ScanLn returns a int for the number of items scanned and an error. It reads the cmd line entry to the variable price
		if err != nil {
			break
		}
		if price == 0 {
			break
		}

		prices = append(prices, price)
	}
	cmd.Prices = prices
	return prices, nil
}

func (cmd *CmdManager) WriteResult(data any) error {

	err := cmd.fmWrite.WriteResult(data)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil

}
