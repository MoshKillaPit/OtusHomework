package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBook_ProtoMarshal(t *testing.T) {
	type fields struct {
		ID     int
		Title  string
		Author string
		Year   int
		Size   int
		Rate   float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Test",
			fields:  fields{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				ID:     tt.fields.ID,
				Title:  tt.fields.Title,
				Author: tt.fields.Author,
				Year:   tt.fields.Year,
				Size:   tt.fields.Size,
				Rate:   tt.fields.Rate,
			}
			_, err := b.ProtoMarshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("ProtoMarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestBook_SravnenieJSON(t *testing.T) {
	book := Book{
		ID:     1,
		Title:  "Sravnenie",
		Author: "Sravnenie",
		Year:   1980,
		Size:   20,
		Rate:   2.2,
	}

	slicebook, _ := book.MarshalJSON()

	booktwo := Book{}

	booktwo.UnmarshalJSON(slicebook)

	assert.Equal(t, book, booktwo)
}

func TestBook_SravneniePROTO(t *testing.T) {
	book := Book{
		ID:     1,
		Title:  "Sravnenie",
		Author: "Sravnenie",
		Year:   1980,
		Size:   20,
		Rate:   2.2,
	}

	slicebook, _ := book.ProtoMarshal()

	booktwo := Book{}

	booktwo.ProtoUnmarshal(slicebook)

	assert.Equal(t, book, booktwo)
}
