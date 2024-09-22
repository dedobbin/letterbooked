package booksapi

import (
	"encoding/json"
	"strconv"
)

// Open Library Search API Result schema for a single document
type Document struct {
	DDC                   []string `json:"ddc"`                               // Relevant Dewey Decimal codes
	CoverIdx              int      `json:"cover_i"`                           // Cover Image reference value
	Fulltext              bool     `json:"has_fulltext"`                      // Does OpenLibrary have a full copy of the text?
	EditionCount          uint     `json:"edition_count"`                     // ?
	Title                 string   `json:"title"`                             // Title of the book
	AuthorName            []string `json:"author_name"`                       // List of all author names
	AuthorAlternativeName []string `json:"author_alternative_name,omitempty"` // List of names that the author would be known by in other regions
	FirstPublishYear      uint     `json:"first_publish_year"`                // Year the book was first published
	Key                   string   `json:"key"`                               // some reference key idk
	Ia                    []string `json:"ia"`                                // No idea
	AuthorKey             []string `json:"author_key"`                        // OLID of the Author
	PublicScanB           bool     `json:"public_scan_b"`                     // ?
	ISBN                  []string `json:"isbn"`                              // Collection of ISBNs that the book is published with
	CoverEditionKey       string   `json:"cover_edition_key"`                 // Edition that the cover is taken from
}

// Stringer interface compliance for document
func (d Document) String() string {
	docJson, _ := json.MarshalIndent(d, "", "\t")
	return string(docJson)
}

// Image URL for a book cover
func (d Document) ImageUrl() string {
	return "https://covers.openlibrary.org/b/id/" + strconv.Itoa(d.CoverIdx) + "-L.jpg"
}

// This struct represents will contain the query results and whatever.
type BookQueryResult struct {
	NumFound      uint       `json:"numFound"`      // Number of results
	Start         uint       `json:"start"`         // Index of the first document in the results
	NumFoundExact bool       `json:"numFoundExact"` // Is NumFound the exhaustive amount
	Docs          []Document `json:"docs"`          // The documents themselves
	Query         string     `json:"q"`             // The Query string
}

// Stringer interface compliance for QueryResult
func (q BookQueryResult) String() string {
	queryJSON, _ := json.MarshalIndent(q, "", "\t")
	return string(queryJSON)
}

// Open Library Search API Result schema for a given author
type Author struct {
	Name                  string   `json:"name"`                      // Authors first name
	BirthDate             string   `json:"birth_date,omitempty"`      // Authors Date of Birth
	DeathDate             string   `json:"death_date,omitempty"`      // Authors Death Date
	Date                  string   `json:"date,omitempty"`            // range of activity?
	AuthorAlternativeName []string `json:"alternate_names,omitempty"` // Alternate names for an author
	TopWork               string   `json:"top_work"`                  // Their top work
	Key                   string   `json:"key"`                       // Author portrait reference key
	WorkCount             uint     `json:"work_count"`                // Number of works published
}

// Stringer compliance
func (a Author) String() string {
	authorJson, _ := json.MarshalIndent(a, "", "\t")
	return string(authorJson)
}

// This constructs an olid API request string for an image of an author
func (a Author) ImageUrl() string {
	return "https://covers.openlibrary.org/a/olid" + a.Key + "-L.jpg"
}

// Struct containing author result query data
type AuthorQueryResult struct {
	NumFound      uint     `json:"numFound"`      // Number of authors that matched a query
	Start         uint     `json:"start"`         // Start idx?
	NumFoundExact bool     `json:"numFoundExact"` // Exact amount
	Docs          []Author `json:"docs"`          // List of authors
}

// Stringer interface compliance
func (q AuthorQueryResult) String() string {
	queryJSON, _ := json.MarshalIndent(q, "", "\t")
	return string(queryJSON)
}
