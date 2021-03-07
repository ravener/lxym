
package pages

const Protocol_Buffer_3 = "# Protocol Buffers\n\n## Why Protocol Buffers\n\nProtocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler.\nYou define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.\nProtocol Buffers are Schema Of Messages. They are language agnostic.\nThey can be converted to binary and converted back to message formats using the code generated by the protoc compiler for various languages.\n\n```\n/*\n* Language Syntax\n*/\n\n/*\n* Specifying Syntax Of Protocol Buffer Version\n* Specifying Which Protocol Buffer Version To Use\n* It can be usually proto3 or proto2\n*/\nsyntax = \"proto3\";\n\n/*\n* Declaring Message In Protocol Buffer:\n* As you can see, each field in the message definition has a unique number.\n* These field numbers are used to identify your fields in the message binary format,\n* and should not be changed once your message type is in use.\n* Note that field numbers in the range 1 through 15 take one byte to encode, including the field number and the field's type (you can find out more about this in Protocol Buffer Encoding).\n* Field numbers in the range 16 through 2047 take two bytes. So you should reserve the numbers 1 through 15 for very frequently occurring message elements.\n* Remember to leave some room for frequently occurring elements that might be added in the future.\n* The smallest field number you can specify is 1, and the largest is 2^29 - 1, or 536,870,911.\n* You also cannot use the numbers 19000 through 19999 (FieldDescriptor::kFirstReservedNumber through FieldDescriptor::kLastReservedNumber),\n* as they are reserved for the Protocol Buffers implementation - the protocol buffer compiler will complain if you use one of these reserved numbers in your .proto.\n* Similarly, you cannot use any previously reserved field numbers.\n*\n*/\n\n/*\nSyntax For Declaring Message:\n    message ${MessageName} {\n        ${Scalar Value Type} ${FieldName1} = ${Tag Number1};\n                .\n                .\n                .\n        ${Scalar Value Type} ${FieldNameN} = ${Tag NumberN};\n    }\n\nDefault Values Will be applied any case if the message doesn't contain a existing field defined\nin the message definition\n*/\n\nmessage MessageTypes {\n    /*\n    * Scalar Value Types\n    */\n    string stringType = 1; // A string must always contain UTF-8 encoded or 7-bit ASCII text. Default value = \"\"\n\n    // Number Types, Default Value = 0\n    int32 int32Type = 2; // Uses Variable Length Encoding. Inefficient For Negative Numbers, Instead Use sint32.\n    int64 int64Type = 3; // Uses Variable Length Encoding. Inefficient For Negative Numbers, Instead Use sint64.\n    uint32 uInt32Type = 4; // Uses Variable Length Encoding\n    uint64 uInt64Type = 5; // Uses Variable Length Encoding\n    sint32 sInt32Type = 6; // Uses Variable Length Encoding. They are efficient in encoding for negative numbers.\n                           // Use this instead of int32 for negative numbers\n    sint64 sInt64Type = 7; // Uses Variable Length Encoding. They are efficient in encoding for negative numbers.\n    // Use this instead of int64 for negative numbers.\n\n    fixed32 fixed32Type = 8; // Always four bytes. More efficient than uint32 if values are often greater than 2^28.\n    fixed64 fixed64Type = 9; // Always eight bytes. More efficient than uint64 if values are often greater than 2^56\n\n    sfixed32 sfixed32Type = 10; // Always four bytes.\n    sfixed64 sfixed64Type = 11; // Always Eight bytes.\n\n    bool boolType = 12; // Boolean Type. Default Value = false\n\n    bytes bytesType = 13; // May contain any arbitrary sequence of bytes. Default Value = Empty Bytes\n\n    double doubleType = 14;\n    float floatType = 15;\n\n    enum Week {\n        UNDEFINED = 0; // Tag 0 is always used as default in case of enum\n        SUNDAY = 1;\n        MONDAY = 2;\n        TUESDAY = 3;\n        WEDNESDAY = 4;\n        THURSDAY = 5;\n        FRIDAY = 6;\n        SATURDAY = 7;\n    }\n    Week wkDayType = 16;\n\n    /*\n    * Defining Collection Of Scalar Value Type\n    * Syntax: repeated ${ScalarType} ${name} = TagValue\n    */\n    repeated string listOfString = 17; // List[String]\n}\n\n/*\n* Defining Defined Message Types In Other Message Definition\n*/\nmessage Person {\n    string fname = 1;\n    string sname = 2;\n}\n\nmessage City {\n    Person p = 1;\n}\n\n/*\n* Nested Message Definitions\n*/\n\nmessage NestedMessages {\n    message FirstLevelNestedMessage {\n        string firstString = 1;\n        message SecondLevelNestedMessage {\n            string secondString = 2;\n        }\n    }\n    FirstLevelNestedMessage msg = 1;\n    FirstLevelNestedMessage.SecondLevelNestedMessage msg2 = 2;\n}\n\n/*\n* Importing Message From A File\n*/\n\n// one.proto\n// message One {\n//     string oneMsg = 1;\n// }\n\n// two.proto\n//  import \"myproject/one.proto\"\n//  message Two {\n//       string twoMsg = 2;\n//  }\n\n\n/*\n* Advanced Topics\n*/\n\n/*\n* Handling Message Type Changes:\n* Never Change/Use The TagNumber Of A Message Field Which Was Removed\n* We should use reserved in case of message definition update.\n* (https://developers.google.com/protocol-buffers/docs/proto3#updating)\n*/\n\n/*\n* Reserved Fields\n* It's used in case if we need to add/remove new fields into message.\n* Using Reserved Backward and Forward Compatibility Of Messages can be achieved\n*/\n\n\nmessage ReservedMessage {\n    reserved 0, 1, 2, 3 to 10; // Set Of Tag Numbers Which Can't be reused.\n    reserved \"firstMsg\", \"secondMsg\", \"thirdMsg\"; // Set Of Labels Which Can't Be reused.\n}\n\n/*\n* Any\n* The Any message type lets you use messages as embedded types without having their .proto definition.\n* An Any contains an arbitrary serialized message as bytes,\n* along with a URL that acts as a globally unique identifier for and resolves to that message's type.\n* For Any to work we need to import it as shown below.\n*/\n/*\n    import \"google/protobuf/any.proto\";\n    message AnySampleMessage {\n        repeated google.protobuf.Any.details = 1;\n    }\n\n*/\n\n\n/*\n*  OneOf\n* There are cases, wherein only one field at-most might be present as part of the message.\n* Note: OneOf messages can't be repeated.\n*/\n\nmessage OneOfMessage {\n    oneof msg {\n        string fname = 1;\n        string sname = 2;\n    };\n}\n\n/*\n* Maps\n* Map fields cannot be repeated.\n* Ordering Of A Map Is Not Guaranteed.\n*/\n\nmessage MessageWithMaps {\n    map<string, string> mapOfMessages = 1;\n}\n\n\n/*\n* Packages\n* Used for preventing name clashes between protocol message types\n* Syntax:\n    package ${packageName};\n\n    To Access the package;\n    ${packageName}.${messageName} = ${tagNumber};\n*/\n\n/*\n* Services\n* Message Types Defined For Using In RPC system.\n*  When protoc compiler generates for various languages it generates stub methods for the services.\n*/\n\nmessage SearchRequest {\n    string queryString = 1;\n}\n\nmessage SearchResponse {\n    string queryResponse = 1;\n}\nservice SearchService {\n    rpc Search (SearchRequest) returns (SearchResponse);\n}\n```\n\n## Generating Classes In Various Languages For Protocol Buffers\n\n```shell\nprotoc --proto_path=IMPORT_PATH --cpp_out=DST_DIR --java_out=DST_DIR --python_out=DST_DIR --go_out=DST_DIR --ruby_out=DST_DIR --objc_out=DST_DIR --csharp_out=DST_DIR path/to/file.proto\n```\n\n## References\n\n[Google Protocol Buffers](https://developers.google.com/protocol-buffers/)"
