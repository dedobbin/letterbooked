package booksapi

import "encoding/json"

// Open Library Search API Result schema for a single document
type Document struct {
	CoverIdx              uint     `json:"cover_i"`
	Fulltext              bool     `json:"has_fulltext"`
	EditionCount          uint     `json:"edition_count"`
	Title                 string   `json:"title"`                             // Title of the book
	AuthorName            []string `json:"author_name"`                       // List of all author names
	AuthorAlternativeName []string `json:"author_alternative_name,omitempty"` // List of names that the author would be known by in other regions
	FirstPublishYear      uint     `json:"first_publish_year"`                // Year the book was first published
	Key                   string   `json:"key"`                               // some reference key idk
	Ia                    []string `json:"ia"`                                // No idea
	AuthorKey             []string `json:"author_key"`
	PublicScanB           bool     `json:"public_scan_b"`
	ISBN                  []string `json:"isbn"` // Collection of ISBNs that the book is published with
}

// Stringer interface compliance for document
func (d Document) String() string {
	docJson, _ := json.MarshalIndent(d, "", "\t")
	return string(docJson)
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
	Key                   string   `json:"key"`                       // some reference key
	WorkCount             uint     `json:"work_count"`                // Number of works published
}

// Stringer compliance
func (a Author) String() string {
	authorJson, _ := json.MarshalIndent(a, "", "\t")
	return string(authorJson)
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
