package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
	"ns/video-cutter/ffmpeg_wrapper"
)

var version = "No Version Provided"

func main() {
	app := &cli.App{
		Name:    "video-cutter",
		Usage:   "Cut video files (supports MP4 and MOV)",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "Start time for cutting (format: HH:MM:SS)",
			},
			&cli.StringFlag{
				Name:    "end",
				Aliases: []string{"e"},
				Usage:   "End time for cutting (format: HH:MM:SS)",
			},
			&cli.StringFlag{
				Name:    "duration",
				Aliases: []string{"d"},
				Usage:   "Duration of the cut (format: HH:MM:SS)",
			},
		},
		Action: runCutVideo,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func runCutVideo(c *cli.Context) error {
	if c.NArg() != 2 {
		return fmt.Errorf("please provide input and output file paths")
	}

	inputFile := c.Args().Get(0)
	outputFile := c.Args().Get(1)
	start := c.String("start")
	end := c.String("end")
	duration := c.String("duration")

	if err := cutVideo(inputFile, outputFile, start, end, duration); err != nil {
		return fmt.Errorf("error cutting video: %v", err)
	}

	fmt.Println("Video cut successfully!")
	return nil
}

func cutVideo(inputFile, outputFile, start, end, duration string) error {
	// Determine output format based on file extension
	outputFormat := "mp4"
	if strings.ToLower(filepath.Ext(outputFile)) == ".mov" {
		outputFormat = "mov"
	}

	var err error
	if duration != "" {
		err = ffmpeg_wrapper.CutVideoWithDuration(inputFile, outputFile, start, duration, outputFormat)
	} else if start != "" && end != "" {
		err = ffmpeg_wrapper.CutVideo(inputFile, outputFile, start, end, outputFormat)
	} else {
		return fmt.Errorf("invalid time parameters")
	}

	if err != nil {
		return fmt.Errorf("ffmpeg error: %v", err)
	}

	return nil
}
