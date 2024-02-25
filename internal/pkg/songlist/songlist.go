package songlist

import (
	"encoding/json"

	"github.com/ShinonomeSetsuna/Pancake50/internal/config"
	"github.com/ShinonomeSetsuna/Pancake50/internal/tools"
	"github.com/ShinonomeSetsuna/Pancake50/internal/types"
)

/*直接获取B50*/
func GetRating() types.Rating {
	var rating types.Rating
	jsonbytes := tools.GetResource("https://v2.otogame.net/api/game/maimai/rating")
	json.Unmarshal(jsonbytes, &rating)
	return rating
}

/*获取乐曲详细成绩*/
func GetRecord(music_id string) types.Record {
	var record types.Record
	if config.QuickMode {
		return record
	}
	jsonbytes := tools.GetResource("https://v2.otogame.net/api/game/maimai/record/" + music_id)
	json.Unmarshal(jsonbytes, &record)
	return record
}
