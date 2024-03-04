initDB:
	sqlite3 tasks.db < init.sql

testProject:
	go test ./...

coverage:
	go test -coverprofile="coverage" ./...
	go tool cover -html="coverage" 
