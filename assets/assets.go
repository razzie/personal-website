package assets

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync/atomic"
	"time"

	"github.com/razzie/gorzsony.com/pkg/content"
)

//go:embed *
var assets embed.FS

func StaticFS() fs.FS {
	return assets
}

func StaticAsset(name string) ([]byte, error) {
	return assets.ReadFile(name)
}

var _ fs.FS = (*AssetLoader)(nil)

type AssetLoader struct {
	cacheDir  string
	remoteDir string
	content   atomic.Value
}

func NewAssetLoader(cacheDir, remoteDir string) *AssetLoader {
	if !strings.HasSuffix(remoteDir, "/") {
		remoteDir += "/"
	}
	loader := &AssetLoader{
		cacheDir:  cacheDir,
		remoteDir: remoteDir,
	}

	contentBytes, err := assets.ReadFile("content.yaml")
	if err != nil {
		panic(err)
	}
	content, err := content.LoadContent(bytes.NewReader(contentBytes))
	if err != nil {
		panic(err)
	}
	loader.content.Store(content)

	go func() {
		ticker := time.NewTicker(time.Hour)
		for {
			<-ticker.C
			content, err := loader.loadContent()
			if err != nil {
				log.Println(err)
				continue
			}
			loader.content.Store(content)
		}
	}()

	return loader
}

func (loader *AssetLoader) Content() *content.Content {
	content, _ := loader.content.Load().(*content.Content)
	return content
}

func (loader *AssetLoader) Open(filename string) (fs.File, error) {
	file, err := assets.Open(filename)
	if errors.Is(err, fs.ErrNotExist) {
		return loader.openCached(filename)
	}
	return file, err
}

func (loader *AssetLoader) openCached(filename string) (fs.File, error) {
	osPath := path.Join(loader.cacheDir, filename)
	file, err := os.Open(osPath)
	if errors.Is(err, os.ErrNotExist) {
		if err := loader.loadRemoteFile(filename); err != nil {
			return nil, err
		}
		return os.Open(osPath)
	}
	return file, err
}

func (loader *AssetLoader) loadRemoteFile(filename string) error {
	return downloadFile(path.Join(loader.cacheDir, filename), loader.remoteDir+filename)
}

func (loader *AssetLoader) loadContent() (*content.Content, error) {
	resp, err := http.Get(loader.remoteDir + "content.yaml")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot download 'content.yaml', bad status: %s", resp.Status)
	}

	return content.LoadContent(resp.Body)
}

func downloadFile(osPath string, url string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("cannot download %q, bad status: %s", url, resp.Status)
	}

	os.MkdirAll(path.Dir(osPath), os.ModePerm)

	out, err := os.Create(osPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
