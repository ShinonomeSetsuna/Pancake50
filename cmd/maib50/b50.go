package main

import (
	"fmt"
	"os"

	"github.com/ShinonomeSetsuna/Pancake50/internal/config"
	"github.com/ShinonomeSetsuna/Pancake50/internal/pkg/maimage"
	"github.com/ShinonomeSetsuna/Pancake50/internal/pkg/songlist"
)

func main() {
	rating := songlist.GetRating()
	maimage.DrawAll(rating.Data.Rating_list, rating.Data.New_rating_list)
}

func init() {
	// Init
	songlist.DS_load()
	for cnt, arg := range os.Args {
		fmt.Println(cnt, arg)
		if arg == "--detail" {
			config.QuickMode = false
		}
	}
}
