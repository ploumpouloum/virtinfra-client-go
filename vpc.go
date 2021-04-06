package virtinfra

type VpcId string
type Cidr string

type Vpc struct {
	Id      VpcId
	Cidr    Cidr
	subnets []Subnet
}

func (client Client) VpcGetList() ([]Vpc, error) {
	return client.account.Vpcs, nil
}
