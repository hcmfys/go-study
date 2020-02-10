package main

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	pref   = "https://www.2717.com/ent/meinvtupian/2019/"
	exitCh chan int
	done   int
	total  int
)

func getBody(url string) (io.ReadCloser, error) {

	resp, e := http.Get(url)

	if e != nil {
		fmt.Println("获取页面失败", e.Error())
		return nil, errors.New("获取页面失败")
	}

	return resp.Body, nil
}

func getDoc(url string) (*goquery.Document, error) {

	resp, e := getBody(url)
	defer resp.Close()

	doc, e := goquery.NewDocumentFromReader(resp)
	if e != nil {
		fmt.Println("获取页面doc失败", e.Error())
		return nil, errors.New("获取页面doc失败")
	}
	return doc, nil
}

func getPage(url string) {
	exitCh <- 0
	done = 0
	doc, e := getDoc(url)
	if e != nil {
		panic(e)
	}

	lst := list.New()
	doc.Find(".articleV4Page >li").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		href, _ := a.Attr("href")

		if strings.Index(href, ".html") > 0 {
			url := pref + href
			lst.PushBack(url)
		}
		fmt.Println(href)
	})

	total = lst.Len()

	for i := lst.Front(); i != nil; i = i.Next() {
		url := i.Value.(string)
		go parseImage(url)
	}
}

func parseImage(url string) {
	doc, e := getDoc(url)
	if e != nil {
		panic(e)
	}

	imgs := doc.Find(".articleV4Body img")
	imgs.Each(func(_ int, s *goquery.Selection) {
		// For each item found, get the band and title
		img := s.Nodes[0]
		src := img.Attr[1].Val
		go saveImg(src)

	})
	fmt.Println("get images")
	fmt.Println(imgs.Nodes)

}

func saveImg(url string) {
	fmt.Println("get img", url)
	todir := "./out/imgs"
	os.MkdirAll(todir, os.ModePerm)
	index := strings.LastIndex(url, "/")
	extName := url
	if index > 0 {
		extName = url[index+1:]
	}
	extName = strings.ReplaceAll(extName, "/", "-")
	f, e := os.Create(todir + "/" + extName)
	defer f.Close()
	if e != nil {
		panic(e)
	}
	body, e := getBody(url)

	if e != nil {
		fmt.Println(" save img error", url)
		done++

	}
	data, e := ioutil.ReadAll(body)
	if e != nil {
		fmt.Println(" write img error", url)
		done++
		return
	}
	done++
	defer body.Close()
	f.Write(data)
	fmt.Println("done ==>", done, "of ", total)

	if done >= total {
		exitCh <- 100
	}

}
func main() {
	exitCh = make(chan int, 100)
	url := pref + "314973.html"
	//https://www.2717.com/ent/meinvtupian/2017/213772.html
	//pref = "https://www.2717.com/ent/meinvtupian/2017/"
	//url = pref + "213772.html"

	//pref = "https://www.2717.com/ent/meinvtupian/2016/"
	//url = pref + "168724.html"
	go getPage(url)

	for {
		i := <-exitCh
		if i == 0 {
			fmt.Println("\r start downloading ....")
		} else if i > 0 {
			break
		}
	}

}
