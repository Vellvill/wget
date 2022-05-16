package scrapper

import "strings"

type helper struct {
	noturl map[string]struct{}
}

func newHelper() helper {
	return helper{
		noturl: map[string]struct{}{
			"skype":      {},
			"mailto":     {},
			"javascript": {},
			"tel":        {},
			"sms":        {},
			"market":     {},
			"whatsapp":   {},
		}}
}

func (h *helper) isFalseUrl(link string) bool {
	format := link[:strings.IndexByte(link, ':')]
	if _, ok := h.noturl[format]; ok {
		return false
	}
	return true
}

func (h *helper) sanitizeUrl(link string) string {
	if h.isFalseUrl(link) {
		link = strings.TrimSpace(link)

		link = strings.Split(link, "?")[0]
		/*
			if string(link[len(link)-1]) != "/" {
				link = link + "/"
			}
		*/
		return link
	}
	return ""
}
