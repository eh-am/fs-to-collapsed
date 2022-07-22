package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	walk(root)
}

func walk(root string) {
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		// TODO: escape ;
		s := fmt.Sprintf("%s %d", strings.Replace(path, "/", ";", -1), info.Size())
		if strings.HasPrefix(s, ";") {
			// TODO: utf8
			s = s[1:]
		}

		fmt.Println(s)

		return nil
	})
}
