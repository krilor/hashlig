# hashlig

Simple file checksum integrity tool. 

## But why yet another hashing tool?

I wanted to create a command-line tool using go, and at the same time I needed to hash a file.

# Use

`hashlig --file [thefilepath]`

# Install

`go install hashlig.go`

# TODOs

* Add flag for hash
* Tests
* Let the tool guess hashing mechanism based on hash
* Compare result to hash and spit out result
