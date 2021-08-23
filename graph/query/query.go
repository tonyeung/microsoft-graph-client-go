package query

import (
	"net/url"
	"strconv"
	"strings"
)

type QueryOption func(url url.URL) url.URL

func Version(version string) QueryOption {
	return func(u url.URL) url.URL {
		u.Path = strings.Replace(u.Path, "v1.0", version, 1)
		return u
	}
}

func Count(b bool) QueryOption {
	return func(u url.URL) url.URL {
		queryString, _ := url.ParseQuery(u.RawQuery)
		queryString.Add("$count", strconv.FormatBool(b))
		u.RawQuery = queryString.Encode()
		return u
	}
}

func Expand(resource string) QueryOption {
	return func(u url.URL) url.URL {
		queryString, _ := url.ParseQuery(u.RawQuery)
		queryString.Add("$expand", resource)
		u.RawQuery = queryString.Encode()
		return u
	}
}

func Filter(filter string) QueryOption {
	return func(u url.URL) url.URL {
		queryString, _ := url.ParseQuery(u.RawQuery)
		queryString.Add("$filter", filter)
		u.RawQuery = queryString.Encode()
		return u
	}
}

func Format(format string) QueryOption {
	return func(u url.URL) url.URL {
		queryString, _ := url.ParseQuery(u.RawQuery)
		queryString.Add("$format", format)
		u.RawQuery = queryString.Encode()
		return u
	}
}

func OrderBy(order string) QueryOption {
	return func(u url.URL) url.URL {
		queryString, _ := url.ParseQuery(u.RawQuery)
		queryString.Add("$orderby", order)
		u.RawQuery = queryString.Encode()
		return u
	}
}

func Search(term string) QueryOption {
	return func(u url.URL) url.URL {
		queryString, _ := url.ParseQuery(u.RawQuery)
		queryString.Add("$search", term)
		u.RawQuery = queryString.Encode()
		return u
	}
}

func Select(fields string) QueryOption {
	return func(u url.URL) url.URL {
		queryString, _ := url.ParseQuery(u.RawQuery)
		queryString.Add("$select", fields)
		u.RawQuery = queryString.Encode()
		return u
	}
}

func Skip(items int) QueryOption {
	return func(u url.URL) url.URL {
		queryString, _ := url.ParseQuery(u.RawQuery)
		queryString.Add("$skip", strconv.Itoa(items))
		u.RawQuery = queryString.Encode()
		return u
	}
}

func SkipToken(token string) QueryOption {
	return func(u url.URL) url.URL {
		queryString, _ := url.ParseQuery(u.RawQuery)
		queryString.Add("$skiptoken", token)
		u.RawQuery = queryString.Encode()
		return u
	}
}

func Top(count int) QueryOption {
	return func(u url.URL) url.URL {
		queryString, _ := url.ParseQuery(u.RawQuery)
		queryString.Add("$top", strconv.Itoa(count))
		u.RawQuery = queryString.Encode()
		return u
	}
}
