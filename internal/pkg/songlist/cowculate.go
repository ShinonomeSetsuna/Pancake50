// 储存定数的文件
package songlist

import (
	"fmt"
	"math"

	"github.com/ShinonomeSetsuna/Pancake50/internal/types"
)

/*
return 定数 rating
*/
func GetSongRating(song types.Song) (string, string) {
	var ds float64
	var level types.DS

	if song.Music.Is_deluxe {
		level = DS[song.Music.Name].DX
	} else {
		level = DS[song.Music.Name].SD
	}

	switch song.Music.Level_info.Difficulty {
	case 0:
		ds = level.Basic
	case 1:
		ds = level.Advanced
	case 2:
		ds = level.Expert
	case 3:
		ds = level.Master
	case 4:
		ds = level.ReMaster
	}

	return fmt.Sprintf("%.1f", ds), fmt.Sprintf("%.0f", Calculate(song.Achievement, ds))
}

func Calculate(score int, ds float64) float64 {
	score = min(score, 1005000)
	factor := 0.0
	if score < 100000 {
		factor = 0
	} else if score < 200000 {
		factor = 1.6
	} else if score < 300000 {
		factor = 3.2
	} else if score < 400000 {
		factor = 4.8
	} else if score < 500000 {
		factor = 6.4
	} else if score < 600000 {
		factor = 8.0
	} else if score < 700000 {
		factor = 9.6
	} else if score < 750000 {
		factor = 11.2
	} else if score < 800000 {
		factor = 12.0
	} else if score < 900000 {
		factor = 12.8
	} else if score < 940000 {
		factor = 13.6
	} else if score < 970000 {
		factor = 16.8
	} else if score < 980000 {
		factor = 20.0
	} else if score < 990000 {
		factor = 20.3
	} else if score < 995000 {
		factor = 20.8
	} else if score < 1000000 {
		factor = 21.1
	} else if score < 1005000 {
		factor = 21.6
	} else {
		factor = 22.4
	}
	return math.Floor(ds * factor * float64(score) / 1000000)
}
