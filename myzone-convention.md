# MyZone Introduction

Myzone is an information exchange convention for IoT devices, services and other information producers and consumers


## Mission statement

Assist inhabitants of a zone by providing relevant information and control. Information from IoT devices and other sources is collected, processed and presented. Collected information can be shared with other
zones. Information is presented through desktop and mobile and wearable devices based on the consumer's situation.

## Information and Control

Information can come from many sources like sensors as well as services that derive new information from existing information. This includes any data captured by IoT devices, camera images, data from the internet,
user input, as well as services that generate new information like analytics and machine learning.

Control of devices or services is a form of information, one that is provided by users and feeds back into the devices. If control is shared with other zones, then inhabitants and services in that zone can also provide control.

## Addressing & Zoning

Anything that produces information is called a 'node', including people. Nodes reside in a zone and each node has a unique address consisting of the zone, publisher and nodeId.

A zone is a physical or virtual area where information resides and can be presented. A zone can be a home, a street, a city, or an area in a virtual game world. All consumers within a zone have access to information published in that zone. Each zone has a unique identifier. A zone ID does not have to be globally unique but must be unqiue between zones it shares information with. It is valid to use a domain name as a zoneId where the dots are replaced with dashes.

Information that is collected within a zone can be shared with other zones of choice. External zones can subscribe to information that is made available to them. Sharing takes place through a secure connection
between zones.

For example, a water level sensor provides water levels to a city monitoring zone. A service within the monitoring zone interprets the water levels from multiple sensors and determines the risk level for flooding. This risk level information is shared with the city's community zone and available to residents and visitors of the city website.

A virtual game uses zones for its street map that are bridged to street zones in the real world. The number of people in the real world is reflected in the zone of the game world; An alarm triggered in the real world shows up in the game world; A message sent in the game world shows up in the real world. Once support for zones is available in the game the limitation is the imagination.

More examples are presented in the MyZone use-cases document.

## Services Derive Information From Data

The amount of data collected easily becomes overwhelming. In itself this raw data is often not immediately useful and can lead to information overload, or simply ignoring the data, which defeats the purpose of
collecting it.

Services in a zone derive information from collected data and turn it into something that is useful to the consumer. For example, a service issues a security alert when a motion sensor triggers when no-one is
home. This derived information is useful information, while the motion sensor trigger on its own is not useful if the consumer only wants to know if there is a security breach. These services can be simple rule
based logic, or more advanced like a neural network that uses image recognition to classify the object seen on camera. Services are also nodes with inputs and outputs.

## Information Identification and Transparency

When consuming information from external sources, trust in the validity of this information and the ability to identify its source is important. Faulty sensors or bad actors can generate information that is unreliable. It should be easy to identify these if needed.

Published information carries metadata like the timestamp, location, and identification of the producer. Producer identification provides transparency as to the source of the information. The consumer can choose to include or exclude information from unverified sources. Identification and location can be omitted in case of privacy concerns.

For example, an alarm company receives a security alert. The alert signature is verified to be from a client and not from a bad actor creating a distraction. The location and timestamp are used to identify where and when the alert was given.

## Presenting information

Inhabitants of a zone can be notified of updates to information based on the information priority and the situation of the inhabitant.

Situational awareness can come from location, time of day, activity and other information. It can be used to filter collected information before it is presented, or delay its presentation until the situation has
changed. The location and activity of an inhabitant can be determined via a portable or wearable device linked to that inhabitant, or derived from cameras or other sensors.

Once information updates are accepted it is presented to the zone inhabitant through the available presentation device. This can simply be shown on all devices associated with the inhabitant, or the device
currently in use.

Presented information has a life span. Stale information that has expired should be removed from presentation, depending on the type of information. This can be a personal preference.

Notification and presentation of notifications are provided by services that combine information and decide when and where to pass the information to the consumer. MyZone provides the infrastructure and ease of interoperability to create reusable building blocks for these features.

## Discovery and Configuration

Publishers of information also provide discovery and configuration metadata of the information.The discovery metadata describes the type of information, its publisher, and its configuration. Standardization of discovery and configuration allows data and information to be managed from a single user interface regardless of the various technologies involved.

Changing configuration and controlling inputs can be limited to specific users as identified by the signature contained in the configuration and control messages.

## Technology

