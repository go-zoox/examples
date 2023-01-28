package main

import (
	"example/orm"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Item struct {
	ID   int64
	Name string
}

func main() {
	db := orm.New("postgres", "postgres://eunomia:7ea1b4f7-e0d1-4bc8-88ef-4e1a142d63e7@127.0.0.1:5432/eunomia?sslmode=disable")
	if err := db.Connect(); err != nil {
		log.Fatal("connect error:", err)
	}

	// rows, headers, err := db.Query("SELECT id,name,status FROM v2_flow_task LIMIT 10")
	rows, headers, err := db.Query("SELECT TABLE_NAME FROM information_schema.tables WHERE TABLE_SCHEMA = 'public' AND TABLE_CATALOG = 'eunomia' AND TABLE_TYPE = 'BASE TABLE'")
	if err != nil {
		log.Fatal("prepare error:", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	tableHeader := table.Row{}
	for _, header := range headers {
		tableHeader = append(tableHeader, header.Name())
	}
	t.AppendHeader(tableHeader)

	tableRows := []table.Row{}
	for _, row := range rows {
		// for k, v := range row {
		// 	fmt.Printf("key = %s, value = %v (type: %s)\n", k, v, reflect.TypeOf(v))
		// }

		tableRow := table.Row{}
		for i := range headers {
			tableRow = append(tableRow, fmt.Sprintf("%s", row[headers[i].Name()]))
		}

		tableRows = append(tableRows, tableRow)
	}
	t.AppendRows(tableRows)

	t.Render()
}

// func main() {
// 	db, err := sql.Open("postgres", "postgres://eunomia:7ea1b4f7-e0d1-4bc8-88ef-4e1a142d63e7@127.0.0.1:5432/eunomia?sslmode=disable")
// 	if err != nil {
// 		log.Fatal("connect error:", err)
// 	}
// 	defer db.Close()

// 	stmt, err := db.Prepare("SELECT * FROM v2_flow_task WHERE id = $1")
// 	if err != nil {
// 		log.Fatal("prepare error:", err)
// 	}
// 	defer stmt.Close()

// 	rows, err := stmt.Query(1)
// 	if err != nil {
// 		log.Fatal("query error:", err)
// 	}

// 	//字段
// 	types, _ := rows.ColumnTypes()
// 	cols, _ := rows.Columns()
// 	for i := range cols {
// 		fmt.Println(cols[i], types[i].Name(), types[i].DatabaseTypeName(), types[i].ScanType())
// 	}

// 	values := make([]sql.RawBytes, len(cols))
// 	scans := make([]interface{}, len(cols))
// 	for i := range values {
// 		scans[i] = &values[i]
// 	}

// 	results := make(map[int]map[string]string)
// 	i := 0
// 	for rows.Next() {
// 		err := rows.Scan(scans...)
// 		if err != nil {
// 			log.Fatal("iterate row error:", err)
// 		}

// 		row := make(map[string]string)
// 		for j, v := range values {
// 			key := cols[j]
// 			row[key] = string(v)
// 		}
// 		results[i] = row

// 		i++
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for i, m := range results {
// 		fmt.Println(i)

// 		for k, v := range m {
// 			fmt.Println(k, " : ", v)
// 		}

// 		fmt.Println("========================")
// 	}
// }
