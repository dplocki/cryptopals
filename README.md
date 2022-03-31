# The Cryptopals crypto challenges

My approach to Cryptopals challenges. Written in ~~Python~~ Go.

## Cryptopals.com

The link: [cryptopals.com](https://cryptopals.com/)

## Warning

I hope I don't have explain this, but just stating the obvious: **do not use any of the code available in this repository in production cryptographic purposes**. The whole point of this repository is for me to learn.

## The language change

Originaly I was implementing this project in Python. I have decided to switch because:

* I was learning Go in the meantime and I feel comfortable enought to switch
* I wanted to learn more about initialization of Go project and Go packages
* The challenges seems to require a lots of byte, bits even, operations, which Go seems to match better
* It was yet another of mine project in Python

## Challenges

### Go

Althoght I have some basic knownalge of Go language, I need to learn fast how excacly structure the project.

Another task was to start the unittest solution, which appear to be quite trivial, due to nice, out of the box integration of go unittest library.

### Scoring function

In order to detect human-readable text, we need to mechanicly recognise English. The easiest way to do this is ofcourse calculating the letters fequency. The only problem is to code it properly.

Original solution, calculating how much the most common letters is in the text seems work for simple task. It was not sufficient for breaking the **set1/challenge**.

I have add letter score value (the more common letter, there more score it provides). But finally the problem was within my key length function.

### AES in ECB mode

The Golang package [crypto](https://pkg.go.dev/crypto) doesn't support the ECB mode. [Wonder why](https://github.com/golang/go/issues/5597). Anyway it is easy to implement... if you know enought about AES algorythm and its implemention.

### Don't implement your own cryptographic

If this challenges taught me anything by now, it will be: you can always make an error in implementing the good algorithm. I am adding the warning to the [README.md](./README.md) file.

### Down of cryptopals.com

In 31 march 2022 page https://cryptopals.com/ seems stop responding.
