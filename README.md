# gocrtsh
A Go-based command-line tool for enumerating subdomains using https://crt.sh/

# About

`gocrtsh` is a lightweight tool that enumerates subdomains for a target by querying [Certificate Transparency](https://de.wikipedia.org/wiki/Certificate_Transparency)
records indexed by crt.sh. It requests certificates issued for the target domain and its subdomains, removes duplicates, and outputs a clean list of discovered subdomains to stdout or a file.

```bash
gocrtsh example.com
```
