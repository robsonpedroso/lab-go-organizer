package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"

	. "github.com/logrusorgru/aurora"
)

type MyFile struct {
	Name       string
	SourcePath string
	DestPath   string
	ModDate    time.Time
}

func main() {
	var sourceFolder string
	var destFolder string
	var withMonth bool

	flag.StringVar(&sourceFolder, "s", "", "Diretório de origem, onde será listado os arquivos")
	flag.StringVar(&destFolder, "o", "", "Diretório de saída, onde será gerado os arquivos")
	flag.BoolVar(&withMonth, "m", true, "Incluí ou não o mês como diretório de saída")

	flag.Parse()

	if sourceFolder == "" {
		fmt.Print(Blue("Diretório a ser listado: "))
		fmt.Scanf("%s", &sourceFolder)
	}

	if destFolder == "" {
		fmt.Print(Green("Diretório onde será gerado os arquivos: "))
		fmt.Scanf("%s", &destFolder)
	}

	var myFiles []MyFile

	err := filepath.Walk(sourceFolder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			myFiles = append(myFiles, MyFile{
				Name:       info.Name(),
				SourcePath: path,
				DestPath:   "",
				ModDate:    info.ModTime(),
			})
		}
		return nil
	})

	if err != nil {
		fmt.Println(Red(err))
		panic(err)
	}

	sort.Slice(myFiles, func(i, j int) bool {
		return myFiles[i].ModDate.Before(myFiles[j].ModDate)
	})

	for _, file := range myFiles {
		year, month, _ := file.ModDate.Date()
		folder := destFolder + "\\" + strconv.Itoa(year)

		if withMonth {
			folder = folder + "\\" + strconv.Itoa(int(month)) + " - " + month.String()
		}

		if _, err := os.Stat(folder); os.IsNotExist(err) {
			os.MkdirAll(folder, os.ModePerm)
		}

		file.DestPath = folder + "\\" + file.Name
		fmt.Print(Blue(file.Name))

		errorCopy := CopyFile(file.SourcePath, file.DestPath)

		if errorCopy != nil {
			fmt.Println(Red(" - Erro"), errorCopy)
		} else {
			fmt.Println(Green(" - Done"))
		}
	}
}

func CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
