package graph_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
	"github.com/tonyeung/microsoft-graph-client-go/graph"
)

func TestApi(t *testing.T) {

	jsonResponse := `{
	  "id": "one"
	 }`

	// create a new reader with that JSON
	reader := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	mockClient := &MockHttpClient{
		mockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       reader,
			}, nil
		},
	}

	msal := confidential.Client{}

	graphClient, err := graph.New(msal, mockClient)
	if err != nil {
		t.Error("create graphRequest failed.")
	}

	result, err := graphClient.Api("/some-resource")
	if err != nil {
		t.Error("TestGet failed.")
		return
	}

	if result != jsonResponse {
		t.Error("TestGet failed, expected " + jsonResponse + ", got:" + result)
		return
	}
}

// mock the http client and store the MockDo lambda
type MockHttpClient struct {
	mockDo MockDo
}

// mock the do function implementation on the http client as a lambda
type MockDo func(req *http.Request) (*http.Response, error)

// mock the do function of the http client
func (m MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.mockDo(req)
}
