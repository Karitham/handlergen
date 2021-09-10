# REST example

Simple example of a rest API for a book store

## Generation

We can simply run

```sh
handlergen -file rest.yaml -pkg main > handler.go
```

To get up and running helpers for querying books and posting them.

## Logging

This example shows the ability to setup a [custom logger](./logger.go)

Simply make an invalid request to the server, for example

```sh
curl localhost:5555/book -d '{"name}'
```

And view the structured logs

## Routing

This example also shows how easy it is to mount the generated handlers to your own router.

## Result

You can simply add books by running

```sh
curl localhost:5555/book -d '{"name":"Vagabond", "author":"Takehiko Inoue"}'
```

And then query them

```sh
curl 'localhost:5555/book?name=Vagabond'
```
