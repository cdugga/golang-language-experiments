
# HTML Replace go module built to run on AWS Lambda

Go module to process uploaded HTML table results and replace with a supplied parameter

## Installation

* Compile into Linux based binary executable binary
* set GOOS=linux go build -o main
* GOOS=linux go build -o main

## Usage 

* Create Handler function ( with request/response structsc)
* Create main using AWS Lambda package to define lambda.Start(Handler)
* Compile into Linux binary executable 
* Zip main executable 
* Upload to AWS console( or use AWS cli)
* Specify handler as main

## Authors

* **Colin Duggan** - *All dev work* - [LinkedIn](https://www.linkedin.com/in/colinduggan/?originalSubdomain=ie)


## Acknowledgments

* None yet!
