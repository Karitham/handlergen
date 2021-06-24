# handlergen

## Install

With the go toolchain, run

```sh
go install github.com/Karitham/handlergen@latest
```

## Usage

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

The tool also works with open API, see the petstore example.

When using path, the program uses go-chi by default

**[More examples here](_example)**
