package virtinfra

import "testing"

func TestClient_VpcAdd(t *testing.T) {
	type fields struct {
		LocalFileLocation string
		Account           Account
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test 1",
			fields: fields{
				LocalFileLocation: "/tmp/blob.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := Client{
				LocalFileLocation: tt.fields.LocalFileLocation,
				Account:           tt.fields.Account,
			}
			if err := client.VpcAdd(); (err != nil) != tt.wantErr {
				t.Errorf("Client.VpcAdd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
