package src

import "testing"

func TestTranslate(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				word: "I have an apple",
			},
		},
		{
			args: args{
				word: "我有一个苹果",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Translate(tt.args.word)
		})
	}
}
