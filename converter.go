package main

import (
	"github.com/floostack/transcoder/ffmpeg"
)

func convertToOgg(recordUrl string) (string, error) {
	outputPath := "/tmp/record.ogg"
	format := "ogg"
	overwrite := true

	opts := ffmpeg.Options{
		OutputFormat: &format,
		Overwrite:    &overwrite,
	}

	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath:  "ffmpeg",
		FfprobeBinPath: "ffprobe",
	}

	_, err := ffmpeg.
		New(ffmpegConf).
		Input(recordUrl).
		Output(outputPath).
		WithOptions(opts).
		Start(opts)

	if err != nil {
		return "", err
	}

	return outputPath, nil
}
