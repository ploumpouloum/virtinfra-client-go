package virtinfra

type AccountId string

type Account struct {
	Id      AccountId
	Vpcs    []Vpc
	Subnets []Subnet
}
