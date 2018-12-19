# hashlig 

Simple file checksum integrity tool.

## But why yet another hashing tool?

I wanted to create a command-line tool using go, and at the same time I needed to hash a file.

# WIP!

The tool is currently in development. Check [enhancements](https://github.com/krilor/hashlig/labels/enhancement) for TODOs and [bugs](https://github.com/krilor/hashlig/labels/bug) for current issues.

# How to use `hashlig`

Provide the input file, the hash or checksum to compare against. Hashlig will indentify the hash algorithm, compute the file hash, compare and output the result.

```hashlig --file some-file.zip --hash 3589E4FAF1495E6EE3F3F538CF4C3F77576DD35AF8B0238AC3B6F916AA483027```

## Flags

| Flag        | Short           | Description  | If missing |
| ------------- |-------------| -----| - |
| file      | f | The relative or absolute path to the file that should be hashed. | Hashlig will look for the newest file in the current directory. |
| hash      | h      | Upper or lower case string representation of hex hash value. | Output hash value if `algorithm` is specified. If not, output all supported hashes. |
| algorithm | a      | One of <ul><li>md5</li><li>sha1</li><li>sha256</li><li>sha512</li></ul>| Guess the hashing algoritm if `hash` is specified. |

If none of the flags are given, but there are arguments, then hashlig will make an educated guess as to what argument is represents each flag.

This means that the example command can be written as

```hashlig some-file.zip 3589E4FAF1495E6EE3F3F538CF4C3F77576DD35AF8B0238AC3B6F916AA483027```

# Install

`go install hashlig.go`
