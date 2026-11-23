package cmdmanager

import ("fmt"
"taxcalculator/filemanager")

type CmdManager struct {
	// empty struct as we're just using this to attach methods
	fmWrite *filemanager.FileManager
}

func New(fm *filemanager.FileManager) *CmdManager {

	return &CmdManager{
		fmWrite: fm,
	}
}

func (cmd *CmdManager) ReadLines() ([]string, error) {

	fmt.Println("Reading lines")
	var prices []string

	for {
		var price string
		_, err := fmt.Scanln(&price)
		if err != nil {
			break
		}
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
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