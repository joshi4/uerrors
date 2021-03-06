#uerrors 

Internal error messages are usually brief and in a format that is optimized for logging and searching. 

An example would be: 

~~~sh
> error=true, service=phantom, txid=12345, msg="user not authorized"
~~~ 

While that's a great error message for developers to search by tags, its not a great error message to display to an end user.

Usually, there's a translation layer that would take this error message and then return a more user friendly message to the end user. It might look something like: 

~~~go 
switch err:
  case ErrInvalidUser:
    umsg := "You are not authorized to perform this action"
  default: 
    umsg := err.Error()
// send umsg down in a response body along with an appropriate statsCode 
~~~


It's often hard to find all the code paths and errors an API request could create and then individually, add them to a swtich statement. 

This is where `uerrors` comes in. It offers an easy interface to attach a user-friendly message while creating the original error itself.  

## Usage 

To create a New uerror 

~~~go 

err := uerrors.New("pkg: invalid token", "Please renew your API token")
~~~ 


To create an uerror with values available only at run time: 

~~~go 

err := uerrors.FromErrors(fmt.Errorf("pkg: txnid=%s, resource=%s", txnid,resource), 
                          fmt.Errorf("Your API request to %s failed. Please contact support if the issue persists", resource)) 
~~~ 


## Constructing an user friendly response: 

~~~go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/joshi4/uerrors"
)

func main() {
	http.HandleFunc("/feature", serveFeature)
	http.HandleFunc("/", serveIndex)
	http.ListenAndServe(":6060", nil)
}

func serveIndex(w http.ResponseWriter, req *http.Request) {
	uerr := uerrors.New("example: redirect failed", "Please visit http://localhost:6060/feature if you aren't redirected automatically")
	log.Println(uerr.Error())
	io.WriteString(w, uerr.UserError())
}

func serveFeature(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	uerr := uerrors.FromErrors(fmt.Errorf("example: error=true, endpoint=%s", path),
		fmt.Errorf("Please visit %s a little later", path))
	log.Println(uerr.Error())
	io.WriteString(w, uerr.UserError())
}
~~~
