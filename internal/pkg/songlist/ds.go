package songlist

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ShinonomeSetsuna/Pancake50/internal/types"
)

//go:embed json/ds.json
var jsonbytes []byte

var DS types.DSMap

func DS_load() {
	if file, err := os.Open("./ds.json"); err == nil {
		if fb, err := io.ReadAll(file); err == nil {
			fmt.Println("使用外部ds.json文件")
			jsonbytes = fb
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

	if err := json.Unmarshal(jsonbytes, &DS); err != nil {
		log.Fatal(err)
	}
}
