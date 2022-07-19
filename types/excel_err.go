package types

import (
	"errors"
)

/**
 * @Author: likangjia
 * @Description:
 * @File: excel_err
 * @Version: 1.0.0
 * @Date: 2022/7/15 11:20 上午
 */

var (
	ExcelReadVariableErr = errors.New("excel read  variable failed")
	ExcelCreateFailed = errors.New("excel create  failed")
	ExcelReadFailed = errors.New("excel read  failed")
)