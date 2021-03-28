package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"

	proc "github.com/yoyrandao/split-gif/pkg/imageProcessing"
)

func main() {
	parser := argparse.NewParser("splitgif", "Split or join gifs simply")

	action := parser.String("a", "action", &argparse.Options{
		Required: true,
		Help:     "Type of action. Must be one of 'split' or 'join'",
	})
	filepath := parser.String("f", "filepath", &argparse.Options{
		Required: true,
		Help:     "Path to file or files",
	})
	outputDir := parser.String("o", "output", &argparse.Options{
		Required: true,
		Help:     "Path to output directory",
		Default:  "./",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	processor := proc.GifProcessor{}
	switch *action {
	case "split":
		processor.SplitGif(*filepath, *outputDir)
		break

	case "join":
		processor.JoinImagesToGif(*filepath, *outputDir)
		break
	}
}
