package assets

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/gobuffalo/packr/v2"
)

type Manifest struct {
	mapLock      *sync.RWMutex
	qualifiedMap map[string]string
	box          *packr.Box
	dev          bool
}

func (m *Manifest) FullyQualified(name string) (string, error) {
	// if we're in dev, we need to call m.parseManifest, so we need to
	// write-lock. Otherwise we need to read-lock
	if m.dev {
		m.mapLock.Lock()
		defer m.mapLock.Unlock()
		if err := m.parseManifest(); err != nil {
			return "", fmt.Errorf("re-parsing manifest (to get %s), failed", name)
		}
	} else {
		m.mapLock.RLock()
		defer m.mapLock.RUnlock()
	}

	filename, ok := m.qualifiedMap[name]
	if !ok {
		return "", fmt.Errorf("no file %s found in manifest", name)
	}
	return filepath.Join("/assets", filename), nil
}

func (m *Manifest) String() string {
	m.mapLock.RLock()
	defer m.mapLock.RUnlock()
	return fmt.Sprintf("%+v", m.qualifiedMap)
}

func ParseManifest(dev bool, box *packr.Box) (*Manifest, error) {
	var mut sync.RWMutex
	ret := &Manifest{dev: dev, mapLock: &mut, box: box}
	if err := ret.parseManifest(); err != nil {
		return nil, err
	}
	return ret, nil
}

// make sure to lock (not RLock) before you do this
func (m *Manifest) parseManifest() error {
	manifestBytes, err := m.box.Find("manifest.json")

	if err != nil {
		manifestBytes, err = m.box.Find("assets/manifest.json")
		if err != nil {
			return fmt.Errorf("Your manifest.json is missing: %s", err)
		}
	}
	fmt.Println("found manifest.json")
	rawMap := map[string]string{}
	if err := json.Unmarshal(manifestBytes, &rawMap); err != nil {
		return fmt.Errorf("Your manifest.json is malformed: %s", err)
	}
	m.qualifiedMap = rawMap
	return nil
}
