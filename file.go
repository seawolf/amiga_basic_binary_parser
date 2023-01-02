package main

import (
	"bytes"
	"fmt"
)

type AmigaBasicFile struct {
	Name       string
	BinaryData []byte

	header AmigaBasicFileHeader
	body   AmigaBasicFileBody
	footer AmigaBasicFileFooter
}

func (bas *AmigaBasicFile) check() {
	if bas.header.isEncryptedBody() {
		err := fmt.Sprintf("!  File is in the encrypted/protected format, which is unsupported. (Header: % x % x)",
			bas.BinaryData[0], bas.BinaryData[1])
		panic(err)
	}

	if !bas.header.isValid() {
		err := fmt.Sprintf("!  File did not start with expected header; expected: %v%v [len:%d] got: % x [len:%d]",
			HEADER[0], HEADER[1], len(HEADER), bas.header.data, len(bas.header.data))
		panic(err)
	}
}

var SPLITTER_BODY_FOOTER = "\x00\x00\x00\x00"

func (bas *AmigaBasicFile) extractBodyAndFooter() {
	headerEnd := len(bas.header.data)

	body, footer, _ := bytes.Cut(bas.BinaryData, []byte(SPLITTER_BODY_FOOTER))

	bas.body.data = body[headerEnd:]
	bas.footer.data = footer
}

func (bas *AmigaBasicFile) transform() {
	bas.body.transformKeywords()
	bas.body.transformLabels(bas.footer.labels)
	bas.body.removeLabelPrefixes()
}
