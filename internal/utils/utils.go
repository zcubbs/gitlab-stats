package utils

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint print the contents of the struct
func PrettyPrint(data interface{}) {
	var p []byte

	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
