import argparse
from pytube import YouTube
import os

def main(args: argparse.Namespace):
	out_file = YouTube(args.url).streams.filter(only_audio=args.audioOnly).first().download(output_path=args.outputPath)
	base, _ = os.path.splitext(out_file)
	new_file = base + '.mp3' if args.audioOnly else '.mp4'
	os.rename(out_file, new_file)


if "__main__" == __name__:
	parser = argparse.ArgumentParser(usage="URL, Audio Only, Output Path flags are needed")
	parser.add_argument("-u", "--url", required=True, help="Youtube URL")
	parser.add_argument("-a", "--audioOnly", required=True,help='Where to only download audio or download the video as well', action="store_true")
	parser.add_argument("-op", "--outputPath", required=True,help="Where to save the downloaded file", default="~")
	main(parser.parse_args())
