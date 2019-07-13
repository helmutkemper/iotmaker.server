package iotmaker_server_json

import (
	"encoding/json"
	"github.com/helmutkemper/util"
	"io/ioutil"
	"math/rand"
)

const kListChar string = "abcdefghijklmnopqrstuvwxyz0123456789"
const KCacheDir string = "./cache"

// pt-br: retorna um novo struct JSonOut para restful
// en: return a new JSonOut struct for restful
func NewJSonOut() Out {
	return Out{}
}

type Out struct {
	jSonOut
	Id string `json:"cacheId"`
}

func (el *Out) Byte() []byte {

	if el.Id == "" {
		err := el.SaveCache()
		if err != nil {
			panic(err)
		}
	}

	if el.Meta.Success != true {
		el.Objects = []int{}
	}

	out, _ := json.Marshal(el)

	return out
}

func (el *Out) SaveCache() error {
	var err error
	// todo: remove this const
	err = util.DirMake(KCacheDir)
	if err != nil {
		return err
	}

	if el.Id == "" {
		el.Id = el.MakeId()
	}

	err = el.save(el.Id)

	return err
}

func (el *Out) LoadCache(id string) error {
	err := el.load(id)
	return err
}

func (el *Out) MakeId() string {

	var id = ""
	for block := 0; block != 4; block += 1 {

		if block != 0 {
			id += "-"
		}

		for digit := 0; digit != 4; digit += 1 {
			var index = rand.Intn(len(kListChar)-0) + 0
			id += kListChar[index : index+1]
		}
	}

	return id
}

func (el Out) save(id string) error {
	buf, err := json.Marshal(el)

	err = ioutil.WriteFile(KCacheDir+"/"+id, buf, 0644)

	return err
}

func (el *Out) load(id string) error {
	var file []byte
	var err error

	file, err = ioutil.ReadFile(KCacheDir + "/" + id)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, el)

	return err
}
