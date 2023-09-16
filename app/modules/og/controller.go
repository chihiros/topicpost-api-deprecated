package og

import (
	"image"
	"image/color"
	"net/http"
	"strings"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"golang.org/x/image/font/opentype"

	_ "embed"
)

//go:embed embed/NotoSansJP-Medium.otf
var fontTitle []byte

//go:embed embed/NotoSansJP-Regular.otf
var fontUserName []byte

//go:embed embed/logo.png
var logo []byte

type OGImage struct {
	dc *gg.Context
}

func NewOGImage() *OGImage {
	return &OGImage{}
}

func GenerateOGImage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		userName := r.URL.Query().Get("user")

		og := NewOGImage()
		og.dc = getBackgroundImage()

		// フォントを読み込む
		if err := og.setFont(fontTitle, opentype.FaceOptions{
			Size: 64,
			DPI:  72,
		}); err != nil {
			http.Error(w, "Failed to parse font", http.StatusInternalServerError)
			return
		}

		// 文字を挿入
		dc.SetRGB(0, 0, 0) // 文字色を黒に設定

		maxWidth := 910.0
		formatTitle := ""

		tmp := 0.0
		for _, word := range title {
			fw, _ := dc.MeasureString(string(word))
			if tmp+fw > maxWidth {
				formatTitle += "\n"
				tmp = 0.0
			}

			formatTitle += string(word)
			tmp += fw
		}

		x := 145.0
		y := 175.0
		for _, line := range strings.Split(formatTitle, "\n") {
			dc.DrawString(line, x, y)
			y += 64
		}

		// サイトのロゴを挿入
		logoImg, _, err := image.Decode(strings.NewReader(string(logo)))
		if err != nil {
			http.Error(w, "Failed to decode logo image", http.StatusInternalServerError)
			return
		}

		// ロゴのサイズを変更
		resizedLogoImg := resize.Resize(0, 50, logoImg, resize.Lanczos3)

		// ロゴを挿入
		dc.DrawImage(resizedLogoImg, 850, 500)

		// フォントを読み込む
		if err := og.setFont(fontUserName, opentype.FaceOptions{
			Size: 48,
			DPI:  72,
		}); err != nil {
			http.Error(w, "Failed to parse font", http.StatusInternalServerError)
			return
		}

		// 投稿者の名前を挿入
		dc.SetRGB(0, 0, 0) // 文字色を黒に設定

		x = 160.0
		y = 520.0
		for _, line := range strings.Split(userName, "\n") {
			dc.DrawString(line, x, y)
			y += 48
		}

		// 画像をレスポンスとして返す
		w.Header().Set("Content-Type", "image/png")
		dc.EncodePNG(w)
	}

}

func getBackgroundImage() *gg.Context {
	// 1200x630の画像を生成
	dc := gg.NewContext(1200, 630)

	// 左上から右下に向かってグラデーション
	grad := gg.NewLinearGradient(0, 0, 1200, 630)
	grad.AddColorStop(0, color.RGBA{255, 197, 193, 255})
	grad.AddColorStop(0.25, color.RGBA{244, 222, 244, 255})
	grad.AddColorStop(0.6943, color.RGBA{255, 249, 195, 255})
	grad.AddColorStop(1, color.RGBA{206, 249, 255, 255})
	dc.SetFillStyle(grad)

	// 画像全体の矩形を描画してグラデーションを適用
	dc.DrawRectangle(0, 0, 1200, 630)
	dc.Fill()

	// 図形のサイズと位置を計算
	rectWidth := 1200 - 2*43
	rectHeight := 630 - 2*41
	rectX := 43
	rectY := 41

	// 背景色を設定
	dc.SetColor(color.RGBA{255, 255, 255, 255})
	dc.DrawRoundedRectangle(float64(rectX), float64(rectY), float64(rectWidth), float64(rectHeight), 16)
	dc.Fill()

	return dc
}

func (og *OGImage) setFont(font []byte, options opentype.FaceOptions) error {
	// フォントを読み込む
	f, err := opentype.Parse(font)
	if err != nil {
		return err
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size: options.Size,
		DPI:  options.DPI,
	})
	if err != nil {
		return err
	}
	og.dc.SetFontFace(face)

	return nil
}

type DrawStringOptions struct {
	Color       color.RGBA // 文字色
	MaxWidth    float64    // 最大幅, 折り返すのPx数
	BetweenLine float64    // 行間
	Text        string     // 描画する文字列
	Position    Position   // 描画する位置
}

type Position struct {
	X float64 // X座標
	Y float64 // Y座標
}

func (og *OGImage) drawString(options DrawStringOptions) {
	og.dc.SetRGBA(0, 0, 0, 255) // 文字色を黒に設定

	formatTitle := ""
	tmp := 0.0
	for _, word := range options.Text {
		fw, _ := og.dc.MeasureString(string(word))
		if tmp+fw > options.MaxWidth {
			formatTitle += "\n"
			tmp = 0.0
		}

		formatTitle += string(word)
		tmp += fw
	}

	for _, line := range strings.Split(formatTitle, "\n") {
		og.dc.DrawString(line, options.Position.X, options.Position.Y)
		options.Position.Y += options.BetweenLine
	}
}
