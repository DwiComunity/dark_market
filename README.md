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
PORT=
SECRET_KEY=
```

this is some code for generate random secret key:
```
package main;import(
"fmt"
"time"
"math/rand")var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890");func randSeq(n int) string {b := make([]rune, n);for i := range b{b[i] = letters[rand.Intn(len(letters))]};return string(b)};func main(){rand.Seed(time.Now().UnixNano());fmt.Println(randSeq(12))}
```


change transaction in [`models/tx.go`](https://github.com/Crownss/dark_market/blob/master/models/tx.go)
default trasanction using BTC

change users requirements in [`models/users.go`](https://github.com/Crownss/dark_market/blob/master/models/users.go)
default using username and password for decentralization and anonymity