MyZone is technology agnostic. It is a convention that describes the information format and exchange, discovery, configuration and zoning, irrespective of the technology used to implement it. Use of different
technologies will serve to improve further integration and makes it easier to expand the information network.

A reference implementation is provided, written in the golang and typescript languages, using the MQTT service bus for publishing information, discovery and configuration in the JSON format.

## Data Format

The information exchange rules must be followed by all implementations of the convention. A JSON based encoding of the data is recommended. Other encodings such as XML or whatever becomes popular tomorrow can
also be used.

The primary requirement is that the information fields described are preserved and all information publishers within a zone use the same format. A zone speaks only one data format language.

Future proofing can be achieved by using different zones for different data formats and using a bridge service to share between zones. The bridge maps between old and new data format. This allows for incremental
improvements while maintaining interoperability.

## Versioning

Future versions of this convention must remain backwards compatible. New fields can be added as long as they remain optional. Implementations must accept and ignore unknown fields.

Publishers include the version of the MyZone convention when publishing their node. See 'discovery' for more information.

# Terminology

| Terminology | Description |
| ----------- |:------------|
| Account | The account used to connect a publisher to an information bus |
| Authentication | Method used to identify the publisher and subscriber with the information bus |
| Configuration | Configuration of the node configuration
| Discovery | Description of nodes, their inputs and outputs
| Information | Anything that can be published by a producer. This can be sensor data, images, 
| Information Bus | A transport for publication of information and control. Information is published by a node onto a message bus. Consumers subscribe to information they are interested in.
| Node | A node is a device or service that provides information and accepts control input. Information from this node can be published by the node itself or published by a (publisher) service that knows how to access the node. 
| Node Input | Input to control the node, for example a switch.
| Node Output| Node Information is provided through outputs. For example, the current temperature.
| Publisher| A service that is responsible for publishing node information and handle configuration updates and control inputs. Publishers are nodes. Publishers can sign their publications to provide source verification.
| Subscriber| Consumer of data or information
| Zone| An area in which information is shared between inhabitants



## Zone Overview

A Zone combines the information from its publishers into an information
bus. Consumers of this information can subscribe to the information and
get notified of updates. Various bus implementations are available. The
reference implementation uses MQTT.

The Zone Bridge is a service that exchanges shared information with
other zones.

\<image\>

A publisher consists of four parts, discovery, configuration, inputs and
outputs.cab

\<image\>

The MyZone convention describes the data exchange format between zone
publishers and consumers of information.

Publishers make information available on addresses, while consumers
subscribe to these addresses to receive the information. The addressing,
structure and security of this information is defined as part of this
convention.

Publishers are responsible for:
1.  Publishing output information 

    This is mandatory. Every publisher must as a minimum publish their information on the value address (see in/output addresses).

2.  Handling requests to update inputs. 
   
    This is only for nodes that have inputs. Publishers can implement constraints that only trusted users can update the inputs.

3.  Publishing discovery information
   
    ... for available nodes, node configuration, node inputs, node outputs. This is optional and intended for environments where the computing power is available.

4.  Update node configuration.

    This is optional and intended for environments where the computing power is available. Publishers can implement constraints that only trusted users can update the inputs.

All publishers in a zone must use the same data interchange format of
the published records. The recommended format is JSON. Other formats
like BSON or XML can be used as long as all publishers of the zone use
the same format. Zone Bridges exchange information in JSON.

# Inputs and Outputs

Information flows between nodes to consumers via publishers. Node
information is published as outputs and control is handled via publisher
inputs.

## In/Output Addresses

The addresses used to publish outputs and control inputs consist of
segments. The address segments include the zone, the publisher of the
information, the node whose information is published or controlled, the
type of information, and the instance of the in- or output.

Segment names consist of alphanumeric, hyphen (-), and underscore (\_)
characters. Reserved keywords start with a dollar ($) character.
Other characters are not recommended to allow for various
publish/subscribe technologies like MQTT, REST or other.

MQTT/REST format example. Depending on the implementation environment,
other methods of separating address segments can be used, as long as
they are consistent within a zone:

* **{zoneId} / {publisherId} / {nodeId} / {type} / {instance} / $value**
* **{zoneId} / {publisherId} / {nodeId} / {type} / {instance} / $latest**
* **{zoneId} / {publisherId} / {nodeId} / {type} / {instance} / $24hours**
* **{zoneId} / {publisherId} / {nodeId} / {type} / {instance} / $set**


