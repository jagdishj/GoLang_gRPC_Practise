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

Installing the goLang : https://www.youtube.com/watch?v=1MXIGYrMk80
protoBuf and gRPC : https://www.youtube.com/watch?v=ES_GI-lmhEU

#### Screenshots
Unary API Greet Example: <img src="Screenshots\UnaryAPI_greetExample.png">
Unary API Calc Sum : <img src="Screenshots\UnaryAPI_Calc_Sum.png">
Server Streaming API Greet Example: <img src="Screenshots\ServerStreamingAPI_greetExample.png">
Server Streaming Calc Prime Decomposition: <img src="Screenshots\ServerStreamingAPI_Calc_PrimeDecomposition.png">
Client Streaming API Greet Example: <img src="Screenshots\ClientStreaming_greetExample.png">
Client Streaming Calc Avg: <img src="Screenshots\ClientStreaming_Calc_Average.png">
Bi-Directional API Greet Example: <img src="Screenshots\Bi-DirectionStreaming_greetExample.png">
Bi-Directional API Calc Find Maximum: <img src="Screenshots\Bi-DirectionStreaming_Calc_FindMaximum.png">
gRPC Error Implementation: <img src="Screenshots\Error_Implementation.png">
Referral Links: <br>
protocol buffers - https://grpc.io/ <br>
HTTP2 vs HTTP1 - https://imagekit.io/demo/http2-vs-http1
Error Codes - https://grpc.io/docs/guides/error/

