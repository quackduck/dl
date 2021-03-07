package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
	//"github.com/gabriel-vasile/mimetype"
	pb "github.com/schollz/progressbar/v3"
)

var (
	version = "dev"
	helpMsg = `Dl - Print, download or copy website content`
)

//TODO: Detect file extension

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
		err, f := getAndWriteNormalizeUrl(os.Args[2], true)
		if err != nil {
			handleErr(err)
			return
		}
		f(os.Stdout)
		return
	}
	if hasOption, _ := argsHaveOption("copy", "c"); hasOption {
		buf := new(bytes.Buffer)
		err, f := getAndWriteNormalizeUrl(os.Args[2], true)
		if err != nil {
			handleErr(err)
			return
		}
		f(buf)
		all, _ := ioutil.ReadAll(buf)
		setLocalClip(string(all))
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
	err, f := getAndWriteNormalizeUrl(os.Args[1], true)
	if err != nil {
		_ = os.Remove(writeTo)
		handleErr(err)
	}
	f(outFile)
	defer outFile.Close()
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

func getAndWriteNormalizeUrl(website string, bar bool) (err error, f func(io.Writer)) {
	err, f = getAndWrite(website, bar)
	//var url string
	//var continueBool bool
	if err != nil && !strings.HasPrefix(website, "https://") && !strings.HasPrefix(website, "http://") { // protocol not already mentioned and error occurred
		err, f = getAndWrite("https://"+website, bar)
		if err != nil {
			stderrln("Resorting to http...")
			err, f = getAndWrite("http://"+website, bar)
			if err != nil {
				return err, nil
			}
		}
	}
	return nil, f
	//for i := 1; err != nil &&
	//	continueBool &&
	//	!strings.HasPrefix(website, "https://") &&
	//	!strings.HasPrefix(website, "http://"); i++ {
	//	switch i {
	//	case 1:
	//		if !strings.HasPrefix(website, "https://") {
	//			url = "https://" + website
	//		}
	//	case 2:
	//		if !strings.HasPrefix(website, "http://") {
	//			url = "http://" + website
	//		}
	//	default:
	//		continueBool = false
	//	}
	//	err = getAndWrite(url, writer)
	//}
	//if err != nil {
	//	handleErr(err)
	//}
}

func getAndWrite(website string, bar bool) (error, func(io.Writer)) {
	response, err := http.Get(website)
	//mime := mimetype.Detect([]byte)
	if err != nil {
		return err, nil
	}
	defer response.Body.Close()
	if bar {
		//bar := pb.DefaultBytes(response.ContentLength)
		pbar := pb.NewOptions64(response.ContentLength,
			pb.OptionEnableColorCodes(true),
			pb.OptionShowBytes(true),
			pb.OptionSetWriter(os.Stderr),
			pb.OptionThrottle(65*time.Millisecond),
			pb.OptionShowCount(),
			pb.OptionOnCompletion(func() {
				stderrln()
			}),
			//pb.OptionClearOnFinish(),
			//pb.OptionSetWidth(15),
			pb.OptionSetDescription("Downloading"),
			pb.OptionFullWidth(),
			pb.OptionSetTheme(pb.Theme{
				Saucer: "=",
				//SaucerHead:    "[green][reset]",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}))
		return nil, func(w io.Writer) { io.Copy(io.MultiWriter(w, pbar), response.Body) }
	} else {
		return nil, func(w io.Writer) { io.Copy(w, response.Body) }
	}
}

func setLocalClip(s string) {
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
		handleErr(err)
		return
	}
	if err = copyCmd.Start(); err != nil {
		handleErr(err)
		return
	}
	if runtime.GOOS != "windows" {
		if _, err = in.Write([]byte(s)); err != nil {
			handleErr(err)
			return
		}
		if err = in.Close(); err != nil {
			handleErr(err)
			return
		}
	}
	if err = copyCmd.Wait(); err != nil {
		handleErr(err)
		return
	}
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
