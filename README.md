# Learn X in Y Minutes CLI
For when you want to learn something new but you are either too lazy to read the documentation or the tutorials are too slow for you.

[learnxinyminutes.com](https://learnxinyminutes.com) offers some great quick tutorials, I wanted to make it easier to access by making a CLI tool so you can read anytime.

It also highlights the code in your terminal so it's easier to see.

## Install
You can download and run a pre-built binary from [GitHub releases](https://github.com/ravener/lxym/releases)

You can also install via Homebrew: `brew install ravener/tap/lxym`

If that's not enough for you, you can build from source easily. You need Golang 1.13+, clone the repository and run `go install`

## Usage
Run the binary without arguments and you will get a list of all topics you can browse. Pass a topic and you will get the docs in your terminal, as easy as that.

e.g
```sh
$ ./lxym python
```

If you have `less` installed in your PATH, the tool will automatically try to detect it and pipe the output to less for easier browsing.

If less is not found it will just print the whole docs in the terminal normally. You can force this behaviour even if less is available using the `-p` flag (e.g `lxym -p java`)

## License
Released under the [MIT License](LICENSE).

I do not own the tutorials in `pages/` they are turned into a Go source from the original learnxinyminutes-docs repository and belongs to their respective contributors.

I had not found a proper way to credit them so if you think something is wrong, please do let me know.
