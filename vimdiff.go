package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"

	"go.bbkane.com/warg/command"
)

func vimdiff(ctx command.Context) error {
	dsts := ctx.Flags["--dst"].([]string)
	file := ctx.Flags["--file"].(string)
	src := ctx.Flags["--src"].(string)

	for _, dst := range dsts {

		dstFile := path.Join(dst, file)
		if _, err := os.Stat(dstFile); err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				fmt.Printf("Dst file not found: %s\n\n", dstFile)
			} else {
				fmt.Printf("Error with dstfile: %s: %v", dstFile, err)
			}
			continue
		}

		srcFile := path.Join(src, file)
		fmt.Printf("vimdiff %s %s\n\n", srcFile, dstFile)
	}
	return nil
}
