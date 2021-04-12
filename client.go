package virtinfra

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Client struct {
	localFileLocation string
	account           Account
	doNotPersist      bool
}

func OpenClientFromLocalStorage(localFileLocation string) (*Client, error) {
	rand.Seed(time.Now().Unix())
	client := &Client{}
	client.localFileLocation = localFileLocation
	_, err := os.Stat(client.localFileLocation)
	if os.IsNotExist(err) {
		return client, nil
	}
	fileContent, err := ioutil.ReadFile(client.localFileLocation)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileContent, &client.account)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (client *Client) persistState() error {
	if client.doNotPersist {
		return nil
	}
	path := filepath.Dir(client.localFileLocation)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	b, err := json.MarshalIndent(client.account, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(client.localFileLocation, b, 0644)
	if err != nil {
		return err
	}
	return nil
}
