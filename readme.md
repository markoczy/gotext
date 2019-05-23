# Golang Text Tools

The programmer's swiss pocket knife of text editing.

## Purpose

Gotext (Suggested Command `tt`) is a simple yet efficient text editing tool for the clipboard.

As a programmer you might know this situation where you just have to do a simple text transformation like replacing tokens or trimming
some prefix or suffix from every line. Usually you would open the text editor of your choice and do the transformation with that editor.
Imagine you could do all this and more without even opening a text editor, that's exactly where Gotext comes into play. Gotext is even
more powerful than any text editor (except vim maybe) giving you access to functions for sorting, filtering, removing duplicate lines
and more.

## Usage

Currently available functions (type `tt` to see them):

```
u                : UPPERCASE
l                : lowercase
c                : Clear formatting
i                : Invert line order
paste            : Paste text to console (use with ">")
sort             : Sort (alt: 'o')
rdup             : remove all duplicates (alt: 'rd')
purge [-y]       : delete all quicksaves
filter [txt]     : Select lines with [txt] (alt: 'f')
filterx [txt]    : Exclude lines with [txt] (alt: 'fx')
pre [txt]        : prefix [txt] by line (alt: 'p')
post [txt]       : suffix [txt] by line (alt: 's')
ts [txt]         : trim start to end of [txt] by line
tsx [txt]        : trim start to start of [txt] by line
te [txt]         : trim end to start of [txt] by line
tex [txt]        : trim end to end of [txt] by line
save [path]      : Clipboard to file (alt 'sv')
load [path]      : File to clipbooard (alt 'ld')
qs [name]        : Quicksave (save to temp file)
ql [name]        : Quickload (load from temp file)
skip [n]         : Skip first [n] lines
skipe [n]        : Skip last [n] lines
r [in] [out]     : replace all [in] with [out]
rx [in] [out]    : replace regex mode
rt [in] [out]    : replace transform backslashes
rxt [in] [out]   : replace regex transform backslashes
```

### Windows

Just place the `tt.exe` executable anywhere in your `%PATH%` then you can just press `Win+R` (Run Shortcut) and type some Gotext 
commands like `tt u` (Uppercase), `tt sort` (Sort Lines) or else.