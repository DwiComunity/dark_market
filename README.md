## Dark Market

.env file: 
```
DBUSER=
DBPASSWORD=
DBNAME=
DBHOST=
DBPORT=
SSLMODE=
DBTIMEZONE=

RUN_HOST=
RUN_PORT=

SECRET_KEY=
```

this is some code for generate random secret key:
```
package main;import("fmt";"time";"math/rand");var letters=[]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890");func randSeq(n int)string{b:=make([]rune, n);for i:=range b{b[i]=letters[rand.Intn(len(letters))]};return string(b)};func main(){rand.Seed(time.Now().UnixNano());fmt.Println(randSeq(6))}
```

change transaction in [`models/tx.go`](https://github.com/crownss/dark_web/blob/master/models/tx.go)
default trasanction using BTC

change users requirements in [`models/users.go`](https://github.com/crownss/dark_web/blob/master/models/users.go)
default using username and password for decentralization and anonymity