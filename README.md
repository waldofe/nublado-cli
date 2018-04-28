# nublado-cli
A Go CLI for checking weather

This is a WIP project with a simple and straightforward purpose. Checking weather via CLI. You can already play with it:

```bash
› ./nublado-cli -key=<api-key> rio de janeiro
It's 29.25°C right now in Rio de Janeiro, BR!

› ./nublado-cli -h
Usage of ./nublado-cli:
  -key string
    	Provide a valid https://openweathermap.org/api API key
```

### Installing

1. Build the binary

```
go build nublado-cli.go
```

2. Done


### API

You'll need a valid API from https://openweathermap.org/api.


### This project targets

- Supporting multiple APIs :two_hearts:
- Having voice output :lips:
- Having emojis :rocket: :guitar:
- Opinioned responses your momma used to say when you'd go out without your jacket :roll_eyes:
