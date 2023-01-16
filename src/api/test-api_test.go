/**
* Go File created on 16.01.23
* by Antonio Masotti (antonio)
* MIT License
 */

package api

import "testing"

func TestSampleMessage_ToJsonString(t *testing.T) {
	tests := []struct {
		name    string
		fields  SampleMessage
		want    string
		wantErr bool
	}{
		{
			name: "Test 1",
			fields: SampleMessage{
				Id:       "5c4b4b9d6b6d4f0f8c8b4b9d",
				CrmId:    "5c4b4b9d6b6d4f0f8c8b4b9d",
				IsActive: true,
				Age:      33,
				Name:     "Tony the Tony",
				Gender:   "male",
				Company:  "Zuplo",
				Email:    "test@test.go",
			},
			want:    `{"_id":"5c4b4b9d6b6d4f0f8c8b4b9d","crmId":"5c4b4b9d6b6d4f0f8c8b4b9d","isActive":true,"age":33,"name":"Tony the Tony","gender":"male","company":"Zuplo","email":"test@test.go","updateTime":0}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &SampleMessage{
				Id:         tt.fields.Id,
				CrmId:      tt.fields.CrmId,
				IsActive:   tt.fields.IsActive,
				Age:        tt.fields.Age,
				Name:       tt.fields.Name,
				Gender:     tt.fields.Gender,
				Company:    tt.fields.Company,
				Email:      tt.fields.Email,
				UpdateTime: tt.fields.UpdateTime,
			}
			got, err := p.ToJsonString()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJsonString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToJsonString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
