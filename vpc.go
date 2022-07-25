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
	var iVpcToDelete = -1
	for i, vpc := range client.account.Vpcs {
		if vpc.Id == id {
			iVpcToDelete = i
		}
	}
	if iVpcToDelete >= 0 {
		client.account.Vpcs = append(client.account.Vpcs[:iVpcToDelete], client.account.Vpcs[iVpcToDelete+1:]...)
	} else {
		return fmt.Errorf("Unable to find VPC Id '%s'", id)
	}
	err := client.persistState()
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) VpcGet(id VpcId) (*Vpc, error) {
	for _, vpc := range client.account.Vpcs {
		if vpc.Id == id {
			return &vpc, nil
		}
	}
	return nil, nil
}

func (client *Client) VpcUpdate(updatedVpc *Vpc) error {
	matchingVpcIdx := slices.IndexFunc(client.account.Vpcs, func(v Vpc) bool { return v.Id == updatedVpc.Id })
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
