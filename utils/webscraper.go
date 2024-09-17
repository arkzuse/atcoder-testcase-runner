package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func ScrapeTestcase(contest string, task string) [][]string {
	res, err := http.Get(fmt.Sprintf("https://atcoder.jp/contests/%s/tasks/%s_%s", contest, contest, task))
	if err != nil {
		return nil
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	enSpan := doc.Find("span .lang-en")
	// mutable array of 2d string
	var samples [][]string
	var t = 1
	var p = 0

	enSpan.Children().Each(func(i int, s *goquery.Selection) {
		// check if element has <section> as child then print it
		if s.Has("section").Length() > 0 {
			section := s.Find("section")

			// check if section has <h3> as child
			if section.Has("h3").Length() > 0 {
				h3 := section.Find("h3")

				if h3.Text() == fmt.Sprintf("Sample %s %d", []string{"Input", "Output"}[p], t) {
					if p == 1 {
						samples[len(samples)-1] = append(samples[len(samples)-1], section.Find("pre").Text())
						t++
					} else {
						samples = append(samples, []string{section.Find("pre").Text()})
					}

					p++
					p %= 2
				}
			}
		}
	})

	return samples
}
