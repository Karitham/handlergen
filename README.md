# handlergen

describe the function you want

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

```go
func example1(handler func(http.ResponseWriter, *http.Request, int, gen.Template)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		query_user_id := query.Get("user_id")
		user_id, err := strconv.Atoi(query_user_id)
		if err != nil {
			http.Error(w, "invalid query", 400)
			return
		}

		var body gen.Template
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid body", 400)
			return
		}

		handler(
			w,
			r,
			user_id,
			body,
		)
	}
}
```

you only have to create your handler with the right function signature.
