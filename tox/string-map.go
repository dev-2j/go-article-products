package tox

import (
	"net/url"
	"strings"
)

// แปลง string ที่ได้จาก body.text ให้เป็น map
func StringMap(s string) map[string]string {
	items := map[string]string{}
	for _, v := range strings.Split(s, `&`) {
		kv := strings.Split(v, `=`)
		if len(kv) > 1 {
			items[kv[0]], _ = url.PathUnescape(kv[1])
		}
	}
	return items
}

func StringMap2(s string) map[string]string {
	items := map[string]string{}
	for _, v := range strings.Split(s, "\n") {
		kv := strings.Split(v, `=`)
		if len(kv) > 1 {
			items[kv[0]], _ = url.PathUnescape(kv[1])
		}
	}
	return items
}
