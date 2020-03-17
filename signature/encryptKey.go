package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//Conf configuration structure
type conf struct {
	APIKey        string `yaml:"api_key"`
	SecretKey     string `yaml:"api_secret"`
	APIKeyTest    string `yaml:"apiKey"`
	SecretKeyTest string `yaml:"secretKey"`
}

//GetConf getter for getting configuration
func (c *conf) getConf() {
	yamlFile, err := ioutil.ReadFile("./signature/conf.yml")
	if err != nil {
		log.Println("Parsing conf.yml get err ", err)
	}
	yaml.Unmarshal(yamlFile, c)
}

//GetSignature : for encrypt signature
func GetSignature(context string) {
	c := conf{}
	c.getConf()
	key := []byte(c.SecretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(context))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
