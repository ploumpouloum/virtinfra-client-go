package virtinfra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_VpcAdd(t *testing.T) {
	client := Client{
		localFileLocation: "dummyLocation",
		account:           Account{},
		doNotPersist:      true,
	}
	tests := []struct {
		name       string
		vpcs       []Vpc
		resultVpcs []Vpc
		vpcToAdd   Vpc
		wantErr    bool
	}{
		{
			name: "Add Vpc to empty list",
			vpcs: []Vpc{},
			resultVpcs: []Vpc{
				{
					Id: "1234",
				},
			},
			vpcToAdd: Vpc{
				Id: "1234",
			},
			wantErr: false,
		},
		{
			name: "Add Vpc to existing list",
			vpcs: []Vpc{
				{
					Id: "1234",
				},
			},
			resultVpcs: []Vpc{
				{
					Id: "1234",
				},
				{
					Id: "12345",
				},
			},
			vpcToAdd: Vpc{
				Id: "12345",
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client.account.Vpcs = tt.vpcs
			if err := client.VpcAdd(tt.vpcToAdd); (err != nil) != tt.wantErr {
				t.Errorf("Client.VpcAdd() error = %v, wantErr %v", err, tt.wantErr)
			}
			for _, resultVpc := range tt.resultVpcs {
				assert.Contains(t, client.account.Vpcs, resultVpc)
			}
			if len(tt.resultVpcs) == 0 && len(client.account.Vpcs) > 0 {
				t.Error("Client.VpcAdd() error: Vpc list is not empty")
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
			for _, resultVpc := range tt.resultVpcs {
				assert.Contains(t, client.account.Vpcs, resultVpc)
			}
			if len(tt.resultVpcs) == 0 && len(client.account.Vpcs) > 0 {
				t.Error("Client.VpcDelete() error: Vpc list is not empty")
			}
		})
	}
}
