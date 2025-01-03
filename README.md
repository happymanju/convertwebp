# convertwebp

Converts images to Google's webp format. Converts any images found in the current folder with any of the following file extensions:
- jpg
- jpeg
- png

This program is really just a wrapper around Google's precompiled `cwebp.exe` tool. It runs `cwebp.exe` for every image found in the current folder, only.

## Requirements

This tool needs the `cwebp` binary on your system PATH to work correctly.

You can get the libwebp tools collection for your system here: [https://storage.googleapis.com/downloads.webmproject.org/releases/webp/index.html](https://storage.googleapis.com/downloads.webmproject.org/releases/webp/index.html)

## Details

`convertwebp` executes `cwebp` as the following: `cwebp -pass 10 -m 6 <input file> -o <output file .webp>`. Quality of compression is prioritized over speed, so this tool will take a long time if many large images are converted.

- `-pass` sets the analysis passes to maximum of 10 to help the tool do the best possible conversion. 
- `-m` sets the compression to maximum of 6