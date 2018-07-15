# Golang Text Tools

## Purpose

Text Operations to modify clipboard text

## Usage

```
Single Commands:

u                : UPPERCASE
l                : lowercase
c                : Clear formatting
i                : Invert line order
sort             : Sort (alt: 'o')
rdup             : remove all duplicates (alt: 'rd')
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
r [in] [out]     : replace all [in] with [out]
rx [in] [out]    : replace regex mode
rt [in] [out]    : replace transform backslashes
rxt [in] [out]   : replace regex transform backslashes
```
