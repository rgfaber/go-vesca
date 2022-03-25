package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
}

func getAssets() GetAssetsRsp {
	url := "https://opensea-data-query.p.rapidapi.com/api/v1/assets?collection=ongakucraft&order_direction=desc&limit=20&offset=0"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "opensea-data-query.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "a90252195amsh7385cc8c5a70e4fp1e5a18jsn4b80044d5e0c")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	err := ioutil.WriteFile("assets.json", body, 0644)
	if err != nil {
		panic(err)
	} //	fmt.Println(res)

	return GetAssetsRsp(body)

}
