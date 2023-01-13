# Go Kafka consumer

This is a simple Kafka consumer written in Go as test for the [OD-69](https://entwicklung-bikeleasing.myjetbrains.com/youtrack/issue/OD-69/Evaluate-techstacks-for-service-brokers) Ticket in the 
*Online-Delivery Project*


## Structure


The project is structured as follows:

* `main.go` - The main file of the project to be executed when starting this app
* `consumer/` - The package containing the consumer logic
* `producer/` - The package containing the producer logic - It simply writes back the same message to another topic
* `api/` - A test API that simply takes the messages and redirects it as it is to a echo API


### Sample Messages

The messages were JSON payload generated using [`JSON Generator`](https://json-generator.com/) and have this structure:

~~~
[
  '{{repeat(5, 7)}}',
  {
    _id: '{{objectId()}}',
    crmId: '{{guid()}}',
    isActive: '{{bool()}}',
    age: '{{integer(20, 80)}}',
    name: '{{firstName()}} {{surname()}}',
    gender: '{{gender()}}',
    company: '{{company().toUpperCase()}}',
    email: '{{email()}}'
  }
]
~~~

**Sample Output**
~~~
{
	"_id": "63c184fc0b003b5184665453",
	"crmId": "81ee5d9b-a92b-4a47-a3f8-c1cb918d2fd9",
	"isActive": true,
	"age": 39,
	"name": "Wilcox Anthony",
	"gender": "male",
	"company": "NEPTIDE",
	"email": "wilcoxanthony@neptide.com"
}
~~~

### Test Echo API

As test API I've used `https://echo.zuplo.io`, which is a dummy web server that
just re-echos back what it got.

An alternative would have been the much more complex [Postman ECHO API](https://www.postman.com/postman/workspace/published-postman-templates/documentation/631643-f695cab7-6878-eb55-7943-ad88e1ccfd65), a REST API service
that allows you to test your REST clients and make sample API calls with any methods and
configuration. It just returns back the request as a response.


### Note

The  Kafka local setup is not included in this project, it's just a copy of [this repo](https://github.com/conduktor/kafka-stack-docker-compose/blob/master/full-stack.yml)