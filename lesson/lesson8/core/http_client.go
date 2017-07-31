package main

import (
	"github.com/PuerkitoBio/goquery"

	"archive/tar"
	"compress/gzip"
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
)

// 清洗URL
func clearUrls(u string, urls []string) ([]string, error) {
	var ret []string
	uri, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	for _, link := range urls {
		switch {
		case strings.HasPrefix(link, "https") || strings.HasPrefix(link, "https"):
			ret = append(ret, link)
		case strings.HasPrefix(link, "//"):
			ret = append(ret, uri.Scheme+":"+link)
		case strings.HasPrefix(link, "/"):
			ret = append(ret, fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, link))
		default:
			p := strings.SplitAfter(uri.Path, "/")
			path := strings.Join(p[:2], "")
			ret = append(ret, fmt.Sprintf("%s://%s%s%s", uri.Scheme, uri.Host, path, link))
		}
	}
	return ret, nil
}

// 抓取URL
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
	//io.Copy(os.Stdout, resp.Body)
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	// Find the review items
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		if link, ok := s.Attr("src"); ok {
			urls = append(urls, link)
		}
	})
	return urls, nil

}

// 下载URL
func downloadImg(urls []string, dir string) error {
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return errors.New(resp.Status)
		}
		////io.Copy(os.Stdout, resp.Body)
		fullpath := filepath.Join(dir, filepath.Base(u))
		f, err := os.Create(fullpath)
		if err != nil {
			return err
		}
		defer f.Close()
		io.Copy(f, resp.Body)

	}
	return nil
}

func maketar(dir string, w io.Writer) error {
	basedir := filepath.Base(dir)
	compress := gzip.NewWriter(w)
	defer compress.Close()
	tr := tar.NewWriter(compress)
	defer tr.Close()
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		handler, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		p, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		handler.Name = filepath.Join(basedir, p)
		tr.WriteHeader(handler)
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
	if err != nil {
		return err
	}
	return nil
}

func fetchImgs(url string) {
	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	n_urls, err := clearUrls(url, urls)
	if err != nil {
		log.Fatal(err)
	}

	tmpdir, err := ioutil.TempDir("", "spider")
	if err != nil {
		log.Panic(err)
	}
	err = downloadImg(n_urls, tmpdir)
	if err != nil {
		log.Println(err)
	}
	f, err := os.Create("img.tar.gz")
	if err != nil {
		log.Fatal(err)
	}
	maketar(tmpdir, f)
}

func main() {
	u := "https://daily.zhihu.com/"
	fetchImgs(u)
}
