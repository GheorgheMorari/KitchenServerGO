package main

type Order struct {
	Id      int `json:"id"`
	TableId int `json:"table_id"`
	WaiterId int   `json:"waiter_id"`
	Items    []int `json:"items"`
	Priority   int   `json:"priority"`
	MaxWait    int   `json:"max_wait"`
	PickUpTime int64 `json:"pick_up_time"`
}
