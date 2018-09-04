BINDIR	= bin

init:
	mkdir -p bin local
	sqlite3 local/db.sqlite < scripts/schema.sql

build:
	go build -o $(BINDIR)/bkpr github.com/colvin/bkpr

clean:
	-rm bin/bkpr
