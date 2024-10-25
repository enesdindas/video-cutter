
# Video Cutter - CLI Tool for Trimming Videos

The **Video Cutter** is a simple, efficient command-line tool designed for trimming videos. Initially created to streamline the process of cutting Zoom recordings, this tool supports a variety of video formats and offers precise control over start and end times.

## Features
- Trim videos to extract only the segments you need.
- Support for both same-format and different-format output files (e.g., `.mov` to `.mp4`).
- Fast and lightweight, ideal for quick video editing directly from the command line.

## Installation

Clone the repository and build the application using `go build`.

```bash
git clone https://github.com/yourusername/video-cutter.git
cd video-cutter
go build -o video-cutter
```

## Usage

### Basic Usage
Trim a video by specifying the start and end times for the output clip. Both start and end times should be provided in the format `HH:MM:SS`.

#### 1. Same Format Output
To trim a video and save it in the same format as the input:

```bash
./video-cutter --start 00:00:05 --end 00:00:10 input.mov output.mov
```

This will take the video `input.mov`, extract the clip from 5 seconds to 10 seconds, and save it as `output.mov`.

#### 2. Different Format Output
You can also export the output clip in a different format from the input video:

```bash
./video-cutter --start 00:00:05 --end 00:00:10 input.mov output.mp4
```

In this case, the video will be trimmed and saved as `output.mp4` instead of the original format.

## Future Enhancements
- **Batch Processing**: Allow trimming multiple videos at once.
- **GUI Version**: A graphical interface for users who prefer not to use the command line.
- **Cloud Storage Integration**: Save output clips directly to cloud storage platforms like Google Drive or Dropbox.

## Contributing
Contributions are welcome! Feel free to submit a pull request or open an issue to report bugs or suggest new features.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
