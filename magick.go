package magick

import (
	"os"
	"os/exec"
	"strings"
)

// Engine represents ImageMagick or GraphicsMagick that special in Path.
//
//   e := &magick.Engine{}
//   icmd := e.Identify("testdata/animation.gif")
//   icmd.Stderr = os.Stderr
//   data, err := icmd.Output()
//   fmt.Println(string(data), err)
//     => testdata/animation.gif[0] GIF 500x281 500x281+0+0 8-bit sRGB 256c 912502B 0.000u 0:00.000
//        testdata/animation.gif[1] GIF 453x281 500x281+0+0 8-bit sRGB 256c 912502B 0.000u 0:00.000
//        ...
//        testdata/animation.gif[19] GIF 453x281 500x281+0+0 8-bit sRGB 256c 912502B 0.000u 0:00.000
//        <nil>
type Engine struct {
	// Path is gm/magick command path.
	// If this value is empty string, value of IPath() or GPath() will be assigned when create a Cmd.
	Path string
}

// Command create Cmd for exec gm/magick command that special in Path.
// Command list in:
//   graphicsmagick: http://www.graphicsmagick.org/utilities.html
//   imagemagick: https://imagemagick.org/script/command-line-tools.php
func (e *Engine) Command(name string, arg ...string) *exec.Cmd {
	if e.Path == "" {
		e.Path, _ = IPath()

		if e.Path == "" {
			e.Path, _ = GPath()
		}

		// fork/exec : no such file or directory
	}

	var cmd *exec.Cmd
	if name == "" {
		exec.Command(e.Path, arg...)
	} else {
		cmd = exec.Command(e.Path, name)
		cmd.Args = append(cmd.Args, arg...)
	}
	return cmd
}

// IPath get pathname of `magick`.
func IPath() (string, error) {
	return Which("magick")
}

// GPath get pathname of `gm`.
func GPath() (string, error) {
	return Which("gm")
}

// see https://imagemagick.org/script/animate.php
func (e *Engine) Animate(arg ...string) *exec.Cmd {
	return e.Command("animate", arg...)
}

// see https://imagemagick.org/script/compare.php
func (e *Engine) Compare(arg ...string) *exec.Cmd {
	return e.Command("compare", arg...)
}

// see https://imagemagick.org/script/composite.php
func (e *Engine) Composite(arg ...string) *exec.Cmd {
	return e.Command("composite", arg...)
}

// see https://imagemagick.org/script/conjure.php
func (e *Engine) Conjure(arg ...string) *exec.Cmd {
	return e.Command("conjure", arg...)
}

// see https://imagemagick.org/script/convert.php
func (e *Engine) Convert(arg ...string) *exec.Cmd {
	return e.Command("convert", arg...)
}

// see https://imagemagick.org/script/display.php
func (e *Engine) Display(arg ...string) *exec.Cmd {
	return e.Command("display", arg...)
}

// see https://imagemagick.org/script/identify.php
func (e *Engine) Identify(arg ...string) *exec.Cmd {
	return e.Command("identify", arg...)
}

// see https://imagemagick.org/script/import.php
func (e *Engine) Importc(arg ...string) *exec.Cmd {
	return e.Command("import", arg...)
}

// see https://imagemagick.org/script/mogrify.php
func (e *Engine) Mogrify(arg ...string) *exec.Cmd {
	return e.Command("mogrify", arg...)
}

// see https://imagemagick.org/script/montage.php
func (e *Engine) Montage(arg ...string) *exec.Cmd {
	return e.Command("montage", arg...)
}

// see https://imagemagick.org/script/stream.php
func (e *Engine) Stream(arg ...string) *exec.Cmd {
	return e.Command("stream", arg...)
}

// Which is cross-platform way of finding an executable.
//   Which("echo") => "/bin/echo", nil
func Which(cmd string) (path string, err error) {
	var exts []string
	env, ok := os.LookupEnv("PATHEXT")
	if ok {
		exts = strings.Split(env, ";")
	} else {
		exts = []string{""}
	}

	for _, ext := range exts {
		path, err = exec.LookPath(cmd + ext)
		if err == nil {
			break
		}
	}
	return path, err
}
