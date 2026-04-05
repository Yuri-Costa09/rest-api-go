package model

import (
	"testing"
)

func TestNewProduct(t *testing.T) {

	// ARRANGE

	desc := "Very good product!"

	tests := []struct {
		Name        string
		ProductName string
		Price       float32
		Description *string
		WantErr     bool
	}{
		{
			Name:        "Should return new Product with description",
			ProductName: "Example Product",
			Price:       55.5,
			Description: &desc,
			WantErr:     false,
		},
		{
			Name:        "Should return new Product without description",
			ProductName: "Example Product",
			Price:       55.5,
			Description: nil,
			WantErr:     false,
		},
		{
			Name:        "Should return error if negative price",
			ProductName: "Example Product",
			Price:       -1.1,
			Description: nil,
			WantErr:     true,
		},
		{
			Name:        "Should return error if price is 0",
			ProductName: "Example Product",
			Price:       0.0,
			Description: nil,
			WantErr:     true,
		},
		{
			Name:        "Should return error if name is less than 3",
			ProductName: "Ex",
			Price:       10.0,
			Description: nil,
			WantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// ACT
			got, err := NewProduct(tt.Price, tt.ProductName, tt.Description)

			// ASSERTIONS

			if tt.WantErr && err == nil {
				t.Fatal("Expected error and there is no error.")
			}

			if !tt.WantErr && err != nil {
				t.Fatal("Didn't expected error and there is an error.")
			}

			if !tt.WantErr {
				if got.ID == "" {
					t.Error("Should have an ID and there is none.")
				}
				if got.Name != tt.ProductName {
					t.Errorf("got %q and want %q", got.Name, tt.ProductName)
				}
				if got.Price != tt.Price {
					t.Errorf("got %v and want %v", got.Price, tt.Price)
				}
				if got.Description == nil && tt.Description != nil {
					t.Error("Description doesn't matches")
				}
				if got.Description != nil && *got.Description != *tt.Description {
					t.Error("Description doesn't matches")
				}
			}
		})
	}
}
