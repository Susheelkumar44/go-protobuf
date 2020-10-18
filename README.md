# go-protobuf

Client- User facing microservice that will accept REST API Request.
Requests:
POST - hit the end point, pass POST request with JSON Body, Client Service calls Server Service through gRPC call, writes the Request Data to file.
GET - hit the end point, pass the GET request, Client srvice will call read server to transfer data read from the file over http rest api call.
