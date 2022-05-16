package siteMap

import (
	"encoding/json"
	"io/ioutil"
)

type SiteMap struct {
	Pages []string `json:"pages"`
}

func NewSiteMap() *SiteMap {
	return &SiteMap{
		Pages: make([]string, 0),
	}
}

func (s *SiteMap) Fill(links []string) {
	for _, v := range links {
		s.Pages = append(s.Pages, v)
	}
}

func (s *SiteMap) Save() error {
	site, err := json.Marshal(s)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("wget/sitemap.json", site, 0644)
	if err != nil {
		return err
	}
	return nil
}
