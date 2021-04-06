package virtinfra

import (
	"encoding/json"
	"io/ioutil"
)

type Client struct {
	LocalFileLocation string
	Account           Account
}

func (client Client) OpenAccountFromLocalStorage(localFileLocation string) error {
	client.LocalFileLocation = localFileLocation
	fileContent, err := ioutil.ReadFile(client.LocalFileLocation)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileContent, &client.Account)
	if err != nil {
		return err
	}
	return nil
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
