package virtinfra

import "fmt"

type VpcId string
type Cidr string

type Vpc struct {
	Id   VpcId `json:"id"`
	Cidr Cidr  `json:"cidr"`
}

func (client Client) VpcGetList() ([]Vpc, error) {
	return client.account.Vpcs, nil
}

func (client Client) VpcAdd(vpc Vpc) error {
	client.account.Vpcs = append(client.account.Vpcs, vpc)
	err := client.persistState()
	if err != nil {
		return err
	}
	return nil
}

func (client Client) VpcGet(id VpcId) (*Vpc, error) {
	for _, vpc := range client.account.Vpcs {
		if vpc.Id == id {
			return &vpc, nil
		}
	}
	return nil, fmt.Errorf("Unable to find VPC Id '%s'", id)
}
