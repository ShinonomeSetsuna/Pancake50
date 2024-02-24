package maimage

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"strconv"

	"github.com/ShinonomeSetsuna/Pancake50/internal/pkg/songlist"
	"github.com/ShinonomeSetsuna/Pancake50/internal/types"
	"github.com/fishtailstudio/imgo"
)

func DrawAll(b35, b15 []types.Song) {
	background := image.NewRGBA(image.Rect(0, 0, 5600, 5300))

	draw.Draw(background,
		background.Rect,
		&image.Uniform{color.White},
		image.Point{},
		draw.Over)

	counter := 0

	b35t := 0
	for _, song := range b35 {
		draw.Draw(background,
			image.Rect(counter%5*1100+100, counter/5*500+100, counter%5*1100+1100, counter/5*500+500),
			DrawOne(song, songlist.GetRecord(song.Music.Music_id)),
			image.Point{},
			draw.Over)
		fmt.Println("已完成：", counter+1, "/50")
		counter += 1
		_, rating := songlist.GetSongRating(song)
		num, _ := strconv.ParseInt(rating, 10, 64)
		b35t += int(num)
	}

	b15t := 0
	for _, song := range b15 {
		// 手动阴影
		draw.Draw(background,
			image.Rect(counter%5*1100+100, counter/5*500+300, counter%5*1100+1100, counter/5*500+700),
			DrawOne(song, songlist.GetRecord(song.Music.Music_id)),
			image.Point{},
			draw.Over)
		fmt.Println("已完成：", counter+1, "/50")
		counter += 1
		_, rating := songlist.GetSongRating(song)
		num, _ := strconv.ParseInt(rating, 10, 64)
		b15t += int(num)
	}

	fmt.Printf("B35: %d, B15: %d, Total: %d", b35t, b15t, b35t+b15t)
	imgo.LoadFromImage(background).Resize(2800, 2700).Save("./b50.jpg")
}
