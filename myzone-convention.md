Information Exchange for IoT devices and Services
=================================================

Myzone is an information exchange standard for IoT devices, services and other information producers and consumers.

# Current status

2020-02-05 Draft

# Author

* H. Spaay

# TOC

[[TOC]]

# Introduction

As connected devices become more and more prevalent, so have the problems surrounding them. These problems fall into multiple categories: 

## Interoperability
   
The use of information produced by these devices is becoming more and more difficult because of the plethoria of different protocol and data formats these devices use. This is apparent for home automation solutions such as OpenHAB and Home Assistant that each implement hundreds of bindings to talk to different devices and services. Each solution has to reimplement these bindings. This implementation then has to be adjusted to different platforms, eg Linux, Windows and MacOS, which adds even more work.

Without a public standard it is unavoidable that manufacturers of IoT devices choose their own protocols. It is in everyone's interest to provide a common open standard that enables an open information interchange so that bindings only have to be implemented once. 

This standard defines the messages for information exchange.

## Discovery

Discovery of connected IoT devices depends on the technology used. There is no standard that describes what and how discovery information is made available to consumers. Application developers often implement solutions specific to their application and the devices that are supported. To facilitate information exchange it must be possible to discover the information that is available independent of the technology used.

This standard defines the process and messaging for discovery of devices and services.

## Configuration

Configuration of IoT devices is often done through a web portal of some sort. These web portals are not always as secure as they should be. They often require a login name and password and lack 2 factor authentication. Passwords are easily reused. Backdoors are sometimes left active. Overall security is lacking.
   
Configuration is not always suited for centralized management by application services. For example, to configure all temperature sensors to report in Celcius the user has to login to the device management portal(s), find the sensor and find the configuration for this. This is difficult to automate.

This standard defines the process and messaging for remote configuration of devices and services.

Nodes that can be configured contain a list of configuration records described in the node discovery. The configuration value can be updated with a configure command as per below.

The configuration of a node can be updated by a consumer by publishing on the '$configure' address. The node publisher listens to this request and processes it after validation.

Only authorized users can modify the configuration of a node.

## Security
   
Security is a major concern with IoT devices. Problems exist in several areas:

  1. It is difficult to design devices for secure access from the internet. The existance of large botnets from hacked computers and devices show how severe this problem is. Good security is hard and each vendor has to reinvent this wheel. This is not likely to change any time soon.
  
  2. Commercial devices that connect to a service provider share personal information without the user understanding what this information is, and without having control on how it is used. While regulations like Europe's [GDPR](https://en.wikipedia.org/wiki/General_Data_Protection_Regulation) attempt to address this ... somewhat, reports of data misuse and breaches remain all too frequent.

  3. There is no easy and secure way to self-serve information over the internet. Often the only option is to trust a 3rd party service provider in this, which leads to the previous two problems. In addition the monthly recurring cost might be out of reach for many users.

This standard defines the security aspects build into the specification.

# Terminology

| Terminology   | Description |
| -----------   |:------------|
| Account       | The account used to connect a publisher to an message bus |
| Address       | Address of the node consisting of zone, publisher and node identifier. Optionally it can include the input or output type and instance.|
| Authentication| Method used to identify the publisher and subscriber with the message bus |
| Bridge        | The service that publishes subscribed information into a different zone. |
| Configuration | Configuration of the node configuration|
| Data          | The term 'data' is used for raw data collected before it is published. Once it is published it is considered information.|
| Discovery     | Description of nodes, their inputs and outputs|
| Information   | Anything that is published by a producer. This can be sensor data, images, discovery, etc|
| Message Bus   | A publish and subscribe capable transport for publication of information. Information is |published by a node onto a message bus. Consumers subscribe to information they are interested in use the information address. |
| Node          | A node is a device or service that provides information and accepts control input. Information from this node can be published by the node itself or published by a (publisher) service that knows how to access the node. |
| Node Input    | Input to control the node, for example a switch.|
| Node Output   | Node Information is published using outputs. For example, the current temperature.|
| Publisher     | A service that is responsible for publishing node information on the message bus and handle configuration updates and control inputs. Publishers are nodes. Publishers sign their publications to provide source verification.|
| Subscriber    | Consumer of information that uses node address to subscribe to information from that node.|
| ZCAS          | Zone Certificate Authority Service. This service manages keys and certificates of zone members |
| ZSS           | Zone Security Service. Monitors publications in a zone and watches for intrusions.|
| Zone          | An area in which information is shared between members of a zone. |


# Versioning

The standard uses semantic versioning in the form v{major}.{minor}[-RC{N}]. Where RC-{N} is only used for release candidates of the final version. 

Future minor version upgrades of this convention must remain backwards compatible. New fields can be added but must be optional. Implementations must accept and ignore unknown fields.

A major version upgrade of this convention is not required to be backwards compatible but **must** be able to co-exists on the same bus. Implementations must ignore messages with a higher major version.

Publishers include their version of the MyZone standard when publishing their node. See 'discovery' for more information.

# Technology Agnostic

MyZone is technology agnostic. It is a standard that describes the information format and exchange for discovery, configuration, inputs and outputs, irrespective of the technology used to implement it. Use of different technologies will actually serve to further improve interoperability with other information sources.

A reference implementation of a publisher is provided for the golang and python languages using the MQTT service bus.


# System Overview

