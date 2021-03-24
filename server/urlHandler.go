package main
import "strings"

func validateurl(url string) bool {
	isValid := false;
	return isValid;
} 

func parseUrl(urlPath string) (string, string) {
	urls := strings.Split(urlPath, "/")
	tableName := urls[0];
	id := ""
	if (len(urls) > 1) {
		id = urls[1];
	} 
	return tableName, id
}