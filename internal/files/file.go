package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ReadFileContent(n string) string {
	path, _ := filepath.Abs(fmt.Sprintf("../resources/%s.json", n))
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	return string(byteValue)

}
