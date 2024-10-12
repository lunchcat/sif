/*
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                              ║
║                                  SIF                                         ║
║                                                                              ║
║        Blazing-fast pentesting suite written in Go                           ║
║                                                                              ║
║        Copyright (c) 2023-2024 vmfunc, xyzeva, lunchcat contributors         ║
║                    and other sif contributors.                               ║
║                                                                              ║
║                                                                              ║
║        Use of this tool is restricted to research and educational            ║
║        purposes only. Usage in a production environment outside              ║
║        of these categories is strictly prohibited.                           ║
║                                                                              ║
║        Any person or entity wishing to use this tool outside of              ║
║        research or educational purposes must purchase a license              ║
║        from https://lunchcat.dev                                             ║
║                                                                              ║
║        For more information, visit: https://github.com/lunchcat/sif          ║ 
║                                                                              ║
╚══════════════════════════════════════════════════════════════════════════════╝
*/

package templates

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

const (
	archive = "https://github.com/projectdiscovery/nuclei-templates/archive/refs/tags/v%s.tar.gz"
	ref     = "9.6.2"
)

func Install(logger *log.Logger) error {
	// Check if already exists
	if _, err := os.Stat("nuclei-templates"); err == nil {
		return nil
	}

	logger.Infof("nuclei-templates directory not found. Installing...")

	resp, err := http.Get(fmt.Sprintf(archive, ref))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tarball, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer tarball.Close()

	data := tar.NewReader(tarball)

	for {
		header, err := data.Next()
		if errors.Is(io.EOF, err) {
			break
		}
		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(header.Name, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			file, err := os.Create(header.Name)
			if err != nil {
				return err
			}
			if _, err := io.Copy(file, data); err != nil {
				return err
			}
			file.Close()
		}
	}

	if err = os.Rename(fmt.Sprintf("nuclei-templates-%s", ref), "nuclei-templates"); err != nil {
		return err
	}

	return nil
}
