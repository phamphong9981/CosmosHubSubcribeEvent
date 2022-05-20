package test

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name   string
	Id     int
	Result struct {
		Data struct {
			Value struct {
				TxResult struct {
					Result struct {
						Log string
					}
				}
			}
		}
		// Data map[string]interface{}
	}
}

func Test2() {
	var message Message
	var data = `{
		"Name": "2.0",
		"Id": 1,
		"result": {
		}}`

	json.Unmarshal([]byte(data), &message)
	fmt.Print(message.Result.Data.Value.TxResult.Result.Log)
}
