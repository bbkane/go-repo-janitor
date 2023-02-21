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

	srcFile := path.Join(src, file)

	for _, dst := range dsts {
		dstFile := path.Join(dst, file)

		var testLine string
		switch file {
		case ".golangci.yml":
			testLine = fmt.Sprintf("cd %s && golangci-lint run && cd -\n\n", dst)
		default:
			testLine = "\n"
		}

		if _, err := os.Stat(dstFile); err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				fmt.Printf("Dst file not found: %s\n", dstFile)
				fmt.Printf("cp %s %s\n", srcFile, dstFile)
				fmt.Print(testLine)
			} else {
				fmt.Printf("Error with dstfile: %s: %v", dstFile, err)
			}
			continue
		}

		fmt.Printf("vimdiff %s %s\n\n", srcFile, dstFile)
	}
	return nil
}
