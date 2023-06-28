package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/opesun/goquery"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком
*/

func wget(url, dirName string) {
	dir := createDir(url, dirName)

	getHtml(url, dir)

	getFiles(url, dir)
}

func createDir(url, dirName string) string {
	if dirName == "" {
		str := strings.Split(url, "/")
		return str[2]
	} else {
		return dirName
	}
}

func getHtml(url, dirName string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	str := strings.Split(url, "/")

	fileName := str[len(str)-1]
	if fileName == "" {
		fileName = str[2]
	}

	if err := os.MkdirAll(dirName, 0744); err != nil {
		panic(err)
	}

	var name string
	if strings.Contains(fileName, ".html") {
		name = filepath.Join(dirName, fileName)
	} else {
		name = filepath.Join(dirName, fileName+".html")
	}

	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func getFiles(url, dirName string) {
	x, err := goquery.ParseUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	n := x.Find("")
	attrs := n.Attrs("href")

	for _, u := range attrs {
		switch {
		case strings.Contains(u, ".css"):
			str := strings.Split(u, "/")
			dR(dirName, str[len(str)-1], u)
		case strings.Contains(u, ".png"):
			str := strings.Split(u, "/")
			dR(dirName, str[len(str)-1], u)
		case strings.Contains(u, ".jpg"):
			str := strings.Split(u, "/")
			dR(dirName, str[len(str)-1], u)
		}
	}
}

func dR(dirName, fileName string, url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	name := filepath.Join(dirName, fileName)

	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dirName := flag.String("d", "", "directory name")

	flag.Parse()

	if ok, err := regexp.MatchString("^(http|https)://", flag.Arg(0)); ok && err == nil {
		wget(flag.Arg(0), *dirName)
	} else {
		log.Fatal("invalid url")
	}
}
