# MyZone Introduction

Myzone is an information exchange convention for IoT devices, services and other information producers and consumers.


## Mission statement

Assist inhabitants of a zone by providing relevant information and control. Information from IoT devices and other sources is collected, processed and presented. Collected information can be shared with other
zones. Information is presented through desktop and mobile and wearable devices based on the consumer's situation.

## Information and Control

Information can come from many sources like sensors as well as services that derive new information from existing information. This includes any data captured by IoT devices, camera images, data from the internet, user input, as well as services that generate new information like analytics and machine learning.

Control of devices or services is a form of information, one that is provided by users and feeds back into the nodes. 

## Addressing of Information

Anything that produces information is called a 'node', including people. Nodes are published by a publisher and reside in a zone. Combined these form the address of the node: the zone ID / the publisher ID / the node ID. Nodes have inputs and outputs that have their own address and whose type and instance are appended to the node address. The separation of the address segments depends on the publication medium. On MQTT this is a '/'.

## Zones
A zone is a physical or virtual area where information resides and can be presented. A zone can be a home, a street, a city, or a virtual area like a sensor network or a game world. All consumers in a zone have access to information published in that zone. Each zone has a unique identifier. A zone ID does not have to be globally unique but must be unqiue between zones it shares information with.

Information that is collected within a zone can be shared with other zones of choice. External zones can subscribe to information that is made available to them. Sharing takes place through a secure connection
between zones.

For example, a water level sensor provides water levels to a city monitoring zone. A service within the monitoring zone interprets the water levels from multiple sensors and determines the risk level for flooding. This risk level information is shared with the city's community zone and available to residents and visitors of the city website.

A virtual game uses zones for its street map that are bridged to street zones in the real world. The number of people in the real world is reflected in the zone of the game world; An alarm triggered in the real world shows up in the game world; A message sent in the game world shows up in the real world. Once support for zones is available in the game the possibilities are endless.

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

Once information updates are accepted it is presented to the zone inhabitant through the available presentation device. This can simply be shown on all devices associated with the inhabitant, or the device currently in use.

Presented information has a life span. Stale information that has expired should be removed from presentation, depending on the type of information. This can be a personal preference.

Notification and presentation of notifications are provided by services that combine information and decide when and where to pass the information to the consumer. MyZone provides the infrastructure and ease of interoperability to create reusable building blocks for these features.

## Discovery and Configuration

Publishers of information also provide discovery and configuration metadata of the information. The discovery metadata describes the type of information, its publisher, and its configuration. Standardization of discovery and configuration allows data and information to be managed from a single user interface regardless of the various technologies involved.

Changing configuration and controlling inputs can be limited to specific users as identified by the signature contained in the configuration and control messages.

## Technology

MyZone is technology agnostic. It is a convention that describes the information format and exchange, discovery, configuration and zoning, irrespective of the technology used to implement it. Use of different technologies will serve to improve further integration and makes it easier to expand the information network.

A reference implementation is provided, written in the golang and typescript languages, using the MQTT service bus for publishing information, discovery and configuration in the JSON format.

## Data Format

The information exchange rules must be followed by all implementations of the convention. A JSON based encoding of the data is recommended. Other encodings such as XML or whatever becomes popular tomorrow can also be used.

The primary requirement is that the information fields described are preserved and all information publishers within a zone use the same format. A zone speaks only one data format language.

Future proofing can be achieved by using different zones for different data formats and using a bridge service to share between zones. The bridge maps between old and new data format. This allows for incremental improvements while maintaining interoperability.

## Versioning

The convention is version using semantic versioning in the form v1.2 where 1 is the major version, 2 is the minor backwards compatible version. 

Future minor version upgrades of this convention must remain backwards compatible. New fields can be added as long as they remain optional. Implementations must accept and ignore unknown fields.

A major version upgrade of this convention is not required to be backwards compatible but **must** be able to co-exists on the same bus.

Publishers include the version of the MyZone convention when publishing their node. See 'discovery' for more information.

# Terminology

