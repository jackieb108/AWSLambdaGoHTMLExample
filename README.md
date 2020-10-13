# AWSLambdaGoHTMLExample
Example to show how to build AWS lambda service with golang and invokde it using AWS API Gateway.

Setup used in this projects
1. AWS API Gateway (HTTP protocol) invokes lambda function. API Gateway is invokded by calling the published URL
2. Lambda code built using golang calls 2 APIs. First REST API call retrieves US population data as JSON. Lambda function extracts years/US population from this API response. Lambda function then retrieves high chart template file from s3. It replaces the xAxia and yAxis data in the chart with JSON extracted years/US population data. Lambda function returns the updated  highchart html file with years/US population filled in
3. AWS API Gateway sends the html file back to the end user


![images](/assets/images/AWSLambdaGoHTMLExampleArchi.jpg)


Steps using GIT project:
1. Clone the git repo

    `$ git clone https://github.com/jackieb108/AWSLambdaGoHTMLExample.git`
    
2. Copy 'HighChartTemplate.html' to s3 bucket of yours. Note the URL for the s3 bucket. Makes sure this URL is accessible to lambda.
  
3. Open the file 'chartUSpopulation.go' and update HIGHCHARTS3URL to point to the s3 location where you copied 'HighChartTemplate.html' in step2. Save the file and exit

4. Navigate to the root directory for the project (<..>/AWSLambdaGoHTMLExample)

5. Note that I am running all the commands on windows 10 with VSCode and Ubuntu 'Windows susbsystem for linux' on windows 10. Couldn't make 'build-lambda-zip' work to cross-compile for linux

6. Set env variables using cmd shell in wondows 10 (SRC: https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)

    `$set GOARCH=amd64`
    
    `$set GOOS = linux`
7. Run the build for go file

    `$go  build  -o chartUSpopulation.out chartUSpopulation.go`
     `
8. zip the file on ubuntu application. I enabled windows subsytem for linuc and then downloaded ubuntu application to run zip for unix flavor

    `$zip -o chartUSpopulation.zip chartUSpopulation`
    
Steps on AWS console 
1. Create 

2. 
