package main

import (
	"log"
	"os"
)

func init() {
	log.Printf("·  Loaded %d keywords.\n", len(KEYWORDS_TABLE))
	for hexPair, keyword := range KEYWORDS_TABLE {
		log.Printf("   ·  For example: % x translates to: %v\n", hexPair, keyword)
		break
	}
}

func main() {
	bas := AmigaBasicFile{
		Name: os.Args[1],
	}

	bas.Load()
	bas.Parse()

	newFileName := bas.newFileName()
	bas.save(newFileName)

	log.Println("·  Done!")
	log.Printf("·  New file content is %d bytes:\n% x\n%s\n", len(bas.body.data), bas.body.data, bas.body.data)
}
