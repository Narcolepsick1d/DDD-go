package customer_test

import (
	"errors"
	"github.com/Narcolepsick1d/tavern/domain/customer"
	"testing"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}
	testCases := []testCase{
		{
			test:        "empty name validation",
			name:        "",
			expectedErr: customer.ErrInvalidPerson,
		}, {
			test:        "valid name",
			name:        "peter griffin",
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := customer.NewCustomer(tc.name)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v,got %v", tc.expectedErr, err)
			}
		})
	}
}
