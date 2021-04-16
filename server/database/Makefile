startpg:
	sudo  pg_ctlcluster 12 main start

stoppg: 
	sudo  pg_ctlcluster 12 main stop

createdb:
	sudo -u postgres createuser -s richie	
	createdb mydb

insert:
	psql -d mydb -a -f testTable.sql
	psql -d mydb -a -f headphones.sql

all:
	make startpg
	make createdb
	make insert

server_test:
	go run httpcallTest.go
