package control

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

func logSave(strContent string) {
	//need to check whether the log folder exists in the project
	exists, _ := pathExists("../../Go_Scan/log/")
	if exists == true {
		filename := "../../Go_Scan/log/" + time.Now().Format("2006-01-02") + ".log"
		fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		fdTime := time.Now().Format("2006-01-02 15:04:05")
		fdContent := strings.Join([]string{"======", fdTime, "=====", strContent, "\n"}, "")
		buf := []byte(fdContent)
		fd.Write(buf)
		fd.Close()
	} else {
		folderName := "log"
		folderPath := "../../GO_Scan/"
		folderPath = filepath.Join(folderPath, folderName)
		os.Mkdir(folderPath, 0777)
		os.Chmod(folderPath, 0777)
		filename := "../../Go_Scan/log/" + time.Now().Format("2006-01-02") + ".log"
		fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		fdTime := time.Now().Format("2006-01-02 15:04:05")
		fdContent := strings.Join([]string{"======", fdTime, "=====", strContent, "\n"}, "")
		buf := []byte(fdContent)
		fd.Write(buf)
		fd.Close()
	}
}
