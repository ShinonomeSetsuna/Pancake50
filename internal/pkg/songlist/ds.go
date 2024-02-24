package songlist

import (
	_ "embed"
	"encoding/json"
	"log"

	"github.com/ShinonomeSetsuna/Pancake50/internal/types"
)

//go:embed json/ds.json
var jsonbytes []byte

var DS types.DSMap

func DS_load() {
	err := json.Unmarshal(jsonbytes, &DS)
	if err != nil {
		log.Fatal(err)
	}
}
