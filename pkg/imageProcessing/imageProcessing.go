package imageProcessing

import (
	"fmt"
	u "github.com/yoyrandao/split-gif/pkg/utils"
	"image"
	"image/gif"
	"image/png"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

type IGifProcessor interface {
	SplitGif(path string, output string)

	JoinImagesToGif(imagePath string, output string)
}

type GifProcessor struct{}

func (processor GifProcessor) SplitGif(path string, output string) {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("error while decoding: %s", r))
		}
	}()

	file, err := os.Open(path)
	if u.CheckWithMessage(err, fmt.Sprintf("Cannot open file - %s", path)) {
		return
	}

	gifImage, _ := gif.DecodeAll(file)

	_ = os.Mkdir(output, 0644)

	for i, sourceImage := range gifImage.Image {
		if u.Average(sourceImage.Pix) == 0 {
			continue
		}

		stats, err := os.Stat(path)
		u.Check(err)

		file, err := os.Create(fmt.Sprintf("%s/%s___%d%s", output, stats.Name(), i, ".png"))
		u.Check(err)

		err = png.Encode(file, sourceImage)
		u.Check(err)

		u.Check(file.Close())
	}
}

func (processor GifProcessor) JoinImagesToGif(imagesPath string, output string) {
	entries, err := os.ReadDir(imagesPath)
	u.Check(err)

	var filePaths []string
	for _, entry := range entries {
		filePaths = append(filePaths, filepath.Join(imagesPath, entry.Name()))
	}

	sort.SliceStable(filePaths, func(i, j int) bool {
		first := filePaths[i]
		second := filePaths[j]

		reg := regexp.MustCompile("___(\\d+)\\.png")

		firstIndex, _ := strconv.Atoi(reg.FindStringSubmatch(first)[1])
		secondIndex, _ := strconv.Atoi(reg.FindStringSubmatch(second)[1])

		return firstIndex < secondIndex
	})

	gifAnimation := &gif.GIF{}

	for _, currentPath := range filePaths {
		handle, err := os.Open(currentPath)
		u.Check(err)

		img, _, err := image.Decode(handle)
		u.Check(handle.Close())
		u.Check(err)

		gifAnimation.Image = append(gifAnimation.Image, img.(*image.Paletted))
		gifAnimation.Delay = append(gifAnimation.Delay, 0)
	}

	out, _ := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, 0600)
	defer out.Close()

	_ = gif.EncodeAll(out, gifAnimation)
}
