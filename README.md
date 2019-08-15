# Magick

A golang wrapper for [ImageMagick](http://imagemagick.org/) or
[GraphicsMagick](http://www.graphicsmagick.org/) command line.

## Requirements

ImageMagick or GraphicsMagick command-line tool has to be installed. You can
check if you have it installed by running

```sh
$ convert -version
Version: ImageMagick 7.0.8-8 Q16 x86_64 2018-08-06 https://www.imagemagick.org
Copyright: © 1999-2018 ImageMagick Studio LLC
License: https://www.imagemagick.org/script/license.php
Features: Cipher DPC HDRI Modules
Delegates (built-in): bzlib fontconfig freetype jng jpeg ltdl lzma png tiff xml zlib
```

## Installation

```sh
$ go get -u -v github.com/tnclong/magick
```

## Usage

Let's see a basic example of resizing an gif image.

```go
package main

import (
	"os"

	"github.com/tnclong/magick"
)

func main() {
	e := &magick.Engine{}
	resize := e.Convert("testdata/animation.gif", "-resize", "128x128>", "-verbose", "output.gif")
	resize.Stdout = os.Stdout
	resize.Stderr = os.Stderr
	err := resize.Run()
	if err != nil {
		panic(err)
	}
}
```

## Doc

https://godoc.org/github.com/tnclong/magick

