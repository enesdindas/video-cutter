package ffmpeg_wrapper

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func GetJsonOutput(reader io.Reader) (string, error) {
	args := []string{
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		"-show_error",
		"-",
	}
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, "ffprobe", args...)
	cmd.Stdin = reader
	buf := bytes.NewBuffer(nil)
	stdErrBuf := bytes.NewBuffer(nil)
	cmd.Stdout = buf
	cmd.Stderr = stdErrBuf
	err := cmd.Run()
	if err != nil {
		log.Error().Msg("Error: " + err.Error())
		return "", fmt.Errorf("[%s] %w", string(stdErrBuf.Bytes()), err)
	}
	bytes := string(buf.Bytes())
	return string(bytes), nil
}

func CutVideo(reader io.Reader, writer io.Writer, start string, end string) error {
	args := []string{
		"-v", "quiet",
		"-ss", start,
		"-to", end,
		"-c:v", "copy",
		"-c:a", "copy",
		"-f", "mp4",
		"-loglevel", "error",
		"-y",
		"./output.mp4",
		// "-i", "/Users/enes.dindas/Downloads/SampleVideo_1280x720_30mb.mp4",
	}
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, "ffmpeg", args...)
	cmd.Stdin = reader
	cmd.Stdout = writer
	stdErrBuf := bytes.NewBuffer(nil)
	cmd.Stderr = stdErrBuf
	log.Debug().Msg("Command: " + cmd.String())
	err := cmd.Run()
	if err != nil {
		log.Error().Msg("Command Execution Error: " + err.Error())
		return fmt.Errorf("[%s] %w", string(stdErrBuf.Bytes()), err)
	}
	return nil
}

func CutVideoWithDuration(reader io.Reader, writer io.Writer, start string, duration string) error {
	args := []string{
		"-v", "quiet",
		"-ss", start,
		"-t", duration,
		"-i", "-",
		"-c", "copy",
		"-f", "mp4",
		"-",
	}
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, "ffmpeg", args...)
	cmd.Stdin = reader
	cmd.Stdout = writer
	stdErrBuf := bytes.NewBuffer(nil)
	cmd.Stderr = stdErrBuf
	err := cmd.Run()
	if err != nil {
		log.Error().Msg("Error: " + err.Error())
		return fmt.Errorf("[%s] %w", string(stdErrBuf.Bytes()), err)
	}
	return nil
}
