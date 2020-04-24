// qrcode
// Copyright 2020 Bryan Stenson

package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	size := flag.Int("s", 256, "image size (pixel)")
	textArt := flag.Bool("t", terminal.IsTerminal(int(os.Stdout.Fd())), "print as text-art on STDOUT")
	negative := flag.Bool("i", false, "invert black and white")
	disableBorder := flag.Bool("d", false, "disable QR Code border")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `qrcode -- QR Code encoder in Go
https://github.com/skip2/go-qrcode

Flags:
`)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
Data read from STDIN, and written to STDOUT.  If STDOUT is a terminal device, print as text-art.
If STDOUT is not a terminal (unix pipe, for example), create a PNG image and write to STDOUT.

Usage:
  1. Output a text-art QR code to the terminal:
       echo "hello world" | qrcode

  2. Save a PNG image to file:
       echo "hello world" | qrcode > out.png

  3. Pipe PNG image to 'display' (on any X server):
       echo "hello world" | qrcode | display

  4. Force output to PNG image (just the bytes), when writing to a terminal-based STDOUT:
       echo "hello world" | qrcode -t=false

`)
	}
	flag.Parse()

	var content string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// scanner.Scan() returns lines, so re-join them here
		content = content + "\n" + scanner.Text()
	}

	var err error
	var q *qrcode.QRCode
	q, err = qrcode.New(content, qrcode.Highest)
	checkError(err)

	q.DisableBorder = *disableBorder

	if *negative {
		q.ForegroundColor, q.BackgroundColor = q.BackgroundColor, q.ForegroundColor
	}

	if *textArt {
		art := q.ToString(*negative)
		fmt.Println(art)
	} else {
		var png []byte
		png, err = q.PNG(*size)
		checkError(err)
		os.Stdout.Write(png)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
