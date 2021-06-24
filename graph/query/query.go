package query

import (
	"net/url"
	"strconv"
	"strings"
)

type QueryOption func(url url.URL) url.URL

func Top(count int) QueryOption {
	return func(u url.URL) url.URL {
		queryString, _ := url.ParseQuery(u.RawQuery)
		queryString.Add("$top", strconv.Itoa(count))
		u.RawQuery = queryString.Encode()
		return u
	}
}

func Version(version string) QueryOption {
	return func(u url.URL) url.URL {
		u.Path = strings.Replace(u.Path, "v1.0", version, 1)
		return u
	}
}
