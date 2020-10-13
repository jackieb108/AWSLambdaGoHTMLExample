# AWSLambdaGoHTMLExample
Example to show how to build AWS lambda service with golang and invokde it using AWS API Gateway.

Setup used in this projects
1. AWS API Gateway (HTTP protocol) invokes lambda function. API Gateway is invokded by calling the published URL
2. Lambda code built using golang call 2 APIs. First REST API call retrieves US population data as JSON. Lambda function extracts years/US population from this API response. Lambda function then retrieves high chart template file from s3. It replaces the xAxia and yAxis data in the chart with JSON extracted years/US population data. Lambda function returns the updated  highchart html file with years/US population
3. AWS API Gateway send the html file back to the end user

![images](/assets/images/AWSLambdaGoHTMLExampleArchi.jpg)
