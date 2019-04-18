# Split string
This activity provides your flogo application the ability to split a path into array


## Installation

```bash
flogo install github.com/ultrafab/stringtoarray_flogo
```
Link for flogo web:
```
https://github.com/ultrafab/stringtoarray_flogo
```

## Schema
Inputs and Outputs:

```json
{
 "input":[
    {
      "name": "input_string",
      "type": "string"
    },
    {
      "name": "split_char",
      "type": "string"
    }
  ],
  "output": [
    {
      "name": "result",
      "type": "array"
    }
  ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| input_string    | the input string to parse |
| split_char    | the char to use for split the string |

## Ouputs
| Output   | Description    |
|:----------|:---------------|
| result    | the array with the strings |
