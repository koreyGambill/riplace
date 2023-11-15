# riplace

`rp` (short for "riplace") is a command-line utility designed for multiline find and replace across files. Inspired by the blazing-fast search capabilities of "ripgrep", rp brings similar speed and efficiency to the task of replacing text.

# examples
Simple Replacement:

``` bash
rp -p "*.txt" -f "foo" -r "bar"
```

Case-Insensitive Replacement:

``` bash
rp -p "*.config" -i -f "example" -r "sample"
```

Including Hidden Files:

``` bash
rp -p "*" --hidden -f "localhost" -r "127.0.0.1"
```

Multi-line Replacement:

``` bash
rp -p "*.sql" -f "SELECT * FROM users" -r "SELECT * FROM active_users"
```
