package main

import (
	"fmt"
	"os"
	"testing"
)

type Test struct {
	Name        TestName
	FileContent TestFileContent
}
type TestName string
type TestFileContent []byte

func (t Test) FileName(artifactType string) string {
	return fmt.Sprintf("./tests/%s.%s.bin", t.Name, artifactType)
}

func (t Test) WriteTestResultFile(artifactType string, content TestFileContent) {
	fileName := t.FileName(artifactType)
	os.WriteFile(fileName, content, 0644)
}

func (t Test) Failed() bool {
	return string(t.Expected()) != string(t.Actual())
}

func (t *Test) Input() TestFileContent {
	content := t.readTestFile("input")

	if t.FileContent == nil {
		t.FileContent = content
	}

	return content
}

func (t Test) Expected() TestFileContent {
	return t.readTestFile("expected")
}
func (t Test) Actual() TestFileContent {
	return t.readTestFile("actual")
}

func (t Test) readTestFile(artifactType string) TestFileContent {
	fileName := t.FileName(artifactType)
	content, _ := os.ReadFile(fileName)

	return content
}

func TestTransformFile(t *testing.T) {
	subject := func(test Test) TestFileContent {
		bas := AmigaBasicFile{
			Name:       string(test.Name),
			BinaryData: test.Input(),
		}

		bas.Parse()

		test.WriteTestResultFile("actual", bas.body.data)
		return bas.body.data
	}

	t.Run("Null", func(t *testing.T) {
		test := Test{Name: "Null"}
		actual := subject(test)

		if test.Failed() {
			t.Fatalf("expected:\n%s\n\nactual:\n%s\n\ninput:\n% x\n\noutput:\n% x",
				test.Expected(), actual, test.Input(), actual)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		test := Test{Name: "Empty"}
		actual := subject(test)

		if test.Failed() {
			t.Fatalf("expected:\n%s\n\nactual:\n%s\n\ninput:\n% x\n\noutput:\n% x",
				test.Expected(), actual, test.Input(), actual)
		}
	})

	t.Run("Hello World", func(t *testing.T) {
		test := Test{Name: "Hello World"}
		actual := subject(test)

		if test.Failed() {
			t.Fatalf("expected:\n%s\n\nactual:\n%s\n\ninput:\n% x\n\noutput:\n% x",
				test.Expected(), actual, test.Input(), actual)
		}
	})

	t.Run("Hello World, Twice", func(t *testing.T) {
		test := Test{Name: "Hello World, Twice"}
		actual := subject(test)

		if test.Failed() {
			t.Fatalf("expected:\n%s\n\nactual:\n%s\n\ninput:\n% x\n\noutput:\n% x",
				test.Expected(), actual, test.Input(), actual)
		}
	})

	t.Run("Comment via REM", func(t *testing.T) {
		t.Skip("TODO")

		test := Test{Name: "Comment via REM"}
		actual := subject(test)

		if test.Failed() {
			t.Fatalf("expected:\n%s\n\nactual:\n%s\n\ninput:\n% x\n\noutput:\n% x",
				test.Expected(), actual, test.Input(), actual)
		}
	})

	t.Run("Comment via Apostrophe", func(t *testing.T) {
		t.Skip("TODO")

		test := Test{Name: "Comment via Apostrophe"}
		actual := subject(test)

		if test.Failed() {
			t.Fatalf("expected:\n%s\n\nactual:\n%s\n\ninput:\n% x\n\noutput:\n% x",
				test.Expected(), actual, test.Input(), actual)
		}
	})

	t.Run("Input", func(t *testing.T) {
		test := Test{Name: "Input"}
		actual := subject(test)

		if test.Failed() {
			t.Fatalf("expected:\n%s\n\nactual:\n%s\n\ninput:\n% x\n\noutput:\n% x",
				test.Expected(), actual, test.Input(), actual)
		}
	})

	t.Run("Inputs", func(t *testing.T) {
		test := Test{Name: "Inputs"}
		actual := subject(test)

		if test.Failed() {
			t.Fatalf("expected:\n%s\n\nactual:\n%s\n\ninput:\n% x\n\noutput:\n% x",
				test.Expected(), actual, test.Input(), actual)
		}
	})

	t.Run("Subroutine", func(t *testing.T) {
		test := Test{Name: "Subroutine"}
		actual := subject(test)

		if test.Failed() {
			t.Fatalf("expected:\n%s\n\nactual:\n%s\n\ninput:\n% x\n\noutput:\n% x",
				test.Expected(), actual, test.Input(), actual)
		}
	})

	t.Run("Subroutine and Input", func(t *testing.T) {
		test := Test{Name: "Subroutine and Input"}
		actual := subject(test)

		if test.Failed() {
			t.Fatalf("expected:\n%s\n\nactual:\n%s\n\ninput:\n% x\n\noutput:\n% x",
				test.Expected(), actual, test.Input(), actual)
		}
	})
}
