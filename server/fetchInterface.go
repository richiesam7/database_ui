package main

import "encoding/json"

func fetchTableDef() {
	return "SELECT column_name, data_type, is_nullable 
		from information_schema.columns where table_name = '" + tableName +"'";
	queryResponse := executeQuery(tableName, null);
	response := marshallResponse(queryResponse);
}

func prepareQuery(tableName string) string {
	return "SELECT * from " + tableName ;
}

func marshallResponse(response []interface {}) string {
	empData, err := json.Marshal(response)	
	if err != nil {
		panic(err)
	}
	jsonStr := string(empData)
	return jsonStr
}

func sendResponse() {}

func queryParams() {}

func headerOpts() {}