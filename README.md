# Dl
Print, download or copy website content

## Usage
```
Dl - Print, download or copy website content
Usage: 
   dl [-n/--no-overwrite] [-p/--print | -c/--copy] <url>
   dl [-h/--help | --version/-v]
Options:
   -o, --no-overwrite   don't overwrite existing files
   -p, --print          print downloaded data on stdout
   -c, --copy           copy data to clipboard (needs one of xclip, xsel,
                        wl-copy or termux-clipboard-set installed on Linux)
Examples:
   dl google.com
   dl http://foo.com
   dl golang.org/dl/go1.16.src.tar.gz
   dl -c raw.githubusercontent.com/quackduck/dl/main/dl.go
   dl -p https://raw.githubusercontent.com/octocat/Hello-World/master/README
Note:
   Without any options, Dl downloads to a file in the working directory
   When downloading data, Dl shows a progress bar on stderr
   If the url doesn't have an http/https prefix Dl will try https and then http
   Dl will follow up to 10 redirects
```
