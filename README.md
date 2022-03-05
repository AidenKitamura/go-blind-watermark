# go-blind-watermark
A lightweight blind watermark encoding-decoding system in Golang.

---

## Functionality
This program can help generate images with blind watermarks. When used in encoding mode, it takes in two images, the original image file and watermark file, and blend them together using fourier transform. When used in decoding mode, it takes in two images, the watermarked image and the original image, to extract the watermark blended.

---

## Usage
First you need to build the project:
```shell
go build
```
Then run the main file, the flags set as follows.
```shell
-encode     # let program run in encoding mode
-decode     # let program run in decoding mode
-o          # Set original file path
-w          # Set watermark file path (encode mode)
            #     image blended with watermark path (decode mode)
-d          # Set output destination
-i          # Set watermark blending intensity, default by 1.0
-c          # Set channels to be blended. Use binary for RGB. 
            # e.g. RGB to be 111, then type 7.
```