# Web Calculator

An overcomplicated calculator built for the cloud

Originally based on an assignment, the intentions of this self-learning project were to familiarise myself with the fundamentals of modern cloud computing by leveraging a variety of cloud technologies to develop an unnecessarily complicated a web-based calculator

## Docker and Kubernetes

Each component/service below was containerised with Docker.

The containers were orchestrated with Docker-Compose during development.

For the Go and Java component containers, I used multi-stage builds to drastically reduce the final image size (From over 800mb to just under 30mb with Go) as well as to automatically compile it.

In production, the containers were orchestrated with Kubernetes and deployed to Rancher.

## Components

Originally, each of these components had their own repository.
However, since the application is no longer in deployment (for financial reasons), all of the components are within this repository.

### APIs

| Component  | Language |  Framework  |             Testing              |   CI   |
| :--------: | :------: | :---------: | :------------------------------: | :----: |
|   [Add]    |   PHP    |      -      |               Unit               |   -    |
|  [Modulo]  |  Python  |    Flask    |              PyTest              | GitLab |
| [Multiply] |   Ruby   |   Sinatra   |              RSpec               | GitLab |
|  [Power]   |   Java   | Spring Boot | JUnit \| SpringRunner \| MockMvc | GitLab |
| [Quotient] |    Go    |     Gin     |  testing \| httptest \| Testify  | GitLab |
| [Subtract] |   Node   |   Express   |          Mocha \| Chai           | GitLab |

[add]: (https://www.github.com/Inaki-McKearney/Web-Calculator/tree/master/webcalc-add)
[subtract]: (https://www.github.com/Inaki-McKearney/Web-Calculator/tree/master/webcalc-subtract)
[modulo]: (https://www.github.com/Inaki-McKearney/Web-Calculator/tree/master/webcalc-modulo)
[multiply]: (https://www.github.com/Inaki-McKearney/Web-Calculator/tree/master/webcalc-multiply)
[power]: (https://www.github.com/Inaki-McKearney/Web-Calculator/tree/master/webcalc-power)
[quotient]: (https://www.github.com/Inaki-McKearney/Web-Calculator/tree/master/webcalc-quoutient)

### Custom Proxy Router

I wrote a proxy in Go incorporating the multiplexer package Gorilla Mux. It first reads in all available URLs to a map so as to allow customising the URLs without having to rebuild the binary.

Using Mux and regex, any requests to the endpoint are checked for an alphabetic endpoint. This operator endpoint is then used as a key in the aforementioned map to see if a URL exists for the provided operator. If one is not found a "404 operator does not exist" is returned, otherwise, the request is passed in its original form to the appropriate URL before its response is returned back through the proxy.

The service response's "Access-Control-Allow-Origin" header is set to "\*" in order to enable CORS across all services via the proxy service alone. This prevented me from having to individually enable it for each service.

### Monitoring and Testing Service

I wrote a monitoring tool using Node and Express. Every sending, the service would select random values and operations before sending a GET request to that endpoint.

The results are logged using Winston should results need analysed. The responses are all timed and the correctness is measured in real time.

Accessing the service URL displays these relevant statistics in tabular format and the average times are displayed using google charts.

![Monitoring Service](https://github.com/Inaki-McKearney/Web-Calculator/blob/master/monitor-screenshot.png?raw=true)
