package utils
import (
	"os"
	"fmt"
	"time"
)
func CheckDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm);
		if err != nil {
			fmt.Println(err);
		}
	}
}
func GetDirectoryDate() string {
	year, month, day := time.Now().Date()
	directory_str := fmt.Sprintf("%v/%v/%v", year, month, day)
	return directory_str
}