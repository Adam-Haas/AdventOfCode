package main

import (
	"testing"
)

func Test_findCalibrationValues(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name     string
		args     args
		wantVal1 int
		wantVal2 int
	}{
		{
			name: "example 1",
			args: args{
				v: "1abc2",
			},
			wantVal1: 1,
			wantVal2: 2,
		},
		{
			name: "example 2",
			args: args{
				v: "pqr3stu8vwx",
			},
			wantVal1: 3,
			wantVal2: 8,
		},
		{
			name: "example 3",
			args: args{
				v: "a1b2c3d4e5f",
			},
			wantVal1: 1,
			wantVal2: 5,
		},
		{
			name: "example 4",
			args: args{
				v: "treb7uchet",
			},
			wantVal1: 7,
			wantVal2: 7,
		},
		{
			name: "both numbers in second half of string",
			args: args{
				v: "sixfdqttpskdnbksqxg9three6bqqpngfhz",
			},
			wantVal1: 6,
			wantVal2: 6,
		},
		{
			name: "part 2 - example 1",
			args: args{
				v: "two1nine",
			},
			wantVal1: 2,
			wantVal2: 9,
		},
		{
			name: "part 2 - example 2",
			args: args{
				v: "eightwothree",
			},
			wantVal1: 8,
			wantVal2: 3,
		},
		{
			name: "part 2 - example 3",
			args: args{
				v: "abcone2threexyz",
			},
			wantVal1: 1,
			wantVal2: 3,
		},
		{
			name: "part 2 - example 4",
			args: args{
				v: "xtwone3four",
			},
			wantVal1: 2,
			wantVal2: 4,
		},
		{
			name: "part 2 - example 5",
			args: args{
				v: "4nineeightseven2",
			},
			wantVal1: 4,
			wantVal2: 2,
		},
		{
			name: "part 2 - example 6",
			args: args{
				v: "zoneight234",
			},
			wantVal1: 1,
			wantVal2: 4,
		},
		{
			name: "part 2 - example 7",
			args: args{
				v: "7pqrstsixteen",
			},
			wantVal1: 7,
			wantVal2: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal1, gotVal2 := findCalibrationValues(tt.args.v)
			if gotVal1 != tt.wantVal1 {
				t.Errorf("fetchCalibrationValues() gotVal1 = %v, want %v", gotVal1, tt.wantVal1)
			}
			if gotVal2 != tt.wantVal2 {
				t.Errorf("fetchCalibrationValues() gotVal2 = %v, want %v", gotVal2, tt.wantVal2)
			}
		})
	}
}

func Test_solveCalibration(t *testing.T) {
	type args struct {
		calibrations []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				calibrations: []string{
					"1abc2",
					"pqr3stu8vwx",
					"a1b2c3d4e5f",
					"treb7uchet",
				},
			},
			want: 142,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveCalibration(tt.args.calibrations); got != tt.want {
				t.Errorf("solveCalibration() = %v, want %v", got, tt.want)
			}
		})
	}
}
