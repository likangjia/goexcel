package test

import (
	"testing"

	"goexcel"
	"goexcel/excel"
)

/**
 * @Author: hubert
 * @Description:
 * @File: simple_excel_test
 * @Version: 1.0.0
 * @Date: 2022/7/15 11:53 上午
 */
var simpleExcelData = []map[string]interface{}{
	{
		"name":"小明",
		"class":"三年二班",
		"score":96,
		"major":"数学",
	},
	{
		"name":"小智",
		"class":"三年一班",
		"score":98,
		"major":"数学",
	},
}

func TestSimpleExcelCreate (t *testing.T) {
	path := "../temp/data_list_example.xlsx"
	savePath := "../temp"
	fileName := "data_list_example_out"
	ext := "xlsx"
	baseForm := goexcel.FileInfo{
		FilePath: path,
		SavePath: savePath,
		FileName: fileName,
		Ext:      ext,
	}

	simpleExcel := excel.NewExcelList(baseForm)
	err := simpleExcel.Create(simpleExcelData)
	if err != nil {
		println(err)
		return
	}
}
