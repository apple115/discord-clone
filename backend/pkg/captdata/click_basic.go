package captdata

import (
	"log"
	"strconv"

	"github.com/golang/freetype/truetype"
	"github.com/wenlng/go-captcha-assets/resources/fonts/fzshengsksjw"
	"github.com/wenlng/go-captcha-assets/resources/images"
	"github.com/wenlng/go-captcha-assets/sourcedata/chars"
	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/click"
)

var TextCapt click.Captcha
var LightTextCapt click.Captcha

func Setup() {
	builder := click.NewBuilder(
		click.WithRangeLen(option.RangeVal{Min: 4, Max: 6}),
		click.WithRangeVerifyLen(option.RangeVal{Min: 2, Max: 4}),
		click.WithRangeThumbBgColors([]string{
			"#1f55c4",
			"#780592",
			"#2f6b00",
			"#910000",
			"#864401",
			"#675901",
			"#016e5c",
		}),
		click.WithRangeColors([]string{
			"#fde98e",
			"#60c1ff",
			"#fcb08e",
			"#fb88ff",
			"#b4fed4",
			"#cbfaa9",
			"#78d6f8",
		}),
	)

	fonts, err := fzshengsksjw.GetFont()
	if err != nil {
		log.Fatalln(err)
	}
	imgs, err := images.GetImages()
	if err != nil {
		log.Fatalln(err)
	}
	builder.SetResources(
		click.WithChars(chars.GetChineseChars()),
		click.WithFonts([]*truetype.Font{fonts}),
		click.WithBackgrounds(imgs),
	)
	TextCapt = builder.Make()

	builder.Clear()
	builder.SetOptions(
		click.WithRangeLen(option.RangeVal{Min: 4, Max: 6}),
		click.WithRangeVerifyLen(option.RangeVal{Min: 2, Max: 4}),
		click.WithRangeThumbColors([]string{
			"#4a85fb",
			"#d93ffb",
			"#56be01",
			"#ee2b2b",
			"#cd6904",
			"#b49b03",
			"#01ad90",
		}),
	)
	builder.SetResources(
		click.WithChars(chars.GetChineseChars()),
		click.WithFonts([]*truetype.Font{fonts}),
		click.WithBackgrounds(imgs),
	)
	LightTextCapt = builder.Make()
}

// verifyDots 验证用户提交的坐标是否与存储的验证码数据匹配
func VerifyDots(dots []string, storedDotsMap map[int]*click.Dot) bool {
	if len(dots)%2 != 0 || (len(dots)/2) != len(storedDotsMap) {
		return false
	}

	for i := 0; i < len(storedDotsMap); i++ {
		dot := storedDotsMap[i]
		j := i * 2
		k := i*2 + 1
		sx, err := strconv.ParseFloat(dots[j], 64)
		if err != nil {
			return false
		}
		sy, err := strconv.ParseFloat(dots[k], 64)
		if err != nil {
			return false
		}
		chkRet := click.CheckPoint(int64(sx), int64(sy), int64(dot.X), int64(dot.Y), int64(dot.Width), int64(dot.Height), 0)
		if !chkRet {
			return false
		}
	}

	return true
}
