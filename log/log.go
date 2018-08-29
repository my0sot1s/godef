package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

// ⓐ ⓑ ⓒ ⓓ ⓔ ⓕ ⓖ ⓗ ⓘ ⓙ ⓚ ⓛ ⓜ ⓝ ⓞ ⓟ ⓠ ⓡ ⓢ ⓣ ⓤ ⓥ ⓦ ⓧ ⓨ ⓩ
// Ⓐ Ⓑ Ⓒ Ⓓ Ⓔ Ⓕ Ⓖ Ⓗ Ⓘ Ⓙ Ⓚ Ⓛ Ⓜ Ⓝ Ⓞ Ⓟ Ⓠ Ⓡ Ⓢ Ⓣ Ⓤ Ⓥ Ⓦ Ⓧ Ⓨ Ⓩ
const (
	DefaultLimit = 20
)

var (
	charCode  = []string{"Ⓐ", "Ⓑ", "Ⓒ", "Ⓓ", "Ⓔ", "Ⓕ", "Ⓖ", "Ⓗ", "Ⓘ", "Ⓙ", "Ⓚ", "Ⓛ", "Ⓜ", "Ⓝ", "Ⓞ", "Ⓟ", "Ⓠ", "Ⓡ", "Ⓢ", "Ⓣ", "Ⓤ", "Ⓥ", "Ⓦ", "Ⓧ", "Ⓨ", "Ⓩ"}
	asciiCode = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

func IndexOfSlide(slice []string, a string) int {
	for index, v := range slice {
		if v == a {
			return index
		}
	}
	return -1
}

func TitleConsole(name string) string {
	lower := strings.ToLower(name)
	var newStr string
	for _, char := range lower {
		if index := IndexOfSlide(asciiCode, string(char)); index > -1 {
			newStr += charCode[index]
		} else {
			newStr += string(char)
		}
	}
	return newStr
}
func Log(logs ...interface{}) {
	for _, log := range logs {
		fmt.Printf("%v", log)
	}
	fmt.Println()
}

func ErrLog(err error) (b bool) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", fn, line, err)
		b = true
	}
	return
}

func ReadFileRoot(path string) ([]byte, error) {
	if &path == nil || path == "" {
		return nil, errors.New("no path")
	}
	absPath, _ := filepath.Abs(path)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		ErrLog(err)
	}
	return data, err
}

func Jsonify(j interface{}) string {
	bytes, err := json.Marshal(j)
	ErrLog(err)
	return string(bytes[:])
}

func LogJson(j interface{}) {
	fmt.Println(Jsonify(j))
}
