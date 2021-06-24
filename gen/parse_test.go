package gen

import (
	"bytes"
	"testing"
)

func TestExecute(t *testing.T) {
	type args struct {
		t Template
	}
	type T struct {
		name string
		args args
	}
	tests := []T{
		{
			name: "basic test",
			args: args{
				t: Template{
					PkgName: "test_example1",
					Functions: []Function{{
						Body: Body{Type: "user.User"},
						Name: "example1",
						QueryParams: []QueryParam{{
							Name: "user_id",
							Type: "int",
						}},
						HasQueryParams: true,
						HasBody:        true,
					}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := Execute(tt.args.t, w); err != nil {
				t.Errorf("Execute() error = %v", err)
				return
			}
		})
	}
}
