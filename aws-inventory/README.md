# Multicloud Manager (Golang)

`Cloudlist` is a simple command line tool providing commands to fetch resource and pricing information for your AWS GCP & Azure accounts. 

**Usage**
---

```
Usage: main.go [OPTIONS]

  Cloud Automation Tool. - Fetch Resource usage and pricing across AWS, GCP and Azure cloud providers
  Developed by Colin -> (Github: cdugga)


Commands:
  pricing, p        Fetch Usage and Pricing information.
  resources, r      Lists your cloud resources
  --help, -h        Shows a list of commands or help for one command

Flags
  --provider, -c    Speifiy cloud provider name [AWS|GCP|Azure]
  
SubCommands
  --service-code    Specify cloud resource for which data is required 

```


**Installation Options**
---

1. Install with [`go get`](https://github.com/cdugga/golang-repos/aws-inventory)
    + `$ go install`
    + `$ cloudlist resoruces --provider aws`

2. Download the `cloudlist` binary.


**Environment/Configuration**
--
1. Local setup
    + export AWS_REGION=eu-west-1 AWS_SDK_LOAD_CONFIG=true
    + export AWS_SDK_LOAD_CONFIG=true
    + export GOOGLE_APPLICATION_CREDENTIALS="/root/gcloud/creds.json"

## Authors

* **Colin Duggan** - *All dev work* - [LinkedIn](https://www.linkedin.com/in/colinduggan/?originalSubdomain=ie)


## Acknowledgments

* None yet!