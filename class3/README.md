# Class 3 - Using some libs

#### How I import some library? Here an example

`$ go get -u github.com/gin-gonic/gin`

#### Gin -> Web Framework

Gin is used to create Endpoints, routes, validations and more.

```
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

#### Godotenv -> Use .env variables

`$ go get https://github.com/Valgard/godotenv`

Example

```
import (
	"github.com/Valgard/godotenv"
)

func main() {
	dotenv := godotenv.New()
	if err := dotenv.Load(".env"); err != nil {
		panic(err)
	}

	// You can also load several files
	if err := dotenv.Load(".env", ".env.dev"); err != nil {
		panic(err)
	}
}
```

#### Go-MySQL-Driver

A MySQL-Driver for Go's database/sql package

`$ go get -u github.com/go-sql-driver/mysql`
Example

```
import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ...

db, err := sql.Open("mysql", "user:password@<host>/dbname")
if err != nil {
	panic(err)
}
// See "Important settings" section.
db.SetConnMaxLifetime(time.Minute * 3)
db.SetMaxOpenConns(10)
db.SetMaxIdleConns(10)
```

#### CORS with GIN

```

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

```
