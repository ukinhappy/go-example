package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"math/rand"
)

func main() {

	file := excelize.NewFile()
	file.SetSheetName("Sheet1", "统计情况")
	streamWriter, _ := file.NewStreamWriter("统计情况")
	styleID, _ := file.NewStyle(`{"font":{"color":"#FF0000"}}`)
	streamWriter.SetRow("A1", []interface{}{excelize.Cell{StyleID: styleID, Value: "uid"}})
	streamWriter.SetRow("B1", []interface{}{excelize.Cell{StyleID: styleID, Value: "1"}})

	for rowID := 2; rowID <= 100; rowID++ {
		row := make([]interface{}, 2)
		for colID := 0; colID < 2; colID++ {
			row[colID] = rand.Intn(640000)
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		if err := streamWriter.SetRow(cell, row); err != nil {
			fmt.Println(err)
		}
	}
	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
	}

	if err := file.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
