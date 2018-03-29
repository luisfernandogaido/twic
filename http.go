package main

import (
	"net/http"
	"io/ioutil"
	"regexp"
)

const (
	site = "http://theweekinchess.com/twic"
)

func listaZipsPng() ([]string, error) {
	res, err := http.Get(site)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	html := string(bytes)
	re, err := regexp.Compile(`<td><a href="([^"]+)">PGN</a></td>`)
	if err != nil {
		return nil, err
	}
	zips := make([]string, 0)
	matches := re.FindAllStringSubmatch(html, -1)
	for _, m := range matches {
		zips = append(zips, m[1])
	}
	return zips, nil
}

func downloadZip(z string) ([]byte, error) {
	res, err := http.Get(z)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
