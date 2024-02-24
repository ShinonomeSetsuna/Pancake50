// resource.go
// 使用embed储存静态文件

package maimage

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed font/adobe_simhei.otf
var otf []byte
var f, _ = opentype.Parse(otf)

/*
	//go:embed ui/diff_basic.png
	var basic []byte

	//go:embed ui/diff_advanced.png
	var advanced []byte

	//go:embed ui/diff_expert.png
	var expert []byte

	//go:embed ui/diff_master.png
	var master []byte

	//go:embed ui/diff_remaster.png
	var remaster []byte
*/

//go:embed ui/music_icon_a.png
var a []byte

//go:embed ui/music_icon_aa.png
var aa []byte

//go:embed ui/music_icon_aaa.png
var aaa []byte

//go:embed ui/music_icon_b.png
var b []byte

//go:embed ui/music_icon_bb.png
var bb []byte

//go:embed ui/music_icon_bbb.png
var bbb []byte

//go:embed ui/music_icon_c.png
var c []byte

//go:embed ui/music_icon_d.png
var d []byte

//go:embed ui/music_icon_fc.png
var fc []byte

//go:embed ui/music_icon_fcplus.png
var fcplus []byte

//go:embed ui/music_icon_ap.png
var ap []byte

//go:embed ui/music_icon_applus.png
var applus []byte

//go:embed ui/music_icon_fs.png
var fs []byte

//go:embed ui/music_icon_fsplus.png
var fsplus []byte

//go:embed ui/music_icon_fdx.png
var fdx []byte

//go:embed ui/music_icon_fdxplus.png
var fdxplus []byte

//go:embed ui/music_icon_s.png
var s []byte

//go:embed ui/music_icon_splus.png
var splus []byte

//go:embed ui/music_icon_ss.png
var ss []byte

//go:embed ui/music_icon_ssplus.png
var ssplus []byte

//go:embed ui/music_icon_sss.png
var sss []byte

//go:embed ui/music_icon_sssplus.png
var sssplus []byte

//go:embed ui/music_icon_sync.png
var sync []byte

//go:embed ui/music_icon_standard.png
var standard []byte

//go:embed ui/music_icon_dx.png
var dx []byte

// 以下为额外添加的素材

//go:embed ui/blank.png
var blank []byte

//go:embed ui/bar_basic.png
var bar_basic []byte

//go:embed ui/bar_advanced.png
var bar_advanced []byte

//go:embed ui/bar_expert.png
var bar_expert []byte

//go:embed ui/bar_master.png
var bar_master []byte

//go:embed ui/bar_remas.png
var bar_remas []byte

//go:embed ui/b35.png
var b35 []byte

//go:embed ui/bg.png
var bg []byte

var score_rank map[int][]byte = map[int][]byte{
	0:  d,
	1:  c,
	2:  b,
	3:  bb,
	4:  bbb,
	5:  a,
	6:  aa,
	7:  aaa,
	8:  s,
	9:  splus,
	10: ss,
	11: ssplus,
	12: sss,
	13: sssplus,
}

var combo_status map[int][]byte = map[int][]byte{
	1: fc,
	2: fcplus,
	3: ap,
	4: applus,
}

var sync_status map[int][]byte = map[int][]byte{
	1: fs,
	2: fsplus,
	3: fdx,
	4: fdxplus,
	5: sync,
}
