package runes

import (
	"bytes"
	"fmt"
	"testing"
)

func TestByteWriterWrite(t *testing.T) {
	runes := []rune("Hello ✈ Writer")
	expectedBytes := []byte(string(runes))
	expectedLen := len(expectedBytes)
	expectedErr := error(nil)

	underlyingWriter := bytes.Buffer{}
	w := Writer(ByteWriter{&underlyingWriter})
	n, err := w.Write(runes)
	bytesWritten := underlyingWriter.Bytes()

	if err != expectedErr {
		t.Fatalf("ByteWriter.Write(%#v) = %d, %#v, expected %d, %#v", runes, n, err, expectedLen, expectedErr)
	}
	if n != expectedLen {
		t.Fatalf("ByteWriter.Write(%#v) = %d, %#v, expected %d, %#v", runes, n, err, expectedLen, expectedErr)
	}
	if !bytes.Equal(bytesWritten, expectedBytes) {
		t.Fatalf("ByteWriter.Write(%#v) wrote %#v, expected %#v", runes, bytesWritten, expectedBytes)
	}
	t.Logf("ByteWriter.Write(%#v) = %d, %#v", runes, n, err)
}

type WriterThatErrors struct{}

func (w WriterThatErrors) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("Error occurred while writing")
}

func TestByteWriterWrite_Error(t *testing.T) {
	runes := []rune("Hello ✈ Writer")
	expectedErr := fmt.Errorf("Error occurred while writing")

	underlyingWriter := WriterThatErrors{}
	w := Writer(ByteWriter{&underlyingWriter})
	_, err := w.Write(runes)

	if err == nil {
		t.Fatalf("ByteWriter.Write(%#v) = _, %#v, expected _, %#v", runes, err, expectedErr)
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("ByteWriter.Write(%#v) = _, %#v, expected _, %#v", runes, err, expectedErr)
	}
	t.Logf("ByteWriter.Write(%#v) = _, %#v", runes, err)
}
