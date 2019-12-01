# TapSearch

> A tapchief challenge

This repo contains the solution to the challenge by TapChief, listed below.

> You are to build a simple program called TapSearch that achieves these objectives.
>
> It takes in multiple paragraphs of text, assigns a unique ID To each paragraph and stores the words to paragraph mappings on an inverted index. This is similar to what elasticsearch does. This paragraph can also be referred to as a ‘document’
>
> Given a word to search for, it lists out the top 10 paragraphs in which the word is present

## Building

```
go build 
./tapchief-chall
```

Go to http://localhost:8080

## API

### /api/index

### /api/search

### /api/clear


## Testing

Tests for search is provided.

```
go test ./search
```