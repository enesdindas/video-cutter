package ffmpeg_wrapper

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"

	"github.com/rs/zerolog/log"
)

// ExecuteFFCommand executes an FFmpeg or FFprobe command with the given arguments
func ExecuteFFCommand(cmdName string, args []string, reader io.Reader, writer io.Writer) error {
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, cmdName, args...)
	cmd.Stdin = reader
	cmd.Stdout = writer
	stdErrBuf := bytes.NewBuffer(nil)
	cmd.Stderr = stdErrBuf

	err := cmd.Run()
	if err != nil {
		log.Error().Str("command", cmdName).Err(err).Msg("Command execution failed")
		return fmt.Errorf("%s error: %w\nStderr: %s", cmdName, err, stdErrBuf.String())
	}
	return nil
}

// GetJSONOutput retrieves JSON output from FFprobe
func GetJSONOutput(reader io.Reader) (string, error) {
	args := []string{
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		"-show_error",
		"-",
	}
	buf := bytes.NewBuffer(nil)
	err := ExecuteFFCommand("ffprobe", args, reader, buf)
	if err != nil {
		return "", fmt.Errorf("failed to get JSON output: %w", err)
	}
	return buf.String(), nil
}

func CutVideo(inputFile, outputFile, start, end, outputFormat string) error {
	args := []string{
		"-ss", start,
		"-i", inputFile,
		"-to", end,
		"-c:v", "libx264", // Use libx264 for video encoding
		"-c:a", "aac",     // Use AAC for audio encoding
		outputFile,
	}
	return ExecuteFFCommand("ffmpeg", args, nil, nil)
}

func CutVideoWithDuration(inputFile, outputFile, start, duration, outputFormat string) error {
	args := []string{
		"-i", inputFile,
		"-ss", start,
		"-t", duration,
		"-c", "copy",
		outputFile,
	}
	return ExecuteFFCommand("ffmpeg", args, nil, nil)
}
