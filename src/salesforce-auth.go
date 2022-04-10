package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/browser"
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
	Result  Result   `xml:"result"`
}

type Result struct {
	MetadataServerURL string `xml:"metadataServerUrl"`
	PasswordExpired   string `xml:"passwordExpired"`
	Sandbox           string `xml:"sandbox"`
	ServerURL         string `xml:"serverUrl"`
	SessionID         string `xml:"sessionId"`
	UserID            string `xml:"userId"`
}

func GetSalesforceSession(credentials Config) string {

	password := credentials.Password + credentials.SecurityToken

	url := credentials.InstnaceUrl + "/services/Soap/u/54.0"
	method := "POST"

	payload := strings.NewReader(`<?xml version="1.0" encoding="utf-8" ?>
  <env:Envelope xmlns:xsd="http://www.w3.org/2001/XMLSchema"
	  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	  xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
	<env:Body>
	  <n1:login xmlns:n1="urn:partner.soap.sforce.com">
		<n1:username><![CDATA[` + credentials.Username + `]]></n1:username>
		<n1:password><![CDATA[` + password + `]]></n1:password>
	  </n1:login>
	</env:Body>
  </env:Envelope>`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("SOAPAction", "login")
	req.Header.Add("charset", "UTF-8")
	req.Header.Add("Accept", "text/xml")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	resSoap := &MyRespEnvelope{}
	errSoap := xml.Unmarshal(body, resSoap)
	if errSoap != nil {
		fmt.Println(errSoap)
		return ""
	}

	if resSoap.Body.GetResponse.Result.SessionID == "" {
		fmt.Println("Please check login credentials")
	}

	return resSoap.Body.GetResponse.Result.SessionID
}

func openSalesforceBrowser(sessionId string, instanceUrl string) {

	var frontDoor = instanceUrl + "/secur/frontdoor.jsp?sid=" + sessionId
	browser.OpenURL(frontDoor)
}
