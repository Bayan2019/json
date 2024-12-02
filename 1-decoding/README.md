# Decoding JSON

When we receive JSON data in the body of an HTTP response, it comes as a stream of bytes. As we saw before, we can just convert the bytes to a string... but in Go there's a better way. It's typically best to decode the JSON data into a struct.

To decode this JSON into a slice of Item structs, we need to know the JSON fields and their types. The standard encoding/json package uses tags to map JSON fields to struct fields.