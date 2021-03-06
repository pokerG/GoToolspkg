[![Build Status](https://travis-ci.org/franela/goreq.png?branch=master)](https://travis-ci.org/franela/goreq)
GoReq
=======

Simple and sane HTTP request library for Go language.

Why GoReq?
==========

Go has very nice native libraries that allows you to do lots of cool things. But sometimes those libraries are too low level, which means that to do a simple thing, like an HTTP Request, it takes some time. And if you want to do something as simple as adding a timeout to a request, you will end up writing several lines of code.

This is why we think GoReq is useful. Because you can do all your HTTP requests in a very simple and comprehensive way, while enabling you to do more advanced stuff by giving you access to the native API.

How do I install it?
====================

```bash
go get github.com/franela/goreq
```

What can I do with it?
======================

## Making requests with different methods

#### GET
```go
res, err := goreq.Request{ Uri: "http://www.google.com" }.Do()
```

GoReq default method is GET.

You can also set value to GET method easily

```go
type Item struct {
        Limit int
        Skip int
        Fields string
}

item := Item {
        Limit: 3,
        Skip: 5,
        Fields: "Value",
}

res, err := goreq.Request{
        Uri: "http://localhost:3000/",
        QueryString: item,
}.Do()
```

The sample above will send `http://localhost:3000/?limit=3&skip=5&fields=Value`

#### POST

```go
res, err := goreq.Request{ Method: "POST", Uri: "http://www.google.com" }.Do()
```

## Sending payloads in the Body

You can send ```string```, ```Reader``` or ```interface{}``` in the body. The first two will be sent as text. The last one will be marshalled to JSON, if possible.

```go
type Item struct {
    Id int
    Name string
}

item := Item{ Id: 1111, Name: "foobar" }

res, err := goreq.Request{ 
    Method: "POST", 
    Uri: "http://www.google.com", 
    Body: item,
}.Do()
```

## Specifiying request headers

We think that most of the times the request headers that you use are: ```Host```, ```Content-Type```, ```Accept``` and ```User-Agent```. This is why we decided to make it very easy to set these headers.

```go
res, err := Request{
    Uri: "http://www.google.com",
    Host: "foobar.com",
    Accept: "application/json",
    ContentType: "application/json",
    UserAgent: "goreq",
}.Do()
```

But sometimes you need to set other headers. You can still do it.

```go
req := Request{ Uri: "http://www.google.com" }

req.AddHeader("X-Custom", "somevalue")

req.Do()
```

## Setting timeouts

GoReq supports 2 kind of timeouts. A general connection timeout and a request specific one. By default the connection timeout is of 1 second. There is no default for request timeout, which means it will wait forever.

You can change the connection timeout doing:

```go
goreq.SetConnectionTimeout(100 * time.Millisecond)
```

And specify the request timeout doing:

```go
res, err := goreq.Request{ 
    Uri: "http://www.google.com",
    Timeout: 500 * time.Millisecond, 
}.Do()
```

## Using the Response and Error

GoReq will always return 2 values: a ```Response``` and an ```Error```.
If ```Error``` is not ```nil``` it means that an error happened while doing the request and you shouldn't use the ```Response``` in any way.
You can check what happened by getting the error message:

```go
fmt.Println(err.Error())
```
And to make it easy to know if it was a timeout error, you can ask the error or return it:

```go
if serr, ok := err.(*goreq.Error); ok {
    if serr.Timeout() {
        ...
    }
}
return err
```

If you don't get an error, you can safely use the ```Response```.

```go
res.StatusCode //return the status code of the response
res.Body // gives you access to the body
res.Body.ToString() // will return the body as a string
res.Header.Get("Content-Type") // gives you access to all the response headers
```

## Receiving JSON

GoReq will help you to receive and unmarshal JSON.

```go
type Item struct {
    Id int
    Name string
}

var item Item

res.Body.FromJsonTo(item)
```

TODO:
-----

We do have a couple of [issues](https://github.com/franela/goreq/issues) pending we'll be addressing soon. But feel free to
contribute and send us PRs (with tests please :smile:).
