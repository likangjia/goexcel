package excel

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"

	"goexcel"
	"goexcel/types"
	"goexcel/util"
)

/**
 * @Author: likangjia
 * @Description:
 * @File: simple_list
 * @Version: 1.0.0
 * @Date: 2022/7/15 11:08 上午
 */

type List struct {
	goexcel.FileInfo
}

func NewExcelList(base goexcel.FileInfo) List {
	return List{
		base,
	}
}

func (s List) Create(data []map[string]interface{}) error {
	// 读取文件
	f, err := excelize.OpenFile(s.FilePath)
	if err != nil {
		fmt.Println("err is", err)
		return err
	}
	// 获取行数
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println("err is", err)
		return err
	}

	// 读取占位符号
	if len(rows) < 2 {
		return types.ExcelReadVariableErr
	}

	// 获取占位符key
	tempKey := util.GetReplaceKey(rows[1])
	fmt.Println("temp key", tempKey)
	// 遍历填充数据
	for cloNum, datum := range data {
		var column = make([]interface{}, len(tempKey))
		for k, v := range tempKey {
			if inputData, ok := datum[v]; ok {
				column[k] = inputData
			} else {
				column[k] = ""
			}
		}
		startPoint := strconv.Itoa(cloNum + 2)
		fmt.Println("column", column)
		// 逐行增加数据
		err = f.SetSheetRow("Sheet1", "A"+startPoint, &column)
		if err != nil {
			fmt.Println(err)
		}
	}
	if err = f.SaveAs(s.getSavePath()); err != nil {
		fmt.Println(err)
	}
	// 保存文件
	return nil
}

func (s List) getSavePath() string {
	return fmt.Sprintf("%s/%s.%s", s.SavePath, s.FileName, s.Ext)
}
