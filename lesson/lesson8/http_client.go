package main

import (
	"archive/tar"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"compress/gzip"
	"github.com/PuerkitoBio/goquery"
)

func CleanUrl(uri *url.URL, link string) string {
	switch {
	case strings.HasPrefix(link, "https") || strings.HasPrefix(link, "http"):
		return link
	case strings.HasPrefix(link, "//"):
		return uri.Scheme + ":" + link
	case strings.HasPrefix(link, "/"):
		return fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, link)
	default:
		p := strings.SplitAfter(uri.Path, "/")
		path := strings.Join(p[:2], "") //一般情况是这样 ,/static/img/logo.png
		return fmt.Sprintf("%s://%s%s%s", uri.Scheme, uri.Host, path, link)
	}
}

func cleanUrls(u string, urls []string) []string {
	var ret []string
	uri, _ := url.Parse(u)
	for i := range urls {
		ret = append(ret, CleanUrl(uri, urls[i]))
	}
	return ret
}

func downloadImg(urls []string, dir string) error {
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return errors.New(resp.Status)
		}
		fullname := filepath.Join(dir, filepath.Base(url))
		f, err := os.Create(fullname)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer f.Close()
		io.Copy(f, resp.Body)
	}
	return nil
}

func fetch(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("src")
		//fmt.Println(link)
		urls = append(urls, link)

	})
	return urls, nil
}

func maketar(dir string, w io.Writer) error {
	basedir := filepath.Base(dir)
	compress := gzip.NewWriter(w)
	defer compress.Close()
	tr := tar.NewWriter(w)
	defer tr.Close()
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		p, _ := filepath.Rel(dir, path)
		header.Name = filepath.Join(basedir, p)
		tr.WriteHeader(header)
		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		io.Copy(tr, f)
		return nil
	})
	return nil
}

func main() {
	//url := "http://daily.zhihu.com/"
	url := os.Args[1]
	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	urls = cleanUrls(url, urls)
	for _, u := range urls {
		fmt.Println(u)
	}
	tmpdir, err := ioutil.TempDir("", "spider")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(tmpdir)
	//defer os.RemoveAll(tmpdir)
	err = downloadImg(urls, tmpdir)
	if err != nil {
		log.Panic(err)
	}

	if f, err := os.Create("img.tar.gz"); err == nil {
		maketar(tmpdir, f)
	}

}
