package main

import (
	"io"
	"os"
)

func copy(sourceFile string, destFile string) {
	javaSrc, errJavaSrc := os.Open(sourceFile)
	checkError(errJavaSrc)
	defer javaSrc.Close()

	JavaDst, errJavaDst := os.Create(destFile)
	checkError(errJavaDst)
	defer JavaDst.Close()

	_, errCopy := io.Copy(JavaDst, javaSrc)
	checkError(errCopy)
}

func open(file string) *os.File {
	jsonFile, err := os.Open(file)
	checkError(err)

	return jsonFile
}
