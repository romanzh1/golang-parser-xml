package file

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/romanzh1/golang-parser-xml/pkg/parse"
)

func PrintJSONData(path string) {
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + xmlFile.Name())
	defer xmlFile.Close()

	projects := parse.FromXML(xmlFile)
	projectJSON, err := json.MarshalIndent(projects, "	", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(projectJSON))
}
