### SQL Server Script Runner
This simple program uses the sqlcmd utility to run a directory of SQL scripts agaisnt a database.
They are executed in a lexicographically sorted order.

##### Requirements
* You must have sqlcmd on your system
* Your SQL Server instance must be setup for Windows Authentication (At the moment)
	
##### Usage

```
go build main.go
SET GOSQLCONNECTION=<databaseLocation>\<instanceName>:<databaseName>
main.exe C:\path\to\directory
```