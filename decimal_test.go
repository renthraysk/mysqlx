package mysqlx

import (
	"fmt"
	"testing"
)

func TestDecimal(t *testing.T) {

	tests := []struct {
		BCD    []byte
		String string
	}{
		{[]byte{0x04, 0x12, 0x34, 0x01, 0xd0}, "-12.3401"},
	}

	var d Decimal

	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if err := d.Unmarshal(tt.BCD); err != nil {
				t.Fatalf("failed to scan :%s", err)
			}
			if s := d.String(); s != tt.String {
				t.Fatalf("expected %s, got %s", tt.String, s)
			}
		})
	}
}
