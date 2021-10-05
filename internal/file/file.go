package file

import (
	"encoding/json"
	"fmt"

	"github.com/romanzh1/golang-parser-xml/pkg/parse"
)

func PrintJSONData(path string) {
	projects := parse.FromXML(path)
	projectJSON, err := json.MarshalIndent(projects, "	", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(projectJSON))
}