|Address segment |Description|
|----------------|-----------|
|{zoneId} | The zone in which publishing takes place. 
| {publisherId} | The service that is publishing the information. A publisher provides its identity when publishing its discovery. The publisher Id is unique within its zone.
| {nodeId} | The node that owns the input or output. This is a device identifier or a service identifier and unique within a publisher.
| {type} | The type of  input or output, eg temperature. This convention includes a list of output types. 
| {instance} | The instance of the type on the node. For example, a node can have two temperature sensors. The combination type + instance is unique for the node. The instance can be a name or number. If only a single instance exists the instance can be shortened to “_”
| Keywords: | |
| *$value* | The “\$value” keyword indicates the publisher provides the latest known value of the output. The payload is the raw data. This is intended for constrained devices and for interoperability with 3rd party consumers.
| *$latest* | The “\$latest” keyword indicates the publisher provides the latest known value of the output including metadata such as node address, timestamp, and the publisher signature. The value is converted to a string. Binary data is converted to base64.
| *$24hours* | The “\$24hours” keyword indicates the publisher provides a record containing a 24 hour history of the values. This is intended to be able to determine a trend without having to store these values. The value is provided in its string format. The content is not required to persist between publisher restarts.
| *$set* | The “\$set” keyword indicates a consumer is providing the value to control a node input. The publisher subscribes to updates published to this address and updates the node input accordingly.


**MQTT Examples:**

1. The value of the first temperature sensor of node 5 published by a service named  *openzwave* is published on an MQTT bus on topic:
    > **myzone/openzwave/5/temperature/1/$latest**

2. To activate a switch on device node with ID "3", a message is published on topic:
    > **myzone/openzwave/3/switch/1/$set**

## '$value' Output

The payload used with the 'value' output is the straight raw data
without metadata.

Publishing information on this address is required. It is primarily
intended for compatibility with 3rd party systems or for use in
environments with limited bandwidth or computing power.
Example:
```
zone-1/openzwave/6/temperature/0/$value: 20.6
```

## '$latest' Output
---------------

The payload used with the '$latest' output includes the address and
timeStamp of the information and optionally the publisher signature to
verify the content. Publishing information on this address is
recommended for environments that are not too limited in bandwidth and
computing power.

The payload structure is as follows:

| Field |      | Data Type   | Required | Description
|-------|------|-------------|----------|------------
| address     || Address  |  optional | Record with the node address
|| zoneId      | string   | required | zone address segment
|| publisherId | string   | required | publisher address segment
|| nodeId      | string   | required | node address segment
| type        || string   | required | The output type of this node
| instance    || string   | required | The input instance of this node. Use 0 if only a single instance exists
| signature   || string   | optional | Signature of this record, signed by the publisher
| timeStamp   || string   | required | ISO8601 "YYYY-MM-DDTHH:MM:SS.sssTZ"
| unit        || string   | optional | unit of value type
| value       || string   | required | value in string format


JSON example of a '$latest' publication:
```
zone-1/openzwave/6/temperature/0/$latest:
{
  "address":{
    "zoneId" : "zone-1",
    "publisherId": "openzwave",
    "nodeId": "5",
  },
  "type": "temperature",
  "instance": "0",
  "signature": "..."
  "timeStamp": "2020-01-16T15:00:01.000PST",
  "unit": "C",
  "value": "20.6",
}
```

## '$24hours' Output

The payload for the '24hours' output contains a history of the values of
the last 24 hours along with address information and signature. It is
updated each time a value changes. 

This publication is optional and intended for environments with plenty of bandwidth and computing power. Consumers can use this history to display a recent trend, like
temperature rising or falling, or presenting a graph of the last 24 hours.

The payload structure is as follows:

| Field      |     | Data Type   | Required | Description |
| ---------- | --- | ----------- | -------- | ------------ |
| address         || Address     | optional | Record with the node address |
| type         | | string      | required | The output type of this node |
| instance     | | string      | required | The input instance of this node |
| history      | | list        | required | eg: [{"timeStamp": "YYYY-MM-DDTHH:MM:SS.sssTZ","value": string}, ...] |
|| timeStamp     | string      | required | ISO8601 "YYYY-MM-DDTHH:MM:SS.sssTZ" |
|| value         | string      | required | Value in string format |
| signature    | | string      | optional | Signature of this record, signed by the producer |
| unit         | | string      | optional | unit of value type |


