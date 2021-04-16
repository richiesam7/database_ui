package main

func connectionParams() map[string]string {
	connectParams := map[string]string{
		"host": "database",
		"database": "mydb",
		"port": "5432",
		"user": "richie",
		"password": "richie",
	}
	return connectParams;
}