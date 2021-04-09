package virtinfra

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type Client struct {
	localFileLocation string
	account           Account
	doNotPersist      bool
}

func OpenClientFromLocalStorage(localFileLocation string) (client *Client, err error) {
	rand.Seed(time.Now().Unix())
	client = &Client{}
	client.localFileLocation = localFileLocation
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
