package file

import (
	"io/ioutil"
	"log"
	"os"
)

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//创建文件
func CreateFile(filename string) (bool, error) {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return true,nil
	}

	return false,nil
}

func CreateModelFile(path string, filename string, ormType string, content string) bool {
	//判断文件夹是否存在
	exist,err := PathExists(path)

	if err != nil {
		log.Fatal(err)
	}

	if !exist {
		//创建文件夹
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	filename = path + "/" + filename
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write([]byte(content))

	if err != nil {
		return false
	}

	return true

}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}
