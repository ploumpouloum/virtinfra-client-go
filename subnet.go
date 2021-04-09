package virtinfra

import (
	"fmt"
	"math/rand"
)

type SubnetId string

type Subnet struct {
	Id    SubnetId `json:"id"`
	Cidr  Cidr     `json:"cidr"`
	VpcId VpcId    `json:"vpcId"`
	vpc   Vpc
}

func (client *Client) SubnetGetList() ([]Subnet, error) {
	return client.account.Subnets, nil
}

func (client *Client) SubnetAdd(subnet *Subnet) error {
	subnet.Id = (SubnetId)(fmt.Sprintf("subnet-%08d", rand.Intn(100000000)))
	client.account.Subnets = append(client.account.Subnets, *subnet)
	err := client.persistState()
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) SubnetDelete(id SubnetId) error {
	var iSubnetToDelete = -1
	for i, subnet := range client.account.Subnets {
		if subnet.Id == id {
			iSubnetToDelete = i
		}
	}
	if iSubnetToDelete >= 0 {
		client.account.Subnets = append(client.account.Subnets[:iSubnetToDelete], client.account.Subnets[iSubnetToDelete+1:]...)
	} else {
		return fmt.Errorf("Unable to find Subnet Id '%s'", id)
	}
	err := client.persistState()
	if err != nil {
		return err
	}
	return nil
}

func (client Client) SubnetGet(id SubnetId) (*Subnet, error) {
	for _, subnet := range client.account.Subnets {
		if subnet.Id == id {
			return &subnet, nil
		}
	}
	return nil, fmt.Errorf("Unable to find Subnet Id '%s'", id)
}

func (client Client) SubnetUpdate(updatedSubnet *Subnet) error {
	for i, subnet := range client.account.Subnets {
		if subnet.Id == updatedSubnet.Id {
			client.account.Subnets[i] = *updatedSubnet
			break
		}
	}
	err := client.persistState()
	if err != nil {
		return err
	}
	return nil
}
