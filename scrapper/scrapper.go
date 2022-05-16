package scrapper

import (
	"golang.org/x/net/html"
	"net/http"
)

type Scrapper struct {
	Helper   helper
	doc      *html.Node
	resp     *http.Response
	url      string
	Links    []string
	HrefsURL []string
}

func NewScrapper(doc *html.Node, url string, resp *http.Response) *Scrapper {
	return &Scrapper{
		Helper: newHelper(),
		doc:    doc,
		resp:   resp,
		url:    url,
	}
}

func (s *Scrapper) Scrap() {
	var f func(node *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := s.resp.Request.URL.Parse(a.Val)
					if err == nil {
						fl := s.Helper.sanitizeUrl(link.String())
						s.HrefsURL = append(s.HrefsURL, fl)
					}
				}
			}
		} else if n.Type == html.ElementNode && n.Data == "link" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := s.resp.Request.URL.Parse(a.Val)
					if err == nil {
						fl := s.Helper.sanitizeUrl(link.String())
						s.Links = append(s.Links, fl)
					}
				}
			}
		} else if n.Type == html.ElementNode && n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					link, err := s.resp.Request.URL.Parse(a.Val)
					if err == nil {
						fl := s.Helper.sanitizeUrl(link.String())
						s.Links = append(s.Links, fl)
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(s.doc)
}
