startpg:
	sudo  pg_ctlcluster 11 main start

stoppg: 
	sudo  pg_ctlcluster 11 main stop

createdb:
	createdb mydb

insert:
	psql -d mydb -a -f testTable.sql
	psql -d mydb -a -f headphones.sql

all:
	startpg
	createdb
	insert