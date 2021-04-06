package virtinfra

type AccountId string

type Account struct {
	Id      AccountId `json:"id"`
	Vpcs    []Vpc     `json:"vpcs"`
	Subnets []Subnet  `json:"subnets"`
}
