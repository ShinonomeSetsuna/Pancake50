// 绘制单张成绩

package maimage

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/ShinonomeSetsuna/Pancake50/internal/pkg/songlist"
	"github.com/ShinonomeSetsuna/Pancake50/internal/tools"
	"github.com/ShinonomeSetsuna/Pancake50/internal/types"
	"github.com/fishtailstudio/imgo"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"golang.org/x/image/webp"
)

/*绘制封面*/
func DrawCover(song types.Song) image.Image {
	cover, err := webp.Decode(bytes.NewBuffer(
		tools.GetResource("https://oss.bemanicn.com/SDEZ/cover/" + song.Music.Jacket + ".webp")))
	if err != nil {
		log.Fatalln("封面解码时出现错误：", err)
	}
	return cover
}

/*添加封面*/
func addCover(background *image.RGBA, song types.Song) {
	cover := image.NewRGBA(image.Rect(0, 0, 400, 400))
	draw.Draw(cover, cover.Rect, DrawCover(song), image.Point{}, draw.Src)
	sub := image.NewRGBA(image.Rect(0, 0, 1000, 400))
	// 左侧封面
	draw.BiLinear.Scale(sub, image.Rect(0, 0, 400, 400), cover, cover.Bounds(), draw.Over, nil)
	// 右侧封面
	blur := imgo.LoadFromImage(cover).GaussianBlur(10, 10).ToImage()
	draw.BiLinear.Scale(sub, image.Rect(400, -100, 1000, 500), blur, blur.Bounds(), draw.Over, nil)
	draw.Draw(sub, sub.Rect, background, image.Point{}, draw.Over)
	draw.Draw(background, background.Rect, sub, image.Point{}, draw.Src)
}

/*addText 在图像上添加文本*/
func addText(img *image.RGBA, label string, x, y int, c color.Color, fontSize float64) {
	// 创建字体
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalln("生成文字时出现错误：", err)
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(c),
		Face: face,
		Dot:  fixed.P(x, y),
	}
	d.DrawString(label)
}

