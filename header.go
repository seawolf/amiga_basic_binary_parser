package main

var HEADER = "\xf5\x00"
var HEADER_ENCRYPTED = "\xf4\xc2"

func (bas *AmigaBasicFile) extractHeader() {
	headerSize := 4
	if len(bas.BinaryData) < 1 {
		headerSize = 0
	} else if len(bas.BinaryData) < 6 {
		headerSize = 1
	}

	bas.header.data = bas.BinaryData[:headerSize]
}

type AmigaBasicFileHeader struct {
	data []byte
}

func (h AmigaBasicFileHeader) isValid() bool {
	if len(h.data) == 0 {
		return true
	}

	return len(h.data) > 0 &&
		h.data[0] == HEADER[0] &&
		h.data[1] == HEADER[1]
}

func (h AmigaBasicFileHeader) isEncryptedBody() bool {
	return len(h.data) > 0 &&
		h.data[0] == HEADER_ENCRYPTED[0] &&
		h.data[1] == HEADER_ENCRYPTED[1]
}
