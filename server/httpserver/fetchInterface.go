package main

import "encoding/json"

func fetchTableDefinition(tableName string)  string {
	query := prepareDescQuery(tableName) 
	queryResponse := executeQuery(query, "")
	response := marshallResponse(queryResponse);
	return response
}

func fetchTableContent(tableName string, id string) string {
	query := prepareFetchQuery(tableName, id)
	queryResponse := executeQuery(query, id);
	response := marshallResponse(queryResponse);
	// cacheSet(tableName + id, string(response))
	return response
}

func prepareFetchQuery(tableName string, id string) string {
	query := ""
	if (id != "") {
		query = "SELECT * from " + tableName + " WHERE id = " + id
	} else {
		query = "SELECT * from " + tableName
	}
	return query;
}

func prepareDescQuery(tableName string) string {
	return "SELECT column_name, data_type, is_nullable from information_schema.columns where table_name = '" + tableName +"'";
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