/*绘制顶部区域*/
func drawTop(song types.Song) *image.RGBA {
	// 空画布
	background := image.NewRGBA(image.Rect(0, 0, 600, 100))

	// 添加难度bar
	var bar image.Image
	switch song.Music.Level_info.Difficulty {
	case 0:
		bar, _, _ = image.Decode(bytes.NewBuffer(bar_basic))
	case 1:
		bar, _, _ = image.Decode(bytes.NewBuffer(bar_advanced))
	case 2:
		bar, _, _ = image.Decode(bytes.NewBuffer(bar_expert))
	case 3:
		bar, _, _ = image.Decode(bytes.NewBuffer(bar_master))
	case 4:
		bar, _, _ = image.Decode(bytes.NewBuffer(bar_remas))
	}
	draw.Draw(background, image.Rect(25, 25, 388, 75), bar, image.Point{}, draw.Over)

	// 添加dx/std标识
	if song.Music.Is_deluxe {
		pic, _, _ := image.Decode(bytes.NewBuffer(dx))
		sub := image.NewRGBA(image.Rect(0, 0, 175, 50))
		draw.BiLinear.Scale(sub, sub.Rect, pic, pic.Bounds().Bounds(), draw.Over, nil)
		draw.Draw(background, image.Rect(25, 25, 200, 75), sub, image.Point{}, draw.Over)
	} else {
		pic, _, _ := image.Decode(bytes.NewBuffer(standard))
		sub := image.NewRGBA(image.Rect(0, 0, 175, 50))
		draw.BiLinear.Scale(sub, sub.Rect, pic, pic.Bounds().Bounds(), draw.Over, nil)
		draw.Draw(background, image.Rect(25, 25, 200, 75), sub, image.Point{}, draw.Over)
	}

	// 等级哈希表
	levelHash := map[int]string{
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

	ds, rating := songlist.GetSongRating(song)
	// 添加定数
	// Todo: 后面要用定数表替换
	if ds != "0.0" {
		addText(background, ds, 245, 67, color.White, 48)
	} else {
		addText(background, levelHash[song.Music.Level_info.Level], 245, 67, color.White, 48)
	}
	// 添加rating
	// Todo: 后面使用定数表替换
	if rating == "0" {
		rating = "Rank"
	}
	addText(background, rating, 415, 67, color.White, 48)

	return background
}

/*添加顶部区域*/
func addTop(background *image.RGBA, song types.Song, x, y int, width, height int) {
	label := drawTop(song)
	draw.BiLinear.Scale(label, image.Rect(0, 0, width, height), label, label.Rect, draw.Over, nil)
	draw.Draw(background, image.Rect(x, y, x+width, y+height), label, image.Point{}, draw.Over)
}

/*生成中间内容*/
func drawMiddle(song types.Song) *image.RGBA {
	// 创建画布
	background := image.NewRGBA(image.Rect(0, 0, 600, 175))
	// 添加标题
	addText(background, song.Music.Name, 25, 55, color.White, 38)
	// 添加成绩
	addText(background, fmt.Sprintf("%d.%.4d%%", song.Achievement/10000, song.Achievement%10000), 25, 159, color.White, 75)
	return background
}

/*添加中间内容*/
func addMiddle(background *image.RGBA, song types.Song, x, y int, width, height int) {
	title := drawMiddle(song)
	draw.BiLinear.Scale(title, image.Rect(0, 0, width, height), title, title.Rect, draw.Src, nil)
	draw.Draw(background, image.Rect(x, y, x+width, y+height), title, image.Point{}, draw.Over)
}

/*生成底部信息*/
func drawBottom(song types.Song, record types.Record) *image.RGBA {
	background := image.NewRGBA(image.Rect(0, 0, 600, 125))
	var current types.MusicDetail
	for _, levelinfo := range record.Data.Music_detail {
		if levelinfo.Difficulty == song.Music.Level_info.Difficulty {
			current = levelinfo
		}
	}
	sss_, _, _ := image.Decode(bytes.NewBuffer(score_rank[current.Score_rank]))
	ap_, _, _ := image.Decode(bytes.NewBuffer(combo_status[current.Combo_status]))
	fdx_, _, _ := image.Decode(bytes.NewBuffer(sync_status[current.Sync_status]))
	draw.BiLinear.Scale(background, image.Rect(0, 18, 200, 107), sss_, sss_.Bounds(), draw.Over, nil)
	if ap_ != nil {
		draw.BiLinear.Scale(background, image.Rect(250, 7, 350, 118), ap_, ap_.Bounds(), draw.Over, nil)
	}
	if fdx_ != nil {
		draw.BiLinear.Scale(background, image.Rect(450, 7, 550, 118), fdx_, fdx_.Bounds(), draw.Over, nil)
	}
	return background
}

/*添加底部信息*/
func addBottom(background *image.RGBA, song types.Song, record types.Record, x, y int, width, height int) {
	bottom := drawBottom(song, record)
	draw.BiLinear.Scale(bottom, image.Rect(0, 0, width, height), bottom, bottom.Rect, draw.Src, nil)
	draw.Draw(background, image.Rect(x, y, x+width, y+height), bottom, image.Point{}, draw.Over)
}

/*绘制一首歌曲的信息*/
func DrawOne(song types.Song, record types.Record) image.Image {
	// 生成画板
	background := image.NewRGBA(image.Rect(0, 0, 1000, 400))

	// 加载并添加底板
	sub, _, _ := image.Decode(bytes.NewBuffer(blank))
	draw.Draw(background, background.Rect, sub, image.Point{}, draw.Over)

	// 添加封面
	addCover(background, song)

	// 添加分数
	addTop(background, song, 400, 0, 600, 100)

	// 添加额外内容
	addMiddle(background, song, 400, 100, 600, 175)

	// 添加底部信息
	addBottom(background, song, record, 400, 275, 600, 125)

	return imgo.LoadFromImage(background).BorderRadius(20).ToImage()
}
