package downloader

func trimFileName(url string) string {
	r := []rune(url)
	var index int
	if r[len(r)-1] == '/' {
		for i := len(r) - 2; i > 0; i-- {
			if r[i] == '/' {
				index = i + 1
				break
			}
		}
		return string(r[index : len(r)-1])
	}
	for i := len(r) - 1; i > 0; i-- {
		if r[i] == '/' {
			index = i + 1
			break
		}
	}
	return string(r[index:])
}
