package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type metadata struct {
	Version   int       `json:"version"`
	Site      string    `json:"site"`
	NumLinks  int       `json:"num_links"`
	Images    int       `json:"images"`
	LastFetch time.Time `json:"last_fetch"`
}

func (m metadata) ToString() string {
	return fmt.Sprintf("site: %s\nnum_links: %d\nimages: %d\nlast_fetch: %s",
		m.Site,
		m.NumLinks,
		m.Images,
		m.LastFetch.Format("Mon Jan 02 2006 15:04 MST "))
}

func ToFileBasename(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	return u.Host + strings.TrimSuffix(u.EscapedPath(), "/")
}

func AddFileExtension(b, e string) string {
	return b + "." + e
}

func ToHtmlFilename(s string) string {
	return AddFileExtension(s, "html")
}

func ToMetadataFilename(s string) string {
	return ".metadata/" + AddFileExtension(s, "json")
}

func ToMetadata(b []byte, s string) metadata {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(b))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	numLink := doc.Find("a").Length()
	numImage := doc.Find("img").Length()
	return metadata{1, s, numLink, numImage, time.Now()}
}

func PrintMetadata(s string) {
	basename := ToFileBasename(s)
	metafile := ToMetadataFilename(basename)
	b, err := ioutil.ReadFile(metafile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	m := &metadata{}
	err = json.Unmarshal(b, m)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(m.ToString())
}

func PrintMetadatas(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: fetch --metadata URLs")
		os.Exit(1)
	}
	for i := 0; i < len(args); i++ {
		PrintMetadata(args[i])
		if i != len(args)-1 {
			fmt.Println("-------")
		}
	}
}

func FetchURL(u string) ([]byte, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func FetchURLs(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: fetch URLs")
		os.Exit(1)
	}
	for i := 0; i < len(args); i++ {
		body, err := FetchURL(args[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		basename := ToFileBasename(args[i])
		err = os.WriteFile(ToHtmlFilename(basename), body, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		m := ToMetadata(body, basename)
		b, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(ToMetadataFilename(basename), b, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	var isMeta bool = false
	var i int = 1
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: fetch URLs")
		os.Exit(1)
	}
	if os.Args[1] == "--metadata" {
		isMeta = true
		i++
	}
	err := os.Mkdir(".metadata", os.ModeDir|os.ModePerm)
	if err != nil && !os.IsExist(err) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if isMeta {
		PrintMetadatas(os.Args[i:])
	} else {
		FetchURLs(os.Args[i:])
	}
}
