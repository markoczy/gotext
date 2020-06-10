# Golang Text Tools

The programmer's swiss pocket knife of text editing.

## Purpose

Gotext (Suggested Command `tt`) is a simple yet efficient text editing tool for the clipboard.

As a programmer you might know this situation where you just have to do a simple text transformation like replacing tokens or trimming
some prefix or suffix from every line. Usually you would open the text editor of your choice and do the transformation with that editor.
Imagine you could do all this and more without even opening a text editor, that's exactly where Gotext comes into play. Gotext is even
more powerful than any text editor (except vim maybe) giving you access to functions for sorting, filtering, removing duplicate lines
and more. All functions are directly applied to the System's clipboard which means that you can instantly continue your work with the
transformed text (instead of pasting, transorming and copying again).

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

### Example

Imagine you have a heavy xml file in front of you containing 100 items in the following format:

```xml
<someobject>
    <id>48593</id>
    <someusefulinfo1>
        <moreusefulinfo>Sometext</moreusefulinfo>
    </someusefulinfo1>
    <someusefulinfo2></someusefulinfo2>
</someobject>

<someobject>
    <id>99424</id>
    <!-- same stucture... -->
</someobject>

<!-- ... -->
```

Now you want to extract the id of every object in this file.
- First you would copy the text of this file in your clipboard to work with Gotext
- Then you would extract all the lines containing the word `<id>` with the following command: `tt f "<id>"`
- The clipboard will be as follows:

```xml
    <id>48593</id>
    <id>99424</id>
    <!-- ... -->
```

- Next you would trim out the `<id>` and `</id>` tokens:
  - Trim from start to the first `>` with the following command: `tt ts ">"`
  - Trim from end to the first `<` with the following command: `tt te "<"`
- The clipboard will be as follows:

```xml
48593
99424
...
```

The job is done. You want to go further? How about making an SQL Select statement for each of these Id's:

- Prefix the Select Clause: `tt p "SELECT * FROM table WHERE id=\""` (note the escaped `"`)
- Suffix the trailing `"` like this: `tt s "\""` (again we escape `"` with `\`)
- In the end, this would be your result stored in the clipboard:

```sql
SELECT * FROM table WHERE id="48593"
SELECT * FROM table WHERE id="99424"
...
```

## Feature Updates

### Encryption

#### Quicksave Encryption (V1.1.0)

Since Version 1.1.0 of Gotext, everything that is stored in the user folder (command `quicksave`) is encrypted by default. The key that is used is composed from machine and user specific data. This feature is integrated seamlessly in the `quicksave` and `quickload` command and is therefore invisible.

The commands `save` and `load` are not affected by this, an optional password encryption will be implemented in future releases.

#### Clipboard password Encryption (V1.1.0)

The new Commands `encrypt` and `decrypt` can be used to encrypt/decrypt the data in the clipboard with a given password. If the optional second parameter is provided, it will be used as password, otherwise the terminal asks to enter password.

#### Optional password Encryption (V1.1.0)

The commands `quicksave` and `quickload` now support the command switch `-p`. When provided, the terminal asks for a password to encrypt the data. An optional third parameter can be provided as a password, otherwise the terminal asks for a password.

> The command switch must be provided as second parameter and the password as third parameter.

## Deployment

### Windows

Just place the `tt.exe` executable anywhere in your `%PATH%` then you can just press `Win+R` (Run Shortcut) and type some Gotext 
commands like `tt u` (Uppercase), `tt sort` (Sort Lines) or else.

## License

Public Domain