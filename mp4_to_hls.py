"""Converts mp4 files to HLS compatible files."""

import os

filename = "demo"
ffmpeg_str = f"ffmpeg -i {filename}.mp4 -codec: copy -start_number 0 -hls_time 10 -hls_list_size 0 -f hls {filename}.m3u8"
cd = os.getcwd()

os.chdir(cd + '/videos/')
os.system(ffmpeg_str)
os.chdir('..')
