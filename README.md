# go-repo-janitor

A small Go CLI to somewhat automate my Go repo tasks

## Use

![./demo.gif](./demo.gif)

```bash
go-repo-janitor hello
```

## Install

- [Homebrew](https://brew.sh/): `brew install bbkane/tap/go-repo-janitor`
- [Scoop](https://scoop.sh/):

```
scoop bucket add bbkane https://github.com/bbkane/scoop-bucket
scoop install bbkane/go-repo-janitor
```

- Download Mac/Linux/Windows executable: [GitHub releases](https://github.com/bbkane/go-repo-janitor/releases)
- Go: `go install go.bbkane.com/go-repo-janitor@latest`
- Build with [goreleaser](https://goreleaser.com/) after cloning: `goreleaser --snapshot --skip-publish --clean`

## Notes

See [Go Developer Tooling](https://www.bbkane.com/blog/go-developer-tooling/) for notes on development tooling.