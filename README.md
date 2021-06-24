
# HandlerGen

A go net/http HandlerFunc code generator.
Parses all your args and give them to you as function arguments,
which you can then validate and use.


## Features

Multiple arguments format supported:
- Query URL
- HTTP headers
- Request Body
- URL Path, using go-chi

Fully net/http complient

Codegen, no slow downs!

  
## Installation 

Install Handlergen with the go toolchain

```bash 
go install github.com/Karitham/handlergen@latest
```
## Usage/Examples

Describe the function you want

```yml
functions:
  example1:
    query:
      user_id:
        type: int
    body:
      type: gen.Template
      import: github.com/Karitham/handlergen/gen
```

and codegen a httphandler

```sh
handlergen -file _example/basic/basic.yaml > _example/basic/generated.go
```

**[View more examples](_example)**
## Documentation

| flag | default value | description |
| ---- | --------- | ----------- |
| file | handlers.yaml  | handlers gen config file |
| format  | handlergen  | input file format, use `openapi` for an open api file |
| pkg  | handlers  | package name |

## Authors

- [@Karitham](https://www.github.com/Karitham)

  
## License

[ISC](https://choosealicense.com/licenses/isc/)

  
