# username checker

script that checks for usernames based on response codes.

### requirements

- requests
- colorama

### usage

`$ python checker.py [URL] [INPUT] [OUTPUT]`

example:

`$ python checker.py http://github.com/{} dict github_example`

### notes

- URL should be formatted like `github.com/{}` or `{}.tumblr.com`.
- websites will often block certain usernames like `github.com/admin`
- dict comes from `https://www.archlinux.org/packages/community/any/words/files/`