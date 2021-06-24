package query_test

import (
	"net/url"
	"testing"

	"github.com/tonyeung/microsoft-graph-client-go/graph/query"
)

func TestTop(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "slb.com",
	}
	q := query.Top(1)
	u = q(u)

	if u.String() != "https://slb.com?%24top=1" {
		t.Error("TestTop failed, expected https://slb.com?%24top=1, got:" + u.String())
		return
	}
}

func TestVersion(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "slb.com",
		Path:   "v1.0",
	}
	q := query.Version("foo")
	u = q(u)

	if u.String() != "https://slb.com/foo" {
		t.Error("TestTop failed, expected https://slb.com/foo, got:" + u.String())
		return
	}
}
