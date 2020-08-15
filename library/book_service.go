package library

func Summary() string {
	book := book{
		title:  "mission possible",
		isbn:   "sdf",
		author: "ram thapa",
	}
	return book.String()
}
