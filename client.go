package virtinfra

import (
	"encoding/json"
	"io/ioutil"
)

type Client struct {
	LocalFileLocation string
	Account           Account
}

func OpenClientFromLocalStorage(localFileLocation string) (client *Client, err error) {
	client.LocalFileLocation = localFileLocation
	fileContent, err := ioutil.ReadFile(client.LocalFileLocation)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileContent, &client.Account)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (client Client) persistState() error {
	b, err := json.MarshalIndent(client.Account, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(client.LocalFileLocation, b, 0644)
	if err != nil {
		return err
	}
	return nil
}
