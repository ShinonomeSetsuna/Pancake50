// 储存定数的文件
package songlist

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ShinonomeSetsuna/Pancake50/internal/types"
)

var levelHash = map[int]string{
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "7+",
	9:  "8",
	10: "8+",
	11: "9",
	12: "9+",
	13: "10",
	14: "10+",
	15: "11",
	16: "11+",
	17: "12",
	18: "12+",
	19: "13",
	20: "13+",
	21: "14",
	22: "14+",
	23: "15",
}

/*
return 定数 rating
*/
func GetSongRating(song types.Song) (string, string) {
	var cur float64 = 0
	var ds types.DS

	if song.Music.Is_deluxe {
		ds = DS[song.Music.Name].DX
	} else {
		ds = DS[song.Music.Name].SD
	}

	switch song.Music.Level_info.Difficulty {
	case 0:
		cur = ds.Basic
	case 1:
		cur = ds.Advanced
	case 2:
		cur = ds.Expert
	case 3:
		cur = ds.Master
	case 4:
		cur = ds.ReMaster
	}
	if cur == 0 {
		level := levelHash[song.Music.Level_info.Level]
		level = strings.Replace(level, "+", ".7", -1)
		cur, _ = strconv.ParseFloat(level, 64)
		return levelHash[song.Music.Level_info.Level], fmt.Sprintf("%.0f+", Calculate(song.Achievement, cur))
	}
	return fmt.Sprintf("%.1f", cur), fmt.Sprintf("%.0f", Calculate(song.Achievement, cur))
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
