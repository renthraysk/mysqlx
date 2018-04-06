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
		if _, err := db.ExecContext(context.Background(), "DO 1"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPreparedExec(b *testing.B) {
	b.ReportAllocs()
	db := NewDBFatalErrors(b)
	defer db.Close()
	stmt, _ := db.PrepareContext(context.Background(), "DO 1")
	for i := 0; i < b.N; i++ {
		stmt.Exec()
	}
}

func BenchmarkSimpleQueryRow(b *testing.B) {
	b.ReportAllocs()
	db := NewDBFatalErrors(b)
	defer db.Close()
	var num int
	for i := 0; i < b.N; i++ {
		db.QueryRowContext(context.Background(), "SELECT 1").Scan(&num)
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

	db := NewDBFatalErrors(b)
	defer db.Close()

	for i := 0; i < b.N; i++ {
		rows, _ := db.QueryContext(context.Background(), SelectAll)
		if err := rows.Err(); err != nil {
			b.Fatalf("rows error: %+v", err)
		}
		rows.Close()
	}
}

func BenchmarkQueryScan(b *testing.B) {
	var f film

	b.ReportAllocs()

	db := NewDBFatalErrors(b)
	defer db.Close()
	for i := 0; i < b.N; i++ {
		rows, _ := db.QueryContext(context.Background(), SelectAll)
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
