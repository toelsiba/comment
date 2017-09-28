# comment
Golang package for use comments in JSON config files

### Motivation
json config file: ``config.json``
```json
{
  "name": "example",
  "address": "192.168.0.1",
  "active": true
}
```
How we can change address?

Create a new file ``config2.json`` or fix in the current one.

Both variants are not convenient.

### What is proposed

Introduce end-line comment. For shielding use the symbol #.

new json-config file with comments ``config.cjson``
```
{
  "name": "example",
  # "address": "192.168.0.1", - this line is commented
  "address": "192.168.0.2", # New address!
  "active": true 
}
```
Easier and convenient than previous both variants.

For parse this file first use the function ``comment.Trim(data)``.
We obtain the following:
```json
{
  "name": "example",
  "address": "192.168.0.2", 
  "active": true 
}
```
Next, we use standard parser json.Unmarshal

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/toelsiba/comment"
)

type Config struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Active  bool   `json:"active"`
}

func main() {

	data, err := ioutil.ReadFile("config.cjson")
	if err != nil {
		log.Fatal(err)
	}

	data = comment.Trim(data)
	fmt.Println(string(data))

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", config)
}
```

If you need have symbol # in line - used doublet ##.
