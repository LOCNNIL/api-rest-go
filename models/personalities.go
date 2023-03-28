package models

type Personality struct {
	Name    string `json:"nome"`
	History string `json:"history"`
}

var Personalities []Personality
