package control

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Scan interface {
	urlScan(url string) *http.Response
	payloads(url string, fileName string)
}

type target struct {
	url        string
	dictionary string
	req        *http.Request
	code       *http.Response
	err        error
}

func (t *target) urlScan(url string) *http.Response {
	client := &http.Client{}
	t.req, t.err = http.NewRequest("Get", url, nil)
	t.req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")
	t.req.ProtoMinor = 0
	if t.err != nil {
		log.Println("The error is:", t.err)
	}
	t.code, _ = client.Do(t.req)
	defer t.code.Body.Close()
	return t.code
}

func (t *target) payloads(url string, fileName string) {
	fileName = "../../Go_Scan/dictionary/" + fileName
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		t.code = t.urlScan(t.url + line)
		if t.code.StatusCode != 200 {
			fmt.Printf("[%c[1;37;31m%d%c[0m]  ", 0x1B, t.code.StatusCode, 0x1B)
			fmt.Printf("[%c[1;37;31m%s%c[0m]  ", 0x1B, "Warn", 0x1B)
			log.Println(t.url + line)
			Log := "[Warn]" + strconv.Itoa(t.code.StatusCode) + t.url + line
			logSave(Log)
			fmt.Printf("--------------------------------------------------------------------------------\n")
		} else {
			fmt.Printf("[%c[1;37;33m%d%c[0m]  ", 0x1B, t.code.StatusCode, 0x1B)
			fmt.Printf("[%c[1;37;33m%s%c[0m]  ", 0x1B, "Info", 0x1B)
			log.Println(t.url + line)
			Log := "[Info]" + strconv.Itoa(t.code.StatusCode) + t.url + line
			logSave(Log)
			fmt.Printf("--------------------------------------------------------------------------------\n")
		}
	}
}
