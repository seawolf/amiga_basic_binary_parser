package main

import (
	"bytes"
)

type KeywordsList map[KeywordReference]Keyword
type KeywordReference string
type Keyword string

/*
	With many thanks to Pepijn de Vries for the "AmigaFFH" package in R:
	https://CRAN.R-project.org/package=AmigaFFH
	https://github.com/pepijn-devries/AmigaFFH/blob/0db81d031c96329476ef7c4107f93bddd1447b71/R/basic.r#L61-L271
*/
var KEYWORDS_TABLE = KeywordsList{
	"\x00\x00": "\n",
	"\x13\x00": "", // Shift In
	"\x14\x00": "", // Shift Out

	"\x18": "", // End-of-Label marker

	"\x80":     "ABS",
	"\x81":     "ASC",
	"\x82":     "ATN",
	"\x83":     "CALL",
	"\x84":     "CDBL",
	"\x85":     "CHR$",
	"\x86":     "CINT",
	"\x87":     "CLOSE",
	"\x88":     "COMMON",
	"\x89":     "COS",
	"\x8a":     "CVD",
	"\x8b":     "CVI",
	"\x8c":     "CVS",
	"\x8d":     "DATA",
	"\x8e":     "ELSE",
	"\x8f":     "EOF",
	"\x90":     "EXP",
	"\x91":     "FIELD",
	"\x92":     "FIX",
	"\x93":     "FN",
	"\x94":     "FOR",
	"\x95":     "GET",
	"\x96":     "GOSUB",
	"\x97":     "GOTO",
	"\x98":     "IF",
	"\x99":     "INKEY$",
	"\x9a":     "INPUT",
	"\x9b":     "INT",
	"\x9c":     "LEFT$",
	"\x9d":     "LEN",
	"\x9e":     "LET",
	"\x9f":     "LINE",
	"\xa1":     "LOC",
	"\xa2":     "LOF",
	"\xa3":     "LOG",
	"\xa4":     "LSET",
	"\xa5":     "MID$",
	"\xa6":     "MKD$",
	"\xa7":     "MKI$",
	"\xa8":     "MKS$",
	"\xa9":     "NEXT",
	"\xaa":     "ON",
	"\xab":     "OPEN",
	"\xac":     "PRINT",
	"\xad":     "PUT",
	"\xae":     "READ",
	"\xaf\x00": "REM",
	"\xaf\xe8": "'",
	"\xb0":     "RETURN",
	"\xb1":     "RIGHT$",
	"\xb2":     "RND",
	"\xb3":     "RSET",
	"\xb4":     "SGN",
	"\xb5":     "SIN",
	"\xb6":     "SPACE$",
	"\xb7":     "SQR",
	"\xb8":     "STR$",
	"\xb9":     "STRING$",
	"\xba":     "TAN",
	"\xbc":     "VAL",
	"\xbd":     "WEND",
	"\xbe\xec": "WHILE",
	"\xbf":     "WRITE",
	"\xc0":     "ELSEIF",
	"\xc1":     "CLNG",
	"\xc2":     "CVL",
	"\xc3":     "MKL$",
	"\xc4":     "AREA",
	"\xe3":     "STATIC",
	"\xe4":     "USING",
	"\xe5":     "TO",
	"\xe6":     "THEN",
	"\xe7":     "NOT",
	"\xe9":     ">",
	"\xea":     "=",
	"\xeb":     "<",
	"\xec":     "+",
	"\xed":     "-",
	"\xee":     "*",
	"\xef":     "/",
	"\xf0":     "^",
	"\xf1":     "AND",
	"\xf2":     "OR",
	"\xf3":     "XOR",
	"\xf4":     "EQV",
	"\xf5":     "IMP",
	"\xf6":     "MOD",
	"\xf7":     `\`,
	"\xf8\x81": "CHAIN",
	"\xf8\x82": "CLEAR",
	"\xf8\x83": "CLS",
	"\xf8\x84": "CONT",
	"\xf8\x85": "CSNG",
	"\xf8\x86": "DATE$",
	"\xf8\x87": "DEFINT",
	"\xf8\x88": "DEFSNG",
	"\xf8\x89": "DEFDBL",
	"\xf8\x8a": "DEFSTR",
	"\xf8\x8b": "DEF",
	"\xf8\x8c": "DELETE",
	"\xf8\x8d": "DIM",
	"\xf8\x8f": "END",
	"\xf8\x90": "ERASE",
	"\xf8\x91": "ERL",
	"\xf8\x92": "ERROR",
	"\xf8\x93": "ERR",
	"\xf8\x94": "FILES",
	"\xf8\x95": "FRE",
	"\xf8\x96": "HEX$",
	"\xf8\x97": "INSTR",
	"\xf8\x98": "KILL",
	"\xf8\x9a": "LLIST",
	"\xf8\x9b": "LOAD",
	"\xf8\x9c": "LPOS",
	"\xf8\x9d": "LPRINT",
	"\xf8\x9e": "MERGE",
	"\xf8\x9f": "NAME",
	"\xf8\xa0": "NEW",
	"\xf8\xa1": "OCT$",
	"\xf8\xa2": "OPTION",
	"\xf8\xa3": "PEEK",
	"\xf8\xa4": "POKE",
	"\xf8\xa5": "POS",
	"\xf8\xa6": "RANDOMIZE",
	"\xf8\xa8": "RESTORE",
	"\xf8\xa9": "RESUME",
	"\xf8\xaa": "RUN",
	"\xf8\xab": "SAVE",
	"\xf8\xad": "STOP",
	"\xf8\xae": "SWAP",
	"\xf8\xaf": "SYSTEM",
	"\xf8\xb0": "TIME$",
	"\xf8\xb1": "TRON",
	"\xf8\xb2": "TROFF",
	"\xf8\xb3": "VARPTR",
	"\xf8\xb4": "WIDTH",
	"\xf8\xb5": "BEEP",
	"\xf8\xb6": "CIRCLE",
	"\xf8\xb8": "MOUSE",
	"\xf8\xb9": "POINT",
	"\xf8\xba": "PRESET",
	"\xf8\xbb": "PSET",
	"\xf8\xbc": "RESET",
	"\xf8\xbd": "TIMER",
	"\xf8\xbe": "SUB",
	"\xf8\xbf": "EXIT",
	"\xf8\xc0": "SOUND",
	"\xf8\xc2": "MENU",
	"\xf8\xc3": "WINDOW",
	"\xf8\xc5": "LOCATE",
	"\xf8\xc6": "CSRLIN",
	"\xf8\xc7": "LBOUND",
	"\xf8\xc8": "UBOUND",
	"\xf8\xc9": "SHARED",
	"\xf8\xca": "UCASE$",
	"\xf8\xcb": "SCROLL",
	"\xf8\xcc": "LIBRARY",
	"\xf8\xd2": "PAINT",
	"\xf8\xd3": "SCREEN",
	"\xf8\xd4": "DECLARE",
	"\xf8\xd5": "FUNCTION",
	"\xf8\xd6": "DEFLNG",
	"\xf8\xd7": "SADD",
	"\xf8\xd8": "AREAFILL",
	"\xf8\xd9": "COLOR",
	"\xf8\xda": "PATTERN",
	"\xf8\xdb": "PALETTE",
	"\xf8\xdc": "SLEEP",
	"\xf8\xdd": "CHDIR",
	"\xf8\xde": "STRIG",
	"\xf8\xdf": "STICK",
	"\xf9\xf4": "OFF",
	"\xf9\xf5": "BREAK",
	"\xf9\xf6": "WAIT",
	"\xf9\xf8": "TAB",
	"\xf9\xf9": "STEP",
	"\xf9\xfa": "SPC",
	"\xf9\xfb": "OUTPUT",
	"\xf9\xfc": "BASE",
	"\xf9\xfd": "AS",
	"\xf9\xfe": "APPEND",
	"\xf9\xff": "ALL",
	"\xfa\x80": "WAVE",
	"\xfa\x81": "POKEW",
	"\xfa\x82": "POKEL",
	"\xfa\x83": "PEEKW",
	"\xfa\x84": "PEEKL",
	"\xfa\x85": "SAY",
	"\xfa\x86": "TRANSLATE$",
	"\xfa\x87": "OBJECT.SHAPE",
	"\xfa\x88": "OBJECT.PRIORITY",
	"\xfa\x89": "OBJECT.X",
	"\xfa\x8a": "OBJECT.Y",
	"\xfa\x8b": "OBJECT.VX",
	"\xfa\x8c": "OBJECT.VY",
	"\xfa\x8d": "OBJECT.AX",
	"\xfa\x8e": "OBJECT.AY",
	"\xfa\x8f": "OBJECT.CLIP",
	"\xfa\x90": "OBJECT.PLANES",
	"\xfa\x91": "OBJECT.HIT",
	"\xfa\x92": "OBJECT.ON",
	"\xfa\x93": "OBJECT.OFF",
	"\xfa\x94": "OBJECT.START",
	"\xfa\x95": "OBJECT.STOP",
	"\xfa\x96": "OBJECT.CLOSE",
	"\xfa\x97": "COLLISION",
	"\xfb\xff": "PTAB",
}

type AmigaBasicFileBody struct {
	data []byte
}

func (b *AmigaBasicFileBody) transformKeywords() {
	for hexPair, keyword := range KEYWORDS_TABLE {
		b.data = bytes.ReplaceAll(b.data, []byte(hexPair), []byte(keyword))
	}
}

func (b *AmigaBasicFileBody) transformLabels(labels [][]byte) {
	for idx, label := range labels {
		labelReference := []byte{1, 0, byte(idx)}
		labelReplacement := append(label, 24)

		// fmt.Printf("-  Substituting [% x] with [%s] ...\n", labelReference, labelReplacement)
		b.data = bytes.ReplaceAll(b.data, labelReference, labelReplacement)
	}
}

func (b *AmigaBasicFileBody) removeLabelPrefixes() {
	// currentPosition := 0
	prefix := "\x0a"
	firstPrefix := bytes.Index(b.data, []byte(prefix))

	// fmt.Printf("  · Found label prefix at position %d \n", firstPrefix)
	if firstPrefix == -1 {
		return
	}
}

func (b *AmigaBasicFileBody) replaceNewLines() {
	bytes.ReplaceAll(b.data, []byte("\x00\x00"), []byte("\n"))
}
