package entities

import "database/sql/driver"

type LoanProductType string

const (
	LoadProductTypeA LoanProductType = "product_a"
	LoadProductTypeB LoanProductType = "product_b"
)

func (s LoanProductType) IsValid() bool {
	return s == LoadProductTypeA || s == LoadProductTypeB
}

// Value implements the driver Valuer interface.
func (s LoanProductType) Value() (driver.Value, error) {
	return s.String(), nil
}

func (s LoanProductType) String() string {
	return string(s)
}
