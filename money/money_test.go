package money

import (
	"testing"
)

func TestFromMilli(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want Amount
	}{
		{
			name: "0.001",
			args: args{
				v: 1,
			},
			want: Amount(1),
		},
		{
			name: "0.010",
			args: args{
				v: 10,
			},
			want: Amount(10),
		},
		{
			name: "0.100",
			args: args{
				v: 100,
			},
			want: Amount(100),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromMilli(tt.args.v); got != tt.want {
				t.Errorf("FromMilli() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_FromMilli(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want Amount
	}{
		{
			name: "0.001",
			args: args{
				v: 1,
			},
			want: Amount(1),
		},
		{
			name: "0.010",
			args: args{
				v: 10,
			},
			want: Amount(10),
		},
		{
			name: "0.100",
			args: args{
				v: 100,
			},
			want: Amount(100),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Amount
			got.FromMilli(tt.args.v)
			if got != tt.want {
				t.Errorf("Amount.FromMilli() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_Sprint(t *testing.T) {
	type fields struct {
		Thousand rune
		Decimal  rune
	}
	type args struct {
		a Amount
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "0.01",
			fields: fields{
				Thousand: ',',
				Decimal:  '.',
			},
			args: args{
				a: Amount(10),
			},
			want: "0.01",
		},
		{
			name: "0,01",
			fields: fields{
				Thousand: '.',
				Decimal:  ',',
			},
			args: args{
				a: Amount(10),
			},
			want: "0,01",
		},
		{
			name: "0,10",
			fields: fields{
				Thousand: '.',
				Decimal:  ',',
			},
			args: args{
				a: Amount(100),
			},
			want: "0,10",
		},
		{
			name: "1,00",
			fields: fields{
				Thousand: '.',
				Decimal:  ',',
			},
			args: args{
				a: Amount(1000),
			},
			want: "1,00",
		},
		{
			name: "-1,00",
			fields: fields{
				Thousand: '.',
				Decimal:  ',',
			},
			args: args{
				a: Amount(-1000),
			},
			want: "-1,00",
		},
		{
			name: "-1.123,00",
			fields: fields{
				Thousand: '.',
				Decimal:  ',',
			},
			args: args{
				a: Amount(-1123000),
			},
			want: "-1.123,00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				Thousand: tt.fields.Thousand,
				Decimal:  tt.fields.Decimal,
			}
			if got := f.Sprint(tt.args.a); got != tt.want {
				t.Errorf("Formatter.Sprint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromCenti(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want Amount
	}{
		{
			name: "0.01",
			args: args{
				v: 1,
			},
			want: Amount(10),
		},
		{
			name: "0.01",
			args: args{
				v: 10,
			},
			want: Amount(100),
		},
		{
			name: "0.10",
			args: args{
				v: 100,
			},
			want: Amount(1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromCenti(tt.args.v); got != tt.want {
				t.Errorf("FromCenti() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_FromCenti(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want Amount
	}{
		{
			name: "0.01",
			args: args{
				v: 1,
			},
			want: Amount(10),
		},
		{
			name: "0.01",
			args: args{
				v: 10,
			},
			want: Amount(100),
		},
		{
			name: "0.10",
			args: args{
				v: 100,
			},
			want: Amount(1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Amount
			got.FromCenti(tt.args.v)
			if got != tt.want {
				t.Errorf("Amount.FromCenti() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_AsMilli(t *testing.T) {
	tests := []struct {
		name string
		a    Amount
		want int64
	}{
		{
			name: "0.001 => 1",
			a:    Amount(1),
			want: 1,
		},
		{
			name: "0.004 => 4",
			a:    Amount(4),
			want: 4,
		},
		{
			name: "0.005 => 5",
			a:    Amount(5),
			want: 5,
		},
		{
			name: "0.009 => 9",
			a:    Amount(9),
			want: 9,
		},
		{
			name: "0.010 => 10",
			a:    Amount(10),
			want: 10,
		},
		{
			name: "0.100 => 100",
			a:    Amount(100),
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AsMilli(); got != tt.want {
				t.Errorf("Amount.AsMilli() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_AsCenti(t *testing.T) {
	tests := []struct {
		name string
		a    Amount
		want int64
	}{
		{
			name: "0.001 => 0",
			a:    Amount(1),
			want: 0,
		},
		{
			name: "0.004 => 0",
			a:    Amount(4),
			want: 0,
		},
		{
			name: "0.005 => 0",
			a:    Amount(5),
			want: 0,
		},
		{
			name: "0.009 => 0",
			a:    Amount(9),
			want: 0,
		},
		{
			name: "0.010 => 1",
			a:    Amount(10),
			want: 1,
		},
		{
			name: "0.100 => 10",
			a:    Amount(100),
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AsCenti(); got != tt.want {
				t.Errorf("Amount.AsCenti() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_AsDeci(t *testing.T) {
	tests := []struct {
		name string
		a    Amount
		want int64
	}{
		{
			name: "0.001 => 0",
			a:    Amount(1),
			want: 0,
		},
		{
			name: "0.004 => 0",
			a:    Amount(4),
			want: 0,
		},
		{
			name: "0.005 => 0",
			a:    Amount(5),
			want: 0,
		},
		{
			name: "0.009 => 0",
			a:    Amount(9),
			want: 0,
		},
		{
			name: "0.010 => 0",
			a:    Amount(10),
			want: 0,
		},
		{
			name: "0.100 => 1",
			a:    Amount(100),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AsDeci(); got != tt.want {
				t.Errorf("Amount.AsDeci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_AsInt(t *testing.T) {
	tests := []struct {
		name string
		a    Amount
		want int64
	}{
		{
			name: "0.001 => 0",
			a:    Amount(1),
			want: 0,
		},
		{
			name: "0.004 => 0",
			a:    Amount(4),
			want: 0,
		},
		{
			name: "0.005 => 0",
			a:    Amount(5),
			want: 0,
		},
		{
			name: "0.009 => 0",
			a:    Amount(9),
			want: 0,
		},
		{
			name: "0.010 => 0",
			a:    Amount(10),
			want: 0,
		},
		{
			name: "0.100 => 0",
			a:    Amount(100),
			want: 0,
		},
		{
			name: "1.999 => 1",
			a:    Amount(1999),
			want: 1,
		},
		{
			name: "1.000 => 1",
			a:    Amount(1000),
			want: 1,
		},
		{
			name: "1.001 => 1",
			a:    Amount(1001),
			want: 1,
		},
		{
			name: "2.999 => 2",
			a:    Amount(2999),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.AsInt(); got != tt.want {
				t.Errorf("Amount.AsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
