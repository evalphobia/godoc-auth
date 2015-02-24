godoc with Basic Auth
========================================

This is a godoc program with basic auth.  
Almost all code from original godoc.  
(see https://code.google.com/p/go/source/browse/?repo=tools#hg%2Fcmd%2Fgodoc)

Using enviroment params, It's suitable running on heroku.
 
Example usage
-------------

```bash
# download source
$ git clone https://github.com/evalphobia/godoc-auth.git

# compile
$ cd godoc-auth.git
$ go build

# set password
$ htpasswd -nb username password
username:$apr1$q7tg6//P$H9Yx9ZozN7CYoMMPCxjJl0

$ export BASIC_AUTH_USERNAME='username'
$ export BASIC_AUTH_PASSWORD='$apr1$q7tg6//P$H9Yx9ZozN7CYoMMPCxjJl0'

# run
$ ./godoc-auth -http=:8080

# test
$ curl http://localhost:8080
$ curl -u username:password http://localhost:8080
$ curl -u username:bad_password http://localhost:8080
```

