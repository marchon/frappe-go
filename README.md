# frappe-go
Frappe framework API wrapper for Golang 

Quickstart : 


```go
package main
import (
"fmt"
"github.com/srajelli/frappe-go"
)
func main(){
	config := frappe.Config{
		Application_url: "https://demo.erpnext.com",
		User:            "Administrator",
		Password:        "password",
  }
  frappe.Connect(config)

// getting user details 

doc := frappe.FrappeInput{}
doc.Doctype = "User"
doc.Resource = "john@gmail.com"

fmt.Println(frappe.Get(doc))

```

Examples : 
```go
// create new user

doc := frappe.FrappeInput{}
doc.Doctype = "User"
user := map[string]string{
	"doctype": "User", 
	"name": "john@gmail.com", 
	"first_name": "John", 
	"email": "john@gmail.com", 
	"new_password": "SuperSecurePassword"
}

c, _ := json.Marshal(user)
urlE := url.Values{}
urlE.Set("data", string(c))
frappe.Post(doc, urlE)

```
```go
// update user details

doc := frappe.FrappeInput{}
doc.Doctype = "User"
user := map[string]string{
	"doctype": "User", 
	"name": "john@gmail.com", 
	"first_name": "New Name For John"
}
c, _ := json.Marshal(user)
urlE := url.Values{}
urlE.Set("data", string(c))
frappe.Put(doc, urlE)

```
