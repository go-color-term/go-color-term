# Go Color Term

Welcome to Go Color Term! the place where your monochromatic terminal meets a colorful world!

This library allows to apply _ANSI Escape Sequences_ to alter the style of the rendered text on a Terminal.

## Supported Operating Systems and Terminals
* macOS: tested and working on iTerm2.
* Linux: not tested yet, but should work.
* Windows: not supported yet, erratic behaviour.

## Examples

Check the [accompanying repository](https://github.com/nelsonghezzi/go-color-term-examples) with a ready-to-run app which recopiles various examples to showcase the library API. 

## How to use

This library provides several ways to add color to your output, tailored for different use cases.

The idea is to provide with the simplest way to render styled text for each situation.

| Method                         | Use Case |
|--------------------------------|----------|
| `coloring.*` utility functions |          |
| `ColourBuilder`                |          |
| `DecoratedText`                |          |
| `SentenceBuilder`              |          |

The next sections describes each of these approaches in more detail.

### `coloring.*` utility functions

This are utility functions that lets you apply simple style attributes to a provided string, like:

```go
coloring.Red("Fire truck")
```

This will produce the following escaped string:

```
\ESC[31;mFire truck\ESC[39;m
```

Which be can decompose for analysis into:

```
 Sets red      Resets text color
text color        to default
    |                  |
    v                  v
---------          ---------
\ESC[31;mFire truck\ESC[39;m
         ----------
             ^
             |
      Provided string
```

Of course you can combine multiple styles by passing the output of one function to another:

```go
coloring.Bold(coloring.Red("Wolf"))
```

This is probably fine for a few function calls, but can become difficult to read if the combination grows larger.

Also, this approach, while it works, doesn't generate the most efficient escape sequences:

```
\ESC[1;m\ESC[31;mWolf\ESC[39;m\ESC[22;m
```

A more compact sequence for the same style would be `\ESC[1;31;mWolf\ESC[0;m`, which is the topic of the next section.

### `ColourBuilder`

TBD

### `DecoratedText`

TBD

### `SentenceBuilder`

This is perhaps the more daunting form of adding style attributes, but in return it gives the more fine grained control of when to start and end each style attribute.

## Licence

TBD

## Contributing

Feel free to add contributions in the form of issues/feature requests, pull requests discussions
