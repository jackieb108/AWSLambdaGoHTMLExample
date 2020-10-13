package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

const (
	POPULATIONAPIURL = "https://datausa.io/api/data?drilldowns=Nation&measures=Population"
	HIGHCHARTS3URL   = "https://<REPLACE WITH S3 PATH>/HighChartTemplate.html" //REPLACE with your link to HighChartTemplate
	//HIGHCHARTS3URL = "Attempt2.html"
)

var (
	finaloutputhtml = "Fill in the value"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type MyEvent struct {
	Name string `json:"name"`
}

type Response struct {
	DataLst []DataList `json:"data"`
}

type DataList struct {
	CountryName string `json:"Nation"`
	Year        string `json:"Year"`
	Population  int    `json:"Population"`
	IDNation    string `json:"ID Nation"`
}

func getAPIData() (categoryLst string, seriesLst string) {
	response, err := http.Get(POPULATIONAPIURL)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	//fmt.Println(responseObject.DataLst)
	//fmt.Println(len(responseObject.YearLst))

	categoryLst = "categories :["
	seriesLst = "name: 'United States', data: ["
	for i := len(responseObject.DataLst) - 1; i >= 0; i-- {
		//fmt.Println(responseObject.DataLst[i].Year)
		if i != len(responseObject.DataLst)-1 {
			categoryLst = categoryLst + ","
			seriesLst = seriesLst + ","
		}
		categoryLst = categoryLst + "'" + responseObject.DataLst[i].Year + "'"
		seriesLst = seriesLst + strconv.Itoa(responseObject.DataLst[i].Population)
		//fmt.Println(responseObject.DataLst[i].Year)
	}
	categoryLst = categoryLst + "]"
	seriesLst = seriesLst + "]"
	//fmt.Println(categoryLst)
	//fmt.Println(seriesLst)
	return categoryLst, seriesLst

}

func getChartTemplateFromS3() (templatehtml string) {
	content, err := http.Get(HIGHCHARTS3URL)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(content.Body)
	if err != nil {
		log.Fatal(err)
	}

	text := string(responseData)
	//fmt.Println(text)
	return text

}

func main() {

	//Call population API and get category (years) and series (population) data
	categoryLst, seriesLst := getAPIData()

	// Now get HTML file from S3
	text := getChartTemplateFromS3()

	//Find the data to be replaces in chart template from s3. Need to replace categories and series with data from population API
	rcat := regexp.MustCompile(`var\s*xAxis\s*=\s*\{(.|\n|\r|\r\n)*?\};`)       // regex to search categories (years) data in the template. First match
	rseries := regexp.MustCompile(`var\s*series\s*=\s*\[\{(.|\n|\r|\r\n)*?\];`) //regex to search serieas (population0) data in the template

	//replace catergories and serires lists in the template from data retrieved from population API
	out := rcat.ReplaceAllString(text, "var xAxis = {"+categoryLst+"};")
	outfinal := rseries.ReplaceAllString(out, "var series = [{"+seriesLst+"}];")
	//fmt.Println(outfinal)
	//fmt.Println("----")
	finaloutputhtml = outfinal

	lambda.Start(handlerRequest)
}

func handlerRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"content-type": "text/html"},
		Body:       string(finaloutputhtml),
		StatusCode: 200,
	}, nil

}