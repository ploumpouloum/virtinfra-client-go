package virtinfra

type SubnetId string

type Subnet struct {
	Id    SubnetId
	Cidr  Cidr
	VpcId VpcId
	vpc   Vpc
}
