Loads environment variables from .env into ENV, automagically.

## Installation

`go get github.com/jpfuentes2/go-env`

## Usage

### Autoload

Autoloading your `$PWD/.env` is as simple as importing the `go-env/autoload` package:

```go
import _ "github.com/jpfuentes2/go-env/autoload"
```

### I want it my way

That's fine. Chillllll, Winston:

```go
package main

import "github.com/jpfuentes2/go-env"

func main() {
  env.ReadEnv("/path/to/my/file")
}
```

See the [examples](examples/) folder.

### Multiple Application Environments

Export `GOENV` from your shell to specify a different env file to be loaded.

For example:

```
export GOENV=test
```

Would result in `.env.test` being loaded instead of `.env`.

Please note that development is considered to map to `.env`. So exporting `GOENV=development` will mean that `.env` is loaded.

#### Local files

Whenever `.env` is autoloaded, `go-env` will also try to load `.env.local` for local overrides.

## Authors

Created and maintained by

* Jacques Fuentes - [@jpfuentes2](https://github.com/jpfuentes2)

## Credit

I've been sourcing `.env` files for quite some time in Ruby projects, but I did copy some of the vebiage/style from [dotenv](https://github.com/bkeepers/dotenv).

Also, I ~~used~~ copied some code from [goreman](https://github.com/mattn/goreman).

Thanks to [Yasuhiro Matsumoto](https://github.com/mattn) and [Brandon Keepers](https://github.com/bkeepers).

## License

MIT
