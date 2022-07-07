# Simple Go Api
This repository has the objective to register a go exercise made for study reasons.

## What is it

This repository is a go exercise made for study reason. The exercise was create a simple go api that can receive a roman numerals, convert it into a regular number and return which the hightest number insite string sent.

## Setup

Clone this repository then inside run `go build` and `go run` after. You'll be able to make a post to localhost using something like the above code (considering you have curl installed):

```bash
curl -d '{"text":"AXXBLXBMD"}' -H "Content-Type: application/json" -X POST http://localhost:8080/search
```
