## gRPC Learning
### Overview
Use ProtoBuf to define the contract(*.proto) between services and message(request/response). The contract interface will be generated and to be implemented by the developer.

#### ProtoBuf Advantages:
- Language agnostic, code can be generated for any target languages;
- Data is in binary format and can be efficiently serialized(payload is small)
- Convenient for large data transportation
- Allows easy API evolution using rules (?)
- gRPC is the future of micro-service api and mobile-server api

#### gRPC Basics
gRPC is used to define the Message(request and response); and Service(endpoint and service name) Sample *.proto contract as below:

<img src="Screenshots\gRPC_proto_syntex.png">

Commands to Generate the code <br>
windows: protoc --go_out=. --go-grpc_out=. greet/greetpb/greet.proto <br>
Linus: protoc greet/greetpb/greet.proto --go_out=plugins=grpc:. <br>

gRPC server is async by default, no blocking threads on request, each server can serve millions of requests in parallel gRPC client can be sync or async, flexible in either mode and can perform client side load balancing (?).

#### Comparison of gRPC and JSON
- Small Size. Reduce CPU consumption for JSON parsing
- ProtoBuf(binary) is more efficient and faster, good for mobile platform.
- ProtoBuf contract is generated automatically on the fly.
- ProtoBuf messages allows communicating other micro services in different languages, there’s no languages restrictions in each micro services.

<img src="Screenshots\JSON_ProtocolBuffer.png">

#### HTTP/2
- HTTP2 is much faster than conventional HTTP<br>
- HTTP1.1 opens new TCP connection to server at each request, does not compress header, only works on req/res pattern, no server push is allowed.
Headers in plain text, each req/res exchange involves many overhead in transmission.<br>
- HTTP2 allows client and server push messages in parallel in same TCP connection, reduces latency. Support server push and header compression , it’s via binary protocol, faster and efficient , It’s also secure.<br>
- In HTTP2, once the single TCP connection is established, server and client starts multiple roundtrip for subsequent communications in same TCP connection.

#### Types of gRPC APIs
<img src="Screenshots\TypesofAPIs.png">
<img src="Screenshots\procto_apis.png">

#### gRPC and REST
<img src="Screenshots\gRPCvsREST.png">

Installing the goLang : https://www.youtube.com/watch?v=1MXIGYrMk80 <br>
protoBuf and gRPC : https://www.youtube.com/watch?v=ES_GI-lmhEU

#### Output Screenshots
Unary API Greet Example: <img src="Screenshots\UnaryAPI_greetExample.png">
Unary API Calc Sum : <img src="Screenshots\UnaryAPI_Calc_Sum.png">
Server Streaming API Greet Example: <img src="Screenshots\ServerStreamingAPI_greetExample.png">
Server Streaming Calc Prime Decomposition: <img src="Screenshots\ServerStreamingAPI_Calc_PrimeDecomposition.png">
Client Streaming API Greet Example: <img src="Screenshots\ClientStreaming_greetExample.png">
Client Streaming Calc Avg: <img src="Screenshots\ClientStreaming_Calc_Average.png">
Bi-Directional API Greet Example: <img src="Screenshots\Bi-DirectionStreaming_greetExample.png">
Bi-Directional API Calc Find Maximum: <img src="Screenshots\Bi-DirectionStreaming_Calc_FindMaximum.png">
gRPC Error Implementation: <img src="Screenshots\Error_Implementation.png">
UnaryAPI-ClientCall with Deadline <img src="Screenshots\UnaryAPI_deadline.png">
gRPC-auth-TLC-implementation <br>
    - Error <img src="Screenshots\gRPC_auth_TLS_inSecure_Error.png">
    - Success <img src="Screenshots\gRPC_auth_TLS_inSecure_Success.png"><br>
gRPC-Reflection - Evans CLI  <img src="Screenshots\EvansCLI_Reflection.png">
blog-crud <img src="Screenshots\blob_crud.png"> 
Referral Links: <br>
protocol buffers - https://grpc.io/ <br>
HTTP2 vs HTTP1 - https://imagekit.io/demo/http2-vs-http1 <br>
Error Codes - https://grpc.io/docs/guides/error/ <br>
RPC Deadline - https://grpc.io/blog/deadlines/ <br>
gRPC Auth - https://grpc.io/docs/guides/auth/ <br>
TLC Implementation - https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md <br>
OAuth2 - https://github.com/grpc/grpc-go/tree/master/examples/features/authentication <br>
OpenSSL - TLC - https://github.com/simplesteph/grpc-go-course/tree/master/ssl <br>
Evans CLI - https://github.com/ktr0731/evans#from-github-releases <br>
Mongodb - https://docs.mongodb.com/manual/tutorial/install-mongodb-on-windows/ <br>
Robo 3T - https://robomongo.org/download <br>
Mongodb driver for GO - https://github.com/mongodb/mongo-go-driver <br>
Google PubSub - https://github.com/googleapis/googleapis/blob/master/google/pubsub/v1/pubsub.proto <br>
Google Spanner - https://github.com/googleapis/googleapis/blob/master/google/spanner/v1/spanner.proto <br>
gRPC Gateway - https://github.com/grpc-ecosystem/grpc-gateway <br>
gogo project - https://github.com/gogo/protobuf <br>
gogo documentation refer - https://jbrandhorst.com/post/gogoproto/ <br>
