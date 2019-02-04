package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// PrintOutput prints data
func PrintOutput(v interface{}, isJSON bool) {
	if isJSON {
		b, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			log.Fatalf("failed to marshal json `%+v`: %s", v, err)
		}
		os.Stdout.Write(b)
		os.Stdout.WriteString("\n")
		return
	}
	switch vv := v.(type) {
	case []string:
		for _, s := range vv {
			fmt.Println(s)
		}
	}
}
