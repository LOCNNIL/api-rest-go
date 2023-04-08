package models

type Personality struct {
	Id 		int 	`json:"id"`
	Name    string 	`json:"nome"`
	History string 	`json:"history"`
}

var Personalities []Personality
