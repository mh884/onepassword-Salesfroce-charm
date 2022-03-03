package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"github.com/pkg/browser"
	"encoding/xml"
	"github.com/spf13/viper"
	"log"
)


type MyRespEnvelope struct {
	XMLName xml.Name
	Body    Body
}

type Body struct {
	XMLName     xml.Name
	GetResponse loginResponse `xml:"loginResponse"`
}

type loginResponse struct {
	XMLName xml.Name `xml:"loginResponse"`
	Result  Result  `xml:"result"`
}

type Result struct {
	MetadataServerURL string   `xml:"metadataServerUrl"`
	PasswordExpired   string   `xml:"passwordExpired"`  
	Sandbox           string   `xml:"sandbox"`          
	ServerURL         string   `xml:"serverUrl"`        
	SessionID         string   `xml:"sessionId"`        
	UserID            string   `xml:"userId"`                
}

// use viper package to read .env file
// return the value of the key
func viperEnvVariable(key string) string {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile(".env")
  
	// Find and read the config file
	err := viper.ReadInConfig()
  
	if err != nil {
	  log.Fatalf("Error while reading config file %s", err)
	}
  
	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)
  
	fmt.Println(value)
	fmt.Println(key)
	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
	  log.Fatalf("Invalid type assertion")
	}
  
	return value
  }

func main() {

	sfprodurl := viperEnvVariable("SF_PROD_URL")
//	sfsandboxurl := viperEnvVariable("SF_SANDBOX_URL")
	username := viperEnvVariable("USERNAME")
	password := viperEnvVariable("PASSWORDANDTOKEN")

	url := sfprodurl + "/services/Soap/u/54.0"
	method := "POST"

	payload := strings.NewReader(`<?xml version="1.0" encoding="utf-8" ?>
  <env:Envelope xmlns:xsd="http://www.w3.org/2001/XMLSchema"
	  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	  xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
	<env:Body>
	  <n1:login xmlns:n1="urn:partner.soap.sforce.com">
		<n1:username><![CDATA[`+username+`]]></n1:username>
		<n1:password><![CDATA[`+password+`]]></n1:password>
	  </n1:login>
	</env:Body>
  </env:Envelope>`)
  
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)
  
	if err != nil {
	  fmt.Println(err)
	  return
	}
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("SOAPAction", "login")
	req.Header.Add("charset", "UTF-8")
	req.Header.Add("Accept", "text/xml")
  
	res, err := client.Do(req)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	defer res.Body.Close()
  
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	  fmt.Println(err)
	  return
	}

	resSoap := &MyRespEnvelope{}
	errSoap := xml.Unmarshal(body, resSoap)
	if errSoap != nil {
		fmt.Println(errSoap)
		return
	}

	var sessionId = resSoap.Body.GetResponse.Result.SessionID
	var serverURL = strings.Split(resSoap.Body.GetResponse.Result.ServerURL, "/")
	var frontDoor = "https://"+serverURL[2]+"/secur/frontdoor.jsp?sid=" + sessionId
	fmt.Println(frontDoor)
	browser.OpenURL(frontDoor)
}
