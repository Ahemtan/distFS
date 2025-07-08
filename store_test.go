package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestTransformFunc(t *testing.T) {
	key := "imageasdasdasd"
	pathkey := CASPathTansformFunc(key)

	expectedOrgi := "29c470179933de61c396ac4f7b3efe89aa29bf8c"
	expectedPathname := "29c47/01799/33de6/1c396/ac4f7/b3efe/89aa2/9bf8c"

	if pathkey.Pathname != expectedPathname {
		t.Errorf("Expected path to be %s, got %s", pathkey.Pathname, expectedPathname)
	}

	if pathkey.Filename != expectedOrgi {
		t.Errorf("Expected path to be %s, got %s", pathkey.Filename, expectedOrgi)
	}

	fmt.Printf("pathNames: %s", pathkey.Pathname)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTansformFunc,
	}

	s := NewStore(opts)
	key := "imageasdasdasd"
	data := []byte("some data")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)

	if string(b) != string(data) {
		t.Errorf("Expected data to be %s, got %s", data, b)
	}

}
