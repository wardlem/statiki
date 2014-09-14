package main

// A taxonomy is an indexable attribute for a page.
// A taxonomy has a name and terms.
type Taxonomy interface {
    Name()
    Terms()
}

// A term is a key in a taxonomy.
type Term interface {
    Taxonomy() Taxonomy
    Name()
    Values() map[string]Page
}

// Combines the term and page interfaces.
type TermPage interface {
    Term
    Page
}

