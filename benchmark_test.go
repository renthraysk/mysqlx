package mysqlx

import (
	"context"
	"testing"
)

func BenchmarkConnectAuthentication(b *testing.B) {
	b.ReportAllocs()
	connector := NewConnector(b)
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		c, err := connector.Connect(ctx)
		if err != nil {
			b.Fatalf("Connected failed: %s", err)
		}
		c.Close()
	}
}
