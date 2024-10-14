package structcomparator

import (
	"reflect"
	"testing"
)

func TestBookComparator_Compare(t *testing.T) {
	type fields struct {
		mode CompareMode
	}
	type args struct {
		b1 Book
		b2 Book
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "Year",
			fields: fields{0},
			args: args{b1: Book{
				id:     0,
				title:  "",
				author: "",
				year:   2000,
				size:   0,
				rate:   0,
			}, b2: Book{
				id:     0,
				title:  "",
				author: "",
				year:   2010,
				size:   0,
				rate:   0,
			}},
			want: false,
		},
		{
			name:   "Size",
			fields: fields{1},
			args: args{b1: Book{
				id:     0,
				title:  "",
				author: "",
				year:   0,
				size:   1,
				rate:   0,
			}, b2: Book{
				id:     0,
				title:  "",
				author: "",
				year:   0,
				size:   1,
				rate:   0,
			}},
			want: true,
		},
		{
			name:   "Rate",
			fields: fields{2},
			args: args{b1: Book{
				id:     0,
				title:  "",
				author: "",
				year:   0,
				size:   0,
				rate:   2,
			}, b2: Book{
				id:     0,
				title:  "",
				author: "",
				year:   0,
				size:   0,
				rate:   1,
			}},
			want: false,
		},

		{
			name:   "Return",
			fields: fields{2},
			args: args{b1: Book{
				id:     0,
				title:  "",
				author: "",
				year:   0,
				size:   0,
				rate:   2,
			}, b2: Book{
				id:     0,
				title:  "",
				author: "",
				year:   0,
				size:   0,
				rate:   2,
			}},
			want: true,
		},
		{
			name:   "Return",
			fields: fields{-1},
			args: args{
				b1: Book{
					id:     0,
					title:  "",
					author: "",
					year:   0,
					size:   0,
					rate:   0,
				}, // Обратите внимание на расстановку пробелов и скобок
				b2: Book{
					id:     0,
					title:  "",
					author: "",
					year:   0,
					size:   0,
					rate:   0,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := BookComparator{
				mode: tt.fields.mode,
			}
			if got := bc.Compare(tt.args.b1, tt.args.b2); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBook_SetID(t *testing.T) {
	type fields struct {
		id     int
		title  string
		author string
		year   int
		size   int
		rate   float64
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Test",
			fields: fields{id: 1},
			args:   args{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			s := &Book{
				id:     tt.fields.id,
				title:  tt.fields.title,
				author: tt.fields.author,
				year:   tt.fields.year,
				size:   tt.fields.size,
				rate:   tt.fields.rate,
			}
			s.SetID(tt.args.id)
		})
	}
}

func TestBook_SetTitle(t *testing.T) {
	type fields struct {
		id     int
		title  string
		author string
		year   int
		size   int
		rate   float64
	}
	type args struct {
		title string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Test",
			fields: fields{title: "Privet"},
			args:   args{"Privet"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			s := &Book{
				id:     tt.fields.id,
				title:  tt.fields.title,
				author: tt.fields.author,
				year:   tt.fields.year,
				size:   tt.fields.size,
				rate:   tt.fields.rate,
			}
			s.SetTitle(tt.args.title)
		})
	}
}

func TestBook_SetAuthor(t *testing.T) {
	type fields struct {
		id     int
		title  string
		author string
		year   int
		size   int
		rate   float64
	}
	type args struct {
		author string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Author",
			fields: fields{author: "How"},
			args:   args{"How"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			s := &Book{
				id:     tt.fields.id,
				title:  tt.fields.title,
				author: tt.fields.author,
				year:   tt.fields.year,
				size:   tt.fields.size,
				rate:   tt.fields.rate,
			}
			s.SetAuthor(tt.args.author)
		})
	}
}

func TestBook_SetYear(t *testing.T) {
	type fields struct {
		id     int
		title  string
		author string
		year   int
		size   int
		rate   float64
	}
	type args struct {
		year int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Year",
			fields: fields{year: 2012},
			args:   args{2012},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			s := &Book{
				id:     tt.fields.id,
				title:  tt.fields.title,
				author: tt.fields.author,
				year:   tt.fields.year,
				size:   tt.fields.size,
				rate:   tt.fields.rate,
			}
			s.SetYear(tt.args.year)
		})
	}
}

func TestBook_SetSize(t *testing.T) {
	type fields struct {
		id     int
		title  string
		author string
		year   int
		size   int
		rate   float64
	}
	type args struct {
		size int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Size",
			fields: fields{size: 22},
			args:   args{22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			s := &Book{
				id:     tt.fields.id,
				title:  tt.fields.title,
				author: tt.fields.author,
				year:   tt.fields.year,
				size:   tt.fields.size,
				rate:   tt.fields.rate,
			}
			s.SetSize(tt.args.size)
		})
	}
}

func TestBook_SetRate(t *testing.T) {
	type fields struct {
		id     int
		title  string
		author string
		year   int
		size   int
		rate   float64
	}
	type args struct {
		rate float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Rate",
			fields: fields{rate: 2000},
			args:   args{2000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			s := &Book{
				id:     tt.fields.id,
				title:  tt.fields.title,
				author: tt.fields.author,
				year:   tt.fields.year,
				size:   tt.fields.size,
				rate:   tt.fields.rate,
			}
			s.SetRate(tt.args.rate)
		})
	}
}

func TestBook_GetID(t *testing.T) {
	type fields struct {
		id     int
		title  string
		author string
		year   int
		size   int
		rate   float64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "GetID",
			fields: fields{id: 2},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Book{
				id:     tt.fields.id,
				title:  tt.fields.title,
				author: tt.fields.author,
				year:   tt.fields.year,
				size:   tt.fields.size,
				rate:   tt.fields.rate,
			}
			if got := s.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBook_GetTitle(t *testing.T) {
	type fields struct {
		id     int
		title  string
		author string
		year   int
		size   int
		rate   float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{
				id:     0,
				title:  "DSA",
				author: "",
				year:   0,
				size:   0,
				rate:   0,
			},
			want: "DSA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Book{
				id:     tt.fields.id,
				title:  tt.fields.title,
				author: tt.fields.author,
				year:   tt.fields.year,
				size:   tt.fields.size,
				rate:   tt.fields.rate,
			}
			if got := s.GetTitle(); got != tt.want {
				t.Errorf("GetTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBook_GetAuthor(t *testing.T) {
	type fields struct {
		id     int
		title  string
		author string
		year   int
		size   int
		rate   float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{
				id:     0,
				title:  "",
				author: "John",
				year:   0,
				size:   0,
				rate:   0,
			},
			want: "John",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Book{
				id:     tt.fields.id,
				title:  tt.fields.title,
				author: tt.fields.author,
				year:   tt.fields.year,
				size:   tt.fields.size,
				rate:   tt.fields.rate,
			}
			if got := s.GetAuthor(); got != tt.want {
				t.Errorf("GetAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewComparator(t *testing.T) {
	type args struct {
		mode CompareMode
	}
	tests := []struct {
		name string
		args args
		want BookComparator
	}{
		{
			name: "NewComparator",
			args: args{1},
			want: BookComparator{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComparator(tt.args.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
