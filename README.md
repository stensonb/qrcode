# qrcode #

A simple QR Code generator, using [https://github.com/skip2/go-qrcode/](https://github.com/skip2/go-qrcode/)

## Install

```
$ go get -u github.com/stensonb/qrcode
```

A command-line tool `qrcode` will be built into `$GOPATH/bin/`.

## Usage

```
$ qrcode -h
qrcode -- QR Code encoder in Go
https://github.com/stensonb/qrcode

Flags:
  -d	disable QR Code border
  -i	invert black and white
  -s int
    	image size (pixel) (default 256)
  -t	print as text-art on STDOUT (default true)

Data read from STDIN, and written to STDOUT.  If STDOUT is a terminal device, print as text-art.
If STDOUT is not a terminal (unix pipe, for example), create a PNG image and write to STDOUT.

Usage:
  1. Output a text-art QR code to the terminal:
       echo "hello world" | qrcode

  2. Save a PNG file:
       echo "hello world" | qrcode > out.png

  3. Pipe PNG byte stream to 'display' (on any X server):
       echo "hello world" | qrcode | display

  4. Force output to PNG byte stream, when writing to a terminal-based STDOUT:
       echo "hello world" | qrcode -t=false
```

## Maximum capacity
The maximum capacity of a QR Code varies according to the content encoded and the error recovery level. The maximum capacity is 2,953 bytes, 4,296 alphanumeric characters, 7,089 numeric digits, or a combination of these.

## Borderless QR Codes

To aid QR Code reading software, QR codes have a built in whitespace border.

If you know what you're doing, and don't want a border, see https://gist.github.com/skip2/7e3d8a82f5317df9be437f8ec8ec0b7d for how to do it. It's still recommended you include a border manually.

## Links

- [http://en.wikipedia.org/wiki/QR_code](http://en.wikipedia.org/wiki/QR_code)
- [ISO/IEC 18004:2006](http://www.iso.org/iso/catalogue_detail.htm?csnumber=43655) - Main QR Code specification (approx CHF 198,00)<br>
- [https://github.com/skip2/go-qrcode/](https://github.com/skip2/go-qrcode/) - library used here
