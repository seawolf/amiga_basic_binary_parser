package main

import (
	"log"
	"os"
	"strings"
)

func (bas *AmigaBasicFile) Load() {
	binaryData, err := os.ReadFile(bas.Name)

	if err != nil {
		panic(err)
	} else {
		bas.BinaryData = binaryData
		log.Printf("·  Read %d bytes from file: %s\n", len(bas.BinaryData), bas.Name)
	}
}

func (bas *AmigaBasicFile) Parse() {
	bas.check()

	bas.extractHeader()
	bas.extractBodyAndFooter()
	bas.footer.extractLabels()

	bas.transform()
}

func (bas AmigaBasicFile) newFileName() string {
	return strings.Join([]string{bas.Name, ".converted"}, "")
}

func (bas AmigaBasicFile) save(newFileName string) {
	os.WriteFile(newFileName, bas.body.data, 0600)
}
