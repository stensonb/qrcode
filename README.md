# qrcode #

A simple QR Code generator, using github.com/skip2/go-qrcode

## Install

```
$ go get -u github.com/stensonb/qrcode
```

A command-line tool `qrcode` will be built into `$GOPATH/bin/`.

## Usage

```
$ qrcode -h
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
