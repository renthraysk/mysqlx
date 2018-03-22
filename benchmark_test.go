package mysqlx

import (
	"context"
	"testing"
)

func BenchmarkSimpleExec(b *testing.B) {
	b.ReportAllocs()
	db := NewDB(b)
	defer db.Close()
	for i := 0; i < b.N; i++ {
		if _, err := db.Exec("DO 1"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPreparedExec(b *testing.B) {
	b.ReportAllocs()
	db := NewDB(b)
	defer db.Close()
	stmt, err := db.Prepare("DO 1")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		if _, err := stmt.Exec(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSimpleQueryRow(b *testing.B) {
	b.ReportAllocs()
	db := NewDB(b)
	defer db.Close()
	var num int
	for i := 0; i < b.N; i++ {
		if err := db.QueryRow("SELECT 1").Scan(&num); err != nil {
			b.Fatal(err)
		}
	}
}

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

func BenchmarkQueryNoScan(b *testing.B) {
	b.ReportAllocs()

	db := NewDB(b)
	defer db.Close()

	for i := 0; i < b.N; i++ {
		rows, err := db.Query(SelectAll)
		if err != nil {
			b.Fatalf("query failed: %s", err)
		}

		if err := rows.Err(); err != nil {
			b.Fatalf("rows error: %+v", err)
		}
		rows.Close()
	}
}

func BenchmarkQueryScan(b *testing.B) {
	b.ReportAllocs()

	db := NewDB(b)
	defer db.Close()
	for i := 0; i < b.N; i++ {
		rows, err := db.Query(SelectAll)
		if err != nil {
			b.Fatalf("query failed: %s", err)
		}

		var f film

		for rows.Next() {
			if err := f.Scan(rows); err != nil {
				b.Fatalf("scan failed: %s", err)
			}
		}
		if err := rows.Err(); err != nil {
			b.Fatalf("rows error: %+v", err)
		}
		rows.Close()
	}
}
