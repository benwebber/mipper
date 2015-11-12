# Mipper

Mipper is a command-line utility for bundling Mozilla addons into a single installable package.

[Multiple item packages](https://developer.mozilla.org/en-US/docs/Multiple_Item_Packaging) are bundles of addons contained in a single `.xpi`. Mipper lets you build multiple item packages from a JSON package manifest.

As well as building packages, Mipper provides a command-line interface for searching [addons.mozilla.org](http://addons.mozilla.org/).

## Examples

Mipper might help you:

* configure Firefox on a new personal computer,
* distribute standard addons to managed workstations or
* further your obsession with automation.

## Usage

### Building packages

The core workflow is to construct a JSON package manifest and run `mipper build`.

1. Add addons to a package manifest.

    ```
    $ mipper add -f package.json vimperator
    $ mipper add -f package.json RESTClient
    ```
2. Build the package. By default, Mipper will save the package to `<manifest>.xpi`.

    ```
    $ mipper build package.json
    ```
3. Open Firefox and install `package.xpi`

### Searching for addons

Mipper lets you search for addons using the [addons.mozilla.org API](https://developer.mozilla.org/en-US/docs/addons.mozilla.org_%28AMO%29_API_Developers%27_Guide/The_generic_AMO_API).

````
$ mipper search vimperator
4891	Vimperator
123891	Vimperator-ja
480100	Vimperator Chinese Help/中文帮助
481488	Vimperator 繁體中文幫助
358057	Vimium
404785	VimFx
235854	Pentadactyl
58678	Clipple
403306	Zutilo Utility for Zotero
```

Currently, Mipper returns the exact results provided by the API. This can be quite noisy, and a definite place for improvement.

You can also retrieve information about a particular addon:

```
$ mipper info tamper
Name: Tamper Data
Version: 11.0.1
Homepage: http://tamperdata.mozdev.org
Summary: Use tamperdata to view and modify HTTP/HTTPS headers and post parameters...
```

## Contributing

Mipper is written in Go. You will need to [install Go](https://golang.org/doc/install) to build it.

```
$ go get github.com/benwebber/mipper
$ cd $GOPATH/src/github.com/benwebber/mipper
$ go build
```

Mipper is a work in progress. Bug reports, feature requests, and pull requests are all welcome.
