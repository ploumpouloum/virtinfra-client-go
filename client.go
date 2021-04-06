package virtinfra

import (
	"encoding/json"
	"io/ioutil"
)

type Client struct {
	localFileLocation string
	account           Account
}

func (client Client) OpenAccountFromLocalStorage(localFileLocation string) error {
	client.localFileLocation = localFileLocation
	fileContent, err := ioutil.ReadFile(client.localFileLocation)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileContent, &client.account)
	if err != nil {
		return err
	}
	return nil
}

func (client Client) persistState() error {
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