![System Overview](./system-overview.png)


# Information Exchange

The problem of interoperability is addressed by defining an open standard for exchanging information with devices and services. IoT devices can implement this standard directly, or be encapsulated by an 'adapter' service that implements this standard. 
  
Devices such as 1-wire, zigbee, openzwave have adapters that convert from their respective protocols to this standard. These adapters can directly implement the MyZone standard and be reused by many different applications. Non-standard wireless devices, such as Philips Hue lightbulb just to name  one, must be isolated by the wireless access point and only be allowed to communicate to their MyZone adapter. 

## Message Addressing

All messages contain the address of the source and destination of the message. The source is the address of the publisher that publishes the message, the destination is the node, input or output, depending on the message command.

Each node has its own unique address. It consists of the zone, the publisher and the node. The inputs and outputs can be addressed by adding the type and instance of the input or output.

Adress segments consist of alphanumeric, hyphen (-), and underscore (\_) characters. Reserved keywords start with a dollar ($) character. The separator is the '/' character. 

> The address of a node has the form: {zone} / {publisher} / {node} 

> The address of an input and output: {zone} / {publisher} / {node} / {type} / {instance}

Some message bus systems might not support the '/' character as address or topic separator. In this case the message bus address use what is specific to the message bus implementation. The payload of the messages however must contain the '/' character in the address.

## Zones

Zones define the area in which information is shared amongst its members. A zone can be a home, a street, a city, or a virtual area like an industrial sensor network or even a game world. Each zone has a globally unique identifier, except for the local zone. 

All members of a zone have access to information published in that zone. The information is not available outside the zone unless intentionally shared. Publication in the zone is limited to members that have the publish permissions. Not surprisingly these are called 'publishers'.

A zone can be closed or open to consumers. An open zone allows any consumer to subscribe to publications in that zone without providing credentials. A closed zone requires consumers to provide valid credentials to connect to the message bus of that zone. Whether a zone is open or closed is determined by the configuration of the message bus for that zone.

A zone can share information with another zone using a bridge service (ZBS). See more information on bridges in the following paragraphs.

Some examples. A water level sensor provides water levels to a city's monitoring zone. A service within the monitoring zone interprets the water levels from multiple sensors and determines the risk level for flooding. A neighboring city shares its water levels with the city, again using a bridge service, to provide more data points in assessing the flood risk. This risk level information is bridged to the city's open community zone where it is available to residents and visitors of the city website. 

A virtual game uses zones for its street map that are bridged to street zones in the real world. The number of people in the real world is reflected in the zone of the game world; An alarm triggered in the real world shows up in the game world; A message sent in the game world shows up in the real world. Once support for zones is available in the game the possibilities are endless and any information can be made available in-game.

More examples are presented in the MyZone use-cases document.

## Nodes

A zone is populated by Nodes. Anything that produces information is a 'node' of the zone. There are some nodes that have special roles, eg publishers, bridges, security, and certificate authority services.

### Inputs and Outputs

Each node has inputs and/or outputs through which information passes. A publisher publishes the discovery, configuration and values of inputs and outputs separate from the node itself. A node can have many inputs and outputs that are connected to the node. Inputs and outputs cannot exist without a node.

Inputs have the same address as outputs. The difference is that inputs are addressed using the $set command suffix, while outputs have the $value or '$latest' command suffix.

Output discovery messages have a flag indicating if an input is available. 

When inputs are updated, the output is only updated after the input is successfully processed.

To provide timely update of values for new consumers, output messages are repeated after a specified interval. The default interval is 1 hour. On message buses that support 'retainment', output messages are also send as a retained message.


## Publishers

Publishers are nodes that publish the information. Only nodes that implement this standard are publishers. A Node can be its own publisher or use a separate publisher to publish its information. For example, a ZWave publisher can obtain sensor data from a ZWave nodes via a ZWave gateway, and publish information from the ZWave nodes that are connected to the gateway. The publisher, the gateway and each zwave device are represented by a node, all published by the same publisher.

Publishers must have credentials to connect to a zone's message bus before they can publish. To publish securely, a publisher must also have to joined the zone through the Zone Authentication Service (ZCAS). More on that later.

Publishers are responsible for:

1.  Signing and publishing output information 

    This is mandatory. Every publisher must as a minimum publish the output values.
    Publishing of output latest and history information is highly recommended but optional for constrained devices or networks.

2.  Handling requests to update inputs. 
   
    Nodes that have inputs must handle requests to update the input value. If the input is related to an output then the output is updated after the input is validated and processed. Publishers can implement constraints that only trusted users can update the inputs.

3.  Publishing discovery information
   
    Publishers publish discovery information of themselves, other nodes, their inputs and outputs. This is optional but highly recommended for environments where the computing power is available.

4.  Update node configuration.

    Nodes can receive requests to update their configuration. This is optional and intended for environments where the computing power is available. Publishers can implement constraints that only trusted users can update the inputs.

5.  Update security keys and certificates when instructed by the ZCAS. More on that in the security section.

### Discovery

Support for discovery lets consumers search for the information they are interested in. The standard defines the messages for discovery of devices, services, and their inputs and outputs.

Publishers of information also publish discovery messages for nodes, their inputs and outputs. The discovery data describes the type, attributes and configurations for nodes, inputs and outputs.

Standardization of discovery and configuration allows data and information to be managed centrally and automated, from a single interface regardless of the various technologies involved.

