package query_test

import (
	"net/url"
	"testing"

	"github.com/tonyeung/microsoft-graph-client-go/graph/query"
)

func TestVersion(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "google.com",
		Path:   "v1.0",
	}
	q := query.Version("foo")
	u = q(u)

	if u.String() != "https://google.com/foo" {
		t.Error("TestTop failed, expected https://google.com/foo, got:" + u.String())
		return
	}
}

func TestCount(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "google.com",
	}
	q := query.Count(true)
	u = q(u)

	if u.String() != "https://google.com?%24count=true" {
		t.Error("TestTop failed, expected https://google.com?%24count=true, got:" + u.String())
		return
	}
}

func TestExpand(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "google.com",
	}
	q := query.Expand("members")
	u = q(u)

	if u.String() != "https://google.com?%24expand=members" {
		t.Error("TestTop failed, expected https://google.com?%24expand=members, got:" + u.String())
		return
	}
}

func TestFilter(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "google.com",
	}
	q := query.Filter("filter")
	u = q(u)

	if u.String() != "https://google.com?%24filter=filter" {
		t.Error("TestTop failed, expected https://google.com?%24filter=filter, got:" + u.String())
		return
	}
}

func TestFormat(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "google.com",
	}
	q := query.Format("json")
	u = q(u)

	if u.String() != "https://google.com?%24format=json" {
		t.Error("TestTop failed, expected https://google.com?%24format=json, got:" + u.String())
		return
	}
}

func TestOrderBy(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "google.com",
	}
	q := query.OrderBy("displayName desc")
	u = q(u)

	if u.String() != "https://google.com?%24orderby=displayName+desc" {
		t.Error("TestTop failed, expected https://google.com?%24orderby=displayName+desc, got:" + u.String())
		return
	}
}

func TestSearch(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "google.com",
	}
	q := query.Search("pizza")
	u = q(u)

	if u.String() != "https://google.com?%24search=pizza" {
		t.Error("TestTop failed, expected https://google.com?%24search=pizza, got:" + u.String())
		return
	}
}

func TestSelect(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "google.com",
	}
	q := query.Select("givenName,surname")
	u = q(u)

	if u.String() != "https://google.com?%24select=givenName%2Csurname" {
		t.Error("TestTop failed, expected https://google.com?%24select=givenName%2Csurname, got:" + u.String())
		return
	}
}

func TestSkip(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "google.com",
	}
	q := query.Skip(11)
	u = q(u)

	if u.String() != "https://google.com?%24skip=11" {
		t.Error("TestTop failed, expected https://google.com?%24skip=11, got:" + u.String())
		return
	}
}

func TestSkipToken(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "google.com",
	}
	q := query.SkipToken("token")
	u = q(u)

	if u.String() != "https://google.com?%24skiptoken=token" {
		t.Error("TestTop failed, expected https://google.com?%24skiptoken=token, got:" + u.String())
		return
	}
}

func TestTop(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "google.com",
	}
	q := query.Top(1)
	u = q(u)

	if u.String() != "https://google.com?%24top=1" {
		t.Error("TestTop failed, expected https://google.com?%24top=1, got:" + u.String())
		return
	}
}
