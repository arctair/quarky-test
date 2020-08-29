package main

import (
	"encoding/json"
	"net/http"
)

// Build ...
type Build interface {
	getSha1() string
	getVersion() string
}

type controller struct {
	build Build
}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(map[string]string{
		"sha1":    c.build.getSha1(),
		"version": c.build.getVersion(),
	})
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}
