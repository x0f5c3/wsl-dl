package internal

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/pterm/pterm"
	"github.com/x0f5c3/manic-go/pkg/downloader"
)

func GetDistroLink(distro string) (string, error) {
	if link, ok := LinkMap[distro]; ok {
		return link, nil
	}
	return "", ErrInvalidDistro
}

func GetDistroFilename(distro string) (string, error) {
	if filename, ok := FilenameMap[distro]; ok {
		return filename, nil
	}
	return "", ErrInvalidDistro
}

func AskForDistro() (string, error) {
	sel := pterm.DefaultInteractiveSelect.WithOptions(LinkKeys)
	sel.DefaultOption = "Ubuntu"
	selection, err := sel.Show("Select a distro to download:")
	if err != nil {
		return "", err
	}
	conf, err := pterm.DefaultInteractiveConfirm.Show("Are you sure you want to download " + selection + "?")
	if err != nil {
		return "", err
	}
	if !conf {
		return "", fmt.Errorf("user cancelled")
	}
	return selection, nil
}

func DownloadDistro(distroName, outputDir string, workers int) error {
	cl := downloader.NewClient()
	link, err := GetDistroLink(distroName)
	if err != nil {
		return err
	}
	filename, err := GetDistroFilename(distroName)
	if err != nil {
		return err
	}
	length, err := downloader.GetLength(link, cl)
	if err != nil {
		return err
	}
	file := downloader.File{
		Url:      link,
		FileName: filename,
		Length:   length,
		Client:   cl,
	}
	outputPath := filepath.Join(outputDir, file.FileName)
	absPath, err := filepath.Abs(outputPath)
	conf := false
	if err != nil {
		conf, err = pterm.DefaultInteractiveConfirm.Show(fmt.Sprintf("Are you sure you want to download %s to %s?", distroName, outputPath))
		if err != nil {
			return err
		}
	} else {
		conf, err = pterm.DefaultInteractiveConfirm.Show(fmt.Sprintf("Are you sure you want to download %s to %s?", distroName, absPath))
		if err != nil {
			return err
		}
	}
	if !conf {
		return fmt.Errorf("user cancelled")
	}
	f, err := file.DownloadWithProgress(workers, runtime.NumCPU())
	if err != nil {
		return err
	}
	pterm.Success.Printfln("Downloaded %s", distroName)
	pterm.Info.Printfln("Saving %s to %s", distroName, outputPath)
	err = f.Save(outputDir)
	if err != nil {
		return err
	}
	pterm.Success.Printfln("Saved %s to %s", distroName, outputPath)
	return nil
}
