package goexcel

/**
 * @Author: likangjia
 * @Description:
 * @File: excel
 * @Version: 1.0.0
 * @Date: 2022/7/15 11:07 上午
 */

type FileInfo struct {
	FilePath string  // 文件路径
	SavePath string	 // 保存路径
	FileName string	 // 文件名
	Ext      string  // 后缀
}

type Offices interface{
	Create(data []map[string]interface{}) error
}
