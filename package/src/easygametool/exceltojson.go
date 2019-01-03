package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/labstack/gommon/log"
	"os"
	"encoding/json"
)

// 目前只能是 xlsx
// 强类型 需要定义 解析的结构 

type xlsData struct {
	fileName string
	English  string
	Russion  string
	Vietnam  string
	French   string
	Germany  string
}

type sheetData struct {
	sheet0 [5000]xlsData
}

const (
	fileName = iota
	English
	Russion
	Vietnam
	French
	Germany
)

var name = [6]string{"fileName", "English", "Russion", "Vietnam", "French", "Germany"}

func main() {
	xlsF := "/Users/admin/Documents/ljworkspace/local/gotest/package/resources/translateImgData.xlsx"
	var dd = getXlsDatabyFileName(xlsF)
	sheetDataToJson(dd)
	fmt.Printf("open failed: %d\n", len(dd.sheet0))
	fmt.Println("main")
}

func getXlsDatabyFileName(fName string) sheetData {
	xlFile, err := xlsx.OpenFile(fName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	var rd sheetData
	for _, sheet := range xlFile.Sheets {
		//sheet.Rows
		var res [5000]xlsData
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for roIdx, row := range sheet.Rows {
			var data xlsData
			for idx, cell := range row.Cells {
				text := cell.String()
				if idx < len(res) {
					switch idx {
					case fileName:
						data.fileName = text
					case English:
						data.English = text
					case Russion:
						data.Russion = text
					case Vietnam:
						data.Vietnam = text
					case French:
						data.French = text
					case Germany:
						data.Germany = text
					default:
						fmt.Printf("not type data: %s\n", text)
					}
				}
				fmt.Printf("%s\n", text)
			}
			res[roIdx] = data
		}
		rd.sheet0 = res
	}
	return rd

}

func sheetDataToJson(d sheetData) {
	//v interface{}  这样 就可以 简单 忽略类型 ？？
	data, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)

	fp, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}

}
