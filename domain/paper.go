package domain

type Paper struct {
	ID         int
	Title      string `json:"title"`
	DOI        string `json:"doi"`
	Supplement string `json:"supplement"`
}

type Papers []Paper
