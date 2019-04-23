package main

import (
	"reflect"
	"testing"
)

func Test_parseArg(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    cmd
		wantErr error
	}{
		{name: "Map<string:Product>int", want: cmd{
			fn: "Map", index: "string", output: "int", input: "Product",
		}, wantErr: nil},
		{name: "Map<Product>int", want: cmd{
			fn: "Map", index: "int", output: "int", input: "Product",
		}, wantErr: nil},
		{name: "Reduce<Product>int", want: cmd{
			fn: "Reduce", index: "int", output: "int", input: "Product",
		}, wantErr: nil},
		{name: "Reduce<string:Product>int", want: cmd{
			fn: "Reduce", index: "string", output: "int", input: "Product",
		}, wantErr: nil},
		{name: "Filter<Product>", want: cmd{
			fn: "Filter", index: "int", output: "", input: "Product",
		}, wantErr: nil},
		{name: "Filter<string:Product>", want: cmd{
			fn: "Filter", index: "string", output: "", input: "Product",
		}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseArg(tt.name)
			if err != tt.wantErr {
				t.Errorf("parseArg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseArg() = %v, want %v", got, tt.want)
			}
		})
	}
}
