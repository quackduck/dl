# Dl

**Print, download or copy website content**

Dl makes your life easy by:
1. Following redirects
2. Showing beautiful progress bars and reports
3. Letting you copy directly to clipboard with the -c option
4. Letting you not specify http/https
5. Being a bit faster than curl


## How it works
You just type in dl, paste a url, hit enter and you now have a downloaded file in your working directory.   
You also have progress bars showing time left, time taken, current speed and total file size.:
```
$ dl speed.hetzner.de/100MB.bin # test file (notice the missing https)
Downloading  49% [=============               ] (49/100 MB, 5.010 MB/s) [12s:10s]
```
And at the end:
```
Done in 22.475s
```

You also have options to print to stdout and _directly copy to your clipboard_, which is useful with raw.githubusercontent.com links (see example four)
## Usage
```
dl [-n/--no-overwrite] [-p/--print | -c/--copy] <url>
dl [-h/--help | --version/-v]
```

### Options
```text
-o, --no-overwrite   don't overwrite existing files
-p, --print          print downloaded data on stdout
-c, --copy           copy data to clipboard (needs one of xclip, xsel,
                     wl-copy or termux-clipboard-set installed on Linux)
```

### Examples

```shell
dl google.com # download google's front page
dl http://foo.com # you can still specify protocol
dl golang.org/dl/go1.16.src.tar.gz # download Go's source code
dl -c raw.githubusercontent.com/quackduck/dl/main/dl.go # copy dl's source
dl -p https://raw.githubusercontent.com/octocat/Hello-World/master/README # try this :)
```

### Caveats
Without any options, Dl downloads to a file in the working directory.  
When downloading data, Dl shows a progress bar on stderr.  
If the url doesn't have an http/https prefix Dl will try https and then http.  
Dl will follow up to 10 redirects.

## Install

```shell
brew install quackduck/tap/dl
```
or get an executable from [releases](https://github.com/quackduck/dl/releases)
