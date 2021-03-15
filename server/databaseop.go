package main

import "fmt";
import "strconv";
import "database/sql";
import _ "github.com/lib/pq";

func executeQuery(tableName string, queryParams string)  []interface {} {
	params := connectionParams();
	portnum, err:= strconv.Atoi(params["port"]);
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    params["host"], portnum, params["user"], params["password"], params["database"])

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
	  panic(err)
	}
	defer db.Close()
	if err != nil {
		panic(err)
	}
    err = db.Ping()
	if err != nil {
	  panic(err)
	}
	rows, err := db.Query(prepareQuery(tableName));
	if err != nil {
		panic(err)
	}

	columnTypes, err := rows.ColumnTypes()
    if err != nil {
		panic(err)
    }

    count := len(columnTypes)
    finalRows := []interface{}{};

    for rows.Next() {
        scanArgs := make([]interface{}, count)
        for i, v := range columnTypes {
            switch v.DatabaseTypeName() {
            case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
                scanArgs[i] = new(sql.NullString)
                break;
            case "BOOL":
                scanArgs[i] = new(sql.NullBool)
                break;
            case "INT4":
                scanArgs[i] = new(sql.NullInt64)
                break;
            default:
                scanArgs[i] = new(sql.NullString)
            }
        }

        err := rows.Scan(scanArgs...)
        if err != nil {
            panic(err)
        }

        masterData := map[string]interface{}{}
        for i, v := range columnTypes {
            if z, ok := (scanArgs[i]).(*sql.NullBool); ok  {
                masterData[v.Name()] = z.Bool
                continue;
            }
            if z, ok := (scanArgs[i]).(*sql.NullString); ok  {
                masterData[v.Name()] = z.String
                continue;
            }
            if z, ok := (scanArgs[i]).(*sql.NullInt64); ok  {
                masterData[v.Name()] = z.Int64
                continue;
            }
            if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok  {
                masterData[v.Name()] = z.Float64
                continue;
            }
            masterData[v.Name()] = scanArgs[i]
        }
        finalRows = append(finalRows, masterData)
    }
    return finalRows
  }