A JSON example:
```
zone-1/openzwave/6/temperature/0/$24hours:
{
  "address":{
    "zoneId" : "zone-1",
    "publisherId": "openzwave",
    "nodeId": "5",
  },
  "type": "temperature",
  "instance": "0",
  "history": [
    {"timeStamp": "2020-01-16T15:20:01.000PST", "value" : "20.4" },
    {"timeStamp": "2020-01-16T15:00:01.000PST", "value" : "20.6" },
    ...
  ],
  "signature": "...",
  "unit": "C",
}
```
# '$set' Input 

Publishers subscribe to the '$set' input address to receive requests to
update the input of a node.

Subscribing to the set address is only for nodes that have inputs.

The payload structure is as follows:

| Field      | | Data Type   | Required | Description
|------------|----|-------------|----------|------|
| address     || Address  |  optional | Record with the node address
| type || string   | required  | The input type to set
| instance    || string   | required  | The input instance to set 
| timeStamp   || string   | required  | Time this request was created, in ISO8601 format, eg YYYY-MM-DDTHH:MM:SS.sssTZ. The timezone is the local timezone where the value was published. If a request was received with a newer timestamp, up to the current time, then this request is ignored.
| sender      || Address  | required | Address of the node representing the sender requesting to update the input
| signature   || string   | optional | Signature of this record, signed by the sender that wants to set the input.
| value       || string   | required | The input value


A JSON example:
```
zone-1/openzwave/6/switch/0/$set:
{
  "address":{
    "zoneId" : "zone-1",
    "publisherId": "openzwave",
    "nodeId": "6",
  },
  "type": "switch",
  "instance": "0",
  "sender": {
    "zoneId" : "zone-1",
    "publisherId": "dashboard",
    "nodeId": "bob",
  },
  "timeStamp": "2020-01-02T22:03:03.000PST",
  "signature": "...",
  "value": "true",
}
```

### In/output Types

The type of data being published is obviously widely varied. To
facilitate interoperability between publishers and consumers the
input and output types are defined as part of this convention in Appendix A:


# Node Discovery 

Node discovery describes the node, its attributes, configuration, inputs
and outputs. It is usually published by the same publisher that is
responsible for publishing the output values and subscribing to input
control values. Discovery messages can be signed by its publisher to
verify its authenticity.

Publishers that publish discovery must also publish a node that
represents themselves. The publisher's node id must be **\$publisher**.
Publishers that have their own sensors can choose to publish the inputs
and outputs under the **\$publisher** node ID, or publish two records, one
for the **\$publisher** and one for the node with the inputs and outputs.

Publishing of node discovery is optional but highly recommended. It
enables auto discovery, configuration management and information
verification. For very resource restricted devices it can be omitted
however.

## Discovery Address

The addresses used to publish node discovery consists of segments that
describe the zone, publisher of the information, and the node being
discovered.

Each segment consists of alphanumeric, hyphen (-), underscore (_), or dollar ($)
characters. Other characters are not recommended to allow for various
publish/subscribe mediums (like MQTT, REST). The $prefix is for reserved words.

MQTT/REST address format:

> **{zoneId} / {publisherId} / {nodeId} / $discover**

| Address segment  | Description |
| :--------------- | ----------- |
| {zoneId} | The zone in which discovery takes place.
| {publisherId} | The service that is publishing the information. A publisher provides its identity when publishing a node discovery. The publisher Id is unique within its zone. The node can its own publisher. In case of incompatible devices the publisher can be an adapter service that publishes the nodes.
| {nodeId} | The node that is discovered. This is a device identifier or a service identifier and unique within a publisher. Two special nodes are defined: “$publisher” is the service that publishes on behalf of the node. “$gateway” represents the device that acts as a gateway to one or more nodes. For example a zwave controller.
| $discover | Keyword for node discovery. 


MQTT example:

The discovery of a node '5' with a temperature sensor, published by a service named 'openzwave' is published on an MQTT bus on topic:
  > **myzone/openzwave/5/$discover**

The payload describes node 5 with its inputs and outputs.

## Discovery Payload

The discovery payload describes in detail the node, its configuration,
and its inputs and outputs. The objective is for the node to be
sufficiently described so consumers can use it without further
information.

**Discovery Record:**

