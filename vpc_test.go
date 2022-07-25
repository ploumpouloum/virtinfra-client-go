package virtinfra

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestClient_VpcAdd(t *testing.T) {
	client := Client{
		localFileLocation: "dummyLocation",
		account:           Account{},
		doNotPersist:      true,
	}
	tests := []struct {
		name         string
		existingVpcs []Vpc
		resultVpcs   []Vpc
		vpcToAdd     Vpc
		wantErr      bool
	}{
		{
			name:         "Add Vpc to empty list",
			existingVpcs: []Vpc{},
			resultVpcs: []Vpc{
				{
					Cidr: "10.0.0.0/16",
				},
			},
			vpcToAdd: Vpc{
				Cidr: "10.0.0.0/16",
			},
			wantErr: false,
		},
		{
			name: "Add Vpc to existing list",
			existingVpcs: []Vpc{
				{
					Cidr: "10.0.0.0/16",
				},
			},
			resultVpcs: []Vpc{
				{
					Cidr: "10.0.0.0/16",
				},
				{
					Cidr: "10.1.0.0/16",
				},
			},
			vpcToAdd: Vpc{
				Cidr: "10.1.0.0/16",
			},
			wantErr: false,
		},
		{
			name:         "Can't add Vpc with empty CIDR block",
			existingVpcs: []Vpc{},
			resultVpcs:   []Vpc{},
			vpcToAdd:     Vpc{},
			wantErr:      true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client.account.Vpcs = tt.existingVpcs
			if err := client.VpcAdd(&tt.vpcToAdd); (err != nil) != tt.wantErr {
				t.Errorf("Client.VpcAdd() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equalf(t, len(tt.resultVpcs), len(client.account.Vpcs), "Unexpected number of Vpcs found")
			for _, resultVpc := range tt.resultVpcs {
				assert.GreaterOrEqual(t, slices.IndexFunc(client.account.Vpcs, func(v Vpc) bool { return v.Cidr == resultVpc.Cidr }), 0, "Vpc with CIDR %v is missing", resultVpc.Cidr)
			}
		})
	}
}

func TestClient_VpcDelete(t *testing.T) {
	client := Client{
		localFileLocation: "dummyLocation",
		account:           Account{},
		doNotPersist:      true,
	}
	tests := []struct {
		name          string
		vpcs          []Vpc
		resultVpcs    []Vpc
		vpcIdToDelete VpcId
		wantErr       bool
	}{
		{
			name: "One single VPC to delete",
			vpcs: []Vpc{
				{
					Id: "1234",
				},
			},
			vpcIdToDelete: "1234",
			wantErr:       false,
		},
		{
			name: "First VPC to delete among two",
			vpcs: []Vpc{
				{
					Id: "1234",
				},
				{
					Id: "12345",
				},
			},
			resultVpcs: []Vpc{
				{
					Id: "12345",
				},
			},
			vpcIdToDelete: "1234",
			wantErr:       false,
		},
		{
			name: "Last VPC to delete among two",
			vpcs: []Vpc{
				{
					Id: "1234",
				},
				{
					Id: "12345",
				},
			},
			resultVpcs: []Vpc{
				{
					Id: "1234",
				},
			},
			vpcIdToDelete: "12345",
			wantErr:       false,
		},
		{
			name: "Missing VPC to delete among three",
			vpcs: []Vpc{
				{
					Id: "123",
				},
				{
					Id: "1234",
				},
				{
					Id: "12345",
				},
			},
			resultVpcs: []Vpc{
				{
					Id: "123",
				},
				{
					Id: "1234",
				},
				{
					Id: "12345",
				},
			},
			vpcIdToDelete: "1235",
			wantErr:       true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client.account.Vpcs = tt.vpcs
			if err := client.VpcDelete(tt.vpcIdToDelete); (err != nil) != tt.wantErr {
				t.Errorf("Client.VpcDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equalf(t, len(tt.resultVpcs), len(client.account.Vpcs), "Unexpected number of Vpcs found")
			for _, resultVpc := range tt.resultVpcs {
				assert.Contains(t, client.account.Vpcs, resultVpc)
			}
		})
	}
}

func TestClient_VpcUpdate(t *testing.T) {
	client := Client{
		localFileLocation: "dummyLocation",
		account:           Account{},
		doNotPersist:      true,
	}
	tests := []struct {
		name         string
		existingVpcs []Vpc
		resultVpcs   []Vpc
		updatedVpc   Vpc
		wantErr      bool
	}{
		{
			name: "One single VPC to update",
			existingVpcs: []Vpc{
				{
					Id:   "1234",
					Cidr: "10.0.0.0/16",
				},
			},
			resultVpcs: []Vpc{
				{
					Id:   "1234",
					Cidr: "10.1.0.0/16",
				},
			},
			updatedVpc: Vpc{
				Id:   "1234",
				Cidr: "10.1.0.0/16",
			},
			wantErr: false,
		},
		{
			name: "One VPC to update among many",
			existingVpcs: []Vpc{
				{
					Id:   "1234",
					Cidr: "10.0.0.0/16",
				},
				{
					Id:   "12345",
					Cidr: "10.1.0.0/16",
				},
				{
					Id:   "123456",
					Cidr: "10.2.0.0/16",
				},
			},
			resultVpcs: []Vpc{
				{
					Id:   "1234",
					Cidr: "10.0.0.0/16",
				},
				{
					Id:   "12345",
					Cidr: "10.3.0.0/16",
				},
				{
					Id:   "123456",
					Cidr: "10.2.0.0/16",
				},
			},
			updatedVpc: Vpc{
				Id:   "12345",
				Cidr: "10.3.0.0/16",
			},
			wantErr: false,
		},
		{
			name: "One VPC to update which does not exists",
			existingVpcs: []Vpc{
				{
					Id:   "1234",
					Cidr: "10.0.0.0/16",
				},
			},
			resultVpcs: []Vpc{},
			updatedVpc: Vpc{
				Id:   "12345",
				Cidr: "10.3.0.0/16",
			},
			wantErr: true,
		},
		{
			name:         "One VPC to update without any existing VPCs",
			existingVpcs: []Vpc{},
			resultVpcs:   []Vpc{},
			updatedVpc: Vpc{
				Id:   "12345",
				Cidr: "10.3.0.0/16",
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client.account.Vpcs = tt.existingVpcs
			if err := client.VpcUpdate(&tt.updatedVpc); (err != nil) != tt.wantErr {
				t.Errorf("Client.VpcUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return // Do not check other arguments if we expect an error
			}
			assert.Equalf(t, len(tt.resultVpcs), len(client.account.Vpcs), "Unexpected number of Vpcs found")
			for _, resultVpc := range tt.resultVpcs {
				matchingVpcIdx := slices.IndexFunc(client.account.Vpcs, func(v Vpc) bool { return v.Id == resultVpc.Id })
				assert.GreaterOrEqual(t, matchingVpcIdx, 0, "Vpc with Id %v is missing", resultVpc.Id)
				assert.Equalf(t, resultVpc.Cidr, client.account.Vpcs[matchingVpcIdx].Cidr, "Vpc with Id %v has incorrect CIDR", resultVpc.Id)
			}
		})
	}
}

func TestClient_VpcGet(t *testing.T) {
	client := Client{
		localFileLocation: "dummyLocation",
		account:           Account{},
		doNotPersist:      true,
	}
	tests := []struct {
		name         string
		existingVpcs []Vpc
		resultVpc    Vpc
		idToGet      VpcId
		wantErr      bool
	}{
		{
			name: "One single VPC to get",
			existingVpcs: []Vpc{
				{
					Id:   "1234",
					Cidr: "10.0.0.0/16",
				},
			},
			resultVpc: Vpc{

				Id:   "1234",
				Cidr: "10.0.0.0/16",
			},
			idToGet: "1234",
			wantErr: false,
		},
		{
			name: "One single VPC to get among many",
			existingVpcs: []Vpc{
				{
					Id:   "1234",
					Cidr: "10.0.0.0/16",
				},
				{
					Id:   "12345",
					Cidr: "10.1.0.0/16",
				},
				{
					Id:   "123456",
					Cidr: "10.2.0.0/16",
				},
			},
			resultVpc: Vpc{

				Id:   "12345",
				Cidr: "10.1.0.0/16",
			},
			idToGet: "12345",
			wantErr: false,
		},
		{
			name: "One single VPC to get which does not exists",
			existingVpcs: []Vpc{
				{
					Id:   "1234",
					Cidr: "10.0.0.0/16",
				},
			},
			idToGet: "12343",
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client.account.Vpcs = tt.existingVpcs
			vpc, err := client.VpcGet(tt.idToGet)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.VpcGet() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return // Do not check other arguments if we expect an error
			}
			assert.Equalf(t, tt.resultVpc, *vpc, "Unexpected Vpc Id found")
		})
	}
}
