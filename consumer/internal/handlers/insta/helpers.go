package insta

import (
	"log"
	"net/url"
	"regexp"
)

func getPostType(rawUrl string) (PostType, bool) {
	url, err := url.Parse(rawUrl)
	if err != nil {
		log.Printf("bad raw insta url provided: %s", rawUrl)
		return 0, false
	}
	patterns := map[PostType]*regexp.Regexp{
		Post:  regexp.MustCompile(`^/p/[a-zA-Z0-9_-]+/?$`),
		Story: regexp.MustCompile(`^/stories/[a-zA-Z0-9._]{1,30}/[0-9]+/?$`),
		Reel:  regexp.MustCompile(`^/reel/[a-zA-Z0-9_-]+/?$`),
	}

	// Check against each pattern
	for urlType, pattern := range patterns {
		if pattern.MatchString(url.Path) {
			return urlType, true
		}
	}

	return 0, false
}