Discovery messages are signed by its publisher to support verification of its authenticity.

Node discovery publishes the devices and services that are producers of information. The discovery shares the attributes of the device such as its name, make and model, and its configuration. Node discovery is intended for presenting a list of available devices and services and manage their configuration. It does not have to be shared with other zones in order to use their outputs.

To provide timely discovery, discovery messages are repeated after a specified interval. The default is 24 hours. On message buses that support 'retainment', discovery messages are also send as a retained message.


### Configuration

Support for remote configuration lets administrators manage the devices and services sources without having to login to each device and service throught their web portals. The standard defines the messages for obtaining and updating the configuration of nodes by authorized users only.

Publishers of information provide information on node configuration and can accept commands to update this configuration.

Changing configuration and controlling inputs can be limited to specific users as identified by the signature contained in the configuration and input control messages.

## Zone Security

Securing a zone means ensuring that the information can be trusted, its source can be verified, and the information is only accessible to the members of that zone. 

This is achieved by including a publisher signature in every publication. The consumer can verify that the signature is valid and trust the information. To this purpose, publishers include a [[digital signature]](https://en.wikipedia.org/wiki/Digital_signature) in their node publication that lets the consumer verify the records originate from the publisher. 

This standard uses **RSA-PSS** as the preferred digital signatures. This is used in OpenSSL and can be used with 'Lets Encrypt' (Needs verification).

-   [[RSA-PSS]](https://en.wikipedia.org/wiki/Probabilistic_signature_scheme)
    > part of [[PKCS#1 v2.1](https://en.wikipedia.org/wiki/PKCS_1) and used in OpenSSL

As security is constantly evolving, different schemes can be supported in the future.

The keys and certificates needed for this are provided by the Zone Certificate Authority Services - ZCAS.

### Message Signing

When consuming information from external sources, trust in the validity of this information and the ability to identify its source is important. Faulty sensors or bad actors can generate information that is unreliable. It should be easy to identify the publisher and node that provide information.

With a few exceptions, all messages include the timestamp and source address of the node whose information is published. Combined with signing of the message it provides a means to validate that the message hasn't been tampered with.

A publisher **must** sign messages that it publishes. In cases where the signature cannot be verified the consumer has the option to ignore unverified sources.

For example, an alarm company receives a security alert. The alert signature is verified to be from a registered client and not from a bad actor creating a distraction. The node location and alert timestamp are used to identify where and when the alert was given.

### Key and Certificate Management with ZCAS

In order to sign a message, a publisher must have a set of keys. Publishers must be registered with the Zone Certificate Authority Service - ZCAS - to receive keys and a certificate. The keys are also be used to encrypt and decrypt messages at the zone level.

The ZCAS can work with a global Trusted Certificate Authority. This is required to verify signatures when sharing messages between zones.

Since the network topology is separate from the zone topology, the ZCAS and members of a zone can be on different networks and behind firewalls. Communication between ZCAS and zone members uses only the zone message bus and is completely asynchroneous. 

Bus communication should be considered unsafe and must be protected against man-in-the-middle and spoofing attacks. All publications must be considered untrusted unless it is correctly signed. This includes publications by the ZCAS itself.

The ZCAS has the reserved publisher ID of '\$zcas' with a reserved node ID of '\$zcas'. It publishes its own certificate (just like any other publisher) with its node on address '{zone}/\$zcas/\$zcas'. In order to verify signatures the consumer must check if node certificates are issued by the ZCAS using the ZCAS certificate.

A ZCAS can be registered with a global Trusted Certificate Authority (TCA) and creates certificates that are chained to the TCA. By default this uses 'Lets Encrypt' but this can be replaced by other public CAs. Use of a TCA is optional for local-only zones but required when briding between zones. The domain name used for the TCA is '{zone}.myzone.world'. zone has to be globally unique. <todo: how to ensure global uniqueness? UUID?, hash?, registration?>

The zone 'myzone' is reserved for local-only zones. In this case the ZCAS generates its own certificate and is considered the highest authority.

**Joining A Secure Zone**

A publisher has to join a secure zone before it is issued a valid certificate by the ZCAS. Subscribers do not need to join the zone to verify the signatures as the ZCAS node includes its public key and certificate in its node publication. 

To join a zone the publisher starts with an initial public/private key pair which can be self generated or pre-installed. This public key is included when the publisher publishes its own node but it is considered unverified as it has no valid certificate. The ZCAS needs to be told that the publisher with public key X is indeed the publisher with the ID it claims to have. Optionally additional credentials can be required such as location, contact email, phone, administrator name and address. This is the process of joining the zone.

The method used to join the zone can vary based on the situation. A secure method is to download the public key of the publisher with its publisher ID on a USB stick and use the USB stick to upload it to the ZCAS. The ZCAS now trusts the public key to belong to the publisher and issues a new set of keys with a valid certificate (see below). Note that physical security of the ZCAS is always required.

Another method to join a publisher to the security zone is to have the ZCAS present a web page where a (human) administrator can log in and upload the publisher's public key with its publisher ID. This moves the trust to that of the administrator login. Two factor authentication should be used. 

Either way, for now a human must verify the credentials and make the decision to accept a publisher joining the zone. 

**Key and Certificate Renewal**

A publisher can not request a new certificate. Instead, it is issued new keys and certificate by the ZCAS automatically when its certificate has less than half its life left. This is triggered when a publisher publishes its own node information using a valid signature.

A publisher is issued its certificate on address: {zone}/\{publisher}/\$zcas. This contains an encrypted payload with the public key, private key and certificate. The payload is encrypted with the last known valid public key of this publisher. Only the publisher for whom it is intended can decypher it. The keys and certificate are valid for a restricted period of time. The default is 30 days. 

Once a publisher uses the newly issued key and certificate, ZCAS removes the old key from its records. This key can no longer be used to obtain a new key and certificate even if they were still valid. It is therefore important that the publisher persists the new key and certificate before publishing using the new keys.

This means that publishers subscribe to address '{zone}/{publisher}/\$zcas' while ZCAS subscribes to the publisher discovery address '{zone}/{publisher}/\$publisher'. It can also use a wildcard for publisher. 

**Verifying Signatures**

When a consumer receives node output, the message is signed by the node publisher. The consumer must verify the validity of the signature using the public key of the publisher. The verification takes place using the message content with the signature left blank and the publisher's public key.
If the verification fails the message is discarded. The number of discarded messages is tracked for each publisher and can be used to show an alert in a dashboard.

The public key of the publisher and its certificate are included with the publisher node discovery message. The public key is used to verify messages while the certificate is used to verify that the public key belongs to this publisher.

When receiving the publisher discovery message, the consumer must therefor verify that the publisher's certificate is valid by checking its signature with the public key of the ZCAS that issued it.

If the message is from another zone, then the certificate must also be verified against the public key of the TCA.

## Zone Bridge Service - ZBS

Bridges are publisher nodes that share information with a remote zone. 

Information that is published within a zone can be shared with other zones. A 'bridge' service can connect to a remote zone and publish information from the local zone into the remote zone. It is configured with credentials to access to the remote zone message bus. A Zone bridge has a node in both zones, joined both zones, and can be discovered like any other node.

A zone can have zero, one or multiple bridge services. A bridge forwards selected publications to another zone without modifying it. The bridge service is a publisher and has a node of $bridge. The publisher id differs per bridge instance. 

A bridge connects to a single remote bus. The connection attributes are part of the bridge configuration.


# MyZone Messages

This section defines the message address, format and content.

## Message Addressing

All messages contain the address of the source or destination of the message. This is either the address of a node, its input, or output.

Each node has its own unique address consisting of segments for the zone, publisher and node identifiers. The inputs and outputs can be addressed by adding the type and instance of the input or output.

Segment names consist of alphanumeric, hyphen (-), and underscore (\_) characters. Reserved keywords start with a dollar ($) character. The separator is the '/' character. 

> The address of a node has the form: {zone} / {publisher} / {node} 

> The address of an input and output: {zone} / {publisher} / {node} / {type} / {instance}

Some message bus systems might not support the '/' character as a path separator. In this case the character can be replaced to what is appropriate for the message bus implementation. The content of the messages however must contain the '/' character in the address.


## Data Format

The information exchange rules must be followed by all implementations of the standard. Data must be JSON encoded to support information sharing between zones.

Message publications contain two fields, "message" and "signature". The signature is created by encrypting a hash of the message content using its private key. It can be verified with the publisher's public key.

The message fields described in the standard must be followed. Required field must be present in messages while optional fields can be left out. All producers and consumers must be able to handle messages with and without optional fields. Messages that are missing required fields must be discarded before being processed. 


## Information Exchange Messages

The address used to publish outputs and control inputs consist of segments followed by a 'command'. The address segments include the zone, the publisher of the information, the node whose information is published or controlled, the type of information, and the instance of the in- or output. The command indicate the purpose of the publication. 

Segment names are separated with a separator token '/'. 

Output publication addresses

* **{zone} / {publisher} / {node} / {type} / {instance} / $value**
* **{zone} / {publisher} / {node} / {type} / {instance} / $latest**
* **{zone} / {publisher} / {node} / {type} / {instance} / $history**
* **{zone} / {publisher} / {node} / $event**
* **{zone} / {publisher} / {node} / $batch**

Input publication addresses

* **{zone} / {publisher} / {node} / {type} / {instance} / $set**


|Address segment |Description|
|----------------|-----------|
| {zone} | The zone in which publishing takes place. 
| {publisher} | The service that is publishing the information. A publisher provides its identity when publishing its discovery. The publisher Id is unique within its zone.
| {node} | The node that owns the input or output. This is a device identifier or a service identifier and unique within a publisher.
| {type} | The type of  input or output, eg temperature. This convention includes a list of output types. 
| {instance} | The instance of the type on the node. For example, a node can have two temperature sensors. The combination type + instance is unique for the node. The instance can be a name or number. If only a single instance exists the instance can be shortened to “_”
| Commands: | |
| *$value* | The “\$value” command indicates the publisher provides the latest known value of the output. The payload is the raw data. This is intended for constrained devices and for interoperability with 3rd party consumers.
| *$latest* | The “\$latest” command indicates the publisher provides the latest known value of the output including metadata such as node address, timestamp, and the publisher signature. The value is converted to a string. Binary data is converted to base64.
| *$event* | The “\$event command indicates the publisher sends multiple output values that relate together as a single event.
| *$history* | The “\$history command indicates the publisher provides a record containing a - usually 24 hour - history of the values. This is intended to be able to determine a trend without having to store these values. The value is provided in its string format. The content is not required to persist between publisher restarts.
| *$set* | The “\$set” command indicates a consumer is providing the value to control a node input. The publisher subscribes to updates published to this address and updates the node input accordingly.


**MQTT Examples:**

1. The latest value of the first temperature sensor of node 5 published by a service named  *openzwave* is published on an MQTT bus on topic:
    > **myzone/openzwave/5/temperature/1/$latest**

2. To set a switch on device node with ID "3", a message is published on topic:
    > **myzone/openzwave/3/switch/1/$set**

### '$value' Output Command

The payload used with the 'value' command is the straight information without metadata such as timestamp and signature.

Publishing information on this address is required. It is primarily intended for compatibility with 3rd party systems or for use in environments with limited bandwidth or computing power.
Example:
```
zone-1/openzwave/6/temperature/0/$value: 20.6
```

### '$latest' Output Command

The payload used with the '$latest' command includes the address and timestamp of the information and optionally the publisher signature to verify the content. Publishing information on this address is
recommended for environments that are not too limited in bandwidth and computing power.

The message structure is as follows:

| Field        | Data Type | Required     | Description |
|--------------|-----------|------------- |------------ |
| address      | string    | **required** | The address on which the message is published |
| sender       | string    | **required** | Address of the publisher node of the message |
| timestamp    | string    | **required** | ISO8601 "YYYY-MM-DDTHH:MM:SS.sssTZ" |
| unit         | string    | optional     | unit of value type, if applicable |
| value        | string    | **required** | value in string format |


JSON example of a '$latest' publication:
```
zone-1/openzwave/6/temperature/0/$latest:
{
  "message": {
    "address": "zone-1/openzwave/6/temperature/0/\$latest",
    "sender": "zone-1/openzwave/$publisher",
    "timestamp": "2020-01-16T15:00:01.000PST",
    "unit": "C",
    "value": "20.6",
  },
  signature: "..."
}

```

### '$history' OutputCommand

The payload for the '$history' command contains an ordered list of the recent values along with address information and signature. The history is published each time a value changes. 

This publication is optional and intended for environments with plenty of bandwidth and computing power. Consumers can use this history to display a recent trend, like temperature rising or falling, or presenting a history graph. The default length is 24 hours.

The message structure is as follows:

| Field        | Data Type | Required     | Description |
| ----------   | --------  | -----------  | ------------ |
| address      | string    | **required** | The address on which the message is published |
| duration     | integer   | optional     | Nr of seconds of history. Default is 24 hours (24*3600 seconds)
| history      | list      | **required** | eg: [{"timestamp": "YYYY-MM-DDTHH:MM:SS.sssTZ","value": string}, ...] |
|| timestamp   | string    | ISO8601 "YYYY-MM-DDTHH:MM:SS.sssTZ" |
|| value       | string    | Value in string format using the node's unit |
| sender       | string    | **required** | Address of the publisher node of the message |
| timestamp    | string    | **required** | timestamp this message was created |
| unit         | string    | optional     | unit of value type |


A JSON example:
```
zone-1/openzwave/6/temperature/0/$history:
{
  "message": {
    "address" : "zone-1/openzwave/6/temperature/0/\$history",
    "duration": "86400",
    "history" : [
      {"timestamp": "2020-01-16T15:20:01.000PST", "value" : "20.4" },
      {"timestamp": "2020-01-16T15:00:01.000PST", "value" : "20.6" },
      ...
    ],
    "sender": "zone-1/openzwave/$publisher",
    "unit": "C",
  },
  "signature": "...",
}
```

### '$event' Output Command

The payload for the event command is a map of output values along with address information and signature. 

This publication is optional and intended for nodes that send events with multiple related output values.

The message structure is as follows:

| Field        | Data Type | Required     | Description |
| ----------   | --------  | -----------  | ------------ |
| address      | string    | **required** | The address on which the message is published, zone/publisher/node/$outputs |
| event        | map       | **required** | Map with {output type/instance : value} 
| sender       | string    | **required** | Address of the publisher node of the message |
| timestamp    | string    | **required** | timestamp this message was created |

A JSON example:
```
zone-1/vehicle-1/$publisher/$event:
{
  "message" : {
    "address" : "zone-1/vehicle-1/$publisher/\$event",
    "event" : [
      {"speed/0": "30" },
      {"heading/0": "165" },
      {"rpm/0": "2000" },
      ...
    ],
    "sender": "zone-1/vehicle-1/$publisher",
    "timestamp": "2020-01-16T15:00:01.000PST",
  },
  "signature": "...",
}
```
### '$batch' Output Command

The batch command contains a batch of $event values. 

This publication is optional and intended to reduce bandwidth in case for high frequency sampling of multiple values with a reduced publication rate. 

The message structure is as follows:

| Field        | Data Type | Required     | Description |
| ----------   | --------  | -----------  | ------------ |
| address      | string    | **required** | The address on which the message is published, eg zone/publisher/node/\$batch |
| batch        | list      | **required** | The list of outputs with their timestamp
|| timestamp   | string    | ISO8601 "YYYY-MM-DDTHH:MM:SS.sssTZ" |
|| event       | map       | Map with {output type/instance : value} |
| sender       | string    | **required** | Address of the publisher node of the message |
| timestamp    | string    | **required** | ISO8601 timestamp this message was created |


### '$set' Input Command 

Publishers subscribe to the '$set' input address to receive requests to update the input of a node.

Subscribing to the set address is only for nodes that have inputs.

The message structure is as follows:

| Field        | Data Type | Required      | Description
|------------- |-----------|----------     |------ 
| address      | string    | **required**  | The address on which the message is published |
| timestamp    | string    | **required**  | Time this request was created, in ISO8601 format, eg |YYYY-MM-DDTHH:MM:SS.sssTZ. The timezone is the local timezone where the value was published. If a request was received with a newer timestamp, up to the current time, then this request is ignored. |
| sender       | string    | **required** | Address of the publisher node of the message |
| value        | string    | **required** | The input value to set |


A JSON example:
```
zone-1/openzwave/6/switch/0/$set:
{
  "message": {
    "address" : "zone-1/openzwave/6/switch/0/$set",
    "sender": "zone-1/mrbob/$publisher",
    "timestamp": "2020-01-02T22:03:03.000PST",
    "value": "true",
  },
  "signature": "...",
}
```

## Discovery Messages

Publishers that publish discovery of nodes they are connected to (like zwave) must also publish a node that represents the publishing service. The publisher's node id must be **\$publisher**. 

Nodes that are also publishers can choose to publish a single node. In this case the node-id is '\$publisher' and inputs and outputs are linked to this node. 

MQTT/REST node discovery publication address:

  > **{zone} / {publisher} / {node} / $discover**

|Address segment| Description |
| ------------- | ----------- |
| {zone}        | The zone in which the node lives |
| {publisher}   | The service that is publishing the information. A publisher provides its identity when publishing a node discovery. The publisher Id is unique within its zone. |
| {node}        | The node that is discovered. This is a device or a service identifier and unique within a publisher. Two special nodes are defined: “$publisher” is a service node that publishes. “$gateway” represents the device that acts as a gateway to one or more nodes. For example a zwave controller.|
| $discover     | Command for node discovery. |


**Node Discovery Payload***

The discovery message describes in detail the node and its configuration. The objective is for the node to be
sufficiently described so consumers can identify and configure it without further information.

Node discovery message structure:

| Field        | Data Type | Required     | Description
| -----------  |---------- |----------    |------------
| address      | string    | **required** | The address on which the message is published |
| attr         | map       | **required** | Attributes describing the node. Collection of key-value string pairs that describe the node. The list of predefined attribute keys are part of the convention. See appendix. |
| config       | List of **Configuration Records** | optional | Node configuration, if any exist. Set of configuration objects that describe the configuration options. These can be modified with a ‘$configure’ message. |
| sender       | string    | **required** | Address of the publisher node of the message |
| timestamp    | string    | optional     | Time the record is created |

**Publisher Discovery Payload***

Publisher nodes include additional discovery fields in their record. Message structure:

| Field        | Data Type | Required     | Description
| -----------  |---------- |----------    |------------
| address      | string    | **required** | The address on which the message is published |
| attr         | map       | **required** | Attributes describing the node. Collection of key-value string pairs that describe the node. The list of predefined attribute keys are part of the convention. See appendix. |
| certificate  | string    | optional     | A certificate from a trusted source like Lets Encrypt. It is included with publishers that have joined the zone |
| config       | List of **Configuration Records** | optional | Node configuration, if any exist. Set of configuration objects that describe the configuration options. These can be modified with a ‘$configure’ message. |
| publicKey    | string    | **required** | Publisher's public key used verify the signature provided with publications of information and for encrypting messages for this publisher.|
| sender       | string    | **required** | Address of the publisher node of the message |
| timestamp    | string    | optional     | Time the record is created |


**Configuration Record**

The configuration record is used in both node configuration and input/output configuration. Each configuration attribute is described in a record as follows:

| Field    | Data Type| Required | Description |
|--------  |----------|----------|------------ |
| name     | string   | **required** | Name of the configuration. This has to be unique within the list of configuration records. |
| datatype | enum     | optional| Type of value. Used to determine the editor to use for the value. One of: bool, enum, float, int, string. Default is ‘string’ |
| default  | string   | optional| Default value for this configuration in string format |
| description| string | optional | Description of the configuration for human use |
| enum     | \[strings] | optional* | List of valid enum values as strings. Required when datatype is enum |
| max      | float    | optional | Optional maximum value for numeric data |
| min      | float    | optional | Optional minimum value for numeric data | 
| secret   | bool     | optional | Optional flag that the configuration value is secret and will be left empty. When a secret configuration is set in $configure, the value is encrypted with the publisher node public key. |
| value    | string   | **required**| The current configuration value in string format. If empty, the default value is used. |

**Predefined Node Attributes**

| Key              | Value Description |
|--------------    |-------------      |
| version          | Publishers include the version of the myzone convention. Eg v1.0 |
| firmware         | Firmware identifier or version |
| localip          | IP address of the node, for nodes that are publishers themselves |
| location         | String with "latitude, longitude" of device location  |
| mac              | Node MAC address for nodes that have an IP interface |
| manufacturer     | Node make or manufacturer |
| model            | Node model |
| type             | Type of node. Eg, multisensor, binary switch, See the nodeTypes list for predefined values |


Example payload for node discovery in JSON format:
```
zone1/openzwave/5/\$discover:
{
  "message": {
    "address": "zone1/openzwave/5/\$discover",
   
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
    "sender": "zone1/openzwave/$publisher",
  },
  "signature": "...",
}
```

### Input/Output Discovery

Inputs and outputs discovery are published separately from the node to allow control over which ones are shared with other zones.

MQTT/REST publication address formats:
  >Input discovery:  **{zone} / {publisher} / {node} / {inputType} / {instance} / $discover**

  >Output discovery: **{zone} / {publisher} / {node} / {outputType} / {instance} / $discover**


| Address segment | Description |
| :-------------- | ----------- |
| {zone}       | The zone in which the node lives
| {publisher}  | The service that is publishing the information. A publisher provides its identity when publishing a node discovery. The publisher Id is unique within its zone. 
| {node}       | The node whose input or output is discovered. This is a device or a service identifier and unique within a publisher.
| {inputType}  | Type identifier of the input. The list of predefined types is part of this convention. 
| {outputType} | Type identifier of the output. The list of predefined types is part of this convention. 
| {instance}   | The instance of the input or output on the node. If only a single instance exists the convention is to use 0 unless a name is used to provide more meaning.
| $discover    | Command for node discovery. 


MQTT example:

The discovery of a single temperature sensor on node '5', published by a service named 'openzwave' is published on an MQTT bus on topic:

  > **myzone/openzwave/5/temperature/0/$discover**


**Message Structure**

| Field       | Data Type | Required     | Description |
|------------ |---------- |----------    |------------ |
| address     | string    | **required** | The address on which the message is published |
| config      | List of **Configuration Records**|optional|List of Configuration Records that describe in/output configuration. Only used when an input or output has their own configuration. See Node configuration record above for the definition
| datatype    | string    | optional     | Value datatype. One of boolean, enum, float, integer, jpeg, png, string, raw. Default is "string".
| default     | string    | optional     | Default output value
| description | string    | optional     | Description of the in/output for humans
| enum        | list      | optional*    | List of possible values. Required when datatype is enum
| max         | number    | optional     | Maximum possible in/output value
| min         | number    | optional     | Minimum possible in/output value
| sender      | string    | **required** | Address of the publisher node of the message |
| timestamp   | string    | optional     | Time the reocrd is created
| unit        | string    | optional     | The unit of the data type
| **value**   | string    | **required** | The input or output value at time of discovery


Example payload for output discovery in JSON format:
```
zone1/openzwave/5/temperature/0/$discover:
{
  "message": {
    "address": "zone1/openzwave/5/temperature/0/$discover",
    "datatype": "float",
    "sender": "zone1/openzwave/$publisher"
    "timestamp": "2020-01-20T23:33:44.999PST",
    "unit": "C",
    "value": "20.5",
  },
  "signature": "...",
}
```

## Configuration Message

Configure Node Publication Address

> {zone}/{publisher}/{node}/$configure


Configuration Message structure:

| Field		     | type	    | required     | Description
|--------------|----------|--------------|-----------
| address      | string   | **required** | The address on which the message is published |
| config 	     | map      | **required** | key-value pairs for configuration to update { key: value, …}
| sender   		 | string   | **required** | Address of the sender submitting the request. This is the zone/publisher/node of the consumer.
| timestamp		 | string   | **required**    | Time this request was created, in ISO8601 format



## Node Status Message

The availability status of a node is published by its publisher when the availability changes or errors are encountered. Publishing of a node status is required.

Status Address:

**{zone} / {publisher} / {node} / $status**

| Address segment | Description
|-----------------|--------------
| {zone}          | The zone in which discovery takes place. 
| {publisher}     | The publisher of the node discovery which is handling the configuration update for that node.
| {node}          | The node whose configuration is updated. 
| $status         | Keyword for node status. Published when the availability of a node changes or new errors are reported. It is published by the publisher of the node.


### Status Message Structure

| Field 		   | type     | required 	   | Description |
|--------------|--------- |------------  |------------ |
| address      | string   | **required** | The address on which the message is published |
| errorCount   | integer  | optional     | Nr of errors since startup |
| errorMessage | string   | optional 	   | Last reported error message |
| errorTime    | string   | optional		 | timestamp of last error message in ISO8601 format |
| interval     | integer  | **required** | Maximum interval of status updates in seconds. If no updated status is received after this interval, the node is considered to be lost. |
| lastSeen     | string   | **required** | timestamp in ISO8601 format that the publisher communicated with the node. |
| sender   		 | string   | **required** | Address of the sender submitting the request. This is the zone/publisher/node of the consumer. |
| status       | enum     | **required** | The node availability status. See below for valid values |
| timestamp    | string   | **required** | Time the status was last updated, in ISO8601 format |

Status values:
* ready: node is ready to perform a task
* asleep: node is a periodically listening and updates to the node can take a while before being processed.
* error: node is in an error state and needs servicing
* lost: communication with the node is lost

# Zone Certificate Authority Service - ZCAS

## Issue New Keys

ZCAS subscribes to node discovery messages on address {zone}/{publisher}/{node}/$discover, and determines if a new certificate must be issued. If so the ZCAS send a message to the node containing the new keys and certificate.

The address used to issue new keys and certificate is:

> {zone} / {publisher} / \$zcas

The message structure is fully encrypted with the existing public key of the node

| Field 		  | type 		 | required 		| Description |
|-------------|----------|------------  |-----------  |
| address     | string   | **required** | The address on which the message is published |
| publicKey   | string   | **required** | The new public key to use by the node |
| privateKey  | string   | **required** | The new private key to use by the node |
| sender   	  | string   | **required** | Address of the sender, eg: zone-1/$zcas/$publisher of the consumer. |
| zoneKey     | string   | optional     | Shared zone key for encrypting zone messages
| certificate | string   | **required** | The new certificate to use by the node 
| timestamp   | string   | **required** | Time the message was created |

The node must decrypt the message, then verify if the timestamp is valid and the signature matches that of the ZCAS service. The signature check is based on the payload without the signature.

# Zone Security Service - ZSS

The ZSS subscribes to published messages and validates that the messages carry a valid signature. If the signature is not valid then the administrator is modified.

# Zone Bridge Service - ZBS

The bridge service has facility to select which nodes, inputs or outputs are bridged. 

## Bridge Configuration

Bridge configuration variables:

| Field 		   | value type   | value 		     | Description
|------------- |----------    |------------  |-----------
| host         | string       | **required** | IP address or hostname of the remote bus
| port         | integer      | optional     | port to connect to. Default is determined by protocol
| protocol     | enum ["MQTT", "REST"] | optional  | Protocol to use, MQTT (default), REST API, ...
| format       | enum ["JSON", "XML"]  | optional  | Publishing format used on the external bus. Default is JSON. Only include this if a different format is needed. This will invalidate the signature.
| clientId     | string       | **required** | ID of the client that is connecting
| loginId      | string       | optional     | Login identifier
| credentials  | string       | optional     | Password to connect


## Add/remove nodes to export

To forward nodes through the bridge, use the pushbuttons for $forward and $remove.

MQTT Example:
* {zone} / {publisher} / \$bridge / $pushbutton / $forward
* {zone} / {publisher} / \$bridge / $pushbutton / $remove

Address segments:

| Field 		     | Description
|--------------- |------------|
| {zone}         | The zone the bridge lives in.
| {publisher}    | The publisher of the bridge.
| $bridge        | Reserved ID of a bridge node.
| $pushbutton    | Input type 
| $forward       | button instance to add a forward. Payload contains the address record.
| $remove        | button instance to remove a forward. Payload is the address record


Message structure:

| Field 		   | type 		| cabvalue 		    | Description |
|------------- |----------|------------  |----------- |
| address      | string   | **required** | The address on which the message is published |
| sender       | string   | **required** | Address of the sender, eg: zone-1/$zcas/$publisher of the consumer. |
| fwdValue     | boolean  | optional     | Forward the output $value publication(s), default=true
| fwdLatest    | boolean  | optional     | Forward the output $latest publication(s), default=true
| fwdHistory   | boolean  | optional     | Forward the output $history publication(s), default=true
| fwdDiscovery | boolean  | optional     | Forward the node/output $discovery publication, default=true
| fwdStatus    | boolean  | optional     | Forward the node $status publication, default=true



# Appendix A: Input and Output Types

When available, units used in publication follow the SI standard 

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
| dewpoint        | C, F    | float     | Dewpoint in degrees Celcius
| distance        | m, yd, ft, in | float | distance in meters
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
| latency         | sec     | float     | 
| location        |         | List of 3 floats | [latitude, longitude, elevation]
| locked          |         | boolean   | Door lock status 
| luminance       | cd/m2, lux | float  | Amount of light in candela/m2 or in lux
| magneticfield   | T, mT, uT, nT, G(auss) | float | Static magnetic field in Tesla or Gauss
| motion          |         | boolean   | Motion detection status |
| power           | W       | float     | Power consumption in watts |
| pushbutton      |         | boolean   | Momentary pushbutton |
| signalstrength  | dBm     | float     |
| speed           | m/s, Km/h, mph | float |
| switch          |         | boolean   |
| temperature     | C, F    | float     | Celcius or Fahrenheit. The default is Celcius if available.
| ultraviolet     | UV      | float     | Radiation index with wavelength from 10 nm to 400 nm, range 0-11+
| voltage         | V       | float     | Volts
| waterlevel      | m(eters), ft, in | float |
| wavelength      | m       | float     |
| weight          | Kg, Lbs | float     | 
| windchill       | C, F    | float     | Apparent temperature based on the air temperature and wind speed,    when lower than the air temperature.
| windspeed       | m/s, km/h, mph | float | 



# Appendix B: Implementation Libraries

### GoLang

-   [[https://golang.org/pkg/crypto/rsa/\#GenerateKey]](https://golang.org/pkg/crypto/rsa/#GenerateKey)

<!-- -->

-   [[https://golang.org/pkg/crypto/rsa/\#SignPSS]](https://golang.org/pkg/crypto/rsa/#SignPSS)

-   [[https://golang.org/pkg/crypto/rsa/\#VerifyPSS]](https://golang.org/pkg/crypto/rsa/#VerifyPSS)

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

-   [[https://github.com/kjur/jsrsasign]](https://github.com/kjur/jsrsasign)

-   [[https://github.com/rzcoder/node-rsa]](https://github.com/rzcoder/node-rsa)

Python
------

Paramiko

import hashlib, paramiko.agent

data\_sha1 = haslib.sha1(recordToPublish).digest()

agent = paramiko.agent.Agent()

key = agent.keys\[0\]

signature = key.sign\_ssh\_data(None, data\_sha1)
