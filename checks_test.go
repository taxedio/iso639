package iso639

import (
	"reflect"
	"testing"
)

var (
	Nil       *string
	NilStruct *ISOEntry

	ENb      string = isocodes3["ENG"].alph3b
	EN2      string = isocodes3["ENG"].alph2
	ENStruct        = ISOEntry{
		enName: isocodes3["ENG"].enName,
		frName: isocodes3["ENG"].frName,
		alph2:  isocodes3["ENG"].alph2,
		alph3b: isocodes3["ENG"].alph3b,
		alph3t: isocodes3["ENG"].alph3t,
	}

	FRb      string = isocodes3["FRE"].alph3b
	FR2      string = isocodes3["FRE"].alph2
	FRStruct        = ISOEntry{
		enName: isocodes3["FRE"].enName,
		frName: isocodes3["FRE"].frName,
		alph2:  isocodes3["FRE"].alph2,
		alph3b: isocodes3["FRE"].alph3b,
		alph3t: isocodes3["FRE"].alph3t,
	}

	// géorgien
	KAb string = isocodes3["GEO"].alph3b
	KA2 string = isocodes3["GEO"].alph2

	KAStruct = ISOEntry{
		enName: isocodes3["GEO"].enName,
		frName: isocodes3["GEO"].frName,
		alph2:  isocodes3["GEO"].alph2,
		alph3b: isocodes3["GEO"].alph3b,
		alph3t: isocodes3["GEO"].alph3t,
	}
)

func TestAlpha3Match(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{"EN Test", args{s: "EN"}, &ENb},
		{"EN English Name Test", args{s: "English"}, &ENb},
		{"EN French Name Test", args{s: "   Anglais   "}, &ENb},
		{"FR upper Test", args{s: "FRE"}, &FRb},
		{"KA accent test", args{s: "géorgien"}, &KAb},
		{"Nothing Test", args{s: ""}, Nil},
		{"XX Test", args{s: "XX"}, Nil},
		{"XXX Test", args{s: "XXX"}, Nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Alpha3Match(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Alpha3Match() = %#v, want %#v", *got, *tt.want)
			}
		})
	}
}

func TestAlpha2Match(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{"EN Test", args{s: "EN"}, &EN2},
		{"EN English Name Test", args{s: "English"}, &EN2},
		{"EN French Name Test", args{s: "   Anglais   "}, &EN2},
		{"FR upper Test", args{s: "FRE"}, &FR2},
		{"KA accent test", args{s: "géorgien"}, &KA2},
		{"Nothing Test", args{s: ""}, Nil},
		{"XX Test", args{s: "XX"}, Nil},
		{"XXX Test", args{s: "XXX"}, Nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Alpha2Match(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Alpha3Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructMatch(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *ISOEntry
	}{
		{"EN Test", args{s: "EN"}, &ENStruct},
		{"EN English Name Test", args{s: "English"}, &ENStruct},
		{"EN French Name Test", args{s: "   Anglais   "}, &ENStruct},
		{"FR upper Test", args{s: "FRE"}, &FRStruct},
		{"KA accent test", args{s: "géorgien"}, &KAStruct},
		{"Nothing Test", args{s: ""}, NilStruct},
		{"XX Test", args{s: "XX"}, NilStruct},
		{"XXX Test", args{s: "XXX"}, NilStruct},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructMatch(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
