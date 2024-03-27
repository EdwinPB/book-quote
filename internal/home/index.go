package home

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/leapkit/core/render"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"

	_ "embed"
)

var (
	//go:embed cards/standard_quote.svg
	svgContent string
)

func Index(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())

	err := rw.Render("home/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Clicked(w http.ResponseWriter, r *http.Request) {
	// Convert SVG to PNG
	err := SVGtoPNG()
	if err != nil {
		fmt.Println("Error converting SVG to PNG:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	file, err := os.Open("card.png")
	if err != nil {
		fmt.Println("Error opening file:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Set appropriate headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename=card.png")
	w.Header().Set("Content-Type", "image/png")

	// Copy the file's contents to the response
	_, err = io.Copy(w, file)
	if err != nil {
		log.Println("Error copying file contents to response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func createFile() (*os.File, error) {
	fmt.Println(svgContent)

	f, err := os.Create("/tmp/card.svg")
	defer f.Close()

	data := []byte(svgContent)
	err = os.WriteFile("/tmp/card.svg", data, 0644)

	fmt.Println("------->Ojo", string(data))

	return f, err
}

func SVGtoPNG() error {
	_, err := createFile()
	if err != nil {
		return err
	}

	in, err := os.Open("/tmp/card.svg")
	if err != nil {
		return err
	}

	icon, err := oksvg.ReadIconStream(in, oksvg.ErrorMode(2))
	if err != nil {
		fmt.Println("------>", err)
		return err
	}

	w := int(icon.ViewBox.W)
	h := int(icon.ViewBox.H)
	x := int(icon.ViewBox.X)
	y := int(icon.ViewBox.Y)

	icon.SetTarget(float64(x), float64(y), float64(w), float64(h))
	rgba := image.NewRGBA(image.Rect(x, y, w, h))
	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)

	rgba.SubImage(rgba.Bounds())

	out, err := os.Create("card.png")
	if err != nil {
		log.Println(err)
		return err
	}
	defer out.Close()

	err = png.Encode(out, rgba)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