| Field  | Data Type | Required | Description
| -------- |----------|----------|------------
| address  | Address record |  optional | Record with the node address
| attr     | Dictionary |required |Node attributes provided by the node. Collection of key-value string pairs that describe the node. For interoperability, attribute keys that are part of the convention are described below.
| config   | List of **Configuration records** | optional | Node configuration. Set of configuration objects that describe the configuration options. These can be modified with a ‘configure’ message.
|inputs    | List of Input Records | optional|List of records describing each available input. See input/output definition below.
|outputs   | List of Output Records | optional|List of records describing each available output. See input/output definition below.
|timeStamp | string | optional | Time the reocrd is created
|signature | string | optional | Signature of this record signed by the publisher

**Configuration Record:**

| Field  | Data Type| Required | Description
|--------|----------|----------|------------
| name | string| required | Unique name of the configuration.
| value| string| required| The current configuration value in string format.
|
| datatype | enum| optional| Type of value. Used to determine the editor to use for the value. One of: bool, enum, float, int, string. Default is ‘string’
| default  | string| optional| Default value for this configuration in string format
| description| string | optional | Description of the configuration for human use
| enum | List of strings | optional* | Required when datatype is enum. List of valid enum values as strings
| max | number | optional | Optional maximum value for numeric data
| min | number | optional | Optional minimum value for numeric data
| secret| bool| optional| Optional flag that the value is secret and will not be published. When a secret configuration is set, the value is encrypted with the node public key. 
| 

**Input/Output Discovery Record:**

| Field  | Data Type| Required | Description
|--------|----------|----------|------------
| type | string | required | Type of input/output. See list below
| instance | string | required | The output instance when multiple instances of the same type exist. Default is ‘0’ when only a single instance exists
| value | string | required | The input or output value at time of discovery
| config | List of **Configuration records**|optional|List of Configuration Records that describe in/output configuration. Only used when an input or output has their own configuration. See Node configuration record above for the definition
| datatype | string | optional | Value datatype. One of boolean, enum, float, integer, jpeg, png, string, raw. Default is "string".
| default | string | optional | Default output value
| description | string | optional | Description of the in/output for humans
| enum | list | optional* | List of possible values. Required when datatype is enum
| max | number | optional | Maximum possible in/output value
| min | number | optional | Minimum possible in/output value
| unit | string | optional | The unit of the data type


Example payload for node discovery in JSON format:
```
zone1/openzwave/5/$discover:
{
   "address": {
      "zoneId": "zone1",
      "publisherId": "openzwave",
      "nodeId": "5",
   },
   "attr": {
     "make": "AeoTec",
     "type": "multisensor",
      ...
   },
   "config": {
      name: {
          datatype: string,
          default: “”,
          description: “Friendly name of the node",
          value: “barn multisensor”,
      },
      …
   },
   "inputs": [{
      ...
   }],
   "outputs": [{
     ...
   }],
   "timestamp": "2020-01-20T23:33:44.999PST",
   "signature": "...",
}
```

**Node Attribute Keys:**

| Key | Value Description  |
|--------------|------------- |
| certificate  | A certificate from a trusted source, like Lets Encrypt. It is included by publishers to provide consumers a means to verify their identity
| localip      | IP address of the node, for nodes that are publishers themselves
| location     | String with "latitude, longitude" of device location
| mac | Node MAC address for nodes that are publishers
| manufacturer | Node make or manufacturer
| model | Node model
| myzone       | Version of the convention this publisher uses. This attribute must be present when a publisher publishes its own node
| publicKey    | Publisher's public key used verify the signature provided with publications of information. Only accept public keys from publishers that are verified through their certificate or other means
| type  | Type of node. Eg, multisensor, binary switch, See the nodeTypes list for predefined values
| version      | Hardware or firmware version


# Node Configuration

Nodes that can be configured contain a list of configuration records
described in the node discovery. The configuration value can be updated
with a configuration command as per below.

The configuration of a node can be updated by a consumer by publishing
on the 'configure' address. The node publisher listens to this request
and processes it after validation.

Only authorized users can modify the configuration of a node.

## Configure Address

> {zoneId}/{publisherId}/{nodeId}/$configure

## Configure Payload