| Terminology | Description |
| ----------- |:------------|
| Account     | The account used to connect a publisher to an message bus |
| Authentication| Method used to identify the publisher and subscriber with the message bus |
| Configuration | Configuration of the node configuration
| Discovery   | Description of nodes, their inputs and outputs
| Information | Anything that can be published by a producer. This can be sensor data, images, 
| message bus | A transport for publication of information and control. Information is published by a node onto a message bus. Consumers subscribe to information they are interested in.
| Node        | A node is a device or service that provides information and accepts control input. Information from this node can be published by the node itself or published by a (publisher) service that knows how to access the node. 
| Node Input  | Input to control the node, for example a switch.
| Node Output | Node Information is provided through outputs. For example, the current temperature.
| Publisher   | A service that is responsible for publishing node information and handle configuration updates and control inputs. Publishers are nodes. Publishers can sign their publications to provide source verification.
| Subscriber  | Consumer of data or information
| Zone        | An area in which information is shared between inhabitants



## Zone Overview

A Zone combines the information from its publishers into an message bus. Consumers of this information can subscribe to the information and get notified of updates. Various bus implementations are available. The reference implementation uses MQTT.

The Zone Bridge is a service that exchanges shared information with other zones.

\<image\>

A publisher consists of four parts, discovery, configuration, inputs and outputs.

\<image\>

The MyZone convention describes the data exchange format between zone publishers and consumers of information.

Publishers make information available on addresses, while consumers subscribe to these addresses to receive the information. The addressing, structure and security of this information is defined as part of this convention.

Publishers are responsible for:
1.  Publishing output information 

    This is mandatory. Every publisher must as a minimum publish their information on the value address (see in/output addresses).

2.  Handling requests to update inputs. 
   
    This is only for nodes that have inputs. Publishers can implement constraints that only trusted users can update the inputs.

3.  Publishing discovery information
   
    ... for available nodes, node configuration, node inputs, node outputs. This is optional and intended for environments where the computing power is available.

4.  Update node configuration.

    This is optional and intended for environments where the computing power is available. Publishers can implement constraints that only trusted users can update the inputs.

All publishers in a zone must use the same data interchange format of the published records. The recommended format is JSON. Other formats such as BSON or XML can be used as long as all publishers of the zone use the same format. Zone Bridges always exchange information in JSON.

# Inputs and Outputs

Information flows between nodes to consumers via publishers. Node output information is published as and control is handled via publisher inputs.

## In/Output Addresses

The addresses used to publish outputs and control inputs consist of segments. The address segments include the zone, the publisher of the information, the node whose information is published or controlled, the type of information, and the instance of the in- or output.

Segment names consist of alphanumeric, hyphen (-), and underscore (\_) characters. Reserved keywords start with a dollar ($) character. Other characters are not recommended to allow for various
publish/subscribe technologies like MQTT, REST or other.

Segment names are separated with a separator token. The token used depends on the communication bus used. By default this is the forward slash (/). Other tokens or methods of separating address segments can be used, but they must be consistent within a zone. 

If the message payload contain a reference to a node, input or output then only the segments must be used, not the separator tokens: 

Address Example for MQTT/REST:

* **{zone} / {publisher} / {node} / {type} / {instance} / $value**
* **{zone} / {publisher} / {node} / {type} / {instance} / $latest**
* **{zone} / {publisher} / {node} / {type} / {instance} / $24hours**
* **{zone} / {publisher} / {node} / {type} / {instance} / $set**


|Address segment |Description|
|----------------|-----------|
|{zone} | The zone in which publishing takes place. 
| {publisher} | The service that is publishing the information. A publisher provides its identity when publishing its discovery. The publisher Id is unique within its zone.
| {node} | The node that owns the input or output. This is a device identifier or a service identifier and unique within a publisher.
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

| Field        | Data Type | Required     | Description
|--------------|-----------|------------- |------------
| zone         | string    | **required** | The zone in which the node lives.
| publisher    | string    | **required** | The service that is publishing the information. 
| node         | string    | **required** | The node whose in/output is discovered. 
| type         | string    | **required** | The output type of this node
| instance     | string    | **required** | The input instance of this node
| signature    | string    | optional     | Signature of this record, signed by the publisher
| timeStamp    | string    | **required** | ISO8601 "YYYY-MM-DDTHH:MM:SS.sssTZ"
| unit         | string    | optional     | unit of value type
| value        | string    | **required** | value in string format


