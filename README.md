# convertwebp

Converts images to Google's webp format. Converts any images found in the current folder with any of the following file extensions:
- jpg
- jpeg
- png

This program is really just a wrapper around Google's precompiled `cwebp.exe` tool. It runs `cwebp.exe` for every image found in the current folder, only.

## Requirements

This tool needs the `cwebp` binary on your system PATH to work correctly.

