# easy-go-optimizations

# strings.Index -> strings.IndexByte
- Search: strings\.Index\(("?[\w]+"?), "([\x20-\x7E]{1})"\)
- Replace: strings\.IndexByte\($1, '$2'\)

# strings.LastIndex -> strings.LastIndexByte
- Search: strings\.LastIndex\(("?[\w]+"?), "([\x20-\x7E]{1})"\)
- Replace: strings\.LastIndexByte\($1, '$2'\)

# bytes.Index -> bytes.IndexByte
- Search: bytes\.Index\(([\w]+), \[\]byte\("([\x20-\x7E]{1})"\)\)
- Replace: bytes\.IndexByte\($1, '$2'\)

# .Write() -- TODO: for bytes.Buffer, replace Write with WriteByte where possible
