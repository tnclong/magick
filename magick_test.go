package magick

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestResize(t *testing.T) {
	e := &Engine{}
	path := imagePath(t, "animation")
	defer os.Remove(path)
	resize := e.Convert(path, "-resize", "128x128>", "-verbose", path)
	t.Log("resize:", *resize)
	resize.Stdout = os.Stdout
	resize.Stderr = os.Stderr
	err := resize.Run()
	if err != nil {
		t.Fatal(err)
	}

	wh := e.Identify("-format", "%w%h", path)
	t.Log("wh:", *wh)
	wh.Stderr = os.Stderr
	data, err := wh.Output()
	if err != nil {
		t.Fatal(err)
	}
	str := string(data)
	if !strings.Contains(str, "128") {
		t.Fatalf("want %s contains 128", str)
	}
}

func TestWhich(t *testing.T) {
	echo := "echo"
	path, err := Which(echo)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("which echo => ", path)
	if !strings.Contains(path, echo) {
		t.Fatalf("want %s contains %s", path, echo)
	}
}

func imagePath(t *testing.T, name string) string {
	switch name {
	case "animation":
		name = "animation.gif"
	}

	f, err := os.Open(filepath.Join("testdata", name))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	tf, err := ioutil.TempFile("", "magick*"+filepath.Ext(name))
	if err != nil {
		t.Fatal(err)
	}
	defer tf.Close()
	_, err = io.Copy(tf, f)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("copy %s to %s", name, tf.Name())
	return tf.Name()
}
