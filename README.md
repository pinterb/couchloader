# couchloader
Load some data into a [Couchbase][1] server

## Description
With any database, there are common use cases where out-of-band data needs to be loaded.
This is a very dumb-downed utility for loading data into a [Couchbase][1] cluster.

**Please note:** This is just some proof-of-concept work.  It's very much pre-Alpha
code.

## Usage
Get to a `shell` prompt:

```bash
$ couchloader load -http-addr=127.0.0.1:8091 -pass=bucketpassword -bucket=mybucket
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/pinterb/couchloader
```

## Contribution

1. Fork ([https://github.com/pinterb/couchloader/fork](https://github.com/pinterb/couchloader/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Building
When I started this project, I didn't have [Go][3] installed locally.  So I
decided to try [golang-builder][2] from CenturyLink.  To use this build process
you'll need a couple of things installed:

- [Docker][4] or [boot2docker][5]
- [make][6]

Now, let's build:

```bash
$ make build
```

Need a container with your build artifact:

```bash
$ make container
```

## Author

[Brad Pinter](https://github.com/pinterb)

[1]: http://www.couchbase.com/
[2]: https://github.com/pinterb/golang-builder
[3]: http://golang.org/
[4]: https://docker.com
[5]: http://boot2docker.io/
[6]: https://www.gnu.org/software/make/