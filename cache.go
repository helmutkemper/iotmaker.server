package iotmaker_server_json

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/util"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const kListChar string = "abcdefghijklmnopqrstuvwxyz0123456789"
const KCacheDir string = "./cache"

var machineName string

// pt-br: retorna um novo struct JSonOut para restful
// en: return a new JSonOut struct for restful
func NewJSonOut() Out {
	var ret = Out{}
	ret.Meta.Success = true

	return ret
}

type Out struct {
	jSonOut
}

func (el *Out) Byte() []byte {

	if el.Meta.Success != true {
		el.Objects = []int{}
	} else if el.Meta.Cache == "" {

		err := el.SaveCache()
		if err != nil {
			panic(err)
		}

	}

	switch converted := el.Objects.(type) {
	case []types.Container:
		el.Meta.TotalCount = len(converted)
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

	if el.Meta.Cache == "" {
		el.Meta.Cache = el.MakeId()
	}

	err = el.save(el.Meta.Cache)

	return err
}

func (el *Out) LoadCache(id string) error {
	err := el.load(id)
	return err
}

func (el *Out) GetTime() string {
	var tm = time.Now().UnixNano()
	return strconv.FormatInt(tm, 16)
}

func (el *Out) GetMachineName() string {
	if machineName != "" {
		return machineName
	}

	machineName = os.Getenv("MACHINE_NAME")
	if machineName == "" {
		machineName = el.GetRandString(10)
		log.Print("iotMaker.server.json.GetMachineName.error: please, set environment var MACHINE_NAME\n")
	}

	err := os.Setenv("MACHINE_NAME", machineName)
	if err != nil {
		log.Printf("iotMaker.server.json.GetMachineName.error: %v\n", err.Error())
	}

	return machineName
}

func (el *Out) GetRandString(l int) string {
	var randString = ""

	for digit := 0; digit != l; digit += 1 {
		var index = rand.Intn(len(kListChar))
		randString += kListChar[index : index+1]
	}

	return randString
}

func (el *Out) MakeId() string {
	return el.GetTime() + "-" + el.GetMachineName() + "-" + el.GetRandString(10)
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
