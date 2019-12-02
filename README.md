# TapSearch

> A tapchief challenge

This repo contains the solution to the challenge by TapChief, listed below.

> You are to build a simple program called TapSearch that achieves these objectives.
>
> It takes in multiple paragraphs of text, assigns a unique ID To each paragraph and stores the words to paragraph mappings on an inverted index. This is similar to what elasticsearch does. This paragraph can also be referred to as a ‘document’
>
> Given a word to search for, it lists out the top 10 paragraphs in which the word is present

Hosted at: https://tapchief-chall.herokuapp.com

## Building

```
go build 
./tapchief-chall
```

Go to http://localhost:8080

## API

### POST: /api/index

#### Params:

para: The paragraph to be indexed.

#### Description

This endpoint indexes a paragraph. A single paragraph can contain multiple documents seperated by 2 newlines (`"\n\n"`).

### GET: /api/search

#### Params:

query: The word to be searched.

#### Description

This endpoint searches for a word in the all the documents and returns the documents which match with the position of the word.

Example:

```json
{
    "status":"success",
    "response": [
        {
            "document":"para 3 aa dddf",
            "position":[2]
        },
        {
            "document":"para 2 aa dddf",
            "position":[2]
        }
    ]
}
```

### GET: /api/clear

#### Params: None

#### Description

This endpoint clears all the indexed documents.


## Testing

Tests for search is provided.

```
go test ./search
```