package commands

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

func HashObject() error {
	filenameToRead := os.Args[3]

	fileContent, err := os.ReadFile(filenameToRead)
	if err != nil {
		return err
	}

	content := []byte(fmt.Sprintf("blob %d%b%s", len(fileContent), '\u0000', fileContent))

	zlibBuf := new(bytes.Buffer)

	w := zlib.NewWriter(zlibBuf)
	_, err = w.Write(content)
	if err != nil {
		return err
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	hash := sha1.Sum(zlibBuf.Bytes())
	hexHash := hex.EncodeToString(hash[:])

	dir := hexHash[:2]
	filename := hexHash[2:]

	err = os.Mkdir(".git/objects/"+dir, 0666)
	if err != nil {
		return err
	}

	f, err := os.Create(".git/objects/" + dir + "/" + filename)
	if err != nil {
		return err
	}

	_, err = f.Write(zlibBuf.Bytes())
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	fmt.Print(hexHash)

	return nil
}
