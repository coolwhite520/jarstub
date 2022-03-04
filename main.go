package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"jarstub/utils"
	"os"
	"os/exec"
)
// PathExists 判断文件是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}


// 实际中应该用更好的变量名
var (
	encrypt bool
	run bool
	inFilePath string
	outFilePath string
)

func init() {
	flag.BoolVar(&encrypt, "encrypt", false, "is encrypt file?")
	flag.BoolVar(&run, "run", false, "is run?")
	flag.StringVar(&inFilePath, "in", "", "input file path")
	flag.StringVar(&outFilePath, "out", "", "output file path")
}

func main() {
	flag.Parse()
	//加密文件
	key := []byte{102, 117, 99, 107, 32, 105, 115, 32, 102, 117, 99, 107, 32, 97, 110, 100, 32, 102, 117, 99, 107, 46, 46, 46}
	if encrypt {
		//加密文件
		file, err := ioutil.ReadFile(inFilePath)
		if err != nil {
			panic(err)
		}
		//[]byte("fuck is fuck and fuck...")
		encrypt, err := utils.AesEncrypt(file, key)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(outFilePath, encrypt, os.ModePerm)
		if err != nil {
			panic(err)
		}
	} else if run {
		//解密文件并运行
		file, err := ioutil.ReadFile(inFilePath)
		if err != nil {
			panic(err)
		}
		encrypt, err := utils.AesDecrypt(file, key)
		if err != nil {
			panic(err)
		}
		//temp
		dir, err := ioutil.TempDir("", "jarstub")
		if err != nil {
			panic(err)
		}
		runPathFile := fmt.Sprintf("%s/app.jar", dir)
		err = ioutil.WriteFile(runPathFile, encrypt, os.ModePerm)
		if err != nil {
			panic(err)
		}
		if PathExists(runPathFile) {
			//go func() {
			//	time.Sleep(time.Second * 30)
			//	err = os.RemoveAll(dir)
			//	if err != nil {
			//		fmt.Println(err)
			//		return
			//	}
			//}()
			//cmdLine := fmt.Sprintf("java -jar %s --server.port=5001", runPathFile)
			command := exec.Command("java","-jar", runPathFile, "--server.port=5001")
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			err = command.Run()
			if err != nil {
				panic(err)
			}
		}
	} else {
		fmt.Println("命令行参数错误")
	}
}
