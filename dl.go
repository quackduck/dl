package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/fatih/color"
	pb "github.com/schollz/progressbar/v3"
)

var (
	version = "dev"
	helpMsg = `Dl - Print, download or copy website content
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
   Dl will follow up to 10 redirects`
	allowOverwrite = true
)

//TODO: Maybe detect file extension
// Let setLocalClip give out a writer

func main() {
	if len(os.Args) == 1 {
		handleErrStr("invalid number of arguments")
		fmt.Println(helpMsg)
		return
	}
	if hasOption, _ := argsHaveOption("help", "h"); hasOption {
		fmt.Println(helpMsg)
		return
	}
	if hasOption, _ := argsHaveOption("version", "v"); hasOption {
		fmt.Println("Dl " + version)
		return
	}
	if hasOption, i := argsHaveOption("no-overwrite", "n"); hasOption {
		allowOverwrite = false
		os.Args = removeKeepOrder(os.Args, i)
		main()
		return
	}
	if hasOption, _ := argsHaveOption("print", "p"); hasOption {
		err := getAndWriteNormalizeURL(os.Args[2], os.Stdout)
		if err != nil {
			handleErr(err)
			return
		}
		return
	}
	if hasOption, _ := argsHaveOption("copy", "c"); hasOption {
		w, err := getClipWriter()
		if err != nil {
			handleErr(err)
			return
		}
		err = getAndWriteNormalizeURL(os.Args[2], w)
		if err != nil {
			handleErr(err)
			return
		}
		w.Close()
		return
	}

	writeTo := filepath.Base(os.Args[1])

	if !allowOverwrite && exists(writeTo) {
		handleErrStr("Not overwriting " + writeTo + ". Exiting.")
		return
	}

	outFile, err := os.Create(writeTo)
	if err != nil {
		handleErr(err)
		return
	}
	defer outFile.Close()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(writeTo)
		os.Exit(0)
	}()

	err = getAndWriteNormalizeURL(os.Args[1], outFile)
	if err != nil {
		os.Remove(writeTo)
		handleErr(err)
		return
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !(os.IsNotExist(err))
}

func argsHaveOption(long string, short string) (hasOption bool, foundAt int) {
	for i, arg := range os.Args {
		if arg == "--"+long || arg == "-"+short {
			return true, i
		}
	}
	return false, 0
}

func getAndWriteNormalizeURL(website string, w io.Writer) error {
	err := getAndWrite(website, w)
	if err != nil && !strings.HasPrefix(website, "https://") && !strings.HasPrefix(website, "http://") { // protocol not already mentioned and error occurred
		err = getAndWrite("https://"+website, w)
		if err != nil {
			stderrln("Resorting to http...")
			err = getAndWrite("http://"+website, w)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getAndWrite(website string, w io.Writer) error {
	start := time.Now()
	response, err := http.Get(website)
	//mime := mimetype.Detect([]byte)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	pbar := pb.NewOptions64(response.ContentLength,
		pb.OptionEnableColorCodes(true),
		pb.OptionShowBytes(true),
		pb.OptionSetWriter(os.Stderr),
		pb.OptionThrottle(65*time.Millisecond),
		pb.OptionShowCount(),
		pb.OptionClearOnFinish(),
		pb.OptionSetDescription("Downloading"),
		pb.OptionFullWidth(),
		pb.OptionSetTheme(pb.Theme{
			Saucer:        "=",
			SaucerPadding: " ",
			BarStart:      "|",
			BarEnd:        "|",
		}))

	io.Copy(io.MultiWriter(w, pbar), response.Body)
	stderrln("Done in", time.Since(start).Round(time.Millisecond))
	return nil
}

func getClipWriter() (*clipboard, error) {
	var copyCmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		copyCmd = exec.Command("pbcopy")
	case "windows":
		copyCmd = exec.Command("powershell.exe", "-command", "Set-Clipboard")
	default:
		if _, err := exec.LookPath("xclip"); err == nil {
			copyCmd = exec.Command("xclip", "-in", "-selection", "clipboard")
		} else if _, err = exec.LookPath("xsel"); err == nil {
			copyCmd = exec.Command("xsel", "--input", "--clipboard")
		} else if _, err = exec.LookPath("wl-copy"); err == nil {
			copyCmd = exec.Command("wl-copy")
		} else if _, err = exec.LookPath("termux-clipboard-set"); err == nil {
			copyCmd = exec.Command("termux-clipboard-set")
		} else {
			handleErrStr("Sorry, --copy won't work if you don't have xsel, xclip, wayland or Termux:API installed")
			os.Exit(2)
		}
	}
	in, err := copyCmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	if err = copyCmd.Start(); err != nil {
		return nil, err
	}
	return &clipboard{
		in,
		copyCmd,
	}, nil
}

type clipboard struct {
	in  io.WriteCloser
	cmd *exec.Cmd
}

func (c *clipboard) Close() error {
	if err := c.in.Close(); err != nil {
		return err
	}
	if err := c.cmd.Wait(); err != nil {
		return err
	}
	return nil
}
func (c *clipboard) Write(p []byte) (n int, err error) {
	return c.in.Write(p)
}

func handleErr(err error) {
	handleErrStr(err.Error())
}

func handleErrStr(str string) {
	_, _ = fmt.Fprintln(os.Stderr, color.RedString("error: ")+str)
}

func stderrln(a ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, a...)
}

func removeKeepOrder(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
