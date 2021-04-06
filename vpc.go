package virtinfra

type VpcId string
type Cidr string

type Vpc struct {
	Id      VpcId `json:"id"`
	Cidr    Cidr  `json:"cidr"`
	subnets []Subnet
}

func (client Client) VpcGetList() ([]Vpc, error) {
	return client.Account.Vpcs, nil
}

func (client Client) VpcAdd() error {
	client.Account.Vpcs = append(client.Account.Vpcs,
		Vpc{
			Id: "1234",
		},
	)
	err := client.persistState()
	if err != nil {
		return err
	}
	return nil
}
