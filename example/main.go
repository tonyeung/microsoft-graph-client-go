package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
	"github.com/tonyeung/microsoft-graph-client-go/graph"
	"github.com/tonyeung/microsoft-graph-client-go/graph/query"
)

var (
	cacheAccessor = &TokenCache{"serialized_cache.json"}
)

func main() {
	msal := getMSAL()
	httpClient := http.Client{}

	client, err := graph.New(msal, &httpClient)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Api("/users", "GET",
		query.Version("beta"),
		query.Top(1),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(result)
}

func getMSAL() confidential.Client {

	config := CreateConfig("config.json")
	confidentialConfig := CreateConfig("confidentialConfig.json")

	cred, err := confidential.NewCredFromSecret(confidentialConfig.ClientSecret)
	if err != nil {
		log.Fatal(err)
	}
	app, err := confidential.New(config.ClientID, cred, confidential.WithAuthority(config.Authority), confidential.WithAccessor(cacheAccessor))
	if err != nil {
		log.Fatal(err)
	}
	return app
}

// Config represents the config.json required to run the samples
type Config struct {
	ClientID            string   `json:"client_id"`
	Authority           string   `json:"authority"`
	Scopes              []string `json:"scopes"`
	Username            string   `json:"username"`
	Password            string   `json:"password"`
	RedirectURI         string   `json:"redirect_uri"`
	CodeChallenge       string   `json:"code_challenge"`
	CodeChallengeMethod string   `json:"code_challenge_method"`
	State               string   `json:"state"`
	ClientSecret        string   `json:"client_secret"`
	Thumbprint          string   `json:"thumbprint"`
	PemData             string   `json:"pem_file"`
}

// CreateConfig creates the Config struct from a json file.
func CreateConfig(fileName string) *Config {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	config := &Config{}
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

//https://github.com/AzureAD/microsoft-authentication-library-for-go/blob/87abbc195aae35793cf90da0efb8b8a443ac30cb/apps/tests/devapps/sample_cache_accessor.go#L14
type TokenCache struct {
	file string
}

func (t *TokenCache) Replace(cache cache.Unmarshaler, key string) {
	jsonFile, err := os.Open(t.file)
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
	}
	err = cache.Unmarshal(data)
	if err != nil {
		log.Println(err)
	}
}

func (t *TokenCache) Export(cache cache.Marshaler, key string) {
	data, err := cache.Marshal()
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(t.file, data, 0600)
	if err != nil {
		log.Println(err)
	}
}
