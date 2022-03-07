package wget

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"path"
	"strings"
)

type Wget struct {
	*Config
	Pages     map[string]bool
	Collector *colly.Collector
}

func NewWget() *Wget {

	return &Wget{
		Config: &Config{},
		Pages:  make(map[string]bool),
	}
}

func (w *Wget) InitConfig() {
	w.Config.Init()
	w.Collector = colly.NewCollector(colly.AllowedDomains(w.Domain.Host))
}

func (w *Wget) Run() {
	if err := w.MakeBasesDir(); err != nil {
		log.Fatalln(err)
	}
	w.ParseLinks()
	w.SaveFiles()

	w.Pages[w.Domain.String()] = true
	if err := w.Collector.Visit(w.Domain.String()); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Site %v is downloaded\n", w.Domain.String())
}

func (w *Wget) MakeBasesDir() error {
	if _, err := os.Stat(w.Dir); err != nil {
		if err := os.Mkdir(w.Dir, os.ModePerm); err != nil {
			return err
		}
	}
	if err := os.Chdir(w.Dir); err != nil {
		return err
	}
	if _, err := os.Stat(w.Domain.Host); err != nil {
		if err := os.Mkdir(w.Domain.Host, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func (w *Wget) IsLinkWithParameters(url string) bool {
	if strings.ContainsAny(url, "?=&") {
		return true
	}
	return false
}

func (w *Wget) ParseLinks() {
	w.Collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		absLink := e.Request.AbsoluteURL(link)
		if !w.Pages[absLink] && !w.IsLinkWithParameters(absLink) {
			if err := w.Collector.Visit(absLink); err != nil {
				log.Printf("%v : %v\n", err, absLink)
			}
			w.Pages[absLink] = true
		}
	})
}

func (w *Wget) GetPathsToSave(urlPath string) (pathDir string, pathFile string) {
	pathFile = w.Domain.Host + urlPath
	pathDir = pathFile

	if path.Ext(urlPath) == "" {
		if pathFile[len(pathFile)-1] != '/' {
			pathFile += "/"
		}
		pathFile += "index.html"
	}
	if index := strings.LastIndex(pathFile, "/"); index != -1 {
		pathDir = pathFile[:index]
	}
	return
}
func (w *Wget) SaveFiles() {
	w.Collector.OnResponse(func(r *colly.Response) {
		pathDir, pathFile := w.GetPathsToSave(r.Request.URL.Path)

		if _, err := os.Stat(pathDir); err != nil {
			if err := os.MkdirAll(pathDir, os.ModePerm); err != nil {
				log.Fatalln(err, 1)
			}
		}
		if err := r.Save(pathFile); err != nil {
			log.Fatalln(err, pathFile)
		}
	})

}
