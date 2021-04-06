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

func (client Client) VpcAdd() error {
	client.account.Vpcs = append(client.account.Vpcs,
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
