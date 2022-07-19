## Dark Market

env file: 
```
DBUSER=
DBPASSWORD=
DBNAME=
DBHOST=
DBPORT=
SSLMODE=
DBTIMEZONE=
```

change transaction in [`models/tx.go`](https://github.com/cronwss/dark_web/models/tx.go)
default trasanction using BTC

change users requirements in [`models/users.go`](https://github.com/cronwss/dark_web/models/users.go)
default using username and password for decentralization and anonymity