package test

import(
	"encoding/json"
	"fmt"
)

type Attribute struct {
	Key   string `json: "key"`
	Value string `json: "value"`
}
type Event struct {
	Type string `json: "type"`
	Attributes []Attribute `json: "attribute"`
}
type Log struct{
	Events []Event `json: "events"`
}
func Test1() {
	var events []Log
	log := "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"cosmos1zy8v468fnwjrgaza7e8rqa6tflw77hfvjyqxs5\"},{\"key\":\"amount\",\"value\":\"2118uatom\"},{\"key\":\"receiver\",\"value\":\"cosmos1tygms3xhhs3yv487phx3dw4a95jn7t7lpm470r\"},{\"key\":\"amount\",\"value\":\"400uatom\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"cosmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd88lyufl\"},{\"key\":\"amount\",\"value\":\"2118uatom\"},{\"key\":\"spender\",\"value\":\"cosmos1fl48vsnmsdzcv85q5d2q4z5ajdha8yu34mf0eh\"},{\"key\":\"amount\",\"value\":\"400uatom\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.staking.v1beta1.MsgUndelegate\"},{\"key\":\"sender\",\"value\":\"cosmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd88lyufl\"},{\"key\":\"sender\",\"value\":\"cosmos1fl48vsnmsdzcv85q5d2q4z5ajdha8yu34mf0eh\"},{\"key\":\"module\",\"value\":\"staking\"},{\"key\":\"sender\",\"value\":\"cosmos1zy8v468fnwjrgaza7e8rqa6tflw77hfvjyqxs5\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cosmos1zy8v468fnwjrgaza7e8rqa6tflw77hfvjyqxs5\"},{\"key\":\"sender\",\"value\":\"cosmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd88lyufl\"},{\"key\":\"amount\",\"value\":\"2118uatom\"},{\"key\":\"recipient\",\"value\":\"cosmos1tygms3xhhs3yv487phx3dw4a95jn7t7lpm470r\"},{\"key\":\"sender\",\"value\":\"cosmos1fl48vsnmsdzcv85q5d2q4z5ajdha8yu34mf0eh\"},{\"key\":\"amount\",\"value\":\"400uatom\"}]},{\"type\":\"unbond\",\"attributes\":[{\"key\":\"validator\",\"value\":\"cosmosvaloper178h4s6at5v9cd8m9n7ew3hg7k9eh0s6wptxpcn\"},{\"key\":\"amount\",\"value\":\"400uatom\"},{\"key\":\"completion_time\",\"value\":\"2022-05-18T06:54:10Z\"}]}]}]"
	err := json.Unmarshal([]byte(log), &events)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", events)
}