|Field 		     |type 		     |required 		     |Description
|--------------|-------------|-----------------|-----------
|address	     |Address      |optional 		     | Address of the node
|config 	     |Dictionary   |required         | key-value pairs for configuration to update { key: value, …}
|sender   		 |Address      |optional 		     | Address of the sender requesting to update the input. This is the publisherId of the consumer, if the consumer publishes.
|signature 		 |string	     |optional 		     | Signature of this configuration record, signed by the consumer 			that wants to modify the configuration. The node publisher can verify if the consumer has permission to modify the configuration of the node.
|timeStamp 		 |string	     |required 		     | Time this request was created, in ISO8601 format


# Node Status

The availability status of a node is published by its publisher when the
availability changes or errors are encountered.

## Status Address

**{zoneId} / {publisherId} / {nodeId} / $status**

| Address segment | Description
|-----------------|--------------
| {zoneId}        | The zone in which discovery takes place. 
| {publisherId}   | The publisher of the node discovery which is handling the configuration update for that node.
| {nodeId}        | The node whose configuration is updated. 
| \$status        | Keyword for node status. Published when the availability of a node changes or new errors are reported. It is published by the publisher of the node.



## Status Payload

| Field 		     |   | type 		     |required 		     |Description
|--------------- |---|----------     |------------     |-----------
|address 		     |   | Address 		   | required 		   |Node address
|status 		     |   | Status record | required 	     |Status record
| | available        | enum (awake, asleep, missing)| required | The node is available
| | errorCount       | integer      | optional        |Nr of errors since startup
| | errorMessage     | string 	     | optional 	     |Last error message
| | errorTime        | string 	     | optional		     |Timestamp of last error message in ISO8601 format
| | lastSeen         | string	       | required		     |Timestamp in ISO8601 format that the publisher received information from the node.
|signature       |   | string        |optional         |Signature of this configuration record, signed by the consumer that wants to modify the configuration. The node publisher can verify if the consumer has permission to modify the configuration of the node.
|timeStamp       |   |string         |required         |Time the status was last updated, in ISO8601 format



# Trust & Digital Signatures

Note: this section needs further review and a reference implementation.

Trust is essential to information exchange between publishers and consumers, especially when the producer and consumer don't know each other directly. In this case trust means that the consumer can be sure
that the publisher is who he claims to be.