JSON example of a '$latest' publication:
```
zone-1/openzwave/6/temperature/0/$latest:
{
  "zone": "zone-1",
  "publisher": "openzwave",
  "node": "5",
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

| Field        | Data Type | Required     | Description |
| ----------   | --------  | -----------  | ------------ |
| zone         | string    | **required** | The zone in which the node lives.
| publisher    | string    | **required** | The service that is publishing the information. 
| node         | string    | **required** | The node whose in/output is discovered. 
| type         | string    | **required** | The output type of this node
| instance     | string    | **required** | The input instance of this node
| history      | list      | **required** | eg: [{"timeStamp": "YYYY-MM-DDTHH:MM:SS.sssTZ","value": string}, ...] |
|| timeStamp   | string    | ISO8601 "YYYY-MM-DDTHH:MM:SS.sssTZ" 
|| value       | string    | Value in string format using the node's unit
| signature    | string    | optional     | Signature of this record, signed by the producer |
| timeStamp    | string    | **required** | timestamp this message was created


A JSON example:
```
zone-1/openzwave/6/temperature/0/$24hours:
{
  "zone" : "zone-1",
  "publisher": "openzwave",
  "node": "5",
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

| Field        | Data Type | Required      | Description
|------------- |-----------|----------     |------ 
| zone         | string    | **required**  | The zone in which the node lives.
| publisher    | string    | **required**  | The publisher
| node         | string    | **required**  | The node 
| type         | string    | **required**  | The input type to set
| instance     | string    | **required**  | The input instance to set 
| timeStamp    | string    | **required**  | Time this request was created, in ISO8601 format, eg YYYY-MM-DDTHH:MM:SS.sssTZ. The timezone is the local timezone where the value was published. If a request was received with a newer timestamp, up to the current time, then this request is ignored.
| sender       | Address   | **required** | Address of the node representing the sender requesting to update the input
| signature    | string    | optional     | Signature of this record, signed by the sender that wants to set the input
| value        | string    | **required** | The input value to set


A JSON example:
```
zone-1/openzwave/6/switch/0/$set:
{
  "zone" : "zone-1",
  "publisher": "openzwave",
  "node": "6",
  "type": "switch",
  "instance": "0",
  "sender": {
    "zone" : "zone-1",
    "publisher": "dashboard",
    "node": "bob",
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


# Discovery 

Discovery describes the node, inputs and outputs. These are usually published by the same publisher that is responsible for publishing the output values and subscribing to input control values. Discovery messages can be signed by its publisher to verify its authenticity.

Publishing of node, input and output discovery is optional but highly recommended. It enables auto discovery,configuration management and information verification. For very resource restricted devices it can be omitted however.

Publishers that publish node discovery must also publish a node that represents themselves. The publisher's node id must be **\$publisher**. Publishers that have their own sensors can choose to publish the inputs
and outputs under the **\$publisher** node ID, or publish two records, one for the **\$publisher** and one for the node with the inputs and outputs.

## Node Discovery
Node discovery publishes the devices and services that are producers of information. The discovery shares the attributes of the device such as its name, make and model, and its configuration. Node discovery is intended for presenting a list of available devices and services and manage their configuration. It does not have to be shared with other zones in order to use their outputs.

The addresses used to publish node discovery consists of segments that describe the zone, publisher of the information, and the node being discovered. 

Each segment consists of alphanumeric, hyphen (-), underscore (_), or dollar (\$) characters. Other characters are not recommended to allow for various publish/subscribe mediums (like MQTT, REST). The '$' prefix is for reserved words.

MQTT/REST node discovery address:
  > **{zone} / {publisher} / {node} / $discover**

### Node discovery address structure

|Address segment| Description |
| ------------- | ----------- |
| {zone}      | The zone in which the node lives
| {publisher} | The service that is publishing the information. A publisher provides its identity when publishing a node discovery. The publisher Id is unique within its zone. 
| {node}      | The node that is discovered. This is a device or a service identifier and unique within a publisher. Two special nodes are defined: “$publisher” is a service node that publishes. “$gateway” represents the device that acts as a gateway to one or more nodes. For example a zwave controller.
| $discover     | Keyword for node discovery. 


### Node Discovery Payload

The discovery payload describes in detail the node and its configuration. The objective is for the node to be
sufficiently described so consumers can identify and configure it without further information.

| Field         | Data Type  | Required     | Description
| -----------   |----------- |----------    |------------
| zone          | string     | **required** | The zone in which the node lives.
| publisher     | string     | **required** | The service that is publishing the information. 
| node          | string     | **required** | The node that is discovered. 
| attr          | dictionary | **required** | Attributes describing the node. Collection of key-value string pairs that describe the node. The list of predefined attribute keys are part of the convention. See appendix.
| config        | List of **Configuration Records** | optional | Node configuration, if any exist. Set of configuration objects that describe the configuration options. These can be modified with a ‘$configure’ message.
| signature     | string | optional | Signature of this record signed by the publisher
| timeStamp     | string | optional | Time the record is created

### Configuration Record

The configuration record is used in both node configuration and input/output configuration. Each configuration attribute is described in a record as follows:

| Field    | Data Type| Required | Description
|--------  |----------|----------|------------
| name     | string   | **required** | Name of the configuration. This has to be unique within the list of configuration records.
| datatype | enum     | optional| Type of value. Used to determine the editor to use for the value. One of: bool, enum, float, int, string. Default is ‘string’
| default  | string   | optional| Default value for this configuration in string format
| description| string | optional | Description of the configuration for human use
| enum     | \[strings] | optional* | List of valid enum values as strings. Required when datatype is enum
| max      | float    | optional | Optional maximum value for numeric data
| min      | float    | optional | Optional minimum value for numeric data
| secret   | bool     | optional | Optional flag that the configuration value is secret and will be left empty. When a secret configuration is set in $configure, the value is encrypted with the publisher node public key. 
| value    | string   | **required**| The current configuration value in string format. If empty, the default value is used.

### Predefined Node Attributes

| Key          | Value Description 
|--------------|------------- 
| certificate  | A certificate from a trusted source, like Lets Encrypt. It is included by publishers to provide consumers a means to verify their identity
| convention   | Publishers include the version of the myzone convention. Eg v1.0
| localip      | IP address of the node, for nodes that are publishers themselves
| location     | String with "latitude, longitude" of device location
| mac          | Node MAC address for nodes that are publishers
| manufacturer | Node make or manufacturer
| model        | Node model
| myzone       | Version of the convention this publisher uses. This attribute must be present when a publisher publishes its own node
| publicKey    | Publisher's public key used verify the signature provided with publications of information. Only accept public keys from publishers that are verified through their certificate or other means
| type         | Type of node. Eg, multisensor, binary switch, See the nodeTypes list for predefined values
| version      | Hardware or firmware version




Example payload for node discovery in JSON format:
```
zone1/openzwave/5/$discover:
{
   "zone": "zone1",
   "publisher": "openzwave",
   "node": "5",
   
   "attr": {
     "make": "AeoTec",
     "type": "multisensor",
      ...
   },
   "config": {
      "name": {
          "datatype": string,
          "description": “Friendly name of the node",
          "value": “barn multisensor”,
      },
      …
   },
   "timestamp": "2020-01-20T23:33:44.999PST",
   "signature": "...",
}
```

## Input/Output Discovery

Inputs and outputs discovery are published separately from the node to allow control over which ones are shared with other zones.

MQTT/REST address formats:
  >Input discovery:  **{zone} / {publisher} / {node} / {inputType} / {instance} / $discover**

  >Output discovery: **{zone} / {publisher} / {node} / {outputType} / {instance} / $discover**

### Input/Output discovery address structure

| Address segment  | Description |
| :--------------- | ----------- |
| {zone}      | The zone in which the node lives
| {publisher} | The service that is publishing the information. A publisher provides its identity when publishing a node discovery. The publisher Id is unique within its zone. 
| {node}      | The node whose input or output is discovered. This is a device or a service identifier and unique within a publisher.
| {inputType} | Type identifier of the input. The list of predefined types is part of this convention. 
| {outputType} | Type identifier of the output. The list of predefined types is part of this convention. 
| {instance} | The instance of the input or output on the node. If only a single instance exists the convention is to use 0 unless a name is used to provide more meaning.
| $discover | Keyword for node discovery. 


MQTT example:

The discovery of a single temperature sensor on node '5', published by a service named 'openzwave' is published on an MQTT bus on topic:
  > **myzone/openzwave/5/temperature/0/$discover**



### Input/Output Discovery Payload

| Field       | Data Type| Required     | Description
|------------ |----------|----------    |------------
| zone      | string   | **required** | The zone in which the node lives.
| publisher | string   | **required** | The service that is publishing the information. 
| node      | string   | **required** | The node whose in/output is discovered. 
| ioType      | string   | **required** | Type identifier of the output. The list of predefined types is part of this convention. 
| instance    | string   | **required** | The instance of the input or output on the node. Use 0 if only a single instance exists and names are not used.
| config      | List of **Configuration Records**|optional|List of Configuration Records that describe in/output configuration. Only used when an input or output has their own configuration. See Node configuration record above for the definition
| datatype  | string         | optional      | Value datatype. One of boolean, enum, float, integer, jpeg, png, string, raw. Default is "string".
| default   | string    | optional | Default output value
| description | string | optional | Description of the in/output for humans
| enum      | list      | optional* | List of possible values. Required when datatype is enum
| **instance** | string         | **required**  | The output instance when multiple instances of the same type
| max       | number    | optional | Maximum possible in/output value
| min       | number    | optional | Minimum possible in/output value
| signature | string    | optional | Signature of this record signed by the publisher
| timeStamp | string    | optional | Time the reocrd is created
| **type**  | string    | **required**  | Type of input/output. See list below
| unit      | string    | optional | The unit of the data type
| **value** | string    | **required**  | The input or output value at time of discovery


Example payload for output discovery in JSON format:
```
zone1/openzwave/5/temperature/0/$discover:
{
   "zone": "zone1",
   "publisher": "openzwave",
   "node": "5",
   "datatype": "float",
   "instance": "0",
   "signature": "...",
   "type": "temperature",
   "timestamp": "2020-01-20T23:33:44.999PST",
   "unit": "C",
   "value": "20.5",
}
```

# Node Configuration

Nodes that can be configured contain a list of configuration records
described in the node discovery. The configuration value can be updated
with a configure command as per below.

The configuration of a node can be updated by a consumer by publishing
on the '$configure' address. The node publisher listens to this request
and processes it after validation.

Only authorized users can modify the configuration of a node.

## Configure Address

> {zone}/{publisher}/{node}/$configure

## Configure Record Payload

|Field 		     |type 		     |required 		     |Description
|--------------|-------------|-----------------|-----------
| zone       | string      | **required**    | The zone in which the node lives.
| publisher  | string      | **required**    | The service that is publishing the information. 
| node       | string      | **required**    | The node whose in/output is discovered. 
| config 	     |Dictionary   | **required**    | key-value pairs for configuration to update { key: value, …}
| sender   		 |Address      | optional 	     | Address of the sender submitting the request. This is the zone/publisher/node of the consumer.
| signature		 |string	     | optional 	     | Signature of the sender of the configuration request. The receiving node publisher verifies if the sender address has permission to modify the configuration of the node before verifying and applying the update.
| timeStamp		 |string	     | **required**    | Time this request was created, in ISO8601 format


# Node Status

The availability status of a node is published by its publisher when the availability changes or errors are encountered. Publishing of a node status is required.

## Status Address

**{zone} / {publisher} / {node} / $status**

| Address segment | Description
|-----------------|--------------
| {zone}        | The zone in which discovery takes place. 
| {publisher}   | The publisher of the node discovery which is handling the configuration update for that node.
| {node}        | The node whose configuration is updated. 
| $status         | Keyword for node status. Published when the availability of a node changes or new errors are reported. It is published by the publisher of the node.



## Status Payload

| Field 		      | type 		    | required 		    | Description
|-----------------|----------   |------------     |-----------
| zone          | string      | **required**    | The zone in which the node lives.
| publisher     | string      | **required**    | The service that is publishing the information. 
| node          | string      | **required**    | The node whose in/output is discovered. 
| errorCount      | integer     | optional        | Nr of errors since startup
| errorMessage    | string 	    | optional 	      | Last reported error message
| errorTime       | string 	    | optional		    | Timestamp of last error message in ISO8601 format
| interval        | integer     | **required**    | Maximum interval of status updates in seconds. If no updated status is received after this interval, the node is considered to be lost. 
| lastSeen        | string	    | **required**	  | Timestamp in ISO8601 format that the publisher communicated with the node.
| signature       | string      | optional        | Publisher signature of this record.
| status          | enum        | **required**    | The node availability status. See below for valid values
| timeStamp       | string      | **required**    | Time the status was last updated, in ISO8601 format

Status values:
* ready: node is ready to perform a task
* asleep: node is a periodically listening and updates to the node can take a while before being processed.
* error: node is in an error state and needs servicing
* lost: communication with the node is lost


# Trust & Digital Signatures

Note: this section needs further review and improvements.

Trust is essential to information exchange between publishers and consumers, especially when the producer and consumer don't know each other directly. In this case trust means that the consumer can be sure
that the publisher is who he claims to be. This is achieved by including a publisher signature in every publication. The consumer can verify that the signature is valid and trust the information. 

To this purpose, publishers include a [[digital signature]](https://en.wikipedia.org/wiki/Digital_signature)
in their node publication that lets the consumer verify the records originate from the publisher. [[This
tutorial]](https://www.tutorialspoint.com/cryptography/cryptography_digital_signatures.htm)
explains it with a picture. This convention uses **RSA-PSS** as the preferred digital signatures. This is used in OpenSSL and can be used with 'Lets Encrypt' (Needs verification).

-   [[RSA-PSS]](https://en.wikipedia.org/wiki/Probabilistic_signature_scheme)
    > part of [[PKCS#1 v2.1](https://en.wikipedia.org/wiki/PKCS_1) and used in OpenSSL

As security is constantly evolving, different schemes can be supported in the future.

## Zone Certificate Authority Service - ZCAS

Zones can be secured using a zone certificate authority service (ZCAS). The ZCAS that is responsible for verifying and issuing keys and certificates. It can work on behalf of a global Trusted Certificate Authority.

Since the network topology is separate from the zone topology, the ZCAS and publishers can be on different networks and behind firewalls. Communication between ZCAS and publisher therefore uses the zone message bus.

Any bus communication is considered unsafe and must be protected against man-in-the-middle and spoofing attacks. All publications must be considered untrusted unless it is correctly signed. This includes publications by the ZCAS itself. 

The ZCAS has the reserved publisher ID of '\$zcas' with a reserved node ID of '\$zcas'. It publishes its own certificate (just like any other publisher) with its node on address '{zone}/\$zcas/\$zcas'

A ZCAS can be registered with a global Trusted Certificate Authority (TCA) and creates certificates that are chained to the TCA. By default this uses 'Lets Encrypt' but this can be replaced by other public CAs. Use of a TCA is optional for local-only zones but required when briding between zones. The domain name used for the TCA is '{zone}.myzone.world'. zone has to be globally unique. 

The zone 'myzone' is reserved for local-only zones. In this case the ZCAS generates its own certificate and is considered the highest authority.

**Joining the Secure Zone**

A publisher has to join the secure zone before it is issued a valid certificate by the ZCAS.

The publisher has an initial public/private key pair which can be self generated or installed. This public key is included when the publisher publishes its own node but it is considered unverified as it has no valid certificate. The ZCAS needs to be told that the publisher with public key X is indeed the publisher with the ID it claims to have. This is the process of joining the zone.

The method used to join the zone can vary based on the situation. A secure method is to download the public key of the publisher with its publisher ID on a USB stick and use the USB stick to upload it to the ZCAS. The ZCAS now trusts the public key to belong to the publisher and issues a new set of keys with a valid certificate (see below). Note that physical security of the ZCAS is always required.

Another method to join a publisher to the security zone is to have the ZCAS present a web page where a (human) administrator can log in and upload the publisher's public key with its publisher ID. This moves the trust to that of the administrator login. Two factor authentication should be used.

**Key and Certificate Renewal**

A publisher can not request a new certificate. Instead, it is issued new keys and certificate by the ZCAS automatically when its certificate has less than half its life left. This is triggered when a publisher publishes its own node information on address '{zone}/{publisher}/\$publisher' using a valid signature.

A publisher is issued its certificate on address: {zone}/\{publisher}/\$zcas. This contains an encrypted payload with the public key, private key and certificate. The payload is encrypted with the last known valid public key of this publisher. Only the publisher for whom it is intended can decypher it. The keys and certificate are valid for a restricted period of time. The default is 30 days. 

Once a publisher uses the newly issued key and certificate, ZCAS removes the old key from its records. This key can no longer be used to obtain a new key and certificate even if they were still valid. It is therefore important that the publisher persists the new key and certificate before publishing using the new keys.

This means that publishers subscribe to address '{zone}/{publisher}/\$zcas' while ZCAS subscribes to the publisher discovery address '{zone}/{publisher}/\$publisher'. It can also use a wildcard for publisher. 

**Verifying Signatures**

When a consumer receives node output, the message is signed by the node publisher. The consumer must verify the validity of the signature using the public key of the publisher. The verification takes place using the message content with the signature left blank and the publisher's public key.
If the verification fails the message is discarded. The number of discarded messages is tracked for each publisher and can be used to show an alert in a dashboard.

The public key of the publisher and its certificate are included with the publisher node discovery message. The public key is used to verify messages while the certificate is used to verify that the public key belongs to this publisher.

When receiving the publisher discovery message, the consumer must therefor verify that the publisher's certificate is valid by checking its signature with the public key of the ZCAS that issued it.

If the message is from another zone, then the certificate must also be verified against the public key of the TCA.

**Policing Publications**

The ZCAS has a secondary function to monitor node publications and verify that the certificate and public key are valid. If an invalid publication is detected then a notice is send to the administrator.

# Zone Bridging

The purpose of a Zone Bridge is to export information from one zone to the message bus of another zone. This is only applicable when the zones are on a different message bus. In the case where multiple zones are publishing on the same bus there is no need for a bridge as consumers can just subscribe to the zone using the bus they are already on.

The bridge subscribes to nodes to be exported and forwards (re-publishes) information onto the message bus of the bridged zone. The original publication and the address remains unchanged. 

The bridge service has facility to select which nodes, inputs or outputs are bridged. 

**Bridge Service**
A zone can have one or multiple bridge services. A bridge forwards selected publications to another zone without modifying it. The bridge service is a publisher and has a node of $bridge. The publisher differs per bridge instance. 

A bridge connects to a single remote bus. The connection attributes are part of the bridge configuration.

Bridge configuration
| Field 		   | value type   | value 		     | Description
|------------- |----------    |------------  |-----------
| host         | string       | **required** | IP address or hostname of the remote bus
| port         | integer      | optional     | port to connect to. Default is determined by protocol
| protocol     | enum ["MQTT", "REST"] | optional  | Protocol to use, MQTT (default), REST API, ...
| format       | enum ["JSON", "XML"]  | optional  | Publishing format used on the external bus. Default is JSON. Only include this if a different format is needed. This will invalidate the signature.
| clientId     | string       | **required** | ID of the client that is connecting
| loginId      | string       | optional     | Login identifier
| credentials  | string       | optional     | Password to connect


**Add/remove nodes to export**
To forward nodes through the bridge, use the pushbuttons for $forward and $remove.

MQTT Example:
* {zone} / {publisher} / \$bridge / $pushbutton / $forward
* {zone} / {publisher} / \$bridge / $pushbutton / $remove

Address segments:

| Field 		     | Description
|--------------- |------------|
| {zone}       | The zone the bridge lives in.
| {publisher}  | The publisher of the bridge.
| $bridge        | Reserved ID of a bridge node.
| $pushbutton    | Input type 
| $forward       | button instance to add a forward. Payload contains the address record.
| $remove        | button instance to remove a forward. Payload is the address record


Payload:
| Field 		   | type 		    |value 		     | Description
|------------- |----------    |------------  |-----------
| zone       | string       | required     | The zone the node lives in
| publisher  | string       | required     | The publisher of the node to export
| node       | string       | required     | The node to export
| type         | string       | optional     | Output type to forward in when only forwarding an output
| instance     | string       | optional     | Output instance to forward in when only forwarding an output
| fwdValue     | boolean      | optional     | Forward the output $value publication(s), default=true
| fwdLatest    | boolean      | optional     | Forward the output $latest publication(s), default=true
| fwdHistory   | boolean      | optional     | Forward the output $history publication(s), default=true
| fwdDiscovery | boolean      | optional     | Forward the node/output $discovery publication, default=true
| fwdStatus    | boolean      | optional     | Forward the node $status publication, default=true




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
