package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	//"github.com/tealeg/xlsx"
	"github.com/360EntSecGroup-Skylar/excelize"
	//"github.com/extrame/xls"
)

var execPath string
var outputDir string
var inputDir string = "input"

func init() {
	execPath = filepath.Dir(os.Args[0])

	now := time.Now()
	_, month, day := now.Date()
	hour, min, sec := now.Clock()
	outputDir = "xlsx-folder" + fmt.Sprintf("-%02d%02d-%02d%02d%02d", month, day, hour, min, sec)
	outputDir = filepath.Join(execPath, outputDir)
	err := os.MkdirAll(outputDir, 0777)
	if err != nil {
		panic(err)
	}

	inputDir := filepath.Join(execPath, inputDir)
	err = os.MkdirAll(inputDir, 0777)
	if err != nil {
		panic(err)
	}
}

func main() {
	files, _ := ioutil.ReadDir(inputDir)
	for _, onefile := range files {
		if onefile.IsDir() {
			continue
		}
		fileSuffix := path.Ext(onefile.Name())
		if strings.ToLower(fileSuffix) != ".xlsx" {
			continue
		}

		fullFromName := filepath.Join(inputDir, onefile.Name())
		fmt.Println("-- fullFromName: " + fullFromName)

		// toName := toDir + "\\" + onefile.Name() + "x"
		// fullToName := filepath.Join(outputDir, toName)
		// fmt.Println("-- --fullToName: " + fullToName)
		// fmt.Println("      ")

		xfile, err := excelize.OpenFile(fullFromName)
		if err != nil {
			panic(err)
		}

		cell := xfile.GetCellValue("Sheet1", "B1")
		if err != nil {
			panic(err)
		}
		fmt.Println(cell)

		rows := xfile.GetRows("Sheet1")
		for _, row := range rows {
			for _, colCell := range row {
				fmt.Print(colCell, "\t")
			}
			fmt.Println()
		}

		//文件另存为xlsx扩展名的文件
		// err = xfile.SaveAs(fullToName)
		// if err != nil {
		// 	panic(err)
		// }

	}
}
