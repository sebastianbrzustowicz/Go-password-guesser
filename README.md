# Password cracker

This CLI tool allows to crack password with brute force algorithm according to provided dictionary.  
It is possible to try variety of permutations with allowed signs.  
Application is using goroutines to boost performance.  
Algorithm is based on linear search O(n).  
Cracking supports many encryption types: <big>raw, MD4, MD5, SHA-1, SHA-226, SHA-256, SHA-384, SHA-512, AES</big>.

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

Importing:

```
go get -u golang.org/x/crypto/md4
```
