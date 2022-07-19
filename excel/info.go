package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"

	"goexcel"
)

/**
 * @Author: likangjia
 * @Description:
 * @File: simple_info
 * @Version: 1.0.0
 * @Date: 2022/7/18 10:03 上午
 */

type Info struct{
	goexcel.FileInfo
}

func NewExcelInfo(info goexcel.FileInfo) Info{
	return Info{
		info,
	}
}

// Create 获取个人数据表单
func (i Info)Create(data map[string]interface{}) error{
	// 读取文件
	f, err := excelize.OpenFile(i.FilePath)
	if err != nil {
		fmt.Println("err is", err)
		return err
	}
	// 获取行数
	rows, err := f.GetRows("Sheet1")
	if err != nil{
		fmt.Println("err is", err)
		return err
	}
	fmt.Println(len(rows[1]))

	return nil
}
