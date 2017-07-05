package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//Post Func
func Post(url string, inputValue interface{}, outputValue interface{}) (err error) {

	return callService(url, "POST", inputValue, outputValue)
}

//Delete Func
func Delete(url string, inputValue interface{}, outputValue interface{}) (err error) {

	return callService(url, "DELETE", inputValue, outputValue)
}

//Get Func
func Get(url string, inputValue interface{}, outputValue interface{}) (err error) {

	return callService(url, "GET", inputValue, outputValue)
}

//Put Func
func Put(url string, inputValue interface{}, outputValue interface{}) (err error) {

	return callService(url, "PUT", inputValue, outputValue)
}

func callService(url string, method string, inputValue interface{}, outputValue interface{}) (err error) {

	body, err := json.Marshal(inputValue)
	checkErr(err)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	checkErr(err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	checkErr(err)
	

	err = json.NewDecoder(resp.Body).Decode(&outputValue)
	defer resp.Body.Close()
	checkErr(err)

	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}