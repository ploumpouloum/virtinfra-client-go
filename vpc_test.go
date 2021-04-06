package virtinfra

import "testing"

func TestClient_VpcAdd(t *testing.T) {
	type fields struct {
		localFileLocation string
		account           Account
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test 1",
			fields: fields{
				localFileLocation: "/tmp/blob.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := Client{
				localFileLocation: tt.fields.localFileLocation,
				account:           tt.fields.account,
			}
			if err := client.VpcAdd(); (err != nil) != tt.wantErr {
				t.Errorf("Client.VpcAdd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
