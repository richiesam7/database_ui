package main
import "fmt";

func validateurl(url: string) {
	bool isValid = false;
	if (url.Path[3:])
		isValid := false;
	return isValid;
} 

func parseurl(url: string) bool {
	fmt.Fprintf('1:' + url.Path[1:] + ',2:' + url.Path[2:]);
	if (url.Path[3:])
		return false;

	return true;
}