To this purpose, publishers include a [[digital signature]](https://en.wikipedia.org/wiki/Digital_signature)
in their publications that lets the consumer verify the records
originate from the publisher. [[This
tutorial]](https://www.tutorialspoint.com/cryptography/cryptography_digital_signatures.htm)
explains it with a picture. Common digital signatures are:

-   [[RSA-PSS]](https://en.wikipedia.org/wiki/Probabilistic_signature_scheme)
    > part of [[PKCS\#1
    > v2.1]{.underline}](https://en.wikipedia.org/wiki/PKCS_1) and used
    > in OpenSSL

-   DSA and [[ECDSA]](https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm)
    > from NIST. NIST produced schemes are considered suspect as there
    > are concerns the NSA has inserted backdoors.

Therefore RSA-PSS is used in this convention. As security is evolving different future schemes allowed if the signature can identify the scheme to use.

## Verification Process

Publishers generate a new key pair on first use and save them in a secure place. The public key is published with the publisher's node information. Consumers can decide to trust a publisher by saving their
public key in a trust store. When information is received, the signature is verified with the public key of that publisher in the trust store. The implementation examples in the following paragraphs point to tools
for each of these steps.

Consumers subscribe to publishers before using their information. During the subscription process, the user can be asked whether to trust this publisher and accept its public key or whether to consider the
information unverified. If accepted, the public key is trusted and the information published by the publisher is verified against this key. If verification fails then the consumer is alerted and the information is
ignored.

Before trusting a publisher's public key, the consumer must perform a verification that this key is indeed from the publisher. There are several options:

1. Manually trust the publisher when subscribing to it.

   This should only be done when there is another way to verify the publisher's identity. For example, the publisher can include a certificate attribute that contains a certificate from Lets Encrypt. On presenting the certificate the consumer can decide to trust the publisher.

2.  Add the publisher's public key using a USB key obtained straight from the publisher.

3.  Add the publisher's public key from a trusted 3rd party.

# Zone Bridging

The purpose of a Zone Bridge is to export information from one zone to the information bus of another zone. This is based on the design that each zone has its own information bus and all subscribers on a bus can read all zones on that bus.

In case multiple zones are on the same bus, and each connection is restricted to its own zone, an access control service is needed. This is beyond the scope of this bridge.

The bridge subscribes to outputs of nodes to be exported and republishes information onto the information bus of the other zone. The original node's information and zone remains unchanged. The zone information now becomes available from the information bus of the second zone.

  * zone-A / publisher-B / node-C [/type/instance]  ->  information bus of zone-B [connection]
  * The information bus of zone-B now publishes information from both zones.

The reference implementation has input pushbuttons for adding and removing inputs and outputs. 

**Bridge (input) control addresses (for managing inputs and outputs)**

Manage bridges by adding and removing them. Each bridge is a node that is configured to connect to another information bus. 
MQTT Example:
* {zoneId} / {publisherId} / $bridgemanager / $pushbutton / $addBridge    - payload is the bridge nodeId
* {zoneId} / {publisherId} / $bridgemanager / $pushbutton / $removeBridge - payload is the bridge nodeId

Address segments:

| Field 		     | Description
|--------------- |-----------
| {zoneId} | The zone the bridge lives in.
| {publisherId}  | The publisher of the bridge. This is the bridge managers Id. Usually there is only a single bridge manager and its ID is "bridge".
| $bridgemanager | Keyword of the node that manages the bridges. There is only 1 bridge manager per publisher
| $pushbutton    | Node input is a push button for adding/removing a bridge instance
| $addBridge     | button to add a new bridge instance. The input payload is its node Id, eg 'bridge-1'
| $removeBridge  | button to remove a bridge instance. The input payload is its node ID, eg 'bridge-1'


**Bridge node configuration fields**
After adding a new bridge, it can be configured with the connection to another information bus. The configuration fields include:

| Field 		   | type 		    |value 		     | Description
|------------- |----------    |------------  |-----------
| address      | string       | required     | IP address or hostname to connect to
| port         | integer      | optional     | port to connect to. Default is determined by protocol
| protocol     | enum ["MQTT", "REST"] | optional  | Protocol to use, MQTT (default), REST API, ...
| format       | enum ["JSON", "XML"]  | optional  | Publishing format used on the external bus. Default is JSON
| clientId     | string       | required     | ID of the client that is connecting
| loginId      | string       | optional     | Login identifier
| credentials  | string       | optional     | Password to connect
| exValue      | boolean      | optional     | Export the node $value publication(s), default=true
| exLatest     | boolean      | optional     | Export the node $latest publication(s), default=true
| exHistory    | boolean      | optional     | Export the node $history publication(s), default=true
| exDiscovery  | boolean      | optional     | Export the node $discovery publication, default=true
| exStatus     | boolean      | optional     | Export the node $status publication, default=true

**Add/remove nodes to export**
To add exported nodes to the bridge use the following pushbuttons for the bridge instance.

MQTT Example:
* {zoneId} / {publisherId} / {bridgeId} / $pushbutton / $addNode
* {zoneId} / {publisherId} / {bridgeId} / $pushbutton / $removeNode

Address segments:

| Field 		     | Description
|--------------- |------------|
| {zoneId} | The zone the bridge lives in.
| {publisherId}  | The publisher of the bridge. This is the bridge managers Id. Usually there is only a single bridge manager and its ID is "bridge".
| $bridgemanager | Keyword of the node that manages the bridges. There is only 1 bridge manager per publisher
| $pushbutton    | Node input is a push button for adding/removing a bridge instance
| $addNode     | button to add a node for export by this bridge. 
| $removeNode  | button to remove a node from export by this bridge. 

**$addNode payload is the input configuration**

| Field 		     | type 		  | required        | Description
|--------------- |----------  |------------     |-----------
| zoneId         | string     | required        | The zone the node lives in
| publisherId    | string     | required        | The publisher of the node to export
| nodeId         | string     | required        | The node to export
| type           | string     | optional        | type of output to export. Default is all node outputs
| instance       | string     | optional        | instance of output type to export. Default is all instances
| exValue        | boolean    | optional        | Export the node $value publication(s), default=true
| exLatest       | boolean    | optional        | Export the node $latest publication(s), default=true
| exHistory      | boolean    | optional        | Export the node $history publication(s), default=true
| exDiscovery    | boolean    | optional        | Export the node $discovery publication, default=true
| exStatus       | boolean    | optional        | Export the node $status publication, default=true



# Appendix A: Input and Output Types

| input/output type| Units  | Value Datatype | description |
|---------------  |---------|------------------|-------------|
| acceleration    | m/s2    | List of floats   | [x,y,z]
| airquality      |         | integer   | Number representing the air quality
| alarm           |         | boolean   | Indicator of alarm status. True is alarm, False is no alarm
| atmosphericpressure | kpa, mbar, Psi, hg | float|  
| avchannel       |         | integer   |
| avmute          |         | boolean   |
| avpause         |         | boolean   |
| avplay          |         | boolean   |
| avvolume        | %       | integer   |
| battery         | %       | integer   |
| co2level        | ppm     | float     |
| colevel         | ppm     | float     |
| color           | rgb     | string    |
| colortemperature| K       | float     |
| compass         | degrees | float     | 0-359 degree compass reading |
| contact         |         | boolean   |
| cpulevel        | %       | integer   |
| current         | A       | float     |
| dewpoint        | C, F    | float     |
| distance        | m, yd, ft, in | float |
| dimmer          | %       | integer   |
| doorwindowsensor|         | boolean   |
| duration        | sec     | float     |
| electricfield   | V/m     | float     | Static electric field in volt per meter
| energy          | KWh     | float     |
| errors          |         | integer   |
| heatindex       | C, F    | float     | Apparent temperature (humiture) based on air temperature and  relative humidity. Typically used when higher than the air temperature. At 20% relative humidity the heatindex is equal to the temperature.
| heading         | Degrees, Radians | float |
| hue             |         |           |
| humidex         | C       | float     | Humidity temperature index (feels like temperature) derived from    dewpoint, in degrees Celcius 
| humidity        | %       | float     | Relative humidity in %
| image           | jpeg, png | bytes   | Image in jpeg or png format
| latency         | sec, msec | float   |
| location        |         | List of floats | [latitude, longitude, elevation]
| locked          |         | boolean   | Door lock status 
| luminance       | cd/m2, lux | float  | Amount of light in candela/m2 or in lux
| magneticfield   | T, mT, uT, nT, G(auss) | float | Static magnetic field in Tesla or Gauss
| motion          |         | boolean   | Motion detection status |
| power           | W       | float     | Power consumption in watts |
| pushbutton      |         | boolean   | Momentary pushbutton |
| signalstrength  | dBm     | float     |
| speed           | m/s, Km/h, mph | float |
| switch          |         | boolean   |
| temperature     | C, F    | float     |
| ultraviolet     | UV      | float     | Radiation with wavelength from 10 nm to 400 nm
| voltage         | V(olt)  | float     |
| waterlevel      | m(eters), ft, in | float |
| wavelength      | m       | float     |
| weight          | Kg, Lbs | float     |
| windchill       | C, F    | float     | Apparent temperature based on the air temperature and wind speed,    when lower than the air temperature.
| windspeed       | m/s, km/h, mph | float | 



# Appendix B: Implementation Libraries

### GoLang

-   [[https://golang.org/pkg/crypto/rsa/\#GenerateKey]{.underline}](https://golang.org/pkg/crypto/rsa/#GenerateKey)

<!-- -->

-   [[https://golang.org/pkg/crypto/rsa/\#SignPSS]{.underline}](https://golang.org/pkg/crypto/rsa/#SignPSS)

-   [[https://golang.org/pkg/crypto/rsa/\#VerifyPSS]{.underline}](https://golang.org/pkg/crypto/rsa/#VerifyPSS)

Command Line
------------

This [[blogpost describes a method of signing data with
ssh-agent]{.underline}](http://blog.oddbit.com/post/2011-05-09-signing-data-with-ssh-agent/).

Openssh has a tool that generates a public and private key pair as
follows:

*ssh-keygen -b 2048 -t rsa -N "" -f myzone.key*

This generates a file "myzone.key" with the private key and
"myzone.key.pub" with the public key. The passphrase is left empty.

Javascript
----------

-   [[https://github.com/kjur/jsrsasign]{.underline}](https://github.com/kjur/jsrsasign)

-   [[https://github.com/rzcoder/node-rsa]{.underline}](https://github.com/rzcoder/node-rsa)

Python
------

Paramiko

import hashlib, paramiko.agent

data\_sha1 = haslib.sha1(recordToPublish).digest()

agent = paramiko.agent.Agent()

key = agent.keys\[0\]

signature = key.sign\_ssh\_data(None, data\_sha1)
