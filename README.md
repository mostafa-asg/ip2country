[![Go Report Card](https://goreportcard.com/badge/github.com/mostafa-asg/ip2country)](https://goreportcard.com/report/github.com/mostafa-asg/ip2country)  
**ip2country** is a golang package to find out IP's origin country. It uses [db-ip.com's csv](https://db-ip.com/db/download/country)
 file to provide answers.
 ### Install
 ```
 go get -u github.com/mostafa-asg/ip2country
 ```
 ### Usage
 ```Go
package main

import (
	"github.com/mostafa-asg/ip2country"
)

func main() {
	ip2country.Load( PATH_TO_DB-IP.COM'S CSV FILE )
	println(ip2country.GetCountry("2.179.6.12"))
	println(ip2country.GetCountry("172.217.18.14"))
	println(ip2country.GetCountry("217.160.123.58"))
}
 ```
 **Tips :** *Load* method should be called once.
