package cmdmanager

import (
	"fmt"
	//"taxcalculator/filemanager"
)

/*Instead of manually writing a wrapper method on CMDManager that forwards calls to fm.WriteResult(data), you can embed *filemanager.FileManager directly inside your CMDManager struct. Go will automatically promote FileManager's methods onto CMDManager. */

/*
type CmdManager struct {
	// Originally an empty struct as we're just using this to attach methods but added a filemanager as I want access to the filemanager. WriteLines() function
	*filemanager.FileManager // This is an embedded struct. It means I don't have to define WriteResult() on cmdmanager because *FileManager is embedded and cmd.WriteResult(data) is promoted automatically
	Prices  []float64
} */

type CmdManager struct {
	Prices []float64
}


func New() *CmdManager {

	return &CmdManager{}

}

/*
func (cmd *CmdManager) SetFileManager(fm *filemanager.FileManager) {
	cmd.fmWrite = fm
}
	*/

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
 // I can remove my definition of WriteResult as I use the WriteResult function from the embedded filemanager.FileManager
/*
func (cmd *CmdManager) WriteResult(data any) error {

	err := cmd.WriteResult(data)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil

}*/
