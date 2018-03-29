package main

import (
	"log"
	"bytes"
	"archive/zip"
	"io/ioutil"
	"os"
)

const (
	arquivoSaida = "./out/twic.pgn"
)

func main() {
	zips, err := listaZipsPng()
	if err != nil {
		log.Fatal(err)
	}
	out, err := os.OpenFile(arquivoSaida, os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	for _, z := range zips {
		data, err := downloadZip(z)
		if err != nil {
			log.Fatal(err)
		}
		reader := bytes.NewReader(data)
		zr, err := zip.NewReader(reader, int64(len(data)))
		if err != nil {
			log.Fatal(err)
		}
		rc, err := zr.File[0].Open()
		if err != nil {
			log.Fatal(err)
		}
		bytes, err := ioutil.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}
		_, err = out.Write(bytes)
		if err != nil {
			log.Fatal(err)
		}
	}
}
