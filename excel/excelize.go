package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"math/rand"
)

func main() {

	file := excelize.NewFile()
	streamWriter, _ := file.NewStreamWriter("Sheet1")
	styleID, _ := file.NewStyle(`{"font":{"color":"#FF0000"}}`)
	streamWriter.SetRow("A1", []interface{}{excelize.Cell{StyleID: styleID, Value: "用户id"}})
	streamWriter.SetRow("B1", []interface{}{excelize.Cell{StyleID: styleID, Value: "用户昵称"}})
	streamWriter.SetRow("C1", []interface{}{excelize.Cell{StyleID: styleID, Value: "用户头像"}})
	streamWriter.SetRow("D1", []interface{}{excelize.Cell{StyleID: styleID, Value: "用户地址"}})
	streamWriter.SetRow("E1", []interface{}{excelize.Cell{StyleID: styleID, Value: "用户大家"}})
	streamWriter.SetRow("F1", []interface{}{excelize.Cell{StyleID: styleID, Value: "用户111"}})
	streamWriter.SetRow("G1", []interface{}{excelize.Cell{StyleID: styleID, Value: "用户222"}})

	for rowID := 2; rowID <= 100; rowID++ {
		row := make([]interface{}, 7)
		for colID := 0; colID < 7; colID++ {
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