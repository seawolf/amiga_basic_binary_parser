package main

import (
	"encoding/binary"
)

type AmigaBasicFileFooter struct {
	data []byte

	labels [][]byte
}

func (f *AmigaBasicFileFooter) extractLabels() {
	footerSize := uint16(len(f.data))

	// fmt.Printf("  · Debugging footer of %d bytes.\n% x\n", footerSize, f.data)
	if footerSize < 2 {
		return
	}

	currentPosition := uint16(0)
	paddingAmount := uint16(0)
	for currentPosition < footerSize {
		shiftAmount := uint16(1)

		paddingAmountBytes := f.data[currentPosition:(currentPosition + shiftAmount)]
		paddingAmountBytes = append([]byte{0}, paddingAmountBytes[0])
		// fmt.Printf("    1. Reading padding amount: % x \n", paddingAmountBytes)
		paddingAmount = binary.BigEndian.Uint16(paddingAmountBytes)
		// fmt.Printf("    2. Padding amount set to: %d \n", paddingAmount)

		currentPosition += paddingAmount
		currentPosition += shiftAmount
		// fmt.Printf("    3. Current position now %d out of %d\n", currentPosition, footerSize)

		labelSizeBytes := f.data[currentPosition:(currentPosition + shiftAmount)]
		labelSizeBytes = append([]byte{0}, labelSizeBytes[0])
		// fmt.Printf("    4. Reading labelSizeBytes amount: % x \n", labelSizeBytes)
		labelSize := binary.BigEndian.Uint16(labelSizeBytes)
		f.labels = append(f.labels, []byte{})

		currentPosition += shiftAmount
		// fmt.Printf("    5. Current position now %d out of %d\n", currentPosition, footerSize)

		shiftAmount = labelSize
		label := f.data[currentPosition:(currentPosition + shiftAmount)]
		// fmt.Printf("    6. Shifting through %v bytes: % x · %s\n", shiftAmount, label, string(label))

		f.labels = append(f.labels, label)

		currentPosition += shiftAmount
		// fmt.Printf("    7. Current position now %d out of %d\n", currentPosition, footerSize)

	}

	// fmt.Printf("  · Found %d labels:\n", len(f.labels))
	// for idx, label := range f.labels {
	// 	fmt.Printf("    · %d ->  %s\n", idx, label)
	// }
}
