package uitl

import (
	"fmt"
	"io/ioutil"
)
type tools struct {

}
// 函数 开头必须 大写  要不 不能用

func RunSubdomains(confFilePath string) {
	fmt.Println("写入文件d")
}

func SaveFile(fileName string, txt string)  {
	name := fileName
	content := txt
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		fmt.Println("写入文件成功:", name)
	}
}