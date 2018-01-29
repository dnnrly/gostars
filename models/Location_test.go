package models_test

import (
	"database/sql/driver"
	"reflect"
	"testing"

	"github.com/dnnrly/gostars/models"
	"github.com/golang/geo/r2"
)

func TestLocation_Value(t *testing.T) {
	tests := []struct {
		name    string
		p       *models.Location
		want    driver.Value
		wantErr bool
	}{
		{name: "Simple convert", p: &models.Location{X: 99.4, Y: 223.0}, want: "(99.400000000000, 223.000000000000)", wantErr: false},
	}
	for _, tt := range tests {
		got, err := tt.p.Value()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Location.Value() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Location.Value() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestLocation_Scan(t *testing.T) {
	var p models.Location
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		p       *models.Location
		args    args
		wantErr bool
	}{
		{name: "String convert", p: &p, args: args{v: "(99.400000000000, 223.000000000000)"}, wantErr: false},
		{name: "[]uint8 convert", p: &p, args: args{v: []uint8("(99.400000000000, 223.000000000000)")}, wantErr: false},
		{name: "Bad string", p: &p, args: args{v: "99.400000000000, 223.000000000000)"}, wantErr: true},
		{name: "Wrong type", p: &p, args: args{v: 5.2}, wantErr: true},
	}
	for _, tt := range tests {
		if err := tt.p.Scan(tt.args.v); (err != nil) != tt.wantErr {
			t.Errorf("%q. Location.Scan() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestLocation_ToPoint(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   r2.Point
	}{
		{name: "Simple convert", fields: fields{X: 99.4, Y: 223.0}, want: r2.Point{X: 99.4, Y: 223.0}},
	}
	for _, tt := range tests {
		l := &models.Location{
			X: tt.fields.X,
			Y: tt.fields.Y,
		}
		if got := l.ToPoint(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Location.ToPoint() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
