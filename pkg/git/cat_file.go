package git

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CatFileOptions struct {
	ShowSHAContent *string
}

func CatFile(opts CatFileOptions) error {
	if opts.ShowSHAContent != nil {
		blobSha := *opts.ShowSHAContent

		if blobSha == "" {
			return errors.New("blob hash is empty")
		}

		header := blobSha[:2]
		hash := blobSha[2:]

		fileContent, err := os.ReadFile(".git/objects/" + header + "/" + hash)
		if err != nil {
			return err
		}

		zlibReader, err := zlib.NewReader(bytes.NewReader(fileContent))
		if err != nil {
			return err
		}

		buf := bufio.NewReader(zlibReader)

		objectType, err := buf.ReadString(' ')
		if err != nil {
			return err
		}

		log.Println("object type:", objectType[:len(objectType)-1])

		contentSizeStr, err := buf.ReadString('\u0000')
		if err != nil {
			return err
		}

		contentSize, err := strconv.Atoi(contentSizeStr[:len(contentSizeStr)-1])
		if err != nil {
			return err
		}

		log.Println("content size:", contentSize)

		if contentSize != buf.Buffered() {
			return errors.New("content size doesn't match")
		}

		content := make([]byte, contentSize)
		n, err := buf.Read(content)
		if err != nil {
			return err
		}
		if contentSize != n {
			return errors.New("content not fully read")
		}

		fmt.Print(string(content))

		err = zlibReader.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
