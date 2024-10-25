# Video Cutter
I am building a tool that I could use for cutting zoom recordings :D

It will be available as CLI application.

## Usage

### Build
go build -o video-cutter

### Run
#### Same Format Output
./video-cutter --start 00:00:05 --end 00:00:10 input.mov output.mov

#### Different Format Output
./video-cutter --start 00:00:05 --end 00:00:10 input.mov output.mp4
