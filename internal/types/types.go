package types

type LevelInfo struct {
	Difficulty int `json:"difficulty"`
	Level      int `json:"level"`
}

type Music struct {
	Name      string `json:"name"`
	Artist    string `json:"artist"`
	Music_id  string `json:"music_id"`
	Is_deluxe bool   `json:"is_deluxe"`
	Jacket    string `json:"jacket"`
}

type Song struct {
	Achievement int       `json:"achievement"`
	Music       Music     `json:"music"`
	Level_info  LevelInfo `json:"level_info"`
}

type SongData struct {
	New_rating_list      []Song `json:"new_rating_list"`
	Rating_list          []Song `json:"rating_list"`
	Next_new_rating_list []Song `json:"next_new_rating_list"`
	Next_rating_list     []Song `json:"next_rating_list"`
}

type Rating struct {
	Code      string   `json:"code"`
	Data      SongData `json:"data"`
	Message   string   `json:"message"`
	Timestamp int      `json:"timestamp"`
}

type MusicDetail struct {
	Achievement      int `json:"achievement"`
	Ap_plus_count    int `json:"ap_plus_count"`
	Combo_status     int `json:"combo_status"`
	Deluxe_score_max int `json:"deluxe_score_max"`
	Difficulty       int `json:"difficulty"`
	Play_count       int `json:"play_count"`
	Score_rank       int `json:"score_rank"`
	Sync_status      int `json:"sync_status"`
}

type LevelData struct {
	Level_info   []LevelInfo   `json:"level_info"`
	Music_detail []MusicDetail `json:"music_detail"`
}

type Record struct {
	Code      string    `json:"code"`
	Data      LevelData `json:"data"`
	Message   string    `json:"message"`
	Timestamp int       `json:"timestamp"`
}

type NetRespone struct {
	Code      string      `json:"code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	Timestamp int         `json:"timestamp"`
}

type DS struct {
	Basic    float64 `json:"Basic"`
	Advanced float64 `json:"Advanced"`
	Expert   float64 `json:"Expert"`
	Master   float64 `json:"Master"`
	ReMaster float64 `json:"ReMaster"`
}

type DSTuple struct {
	DX DS `json:"DX"`
	SD DS `json:"SD"`
}

type DSMap map[string]DSTuple
