package openjtalk

import (
	"os"
	"os/exec"
)

func Jsay(text string) {
	path := "voice.txt"
	_, e := os.Stat(path)
	if e == nil {
		os.Remove(path)
	}
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, er := file.WriteString(text)
	if er != nil {
		panic(er)
	}
	cmd := exec.Command("./jsay.sh")
	cmd.Start()
	cmd.Wait()
}
