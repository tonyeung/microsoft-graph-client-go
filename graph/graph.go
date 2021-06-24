package graph

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
	"github.com/tonyeung/microsoft-graph-client-go/graph/query"
)

type HttpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	msal        confidential.Client
	graphClient HttpDoer
	scopes      []string
	graphHost   string
	version     string
}

func New(app confidential.Client, httpClient HttpDoer) (*Client, error) {
	return &Client{
		msal:        app,
		graphClient: httpClient,
		scopes:      []string{"https://graph.microsoft.com/.default"},
		graphHost:   "graph.microsoft.com",
		version:     "v1.0",
	}, nil
}

func (g Client) Api(resource string, opts ...query.QueryOption) (string, error) {
	url := url.URL{
		Scheme: "https",
		Host:   g.graphHost,
		Path:   g.version + resource,
	}

	for _, opt := range opts {
		url = opt(url)
	}

	return g.get(url.String())
}

func (g Client) get(url string) (string, error) {

	authResult, err := g.msal.AcquireTokenSilent(context.Background(), g.scopes)
	if err != nil {
		authResult, err = g.msal.AcquireTokenByCredential(context.Background(), g.scopes)
		if err != nil {
			log.Fatal(err)
		}
	}
	bearer := "Bearer " + authResult.AccessToken
	fmt.Print(bearer)

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Authorization", bearer)

	response, _ := g.graphClient.Do(request)

	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b), nil
}
