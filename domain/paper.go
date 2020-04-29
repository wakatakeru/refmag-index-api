package domain

type Paper struct {
	ID         int
	Title      string
	DOI        string
	Supplement string
}

type Papers []Paper
