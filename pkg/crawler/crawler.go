package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Tuna_press() chan string {

	c := colly.NewCollector()
	ch := make(chan string)

	c.OnHTML("a[class=featured-thumbnail]", func(e *colly.HTMLElement) {

		go func() {

			ch <- e.Attr("href")
		}()
	})

	c.Visit("https://tuna.press/")
	return ch
}

func Business_next() chan string {

	c := colly.NewCollector()
	ch := make(chan string)

	c.OnHTML("a.item_img", func(e *colly.HTMLElement) {

		go func() {

			ch <- e.Attr("href")
		}()
	})

	c.Visit("https://www.bnext.com.tw")
	return ch
}

func Meet_bnext() chan string {

	c := colly.NewCollector()
	ch := make(chan string)

	c.OnHTML("a.item_img_link", func(e *colly.HTMLElement) {

		go func() {

			ch <- e.Attr("href")
			fmt.Println(e.Attr("href"))
		}()
	})

	c.Visit("https://meet.bnext.com.tw/")
	return ch
}
