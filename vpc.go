package virtinfra

import (
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/exp/slices"
)

type VpcId string
type Cidr string

type Vpc struct {
	Id   VpcId `json:"id"`
	Cidr Cidr  `json:"cidr"`
}

func (client *Client) VpcGetList() ([]Vpc, error) {
	return client.account.Vpcs, nil
}

func (client *Client) getVpcIdxFromId(id VpcId) int {
	return slices.IndexFunc(client.account.Vpcs, func(v Vpc) bool { return v.Id == id })
}

func (client *Client) VpcAdd(vpc *Vpc) error {
	if strings.TrimSpace(string(vpc.Cidr)) == "" {
		return fmt.Errorf("VPC Cidr block cannot be blank or empty")
	}
	vpc.Id = (VpcId)(fmt.Sprintf("vpc-%08d", rand.Intn(100000000)))
	client.account.Vpcs = append(client.account.Vpcs, *vpc)
	err := client.persistState()
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) VpcDelete(id VpcId) error {
	matchingVpcIdx := client.getVpcIdxFromId(id)
	if matchingVpcIdx < 0 {
		return fmt.Errorf("VPC Id %v does not exists", id)
	}
	client.account.Vpcs = append(client.account.Vpcs[:matchingVpcIdx], client.account.Vpcs[matchingVpcIdx+1:]...)
	err := client.persistState()
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) VpcGet(id VpcId) (*Vpc, error) {
	matchingVpcIdx := client.getVpcIdxFromId(id)
	if matchingVpcIdx < 0 {
		return nil, fmt.Errorf("VPC Id %v does not exists", id)
	}
	return &client.account.Vpcs[matchingVpcIdx], nil
}

func (client *Client) VpcUpdate(updatedVpc *Vpc) error {
	matchingVpcIdx := client.getVpcIdxFromId(updatedVpc.Id)
	if matchingVpcIdx < 0 {
		return fmt.Errorf("VPC Id %v does not exists", updatedVpc.Id)
	}
	client.account.Vpcs[matchingVpcIdx] = *updatedVpc
	err := client.persistState()
	if err != nil {
		return err
	}
	return nil
}
