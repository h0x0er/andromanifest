package andromanifest

import (
	"encoding/xml"
	"io/ioutil"
)

func NewFromFile(manifest string) (*Manifest, error) {

	bytes, _ := ioutil.ReadFile(manifest)
	mf := new(Manifest)
	err := xml.Unmarshal(bytes, mf)

	return mf, err

}
