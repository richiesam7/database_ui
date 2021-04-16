install:
	sudo apt-get install git
	sudo apt install curl
	sudo apt-get install postgresql

database:
	cd test; make all;

install_server:
	cd server; make install;

install_client:
	cd client; make install;

server:
	cd server; make start;

client:
	cd client; make start
