Characters in a string are called runes in go, take 32bits thus utf-32 compatible.

Depending on conversion upon printing, results may vary

Go source code is all encoded in utf8. Strings can contain arbitrary bytes, but when constructed from string literals, those bytes are (if not byte-level escaped, usually) UTF-8.
