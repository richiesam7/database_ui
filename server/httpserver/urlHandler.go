package main
import "strings"

func validateurl(urls []string) bool {
	if (len(urls) > 2) {
		return false;
		} else {
		return true;
	}
} 

func parseUrl(urlPath string) (string, string) {
	urls := strings.Split(urlPath, "/")
	validateurl(urls)
	tableName := urls[0];
	id := ""
	if (len(urls) > 1) {
		id = urls[1];
	} 
	return tableName, id
}