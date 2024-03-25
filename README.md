# Go NetSuite ODBC connector for unixodbc

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Build](https://github.com/ninthclowd/unixodbc/actions/workflows/build.yml/badge.svg?branch=main)
![CodeQL](https://github.com/ninthclowd/unixodbc/actions/workflows/codeql.yml/badge.svg?branch=main)


This topic provides instructions for installing, running, and modifying the [go netsuiteodbc driver](https://github.com/ninthclowd/netsuiteodbc) for connecting to 
NetSuite ODBC databases through [unixodbc](https://www.unixodbc.org/). 

# Notice of Non-Affiliation and Disclaimer

This project is not affiliated, associated, authorized, endorsed by, or in any way officially connected with NetSuite,
Oracle, or any of its subsidiaries or its affiliates. The official NetSuite website can be found at
https://www.netsuite.com.

The names NetSuite and Oracle as well as related names, marks, emblems and images are registered trademarks of their
respective owners.

# Prerequisites

The following software packages are required to use the go unixodbc driver.
## Go

The latest driver requires the [Go language](https://golang.org/) 1.20 or higher.

## [go unixodbc](https://github.com/ninthclowd/unixodbc)
This connector requires [go unixodbc](https://github.com/ninthclowd/unixodbc) and it's corresponding dependencies.

## NetSuite ODBC Driver
The NetSuite ODBC driver must be installed and configured with [unixodbc](https://www.unixodbc.org/) according to NetSuite 
documentation.

# Configuration

## Token Based Authentication
To connect to the ODBC database with token based authentication, use
[netsuiteodbc.NewConnectionStringer](https://pkg.go.dev/github.com/ninthclowd/netsuiteodbc#NewConnectionStringer) with a 
[netsuiteodbc.Config](https://pkg.go.dev/github.com/ninthclowd/netsuiteodbc#Config) to create a dynamic connection string for use in 
[unixodbc.Connector](https://pkg.go.dev/github.com/ninthclowd/unixodbc#Connector):

```go
package main

import (
	"github.com/ninthclowd/unixodbc"
	"github.com/ninthclowd/netsuiteodbc"
	"database/sql"
	"time"
)

func main() {
	//connect to the database using TBA
	db = sql.OpenDB(&unixodbc.Connector{
		ConnectionString: netsuiteodbc.NewConnectionStringer(netsuiteodbc.Config{
			ConnectionString: "DSN=NetSuite",
			ConsumerKey:    "<hex consumer key>",
			ConsumerSecret: "<hex consumer secret>",
			TokenId:        "<hex token id>",
			TokenSecret:    "<hex token secret>",
			AccountId:      "123456_SB1",
		}),
	})
	// set connection timeout to expected lifetime of token
	db.SetConnMaxLifetime(5 * time.Minute)
}
```

# Development

The developer notes are hosted with the source code on [GitHub](https://github.com/ninthclowd/unixodbc). 

