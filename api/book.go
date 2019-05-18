package api

import "net/http"

// Book type with Name, Author and ISBN
type Book struct {
	Title  string
	Author string
	ISBN   string
}

// ToJSON to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
	return nil
}

// FromJSON to be used for unmarshalling of Book type
func FromJSON(data []byte) Book {
	return Book{}
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {

}
