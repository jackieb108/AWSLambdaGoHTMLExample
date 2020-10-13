# AWSLambdaGoHTMLExample
Example to show how to build AWS lambda service with golang and invokde it using AWS API Gateway.

Setup used in this projects
1. AWS API Gateway (HTTP protocol) invokes lambda function. API Gateway is invokded by calling the published URL
2. Lambda code built using golang calls 2 APIs. First REST API call retrieves US population data as JSON. Lambda function extracts years/US population from this API response. Lambda function then retrieves high chart template file from s3. It replaces the xAxia and yAxis data in the chart with JSON extracted years/US population data. Lambda function returns the updated  highchart html file with years/US population filled in
3. AWS API Gateway sends the html file back to the end user


![images](/assets/images/AWSLambdaGoHTMLExampleArchi.jpg)


## Steps for GIT project:
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
    
## Steps for AWS lambda and API Gateway setup 
1. Create AWS Lambda function `AWS Lambda -> Create Function`
    > Add chart name as 'chartpopulation'
    > Select 'Runime' as 'Go 1.x'
    > Hit 'Create Function' button
2. Upload zipped go file from the laptop or VM (chartUSpopulation.zip) from step 8 in above section
    > `Lambda -> Functions -> chartpopulation` (Configuration tab)
    > Scroll down on the page to `Function Code`
    > Click on `Actions->Upload zip file`
    > Upload the `chartUSpopulation.zip`
    > scroll down on same page `Lambda -> Functions -> chartpopulation` (Configuration tab) to section `Basic`
    > Hit edit and update the `Handler` to `chartUSpopulation` and save
3. Navigate to 'API Gateway`(Setting up API Gateway to service HTTP call adn return HTML file)
    > Hit `Create API` button
    > Choose `HTTP API` option and click `Build`
    > On `Create API` page 1) `Add Integrations`. Choose `Lambda` and provide the lambda service name `chartpopulation`. Lambda names should be avaialble as we loaded the file earlier 2) Pick a name for the API - `LambdaUSPopulationProxyAPI`
    > On 'Configure Routes`, keep method as `ANY` and add resource `chartpopulation`. Click `Next` OR navigate to `Stages`
    > Add `test` as a stage. Click `Next` and hit `Create`
    > Check teh URL on API page for this API - https://XXXXXX.us-east-1.amazonaws.com/test . I have added XXXXXX here in teh path, please be sure to check the path for your API and use it apporpirately
    > Deploy the API using 'test' as stage
    > Test URL by adding the resource path https://XXXXXX.us-east-1.amazonaws.com/test/chartpopulation
    
Calling the API page https://XXXXXX.us-east-1.amazonaws.com/test/chartpopulation will invoke lambda and lambda will do the processing calling APIs and generating highchart graph to display US population. You should see a similar char at below


