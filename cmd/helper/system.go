package helper

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
)

func GoModTidy() error {
	var out bytes.Buffer

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = &out

	return cmd.Run()
}

func Wire(pkg string) error {
	var out bytes.Buffer

	cmd := exec.Command("wire", pkg)
	cmd.Stdout = &out

	return cmd.Run()
}

func SaveFile(fileBytes []byte, outdir string, fname string) error {
	if err := os.MkdirAll(outdir, os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(outdir + "/" + fname)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)

	_, err = w.WriteString(string(fileBytes))
	if err != nil {
		return err
	}

	return w.Flush()
}
