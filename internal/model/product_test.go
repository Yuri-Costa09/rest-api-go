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
			t.Parallel()

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

func TestValidateName(t *testing.T) {
	tests := []struct {
		Name        string
		ProductName string
		WantErr     bool
	}{
		{
			Name:        "Error should be nil given valid Name",
			ProductName: "Valid Name",
			WantErr:     false,
		},
		{
			Name:        "Error should be nil given valid 3 characters Name",
			ProductName: "Valid Name",
			WantErr:     false,
		},
		{
			Name:        "Should return err given name less than 3 characters",
			ProductName: "In",
			WantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel() // run tests in parallel

			err := validateName(tt.ProductName)

			if tt.WantErr && err == nil {
				t.Fatal("Expected error and there is no error")
			}

			if !tt.WantErr && err != nil {
				t.Fatal("Didn't expected error and there is an error")
			}
		})
	}
}

func TestValidatePrice(t *testing.T) {

	tests := []struct {
		Name    string
		Price   float32
		WantErr bool
	}{
		{
			Name:    "Error should be nil given valid price",
			Price:   -1.0,
			WantErr: true,
		},
		{
			Name:    "Should return err when price is negative",
			Price:   -1.0,
			WantErr: true,
		},
		{
			Name:    "Should return err when price is zero",
			Price:   0.0,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel() // run tests in parallel

			err := validatePrice(tt.Price)

			if tt.WantErr && err == nil {
				t.Fatal("Expected error and there is no error")
			}

			if !tt.WantErr && err != nil {
				t.Fatal("Didn't expected error and there is an error")
			}
		})
	}

}

func TestProduct_UpdatePrice(t *testing.T) {

	desc := "Description"
	var validPrice float32 = 10.0

	tests := []struct {
		Name    string
		Price   float32
		WantErr bool
	}{
		{
			Name:    "Error should be nil given valid price",
			Price:   10.0,
			WantErr: false,
		},
		{
			Name:    "Should return err when price is negative",
			Price:   -1.0,
			WantErr: true,
		},
		{
			Name:    "Should return err when price is zero",
			Price:   0.0,
			WantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			got, _ := NewProduct(validPrice, "Cool Product", &desc)

			err := got.UpdatePrice(tt.Price)

			if tt.WantErr && err == nil {
				t.Fatal("Expected error and there is no error")
			}

			if !tt.WantErr && err != nil {
				t.Fatal("Didn't expected error and there is an error")
			}

			if !tt.WantErr && tt.Price != got.Price {
				t.Errorf("got %v, want %v", got.Price, tt.Price)
			}

		})
	}
}

func TestProduct_UpdateDesc(t *testing.T) {
	desc := "Description"
	newDesc := "New description"

	tests := []struct {
		Name          string
		InitialDesc   *string
		NewDesc       *string
		WantDescIsNil bool
		WantDescValue string
	}{
		{
			Name:          "should update description when a new description is provided",
			InitialDesc:   &desc,
			NewDesc:       &newDesc,
			WantDescIsNil: false,
			WantDescValue: newDesc,
		},
		{
			Name:          "should update description to nil",
			InitialDesc:   &desc,
			NewDesc:       nil,
			WantDescIsNil: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			p, err := NewProduct(10.0, "Cool Product", tt.InitialDesc)
			if err != nil {
				t.Fatalf("unexpected error creating product: %v", err)
			}

			initialID := p.ID
			initialName := p.Name
			initialPrice := p.Price

			if err := p.UpdateDesc(tt.NewDesc); err != nil {
				t.Fatalf("unexpected error updating description: %v", err)
			}

			if p.ID != initialID {
				t.Errorf("got ID %q, want %q", p.ID, initialID)
			}

			if p.Name != initialName {
				t.Errorf("got Name %q, want %q", p.Name, initialName)
			}

			if p.Price != initialPrice {
				t.Errorf("got Price %v, want %v", p.Price, initialPrice)
			}

			if tt.WantDescIsNil {
				if p.Description != nil {
					t.Fatal("expected description to be nil")
				}
				return
			}

			if p.Description == nil {
				t.Fatal("expected description to be set")
			}

			if *p.Description != tt.WantDescValue {
				t.Errorf("got %q, want %q", *p.Description, tt.WantDescValue)
			}
		})
	}

}

func TestProduct_UpdateName(t *testing.T) {
	desc := "Description"

	tests := []struct {
		Name               string
		InitialProduct     *Product
		NewName            string
		WantErr            bool
		WantName           string
		WantIDUnchanged    bool
		WantPriceUnchanged bool
	}{
		{
			Name: "should update name when the new name is valid",
			InitialProduct: func() *Product {
				p, _ := NewProduct(10.0, "Old Name", &desc)
				return p
			}(),
			NewName:            "New Product Name",
			WantErr:            false,
			WantName:           "New Product Name",
			WantIDUnchanged:    true,
			WantPriceUnchanged: true,
		},
		{
			Name: "should return an error when the new name is too short",
			InitialProduct: func() *Product {
				p, _ := NewProduct(10.0, "Old Name", &desc)
				return p
			}(),
			NewName:            "AB",
			WantErr:            true,
			WantName:           "Old Name",
			WantIDUnchanged:    true,
			WantPriceUnchanged: true,
		},
		{
			Name: "should validate the incoming name even if the current name is invalid",
			InitialProduct: &Product{
				ID:          "existing-id",
				Name:        "AB",
				Price:       10.0,
				Description: &desc,
			},
			NewName:            "Valid Name",
			WantErr:            false,
			WantName:           "Valid Name",
			WantIDUnchanged:    true,
			WantPriceUnchanged: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			p := tt.InitialProduct
			initialID := p.ID
			initialPrice := p.Price
			initialName := p.Name

			err := p.UpdateName(tt.NewName)

			if tt.WantErr && err == nil {
				t.Fatal("expected error and there is no error")
			}

			if !tt.WantErr && err != nil {
				t.Fatalf("didn't expect error and there is an error: %v", err)
			}

			if tt.WantErr {
				if p.Name != initialName {
					t.Errorf("got Name %q, want %q", p.Name, initialName)
				}
				return
			}

			if tt.WantIDUnchanged && p.ID != initialID {
				t.Errorf("got ID %q, want %q", p.ID, initialID)
			}

			if tt.WantPriceUnchanged && p.Price != initialPrice {
				t.Errorf("got Price %v, want %v", p.Price, initialPrice)
			}

			if p.Name != tt.WantName {
				t.Errorf("got Name %q, want %q", p.Name, tt.WantName)
			}
		})
	}

}
