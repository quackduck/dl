package main

import (
	"errors"
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
	//"github.com/gabriel-vasile/mimetype"
	pb "github.com/schollz/progressbar/v3"
)

var (
	version = "dev"
	helpMsg = `Dl - Print, download or copy website content`
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
	if hasOption, _ := argsHaveOption("print", "p"); hasOption {
		err := getAndWriteNormalizeUrl(os.Args[2], os.Stdout, true)
		if err != nil {
			handleErr(err)
			return
		}
		return
	}
	if hasOption, _ := argsHaveOption("copy", "c"); hasOption {
		err, w := getClipWriter()
		if err != nil {
			handleErr(err)
			return
		}
		err = getAndWriteNormalizeUrl(os.Args[2], w, true)
		if err != nil {
			handleErr(err)
			return
		}
		w.Close()
		return
	}

	writeTo := filepath.Base(os.Args[1])
	if exists(writeTo) {
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

	err = getAndWriteNormalizeUrl(os.Args[1], outFile, true)
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

func getAndWriteNormalizeUrl(website string, w io.Writer, bar bool) error {
	err := getAndWrite(website, w, bar)
	if err != nil && !strings.HasPrefix(website, "https://") && !strings.HasPrefix(website, "http://") { // protocol not already mentioned and error occurred
		err = getAndWrite("https://"+website, w, bar)
		if err != nil {
			stderrln("Resorting to http...")
			err = getAndWrite("http://"+website, w, bar)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getAndWrite(website string, w io.Writer, bar bool) error {
	response, err := http.Get(website)
	//mime := mimetype.Detect([]byte)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if bar {
		pbar := pb.NewOptions64(response.ContentLength,
			pb.OptionEnableColorCodes(true),
			pb.OptionShowBytes(true),
			pb.OptionSetWriter(os.Stderr),
			pb.OptionThrottle(65*time.Millisecond),
			pb.OptionShowCount(),
			pb.OptionOnCompletion(func() {
				stderrln()
			}),
			pb.OptionSetDescription("Downloading"),
			pb.OptionFullWidth(),
			pb.OptionSetTheme(pb.Theme{
				Saucer:        "=",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}))
		io.Copy(io.MultiWriter(w, pbar), response.Body)
		return nil
	} else {
		io.Copy(w, response.Body)
		return nil
	}
}

func getClipWriter() (error, *clipboard) {
	var copyCmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		copyCmd = exec.Command("pbcopy")
	case "windows":
		copyCmd = exec.Command("powershell.exe", "-command", "Set-Clipboard") //-Value "+"\""+s+"\"")
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
			handleErr(errors.New("sorry, uniclip won't work if you don't have xsel, xclip, wayland or Termux:API installed :(\nyou can create an issue at https://github.com/quackduck/uniclip/issues"))
			os.Exit(2)
		}
	}
	in, err := copyCmd.StdinPipe()
	if err != nil {
		return err, nil
	}
	if err = copyCmd.Start(); err != nil {
		return err, nil
	}
	return nil, &clipboard{
		in,
		copyCmd,
	}
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
