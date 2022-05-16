package main

import (
	"flag"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"wget/downloader"
	"wget/scrapper"
	"wget/siteMap"
)

var (
	t int
	d string
)

func init() {
	flag.StringVar(&d, "d", "", "set your url")
	flag.IntVar(&t, "t", 10, "set timeout to connect to url")
}

func main() {
	flag.Parse()

	client := &http.Client{
		Timeout: time.Duration(t) * time.Second,
	}

	resp, err := client.Get(d)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := scrapper.NewScrapper(doc, strings.TrimSuffix(d, "/"), resp)
	s.Scrap()

	err = os.Mkdir("wget", 0755)
	if err != nil {
		log.Fatal(err)
	}

	siteM := siteMap.NewSiteMap()
	siteM.Fill(s.HrefsURL)
	err = siteM.Save()
	if err != nil {
		log.Fatal(err)
	}

	down := downloader.NewDownloader()
	down.Download(s.Links)
}
