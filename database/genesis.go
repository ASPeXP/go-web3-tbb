package main 

import (
	"io/ioutil"
	"encoding/json"
)

type genesis struct {
	Balances map[Account]uint `json:"balances"`
}

func loadGenesis