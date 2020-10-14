### Go-examples

Here are some exercises I did while learning the Go language。

#### Map

- syntax

  `map[k]v`
  
  ```go
  m := map[string]string {
      "name": "evan",
      "site": "evantian.me"
  }  
  
  m1 := make(map[string]int)
  
  // get the value
  m["name"]
  
  // if key doesn't exist, return the initial value
  m["age"]
  
  // to check if key exists
  value, ok = m[key]
  
  // to delete a key
  delete(m, key)
  ```