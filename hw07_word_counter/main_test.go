package main

import (
	"reflect"
	"testing"
)

func Test_punkDelete(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestPunk",
			args: args{"Privet , kak dela , !"},
			want: "Privet  kak dela  ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := punkDelete(tt.args.text); got != tt.want {
				t.Errorf("punkDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countWords(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Count",
			args: args{"I am Dima Am"},
			want: map[string]int{
				"am":   2,
				"dima": 1,
				"i":    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countWords(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowClear(t *testing.T) {
	type args struct {
		stringClear string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestLow",
			args: args{"A b C A"},
			want: []string{"a", "b", "c", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowClear(tt.args.stringClear); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowClear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countWords1(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Done",
			args: args{"Dima Privet dima !"},
			want: map[string]int{
				"dima":   2,
				"privet": 1,
			},
		},
		{
			name: "Empty",
			args: args{""},
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countWords(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
