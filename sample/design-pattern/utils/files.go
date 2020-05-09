package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// ReadData - reads file pointed path and return file info as map
// Attention - We need to build main file by go build command to use this function
func ReadData(path string) (map[string]int, error) {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)

	fmt.Println(ex)            // /Users/hono/go/src/github.com/ono5/money-boy/sample/design-pattern/singleton/singleton
	fmt.Println(exPath + path) // /Users/hono/go/src/github.com/ono5/money-boy/sample/design-pattern/singleton/capitals.txt

	file, err := os.Open(exPath + path)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}
	return result, nil
}
