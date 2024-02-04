# Go password guesser

This CLI tool allows to match password with brute force algorithm according to provided dictionary.  
Algorithm is based on linear search O(n).  
Application is leveraging goroutines to boost performance.  
For 'all' option, each method is split into goroutines, and these are then split into smaller goroutines for each dictionary keyword.
Cracking supports many encryption types:  
<big>raw, MD4, MD5, SHA-1, SHA-226, SHA-256, SHA-384, SHA-512, AES</big>.

# Signarure

To run the script you have to follow this command:

```
[program] [hash] [method] [filePath]
```

- `program` is a way you execute program i.e. `go run .` or `main.exe`
- `hash` is hashed or raw password i.e. `myPassword` or `7c6a180b36896a0a8c02787eeafb0e4c`
- `method` is one of these: `all`, `raw`, `MD4`, `MD5`, `SHA-1`, `SHA-226`, `SHA-256`, `SHA-384`, `SHA-512`, `AES` (for AES you need to specify key in code)
- `filePath` is path to dictionary, i.e. `passwords.txt`. (one keyword per line)

# Example

We know that password is encrypted with MD5 method and somehow we have passwords dictionary.

```console
go run . 7c6a180b36896a0a8c02787eeafb0e4c MD5 passwords.txt
```

If password is stored in passwords.txt we will obtain this result:

```
Cracking process initialized ...
Password cracked succesfully:
password1
```

Another case:  
We dont know encryption method and we only have encrypted keyword and passwords dictionary.

```console
go run . 450ad03db9395dfccb5e03066fd7f16cfba2b61e23d516373714471459052ec90a9a4bf3a151e600ea8aaed36e3b8c21a3d38ab1705839749d130da4380f1448 all passwords.txt
```

If password is stored in passwords.txt we will obtain this result:

```
Cracking process initialized ...
Trying methods: raw, MD4, MD5, SHA-1, SHA-224, SHA-256, SHA-384, SHA-512
Password cracked succesfully with SHA-512 method:
myPassword
```

### External importing:

```
go get -u golang.org/x/crypto/md4
```

# License

Go password guesser is released under the MIT license.

# Author

Sebastian Brzustowicz &lt;Se.Brzustowicz@gmail.com&gt;
