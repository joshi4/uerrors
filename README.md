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

## Example 

~~~go 
~~~ 
