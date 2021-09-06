# prettypar

[![Go Reference](https://pkg.go.dev/badge/github.com/9seconds/prettypar.svg)](https://pkg.go.dev/github.com/9seconds/prettypar)

This is a simple but highly opinionated alternative to
[par](https://bitbucket.org/amc-nicemice/par/src/master/) utility that fixes a
couple of my irritants and removes tons of complexity.

Par is a great and robust tool that works great with Vim as its
[`formatprg`](https://vimhelp.org/options.txt.html#%27formatprg%27). But it was
created back in the days when people loved to make ASCII boxes and
pseudo-graphics in files and readjust these decorations on where a new character
is added that lead to funny code review confusion. I hate these boxes. They
bring nothing but constant frustration.

These boxes and an attempt to support decoration made par 3000x times more
complex than it should. When I've tried to fix a bug (if each line is a sentence
that ends with a dot, it _aligns a text as a dot as a border_), I was kinda
shocked by how complex it is. I wanted to have something simple and opinionated.
I know myself and I know how I format text so I didn't want to fight with a
tool, I wanted my copilot.

To make this tool simple and robust, I decided to split a task of text
formatting into 2 steps:

1. Preparation
2. Formatting

Preparation is important because I do not want to reformat something from
scratch, I want to give it some shape and _complete_ formatting with a tool,
instead of throwing it a set of letters with an order "do something". I tend to
format text like this:

```
Preparation is important because I do not want to reformat something from scratch,
I want to give it some shape and _complete_ formatting with a tool, instead of
throwing it a set of letters with an order "do something". I tend to format text
like this:

    1. lalala
       and this continues a line. Pay attention to indentation.
    2. And this is a second bullet

    # btw, this tool still needs to treat comments. But not decorations
    #
    # * and remember
    # * that lists
    #   can be nested
```

prettypar transforms this text literally to:

```
Preparation is important because I do not want to reformat something from
scratch, I want to give it some shape and _complete_ formatting with a tool,
instead of throwing it a set of letters with an order "do something". I tend to
format text like this:

    1. lalala and this continues a line. Pay attention to indentation.
    2. And this is a second bullet

    # btw, this tool still needs to treat comments. But not decorations
    #
    # * and remember
    # * that lists can be nested
```

Also, since its main usage is running with no CLI arguments, prettypar has
almost no options. It has 2, but you can set up them with environment
variables.
