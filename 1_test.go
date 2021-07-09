package main

import "testing"

func TestCheckCharCount(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name: "Checking the winner and loser",
			args: args{
				str: "aaabbbcceffg",
			},
			want:  "There is no loser here",
			want1: "There is no Winner here",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CheckCharCount(tt.args.str)
			if got != tt.want {
				t.Errorf("CheckCharCount() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckCharCount() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
