package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//Conf configuration structure
type Conf struct {
	APIKey    string `yaml:"api_key"`
	SecretKey string `yaml:"api_secret"`
	// APIKeyTest    string `yaml:"apiKey"`
	// SecretKeyTest string `yaml:"secretKey"`
}

//GetConf getter for getting configuration
func (c *Conf) GetConf() {
	yamlFile, err := ioutil.ReadFile("./signature/conf.yml")
	if err != nil {
		log.Println("Parsing conf.yml get err ", err)
	}
	yaml.Unmarshal(yamlFile, c)
}

//GetSignature : for encrypt signature
func (c *Conf) GetSignature(context string) string {

	key := []byte(c.SecretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(context))
	return hex.EncodeToString(h.Sum(nil))
}

//GetAPIKey for getting API Key
func (c *Conf) GetAPIKey() string {
	return c.APIKey
}
