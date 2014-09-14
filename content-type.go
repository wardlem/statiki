package main

// ContentType is a special, default taxonomy.
// Content types are created by the directory structure in the input directory,
// though this can be overwritten in a configuration file within a directory,
// and through the frontmatter in a document.
// Each page has exactly one content type.
// Content types can have super and sub types.
// All content types descend from the base page type.
type ContentType interface {
    Term
    Super() ContentType
    Sub() []ContentType
}
