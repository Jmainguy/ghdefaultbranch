package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func renameBranch(repo, token, oldName, newName string) {

	data := payload{NewName: newName}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body := bytes.NewReader(payloadBytes)

	uri := fmt.Sprintf("https://api.github.com/repos/%s/branches/%s/rename", repo, oldName)

	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	token = fmt.Sprintf("token %s", token)
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

}
