package main

import (
	"TIL-Helper/utils"
	"log"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			utils.AddTilFiles(path)
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	utils.SetTils(utils.GetTilFiles())
	utils.SetIntroduction("Today I Learned")
	utils.CreateDocs(utils.GetTilFiles())
}
