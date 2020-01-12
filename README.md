# godict

[![Build Status](https://travis-ci.org/my0k/godict.svg?branch=master)](https://travis-ci.org/my0k/godict)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](./LICENSE)

`godict` is a simple dictionary in terminals written in Go.
The main target is translating Japanese words to the most matched English word.

[EJDict](https://github.com/kujirahand/EJDict) is used in `godict`.

## Usage

`godict word1 word2 ..`

Space delimited words are used for AND search.

## Demo

![godict demo](./screenshots/godict.gif)

## Install

If you have a Go environment, you can use `go get`.

`go get -u github.com/my0k/godict`

Or you can download binaries for each platform from [release](https://github.com/my0k/godict/releases).

## Screenshots

![dark theme](./screenshots/light-theme.png)

![light theme](./screenshots/dark-theme.png)

## License

This software is released under the MIT License, see [LICENSE](./LICENSE).
