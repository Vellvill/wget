package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

//Downloader ...
type Downloader struct {
}

//NewDownloader ...
func NewDownloader() *Downloader {
	return &Downloader{}
}

//Download ...
func (d *Downloader) Download(urls []string) {
	if len(urls) == 0 {
		return
	}

	err := os.Mkdir("wget/files", 0755)
	if err != nil {
		log.Fatal(err)
	}

	f := func(url string) error {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return err
		}
		name := trimFileName(url)

		f, err := os.Create(fmt.Sprintf("wget/files/%s", name))
		if err != nil {
			return err
		}
		defer func() {
			err = f.Close()
			if err != nil {
				panic(err)
			}
		}()

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			return err
		}
		return nil
	}

	log.Println("Starting downloading")
	for i, url := range urls {
		fmt.Printf("%.0f%s\n", float64(float64(i+1)/float64(len(urls)))*100, "%")
		err := f(url)
		if err != nil {
			fmt.Println(err)
		}
	}
	log.Println("Downloading finished")
}
