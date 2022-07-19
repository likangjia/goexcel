package excel

import (
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

/**
 * @Author: likangjia
 * @Description:
 * @File: excel
 * @Version: 1.0.0
 * @Date: 2022/7/18 10:48 上午
 */

type ExcelFile struct {
	f *excelize.File `json:"f"`
}

func OpenFile(filename string, opt ...excelize.Options) (*ExcelFile, error) {
	var excel ExcelFile
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return nil, err
	}
	excel.f, err = excelize.OpenReader(file, opt...)
	if err != nil {
		closeErr := file.Close()
		if closeErr == nil {
			return &excel, err
		}
		return &excel, closeErr
	}
	excel.f.Path = filename
	return &excel, file.Close()
}

func (e *ExcelFile)GetRows(sheet string, opts ...excelize.Options)([][]string, error){
	rows, err := e.f.Rows(sheet)
	if err != nil {
		return nil, err
	}
	results, cur, max := make([][]string, 0, 64), 0, 0
	for rows.Next() {
		cur++
		row, err := rows.Columns(opts...)
		if err != nil {
			break
		}
		results = append(results, row)
		if len(row) > 0 {
			max = cur
		}
	}
	return results[:max], rows.Close()

}