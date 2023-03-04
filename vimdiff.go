package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"

	"go.bbkane.com/warg/command"
	"go.bbkane.com/warg/help/common"
)

func vimdiff(ctx command.Context) error {
	dsts := ctx.Flags["--dst"].([]string)
	file := ctx.Flags["--file"].(string)
	src := ctx.Flags["--src"].(string)

	col, err := common.ConditionallyEnableColor(ctx.Flags, os.Stdout)
	if err != nil {
		return fmt.Errorf("color init failure: %w", err)
	}

	srcFile := path.Join(src, file)

	for _, dst := range dsts {
		dstFile := path.Join(dst, file)

		fmt.Printf(
			col.Add(col.FgCyan+col.Bold, "# "+dst) + "\n",
		)

		var testLine string
		switch file {
		case ".golangci.yml":
			testLine = fmt.Sprintf("cd %s && golangci-lint run && cd -\n\n", dst)
		case ".goreleaser.yml":
			testLine = fmt.Sprintf("cd %s && goreleaser --snapshot --skip-publish --clean && cd -\n\n", dst)
		case ".vscode/settings.json":
			testLine = fmt.Sprintf("mkdir %s\n\n", path.Dir(dstFile))
		case "lefthook.yml":
			testLine = fmt.Sprintf("cd %s && lefthook install && lefthook run pre-commit && cd -\n\n", dst)
		case ".github/workflows/release.yml":
			// there might already be a gorelease file
			testLine = fmt.Sprintf("cd %s/.github/workflows; ls;\ngit add . && git commit -m 'Add release.yml'\n\n", dst)
		default:
			testLine = "\n"
		}

		if _, err := os.Stat(srcFile); err != nil {
			errMsg := fmt.Sprintf("Error with srcFile: %s: %v", srcFile, err)
			fmt.Printf(
				col.Add(col.FgRedBright, errMsg) + "\n\n",
			)
			return errors.New("error with srcFile")
		}

		if _, err := os.Stat(dstFile); err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				infoMsg := fmt.Sprintf("Dst file not found: %s", dstFile)
				fmt.Printf(
					col.Add(col.FgGreen, infoMsg) + "\n",
				)
				fmt.Printf("cp %s %s\n", srcFile, dstFile)
				fmt.Print(testLine)
			} else {
				errMsg := fmt.Sprintf("Error with dstfile: %s: %v", dstFile, err)
				fmt.Printf(
					col.Add(col.FgRedBright, errMsg) + "\n\n",
				)
			}
			continue
		}

		srcBytes, err := os.ReadFile(srcFile)
		if err != nil {
			return fmt.Errorf("error reading src file: %w", err)
		}
		dstBytes, err := os.ReadFile(dstFile)
		if err != nil {
			return fmt.Errorf("error reading dst file: %w", err)
		}

		if bytes.Equal(srcBytes, dstBytes) {
			fmt.Printf(
				col.Add(col.FgGreenBright, "Files are equal!") + "\n\n",
			)
			continue
		}

		fmt.Printf("vimdiff %s %s\n\n", srcFile, dstFile)
	}
	return nil
}
