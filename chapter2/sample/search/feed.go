package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed 包含我们需要处理的数据源的信息
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

/* RetrieveFeeds 读取并反序列化源数据文件
   这个函数没有参数，会返回两个值。
   第一个返回值是一个切片，其中每一项指向一个 Feed 类型的值。
   第二个返回值是一个 error 类型的值，用来表 示函数是否调用成功。
 */
func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// 当函数返回时
	// 关闭文件
	defer file.Close()

	// 将文件解码到一个切片里
	// 这个切片的每一项是一个指向一个 Feed 类型值的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// 这个函数不需要检查错误，调用者会做这件事
	return feeds, err
}
