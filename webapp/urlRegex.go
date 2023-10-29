package webapp

import (
	"fmt"
	"regexp"
)

func ParseUrl() {

	regx := regexp.MustCompile(`:([^/]+)`)

	str := "/users/:userId/blogs/:blogId"
	url := "/users/123/blogs/456"
	newStr := regx.ReplaceAllString(str, "([a-zA-Z0-9]+)")

	urlRegex := regexp.MustCompile(newStr)

	g1 := regx.FindAllStringSubmatch(str, -1)
	g2 := urlRegex.FindStringSubmatch(url)

	fmt.Printf("g1 : %v \n g2: %v", g1, g2)
}
