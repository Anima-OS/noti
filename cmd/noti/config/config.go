package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/variadico/yaml"
)

const (
	Filename = ".noti.yaml"
)

func File() (Options, error) {
	ds, err := dirs()
	if err != nil {
		return NewOptions(), err
	}

	var data []byte
	for _, d := range ds {
		data, err = ioutil.ReadFile(filepath.Join(d, Filename))
		if err == nil {
			break
		}
	}
	if err != nil {
		return NewOptions(), fmt.Errorf("config not found in: %s", ds)
	}

	opts := NewOptions()
	if err := yaml.Unmarshal(data, &opts); err != nil {
		return Options{}, err
	}

	return opts, nil
}

func dirs() ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return decompose(dir), nil
}

func decompose(p string) []string {
	var out []string

	if !strings.HasSuffix(p, string(os.PathSeparator)) {
		p += string(os.PathSeparator)
	}

	for i := len(p) - 1; i >= 0; i-- {
		if p[i] == '/' {
			out = append(out, p[:i+1])
		}
	}

	return out
}
