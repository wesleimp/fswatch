# fswatch

Run commands when file changes

## Install

### Compiling from source

**clone**

```sh
git clone git@github.com:wesleimp/fswatch.git
```

**dependencies**

```sh
go mod download
```

**install**

```sh
make install
```

**verify it works**

```sh
fswatch -v
```

## Example usage

**go**

Run tests when some go file changes

```sh
fswatch go test ./...
```

**node**

Restart server when some file changes

```sh
fswatch node index.js
```

And much more to come...
