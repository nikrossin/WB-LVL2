package tests

import (
	wget "lvl2/develop/dev09/internal"
	"net/url"
	"os"
	"testing"
)

func TestMakeBasesDir(t *testing.T) {
	uri, _ := url.Parse("http://domaintest.ru")
	c1 := &wget.Config{
		Domain: uri,
		Dir:    "dirtest",
	}
	c2 := &wget.Config{
		Domain: uri, // not created
		Dir:    "dirtest2",
	}
	c3 := &wget.Config{
		Domain: uri,        // not created
		Dir:    "dirtest3", //not created
	}
	wg := &wget.Wget{}

	wg.Config = c1
	if err := wg.MakeBasesDir(); err != nil {
		t.Errorf("Incorrect get error with %v %v", wg.Dir, err)
	}
	os.Chdir("..")
	wg.Config = c2
	if err := wg.MakeBasesDir(); err != nil {
		t.Errorf("Incorrect get error with %v %v", wg.Dir, err)
	}
	os.Chdir("..")
	wg.Config = c3
	if err := wg.MakeBasesDir(); err != nil {
		t.Errorf("Incorrect get error with %v %v", wg.Dir, err)
	}

}

func TestIsLinkWithParameters(t *testing.T) {
	wg := new(wget.Wget)
	testURL := []string{
		"http://test.ru/one",
		"http://test.ru/one/two",
		"http://test.ru/one?q=5",
		"http://test.ru/one?q=5&b=6",
	}
	if wg.IsLinkWithParameters(testURL[0]) {
		t.Errorf("Incorrect parse link for params %v", testURL[0])
	}
	if wg.IsLinkWithParameters(testURL[1]) {
		t.Errorf("Incorrect parse link for params %v", testURL[1])
	}
	if !wg.IsLinkWithParameters(testURL[2]) {
		t.Errorf("Incorrect parse link for params %v", testURL[2])
	}
	if !wg.IsLinkWithParameters(testURL[3]) {
		t.Errorf("Incorrect parse link for params %v", testURL[3])
	}
}

func TestGetPathsToSave(t *testing.T) {
	uri, _ := url.Parse("http://domaintest.ru")
	c := &wget.Config{
		Domain: uri,
	}
	wg := &wget.Wget{}
	wg.Config = c

	urlPath := "/test"
	pathD, pathF := wg.GetPathsToSave(urlPath)
	if pathF != "domaintest.ru/test/index.html" || pathD != "domaintest.ru/test" {
		t.Errorf("Incorrect make Paths for domain %v and path %v\n", wg.Domain.Host, urlPath)
	}

	urlPath = "/test/"
	pathD, pathF = wg.GetPathsToSave(urlPath)
	if pathF != "domaintest.ru/test/index.html" || pathD != "domaintest.ru/test" {
		t.Errorf("Incorrect make Paths for domain %v and path %v\n", wg.Domain.Host, urlPath)
	}

	urlPath = "/test/login.php"
	pathD, pathF = wg.GetPathsToSave(urlPath)
	if pathF != "domaintest.ru/test/login.php" || pathD != "domaintest.ru/test" {
		t.Errorf("Incorrect make Paths for domain %v and path %v\n", wg.Domain.Host, urlPath)
	}

	urlPath = "/test/cat/game.html"
	pathD, pathF = wg.GetPathsToSave(urlPath)
	if pathF != "domaintest.ru/test/cat/game.html" || pathD != "domaintest.ru/test/cat" {
		t.Errorf("Incorrect make Paths for domain %v and path %v\n", wg.Domain.Host, urlPath)
	}
}
