package main

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"quote"`
}
