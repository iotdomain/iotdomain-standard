The IoTDomain Standard
=======================

The IoTDomain standard defines a simple and easy to use information exchange method between IoT publishers and consumers.


[[TOC]]

# Introduction

As connected devices become more and more prevalent, so have the problems surrounding them. These problems fall into multiple categories: 

## Interoperability
   
The use of information produced by various devices is challenging because of the plethoria of different protocol and data formats in use. This is apparent in home automation solutions such as OpenHAB and Home Assistant that each implement hundreds of bindings to talk to different devices and services. Each solution has to reimplement these bindings. This implementation then has to be adjusted to different platforms, eg Linux, Windows and MacOS, which adds even more work.

Without a common standard it is unavoidable that manufacturers of IoT devices choose their own protocols. It is in everyone's interest to provide a standard that enables an open information interchange so that bindings only have to be implemented once. 

This standard defines the messages for information exchange.

## Discovery

Discovery of connected IoT devices often depends on the technology used. There is no easy to use standard that describes what and how discovery information is made available to consumers independent of their implementation. 

Application developers often implement solutions specific to their application and the devices that are supported. To facilitate information exchange it must be possible to discover the information that is available independent of the technology used.

This standard defines the process and messaging for discovery of devices and services without the need for a central resource directory. 

Note. The IETF draft "CoRE Resource Directory (draft-ietf-core-resource-directory-20) takes the approach where a centralized service provides a directory of resources. This approach works fine but does not fit within the concept of this standard for several reasons: Mainly it requires and additional direct connection between client and directory which adds an additional protocol with its own encoding. This standard is based on only requiring connections between client and message broker, keeping the attack footprint to a minimum. The protocol is JSON based for all messages. Second, it does not support sharing of information between domains. Last, it does not support the concept of last will and testament when a publisher unexpectedly disconnects from the network. Like most things it is possible to make it work but it is not the best fitting solution.

## Configuration

Configuration of IoT devices is often done through a web portal of some sort from the device itself or a gateway. These web portals are not always as secure as they should be. They often require a login name and password and lack 2 factor authentication. Passwords are easily reused. Backdoors are sometimes left active. Overall security is lacking.
   
Configuration is not always suited for centralized management by application services. For example, to configure all temperature sensors to report in Celcius the user has to login to the device management portal(s), find the sensor and find the configuration for this. This is difficult to automate.

This standard defines the process and messaging for remote configuration of devices and services.

Nodes that can be configured contain a list of configuration records described in the node discovery. The configuration value can be updated with a configure command. This is described further in the configuration section. 

## Security And Privacy
   
Security is a major concern with IoT devices. Problems exist in several areas:

  1. It is difficult to design devices for secure access from the internet. The existance of large botnets consisting of hacked cameras and other devices show how severe this problem is. Good security is hard and each vendor has to reinvent this wheel. This is not likely to change any time soon.
  
  The solution chosen in this standard is to simply assume that the IoT device itself is insecure. All communication with the device is going through a secure publisher that follows this standard.
  
  2. Commercial devices that connect to a service provider share personal information without the user understanding what this information is, and without having control on how it is used. While regulations like Europe's [GDPR](https://en.wikipedia.org/wiki/General_Data_Protection_Regulation) attempt to address this ... somewhat, reports of data misuse and breaches remain all too frequent.

  3. There is no easy secure way to serve information over the internet. It either requires opening a port in the firewall or use a 3rd party service provider, which leads to the previous two problems.

  The solution chosen in this standard is to share information using a secure message bus. 

  4. Information can be tampered with and its source cannot be verified. The amount of fake news and information shows this is problem is pervasive.

  This standard uses JWS message signing and identity verification to ensure the information is not tampered with and the information source can be veried. 
  

# Terminology

| Terminology   | Description |
|:-----------   |:------------|
| Account       | The account used to connect a publisher to an message bus |
| Address       | Address of the node consisting of domain, publisher and node identifier. Optionally it can include the input or output type and instance.|
| Authentication| Method used to identify the publisher and subscriber with the message bus |
| Data          | The term 'data' is used for raw data collected before it is published. Once it is published it is considered information.|
| DBM           | Domain Bridge Manager service that creates bridges with other domains |
| Discovery     | Description of nodes, their inputs and outputs|
| DSS           | IoT Domain Security Service. This service manages keys and certificates of IoT domain members and guards for invalid publishers |
| Information   | Anything that is published by a producer. This can be sensor data, images, discovery, etc|
| IoT Domain    | An area in which information is shared between members |
| JWS           | JSON Web Signature, used to sign messages |
| Message Bus   | A publish and subscribe capable transport for publication of information. Information is published by a node onto a message bus. Consumers subscribe to information they are interested in use the information address. |
| Node          | An IoT node is a device or service that provides information and accepts control input. Information from this node can be published by the node itself or published by a (publisher) service that knows how to access the node. |
| Node Input    | Input to control an IoT node, for example a switch.|
| Node Output   | Node Information is published using outputs. For example, the current temperature.|
| Publisher     | A service that is responsible for publishing node information on the message bus and handle configuration updates and control inputs. Publishers are nodes. Publishers sign their publications to provide source verification.|
| Retainment    | A feature of a message bus that remembers that last published message. Not all message busses support retainment. It is used in publishing the values and discovery messages so new clients receive an instant update of the latest information |
| Signature     | A JWS message signature for verification the message hasn't been tampered with|
| Subscriber    | Consumer of information that uses node address to subscribe to information from that node.|


# Versioning

The standard uses semantic versioning in the form v\{major}.\{minor}.

Future minor version upgrades of this standard must remain backwards compatible. New fields can be added but MUST be optional. Implementations MUST accept and ignore unknown fields and in general follow the [robustness principle](https://engineering.klarna.com/why-you-should-follow-the-robustness-principle-in-your-apis-b77bd9393e4b)

A major version upgrade of this standard is not required to be backwards compatible but **MUST** be able to co-exists on the same bus. Implementations must ignore messages with a higher major version.

Publishers include their version of the standard when publishing their node. See 'discovery' for more information.

# Implementation Agnostic

This standard is implementation agnostic. It is a standard that describes the information format and exchange for discovery, configuration, inputs and outputs, irrespective of the technology used to implement it. Use of different technologies will actually serve to further improve interoperability with other information sources.

A reference implementation of a publisher is provided for the golang and python languages using the MQTT service bus.


# System Overview

![System Overview](./system-overview.png)


## IoT Domain

An IoT Domain defines a physical or virtual area in which information is shared amongst its members. An IoT domain can be a home, a street, a city, or a virtual area like an industrial sensor network or even a game world. Each domain has a globally unique domain name, except for the local domain called 'local'. Local domains cannot share information with other domains.

An IoT domain has members which are publishers or subscribers (consumers). All members have access to information published in that domain. The information is not available outside the domain unless intentionally shared. Publication in the domain is limited to members that have the publish permissions. Not surprisingly these are called 'publishers'.

An IoT Domain can be closed or open to consumers. An open domain allows any consumer to subscribe to publications without providing credentials. A closed domain requires consumers to provide valid credentials to connect to the message bus of that domain. Whether a domain is open or closed is determined by the configuration of the message bus.

For internet accessible IoT domains, the IoT domain name can be the same as the domain name of the message bus. The message bus must be configured with a valid certificate just like any other internet service.

IoT domains can also operate on a local area network in which case the do not need a registered domain name.

## Message Bus

The use of publish/subscribe message bus has a key role in exchange and security of information. It not only routes all messages but also secures publishers and consumers by allowing them to reside behind a firewall, isolated from any other internet access.

A message bus carries only publications for the domain it is intended for. Multi-domain or multi-tenant message busses can be used but each domain must be fully isolated from others. Note that a bridge can publish messages from one domain into another. More on this below.

As the network topology is separate from the message bus topology, publishers and subscribers can be on different networks and behind firewalls. This reduces the attack footprint as none of the publishers or subscribers need to be accessible from the internet. The message bus is the only directly exposed part of the system. It is key to make sure the message bus is properly secured and hardened. For more on securing communication see the security section.

The message bus must be configured to require proper credentials of publishers. Open IoT domains can allow subscribers to omit credentials. Obviously this should only be done in a secured environment.

### Message Bus Protocols

This standard is agnostic to the message bus implementation and protocol. The minimum requirement is support for publishing and subscribing using addresses. 

It is highly recommended that publisher implementations support the MQTT protocol to allow their use on MQTT message busses. Support for additional protocols such as AMQP or HTTP with websockets is optional. 

The reason to choose MQTT as the defacto default is because a common standard is needed for interoperability. MQTT is low overhead, well supported, supports LWT (Last Will & Testament), has QOS, and clients can operate on constrained devices. It is by no means the ideal choice as explained in [this article by Clemens Vasters](https://vasters.com/archive/MQTT-An-Implementers-Perspective.html)

If in future a better protocol becomes the defacto standard, the MQTT protocol will remain supported as a fallback option until this changes in a future version of this standard.

### Guaranteed Delivery (or lack thereof)

The use of a simple message bus, like MQTT, brings with it certain limitations, the main one being the lack of guaranteed delivery. The role of the message bus is to deliver a message to subscribers **that are connected**. While this simplifies the implementation, it pushes the problem of guaranteed delivery to the application. It is effectively a lossy transport between publishers and subscribers. 

MQTT supports 'retainment' messages where the last value of a publication is retained. When a consumer connects, it receives the most recent message for all addresses it subscribes to. To receive the most recent information subscribers do not have to be connected at the same time as the publishers. Note that not all MQTT implementations support retainment. 

This usage of the message bus will do fine in cases where the goal is to get the most recent value. This will work fine if the loss of a occasional output value is not critical. In addition, the use of the history publication can be used to fill in any gaps if needed. It is well suited for monitoring environmental sensors.

In cases of critical messages, such as emergency alerts, a confirmation or failover mechanism might be needed. To guarantee end-to-end delivery requires application level support. For example, if an alert is send to an input a confirmation can be published to the return address included in the initial alert.

Based on these considerations the use of simple message bus, like MQTT, should be sufficient for most use-cases. 

### Severely Constrained Clients

For severely constrained devices such as micro-controller, a message bus client might simply be too complicated to implement. While the JSON message format is easy to generate, it is not as easy to parse. In these cases a publisher service can be translates between the native protocol and this standard.

### Severely Constraint Bandwidth

In the IoT space, bandwidth can be quite limited. The use of LTE Cat M1, NB-IoT, or LPWAN LTE restricts the bandwidth due to cost of the data plan. For example, some plans are limited to 10MB per month. If a sensor reports every minute then a single message is limited to approx 1KB per message including handshake. This almost certainly requires some form of compression or other optimization. Just establishing a TLS connection can take up this much.

The objective of this standard is to support interoperability, not low bandwidth. These are two different concerns that are addressed separately. The use of adapters make it very easy to work with low bandwidth devices using a compressed or native protocol.

## Nodes

IoT Nodes are the sources and destination of information through their inputs and outputs. A node can be a hardware device, a service, or a combination of both. A node has inputs and/or outputs through which information passes. A node can have many as inputs and outputs that are connected to the node. Inputs and outputs are part of their node and cannot exist without it. 

Gateway devices that connect to nodes are nodes themselves. They can have inputs or outputs but this is optional. Their role is to relay information from other nodes. For example, a ZWave USB stick is considered a gateway and connects to ZWave devices through the ZWave mesh network.

## Publishers

Publishers are the service that connect to the message bus and publish node information. Only publishers are allowed to publish information on the message bus.

For example, a ZWave publisher obtains the temperature from a ZWave temperature sensor via a ZWave controller (gateway) and publishes this information on the message bus as an output of the sensor node. 

Publishers must use credentials to connect to a message bus before they can publish. To publish securely, a publisher must also have to join the IoT domain by registering with the domain's Security Service (DSS). More on that later.

Nodes can also publish information according to this standard directly, in which case they are also a publisher. These devices identify themselves as both a publisher and as a node with inputs and/or outputs. 

Publishers:

1. Publish their own verified identity
2. Handle updates to their own security keys and certificates 
3. Publish node discovery information
4. Publish node input and output discovery information
5. Publish node output values 
6. Handling requests to update node inputs
7. Handle updates to node configurations

These tasks are discussed in more detail in following sections.

### Addressing

Information is published using an address on the message bus. This address consists of the node address to whom the information relates \{domain}/\{publisherID}/\{nodeID}, an input or output type and instance, and a message type. The input/output type and instance are used to publish input and output information.

Addresses can only contain alphanumeric, hyphen (-), and underscore (\_) characters. Reserved words start with a dollar ($) character. The separator is the '/' character. All addresses end with a message type indicating the content of the message.


> Address format:
>  **\{domain} / \{publisherid} / \[ \{nodeid} \[ / \{ioType} / \{instance} \] \] / \{messagetype\}**

Where:
* \{domain} is the domain in which the node information is published. This can be "local" or an internet domain.
* \{publisherid} is the ID of the service that publishes the node, input and output information
* \{nodeid} is the ID of the node that is being published
* \{ioType} and \{instance} refers to a particular input or output of the node. 
* \{messagetype} indicates the content of the message, be it publishing of discovery or values.
For message bus systems that do not support the '/' character as address separator, the separator character of the message bus implementation can be used. However, the message itself must contain the original address using the '/' character as the separator to allow for interoperability between different message bus implementations.

**Reserved message types:**

The standard predefines the following message types.

| Publisher message type | Purpose |
|:--------     |:--------|
| \$identity   | Publication of a publisher's identity (domain/publisherid/\$identity)
| \$set        | Renew a publisher's identity by the DSS (domain/publisherid/\$set)


| Node message type | Purpose |
|:--------     |:--------|
| \$alias      | Command to set the node alias |
| \$batch      | Publication of a batch of events |
| \$configure  | Command to update the node configuration |
| \$create     | Command to create a node. Only usable with publishers that can create/delete nodes |
| \$delete     | Command to delete a node. Only usable with publishers that can create/delete nodes |
| \$event      | Publication of all output values at once using a single event message |
| \$node       | Publication of a node discovery |
| \$lwt        | Publisher last will and testament if supported |

| Output message type | Purpose |
|:--------     |:--------|
| \$forecast   | publication of a list of projected output values |
| \$history    | publication of a list of historical output values |
| \$latest     | Publication of a single output value including metadata |
| \$output     | Publication of a node output discovery |
| \$raw        | Publication of raw output sensor value without any signature or metadata |

| Input message type | Purpose |
|:--------     |:--------|
| \$input      | Publication of an input discovery |
| \$set        | Set the input value |


### Message Publication

This standard supports multiple modes of message publication: 
* plain text JSON
* JWS signed messages as described in RFC 7515.
* Compact JWS

*  Use of plain JSON is intended for a trusted environment while JWS is used in a secured domain. While it is recommended to use only a single publication method, they can be mixed.

See the security section for more details.

### Node Aliases

Node Aliases are intended to support replacing devices while retaining node, input and output addresses. When devices are replaced, the hardware address or ID of the replacement can differ from the original. ZWave for example generates a new internal node address each time a node is added to the network. This leads to the problem that when replacing a node, all consumers must be updated to use the replacement node ID, which can take quite a bit of effort. 

To address this problem, a nodeID can be replaced with an 'alias'. When a node alias is set it replaces the nodeID and all publications of the node, its inputs and outputs use the alias  as the node ID.

The node alias can be set through the node $alias command and viewed in the node discovery message.

## Subscribers

Consumers that are only subscribers do not have to be registed as publishers. They need credentials to access the message bus and can simply subscribe to publisher addresses. To control configuration and inputs however consumers must be registered as a publisher.

## Domain Security Service - DSS

The Domain Security service issues keys and identity signatures to publishers. 

Publishers use the issued keys to sign messages using JWS (JSON Web Signing). The signature is used to verify that the message is sent by the publisher and hasn't been tampered with. 

The publisher identity itself is signed by the DSS. It is used to verify that the publisher is who it claims to be and can be trusted. For public domains the DSS includes a certificate from an internet CA with its identity so that the DSS identity can be verified.

For more detail, see the security section


# Discovery

Support for discovery lets consumers find nodes, inputs or outputs they are interested in. The objective is for the node to be sufficiently described so consumers can identify and configure it without further information.

Publishers are responsible for publishing discovery messages for nodes, their inputs and outputs. The discovery data describes the nodes in detail, including their type, attributes and configurations.

Just like publications of the various values, the discovery publications consist of a JSON object with two fields: "message" and "signature". Creation and verification of the base64 encoded signature is described in the 'signing' section.

**Retainment:**

Where supported, discovery messages are published with retainment. When connection to the message bus was lost and is re-established, the discovery messsages are re-published in case the retainment cache was cleared.

When retainment is not available on the message bus, it can be simulated using a discovery service. When subscribers connect to the message bus they send the discovery service a subscription request. The discovery service republishes the most recent discovery messages it received within the last 24 hours 1 minute after receiving discovery requests. This is a very simple service that simply republishes what it received. The 1 minute period is intended to prevent a message storm when multiple publishers connect to the bus at the same time.

In all cases discovery messages are re-published periodically by the publisher to indicate it is still alive and its nodes are available. The default interval is once a day but can be changed through publisher configuration. Without retainment this is best combined with a discovery service. 

## Publisher Discovery

Publisher identity messages contain the publisher's public identity. The identity includes the public key used to verify the JWS signature of the messages published, and the issuer and signature of the identity: DSS, self signed or CA. Messages that fail signature verification MUST be discarded. The public key is also used to encrypt input and configuration messages to this publisher. The process of encryption uses JWE as described in the security section.

Publisher discovery:

  >  **\{domain}/\{publisherid}/\$identity**

Message structure

| Field           | type     | Description |
|:--------------- |:-------  | ------ |
| address         | string   | **required** | The address of the publication
| certificate     | string   | optional | Optional x509 certificate, base64 encoded. Included with the DSS identity to be able to verify it with a 3rd party. |
| domain          | string   | **required** | IoT domain this publisher belongs to. "local" or "test" for local domains |
| issuerId        | string   | **required** | ID of issuer identity, like the DSS ($dss), self signed (publisherID) or the CA name. |
| location        | string   | optional | Optional location of the publisher, city, province/state, country |
| organization    | string   | optional | Organization the publisher belongs to |
| publicKey       | string   | **required** | PEM encoded public key for verifying publisher signatures and encrypting messages to this publisher |
| publisherId     | string   | **required** | ID of this publisher 
| signature       | string   | **required** | base64 encoded signature from issuer of this identity record with the signature field blank
| timestamp       | string   | **required** | Time the identity was signed |
| validUntil      | string   | **required** | ISO8601 Date this identity is valid until |

## \$lwt: Publisher Last Will & Testament (MQTT)

This message only applies when using a message bus that supports LWT (last will & testament) .

The LWT option lets the message bus publish a message when the publisher connection is lost unexpectedly. A message with status "connected" and "disconnected" is sent by the publisher when connecting or gracefully disconnecting. The status "lost" is set through last will & testament feature and send by the message bus if the publisher unexpectedly disconnects. 

Address:  **\{domain}/\{publisherid}/\$lwt**

Message structure:
The message structure:

| Field      | Data Type | Required     | Description  |
|:---------- |:--------  |:-----------  |:------------ |
| address    | string    | **required** | Address of the publication |
| status     | string    | **required** | LWT status: "connected", "disconnected", "lost" |


## Discover Nodes

Node discovery messages contain a detailed description of the node. It does not contain information on inputs and outputs as these are published separately. This reduces the amount of traffic for simple consumers that only subscribe to certain types of inputs or outputs.

Node discovery address:

  >  **\{domain}/\{publisherid}/\{nodeid}/\$node**

Where:
* {domain} is the IoT domain in which the node lives
* {publisherid} is the ID of the publisher of the information. The publisher Id is unique within its domain
* {nodeid} is the ID of the node. This is a device or a service identifier and unique within a publisher. 
* $node message type for node discovery

Node discovery message structure:

| Field        | Data Type | Required     | Description
|:-----------  |:--------- |:----------   |:------------
| address      | string    | **required** | The address of the publication|
| attr         | map       | **required** | Key value pairs describing the node. The list of predefined attribute keys are part of the standard. See appendix B: Predefined Node Attributes. | 
| config       | map of **Configuration Records** | optional | Map of attribute configuration by attribute name. Each record describes the configuration constraints. The attribute value can be set with a ‘$configure’ message based on the configuration description.|
| deviceID     | string    | **required** | The node device ID used as nodeID if no alias is set|
| nodeID       | string    | **required** | The node ID as generated by the device|
| nodeType     | string    | **required** | Description of the type of node, see Appendix B for predefined types|
| status       | map       | optional     | key-value pairs describing node performance status|
| timestamp    | string    | **required** | Time the record is last updated |

**Configuration Record**

The configuration record describes the constraints of the configuration. :

| Field    | Data Type | Required  | Description |
|:-------- |:--------- |:--------- |:----------- |
| dataType | enum      | optional| Type of value. Used to determine the editor to use for the value. One of: bool, enum, float, int, string. Default is ‘string’ |
| default  | string    | optional| Default value for this configuration in string format |
| description| string  | optional | Description of the configuration for human use |
| enum     | \[strings] | optional* | List of valid enum values as strings. Required when dataType is enum |
| max      | float     | optional | Optional maximum value for numeric data |
| min      | float     | optional | Optional minimum value for numeric data | 
| secret   | bool      | optional | Optional flag that the configuration value is secret and its value will not be included in publications. 

**Node Configuration**
some node configuration attributes are standardized. The following attributes are optional:

| Config         | Data Type | Default      | Description |
|:-------------  |:--------- |:----------   |:----------- |
| name           | string    | ""           | Node friendly name|
| publishBatch   | int       | 0            | publish $batch messages containing N events. 0 to ignore|
| publishEvent   | bool      | false        | publish $event messages containing event output values. Only outputs that have their event configuration enabled are included.|
| publishHistory | bool      | true | enable publishing the history of outputs if also enabled in the output itself. Set to false to disable for all outputs.|
| publishLatest  | bool      | true | enable publishing the latest value of outputs if also enabled in the output itself, Set to false to disable for all outputs. |
| publishRaw     | bool      | true | enable publishing the raw value of outputs if also enabled in the output itself, Set to false to disable for all outputs. |


Example payload for node discovery

~~~json
{
  "address": "local/openzwave/5/$node",
  "nodeID": "5",
  
  "attr": {
    "make": "AeoTec",
    "name": "Garage Sensor",
    "type": "multisensor",
    "publishHistory": "false",
  },
  "config": {
    "name": {
      "dataType": "string",
      "description": "Friendly name of the device or service",
    }, 
    "publishHistory": {
      "dataType": "bool",
      "description": "Enable publishing of the output history when also enabled in the output.",
    }, 
  ],
  "timestamp": "2020-01-20T23:33:44.999PST",
}
~~~

## Discover Outputs

Discovered outputs are published separately from the node. This facilitates control over which outputs are shared with other domains. 

Address of output discovery:

> **\{domain}/\{publisherid}/\{nodeid}/\{outputtype}/\{instance}/\$output**

| Address segment | Description |
| :-------------- | :---------- |
| {domain}        | The IoT domain in which the node lives, or "local" for local domains |
| {publisherid}   | The service that is publishing the information |
| {nodeid}  | ID of the node that owns the input or output |
| {outputtype}    | Type identifier of the output. For a list of predefined types see Appendix D |
| {instance}      | The instance of the input or output on the node. If only a single instance exists the standard is to use 0 unless a name is used to provide more meaning|
| \$output        | Message type for output discovery |

For example, the discovery of a temperature sensor on node '5', published by a service named 'openzwave', is published on address:

  > **local/openzwave/5/temperature/0/\$output**

The output message structure:

| Field       | Data Type | Required     | Description |
|:----------- |:--------- |:---------    |:----------- |
| address     | string    | **required** | Address of the publication |
| attr        | map       | optional     | Attributes describing the output, including config values |
| config      | List of **Configuration Records**|optional| See node configuration for details |
| dataType    | string    | optional     | Value data type. See appending for dataTypes, default is string
| enumValues  | list      | optional*    | List of possible values. Required when dataType is enum |
| max         | number    | optional     | Maximum possible in/output value |
| min         | number    | optional     | Minimum possible in/output value |
| timestamp   | string    | **required** | Time the record is last updated |
| unit        | string    | optional     | The unit of the output value |


**Output Destination Configuration**
Where supported, outputs can be configured with methods of publishing their output value to a specific destination. These attributes with their configuration are optional. Standard configuration settings for configuring destinations are:

| Config         | Data Type | Default   | Description |
|:-------------- |:--------- |:--------  |:----------- |
| publishEvent   | bool      | false     | include output in the $event publication (when enabled)
| publishFile    | string    | ""        | save output value to a local file, ignored if filename is ""
| publishHistory | bool      | false     | publish on the output $history address
| publishLatest  | bool      | true      | publish on the output $latest address
| publishRaw     | bool      | true      | publish on the output $raw address

Example payload for output discovery. In this case the published output methods are not configurable:

~~~json
{
  "address": "local/openzwave/5/temperature/0/$output",
  "attr": {
    "publishHistory": "true",
    "publishLatest": "true", 
    "publishRaw": "true",
  },
  "dataType": "float",
  "timestamp": "2020-01-20T23:33:44.999PST",
  "unit": "C"
}
~~~   

## Discover Inputs 

Discovered inputs are published and configured separately from the node and outputs. This facilitates control over which inputs and outputs are enabled and which are shared with other domains.

Address of input discovery:

> **\{domain}/\{publisherid}/\{nodeid}/\{inputtype}/\{instance}/\$input/**

| Address segment | Description |
| :-------------- | :---------- |
| {domain}        | The IoT domain in which the node lives, or "local" for local domains |
| {publisherid}   | The service that is publishing the information |
| {nodeid}        | ID of the node that owns the input |
| {inputtype}     | Type identifier of the input. For a list of predefined types see Appendix D |
| {instance}      | The instance of the input on the node. If only a single instance exists the standard is to use 0 unless a name is used to provide more meaning|
| \$input         | Message type for input discovery |

For example, the discovery of a switch on node '5', published by a service named 'openzwave', is published on address:

  > **local/openzwave/5/switch/0/\$input**

The input discovery message structure:

| Field       | Data Type | Required     | Description |
|:----------- |:--------- |:---------    |:----------- |
| address     | string    | **required** | Input discovery address |
| attr        | map       | optional     | Attributes describing the input, including config values |
| config      | List of **Configuration Records**|optional| See node configuration for details |
| dataType    | string    | optional     | Value data type. See appending for dataTypes, default is string
| enumValues  | list      | optional*    | List of possible values. Required when dataType is enum |
| max         | number    | optional     | Maximum possible in value |
| min         | number    | optional     | Minimum possible in value |
| timestamp   | string    | **required** | Time the record is last updated |
| unit        | string    | optional     | The unit of the input value |

**Input Configuration**
Where supported, inputs can be configured with methods of receiving their input value from a specific source. These attributes with their configuration are optional. 

Standardized configuration settings for configuring input sources are:

| Attr/Config   | Data Type | Default | Description |
|:------------- |:--------- |:--------|:----------- |
| pollInterval  | int       | 0       | interval in seconds to poll source (only for rest endpoints). 0 is disabled |
| login         | string    | ""      | Basic Auth login from rest endpoints, use secret=true |
| password      | string    | ""      | Basic Auth login from rest endpoints, use secret=true |
| setEnabled    | bool      | true    | When enabled, the input can be set with a $set command |
| source        | string    | ""      | Source to read input from, subscription, file://filename or http://host|

Example payload for input discovery. The input set command is enabled:

~~~json
{
  "address": "local/openzwave/5/switch/0/$input",
  "attr": {
    "setEnabled": "true",
  },
  "dataType": "bool",
  "timestamp": "2020-01-20T23:33:44.999PST",
}
~~~   


# Publishing Output Values

Publishers monitor the outputs of their nodes and publish updates to node output values when there is a change. Output values are published using various commands depending on the content, as described in the following paragraphs.

>The general output value address is:
>  **\{domain}/\{publisherid}/\{nodeid}/\{type}/\{instance}/\{$messageType}**

| Address segment | Description|
|:--------------- |:-----------|
| {domain}        | The global IoT domain in which publishing takes place, or "local" |
| {publisherid}   | ID of the publisher of the information |
| {nodeid}        | ID of the node that manages the output |
| {type}          | The type of output, for example "temperature". This standard includes a list of output types |
| {instance}      | The instance of the type on the node |
| {$messageType}  | Type of output value publication as described in the following paragraphs: $raw, $latest, ...|

With exception of the \$raw command, all publications contain a payload consisting of a JSON object with the value and additional metadata.


## \$raw: Publish Single 'no frills' Raw Output Value

The payload used with the '\$raw' message type is the pure information as text, without any signature or other metadata. It is the only message without the json message and signature as described above.

The \$raw publication is the fallback that is enabled by default. It can be disabled in the node or output 'publishRaw' configuration. It is intended for interoperability with highly constrained devices or 3rd party software that do not support JSON parsing. The payload is therefore the straight value. 

Address:  **\{domain}/\{publisherid}/\{nodeid}/\{type}/\{instance}/\$raw**

Payload: Output value, converted to string. There is no message JSON and no signature.

Example:
~~~
local/openzwave/6/temperature/0/\$raw: "20.6"
~~~

## \$latest: Publish Latest Output With Metadata

The \$latest publication contains the latest known value of the output including metadata such as the unit and timestamp. The value is represented as a string. Binary data is converted to base64. It is enabled by default and can be disabled with the node or output 'publishLatest' configuration.

This is the recommended publication publishing updates to single value sensors. See also the \$event publication for multiple values that are related. 

Address:  **\{domain}/\{publisherid}/\{nodeid}/\{type}/\{instance}/\$latest**

The message structure is as follows:

| Field        | Data Type | Required     | Description |
|:-------------|:----------|:------------ |:----------- |
| address      | string    | **required** | Address of the publication |
| timestamp    | string    | **required** | timestamp of the value ISO8601 "YYYY-MM-DDTHH:MM:SS.sssTZ" |
| unit         | string    | optional     | unit of value type, if applicable |
| value        | string    | **required** | value in string format |

Example of a publication:

~~~json
{
  "address": "local/openzwave/6/temperature/0/$latest",
  "timestamp": "2020-01-16T15:00:01.000PST",
  "unit": "C",
  "value": "20.6",
}
~~~

## \$forecast: Publish Forecasted Output Values

The payload for the '\$forecast' command contains an ordered list of the projected future values along with address information and signature. The forecast is published each time a value changes. 

Address:  **\{domain}/\{publisherid}/\{nodeid}/\{type}/\{instance}/\$forecast**

The message structure:

| Field        | Data Type | Required     | Description |
|:----------   |:--------  |:-----------  |:------------ |
| address      | string    | **required** | Address of the publication |
| duration     | integer   | optional     | Nr of seconds of forecast
| forecast     | list      | **required** | eg: \[\{"timestamp": "YYYY-MM-DDTHH:MM:SS.sssTZ","value": string}, ...] |
| | timestamp  | string    | ISO8601 "YYYY-MM-DDTHH:MM:SS.sssTZ" |
| | value      | string    | Value in string format using the node's unit |
| timestamp    | string    | **required** | timestamp the forecast was created |
| unit         | string    | optional     | unit of value type |

For example:

~~~json
{
  "address" : "local/openzwave/6/temperature/0/$forecast",
  "duration": "86400",
  "forecast" : [
    {"timestamp": "2020-01-16T16:00:01.000PST", "value" : "20.4" },
    {"timestamp": "2020-01-16T17:00:01.000PST", "value" : "20.6" },
    ...
  ],
  "timestamp": "2020-01-16T15:00:01.000PST",
  "unit": "C",
}
~~~

## \$history: Publish History of Recent Output Values

The payload for the '\$history' command contains an ordered list of the recent values. The history is 
published each time a value changes. The history publication is optional and can be enabled with the 
node or output 'publishHistory' configuration. It is intended for users that like to view a 24 hour trend.
It can also be used to check for missing values in case transport reliability is untrusted. The content 
is not required to persist between publisher restarts.

Address:  **\{domain}/\{publisherid}/\{nodeid}/\{type}/\{instance}/\$history**

The message structure:

| Field        | Data Type | Required     | Description |
|:----------   |:--------  |:-----------  |:------------ |
| address      | string    | **required** | Address of the publication |
| duration     | integer   | optional     | Nr of seconds of history. Default is 24 hours (24*3600 seconds)
| history      | list      | **required** | eg: \[\{"timestamp": "YYYY-MM-DDTHH:MM:SS.sssTZ","value": string}, ...] |
|| timestamp   | string    | ISO8601 "YYYY-MM-DDTHH:MM:SS.sssTZ" |
|| value       | string    | Value in string format using the node's unit |
| timestamp    | string    | **required** | timestamp of the message |
| unit         | string    | optional     | unit of value type |

For example:

~~~json
{
  "address" : "local/openzwave/6/temperature/0/$history",
  "duration": "86400",
  "history" : [
    {"timestamp": "2020-01-16T15:20:01.000PST", "value" : "20.4" },
    {"timestamp": "2020-01-16T15:00:01.000PST", "value" : "20.6" },
    ...
  ],
  "unit": "C",
}
~~~

## \$event: Publish Event With Multiple Output Values

The optional \$event publication indicates the publisher provides multiple output values with the same
timestamp as a single event. This can be used in lieu of publishing output values separately and thus 
reduce bandwidth. It can also be useful to publish multiple values that are highly correlated. 

This is disabled by default but can be enabled with the node and output 'publishEvent' configuration.
The node configuration enables/disables the publication while the output configuration determines if 
the output value is included in the event publication.

The event value can include one, multiple or all node outputs. 

Address:  **\{domain}/\{publisherid}/\{nodeid}/\$event**

The message structure:

| Field        | Data Type | Required     | Description |
|:----------   |:--------  |:-----------  |:------------ |
| address      | string    | **required** | Address of the publication |
| event        | map       | **required** | Map with one or more {output type/instance : value} 
| timestamp    | string    | **required** | timestamp of the event in ISO8601 format |

For Example:

~~~json
{
  "address" : "local/vehicle-1/{nodeid}/$event",
  "event" : [
    {"speed/0": "30.2" },
    {"heading/0": "165" },
    {"rpm/0": "2000" },
    {"odometer/ecu": "2514333222" },
    ...
  ],
  "timestamp": "2020-01-16T15:00:01.000PST",
}
~~~   

## \$batch: Publish Batch With Multiple Events

The optional \$batch publication indicates the publisher provides multiple events. This is intended to 
reduce bandwidth in case for high frequency sampling of multiple values. Consumers must process the events
in the provided order, as if they were sent one at a time.

Address:  **\{domain}/\{publisherid}/\{nodeid}/\$batch**

The message structure:

| Field        | Data Type | Required     | Description |
|:----------   |:--------  |:-----------  |:------------ |
| address      | string    | **required** | Address of the publication |
| batch        | list      | **required** | Time ordered list of events with their timestamp, oldest first and newest last.|
| | timestamp   | string    | timestamp of the event in ISO8601 format "YYYY-MM-DDTHH:MM:SS.sssTZ" |
| | event       | map       | Map with {output type/instance : value} |
| timestamp    | string    | **required** | ISO8601 timestamp this message was created |


# Node Commands

Node commands are send by other publishers to control a node or set one of its inputs. The messages of all commands contain the address of the sender to be able to verify the signature. 

In secured domains, commands are only accepted if the message is encrypted, properly signed, and is sent by publishers whose identity is signed by the DSS.

Additional restrictions can be imposed by limiting updates to specific publishers.

## \$alias: Set Node Alias

A node alias replaces the nodeID in publications of the node, its inputs and outputs. The purpose is to be able to replace a node without requiring changing all its consumers. Once an alias is set, all node publications and commands use the alias as the node ID in their publications. To clear a node alias set an empty alias.

Address:  **\{domain}/\{publisherid}/\{nodeid}/\$alias**

The message structure:

| Field        | Data Type | Required     | Description
|:------------ |:--------- |:----------   |:-----------
| address      | string    | **required** | Address of the publication |
| alias        | string    | **required** | The new node ID|
| sender       | string    | **required** | domain/publisherId of the message |
| timestamp    | string    | **required** | Time this request was created, in ISO8601 format, eg: YYYY-MM-DDTHH:MM:SS.sssTZ. The timezone is the local timezone where the value was published. If a request was received with a newer timestamp, up to the current time, then this request is ignored. |

For example, to change openzwave node 5 to use 'deck':

~~~json
{
  "address" : "local/openzwave/5/$alias",
  "alias": "deck",
  "sender": "local/mrbob",
  "timestamp": "2020-01-02T22:03:03.000PST",
}
~~~


## \$create: Create Node

Some publishers lets users create and delete nodes. For example to add a new ip camera, the ip camera publisher can be told to create a new node for a new camera where nodeId is the new camera ID.

Address:  **\{domain}/\{publisherid}/\{nodeid}/\$create**

The message structure:

| Field        | Data Type | Required     | Description
|:------------ |:--------- |:----------   |:-----------
| address      | string    | **required** | Address of the publication |
| configure    | map       | **required** | key-value pairs for configuration of the node. This is the same content as the config field in the $configure command
| sender       | string    | **required** | domain/publisherId of the message |
| timestamp    | string    | **required** | Time this request was created, in ISO8601 format, eg: YYYY-MM-DDTHH:MM:SS.sssTZ. The timezone is the local timezone where the value was published. If a request was received with a newer timestamp, up to the current time, then this request is ignored. |

For example, to create a new camera node with the ipcam publisher:

~~~json
{
  "address" : "local/ipcam/kelowna-bennet/$create",
  "configure" : { 
    "url":"https://images.drivebc.ca/bchighwaycam/pub/cameras/149.jpg",
    "name": "Kelowna Bennett bridge",
    },
  "sender": "local/mrbob",
  "timestamp": "2020-01-02T22:03:03.000PST",
}
~~~


## \$delete: Delete Node

Publishers that support creation of nodes, also support deleting these nodes. For example to delete an ip camera, the ip camera publisher can be told to delete the camera node.

Address:  **\{domain}/\{publisherid}/\{nodeid}/\$delete**

The message structure:

| Field        | Data Type | Required     | Description
|:------------ |:--------- |:----------   |:-----------
| address      | string    | **required** | Address of the publication |
| sender       | string    | **required** | Address of the publisher |
| timestamp    | string    | **required** | Time this request was created, in ISO8601 format, eg: YYYY-MM-DDTHH:MM:SS.sssTZ. The timezone is the local timezone where the value was published. If a request was received with a newer timestamp, up to the current time, then this request is ignored. | 

For Example, To delete a previously created camera node:

~~~json
{
  "address" : "local/ipcam/kelowna-bennet/$delete",
  "sender": "local/mrbob",
  "timestamp": "2020-01-02T22:03:03.000PST",
}
~~~

## \$set: Set Input Value
Publishers subscribe to receive commands to update the inputs of the node they manage.

Address:  **\{domain}/\{publisherid}/\{nodeid}/\{type}/\{instance}/\$set**

The message structure:

| Field        | Data Type | Required     | Description
|:------------ |:--------- |:----------   |:-----------
| address      | string    | **required** | Address of the publication |
| sender       | string    | **required** | domain/publisherId the sender of this message |
| timestamp    | string    | **required** | Time this request was created, in ISO8601 format, eg: YYYY-MM-DDTHH:MM:SS.sssTZ. The timezone is the local timezone where the value was published. If a request was received with a newer timestamp, up to the current time, then this request is ignored. |
| value        | string    | **required** | The control input value to set |

For Example:

~~~json
{
  "address" : "local/openzwave/6/switch/0/\$set",
  "sender": "local/mrbob",
  "timestamp": "2020-01-02T22:03:03.000PST",
  "value": "true",
}
~~~

# Configuring A Node  

Support for remote configuration of node attributes enables administrators manage devices and services over the message bus. Publishers of node discovery information include the available configurations for the published nodes. These publishers handle the configuration update messages for the nodes they publish. 


Input commands are send by other publishers to control inputs of a node. The messages of all input commands contain the address of the sender to be able to verify the signature.

In secured domains the message must be correctly signed by the sender, be encrypted with the receiver's public key (see section on encryption using JWE), and the sender must have a valid identity signature issued by the DSS.

Additional restrictions can apply to only allow certain senders to update node configurations.

If one of the above verification steps fail then the message is discarded and the request is logged.

Address:  **\{domain}/\{publisherid}/\{nodeid}/\$configure**

Configuration Message structure:

| Field       | type     | required     | Description
|:----------- |:-------- |:------------ |:-----------
| address     | string   | **required** | Address of the publication |
| attr        | map      | **required** | key-value pairs for attributes to configure { key: value, …}. **Only fields that require change should be included**. Existing fields remain unchanged.
| sender      | string   | **required** | Address of the sender node of the message |
| timestamp   | string   | **required** | Time this request was created, in ISO8601 format


Example payload for node configuration:

~~~json
{
  "address" : "local/openzwave/5/$configure",
  "attr": {
    "name": "My new name"
  },
  "sender": "local/mrbob",
  "timestamp": "2020-01-20T23:33:44.999PST",
}
~~~

In this example, the publisher mrbob must first have published its node discovery containing the identity 
attribute signed by the DSS before the message is accepted.

# Message Signing

By default all messages MUST be signed using JWS with compact serialization, except for local and test 
domains where plaintext JSON is allowed. Outside the local and test domains, unsigned messages MUST be 
discarded. In secured domains, the publisher identity itself must be signed by the DSS using ECDSA. 

## Plaintext JSON - unsigned messages

In local and test domains messages can be published in JSON serialized UTF-8 plain text, except for the 
$raw publication whose content is not JSON serialized.

This mode allows inspection of the data directly on the message bus and is interoperable with consumers 
that understand JSON but don't support signatures. It should however only be used in trusted environments.

## Compact JWS JSON Serialization Signing

Compact JWS JSON serialization, serializes a message consists of three parts concatenated and separated 
by a dot '.'. Eg:
 Part 1 consists of the base64url encoded protected header. This header contains the algorithm claim as described in JWS JSON header specification
 Part 2 consists of the base64url encoded payload. 
 Part 3 consists of the JWS signature, which is the base64url encoded encrypted hash of: \<base64url protected header> . \<base64url encoded payload>. 

> Base64URL(UTF8(protected header)) . Base64URL(payload) . Base64URL(JWS Signature)

For Example: 
  "eyJhbGciOiJIUzI1NiJ9.
   SXTigJlzIGEgZGFuZ2Vyb3VzIGJ1c2luZXNzLCBGcm9kbywgZ29pbmcgb3V0IHlvdXIgZG9vci4gWW
   91IHN0ZXAgb250byB0aGUgcm9hZCwgYW5kIGlmIHlvdSBkb24ndCBrZWVwIHlvdXIgZmVldCwgdGhl
   cmXigJlzIG5vIGtub3dpbmcgd2hlcmUgeW91IG1pZ2h0IGJlIHN3ZXB0IG9mZiB0by4.
   bWUSVaxorn7bEF1djytBd0kHv70Ly5pvbomzMWSOr20"

The signature is generated using [ECDSA Elliptic Curve Cryptography](https://blog.cloudflare.com/ecdsa-the-digital-signature-algorithm-of-a-better-internet). Its keys are shorter than RSA, it has not (yet - May 2020) been broken and it is claimed to be [more secure than RSA](https://blog.cloudflare.com/a-relatively-easy-to-understand-primer-on-elliptic-curve-cryptography/).

See the example code for generating and verifying signatures:
* golang: https://github.com/iotdomain/iotd.standard/tree/master/examples/edcsa_text.go. See also the go-jose library.
* python: https://github.com/iotdomain/iotd.standard/tree/master/examples/example.py
* javascript: https://github.com/iotdomain/iotd.standard/tree/master/examples/example.js


## Updating The Publisher Identity

When a publisher re-publishes its own public identity, the signature verification must use the public key contained in the identity to verify the signature. To close the gaping security hole this opens, the identity MUST be correctly signed by the DSS in secured domains. In non-secured domains the message bus must be configured with ACL to only allow the publisher to publish on its identity address.


# Encrypted Messaging - JWE

In secured domains, messages with potentially sensitive content, eg input commands and configuration updates, must be sent encrypted. The sender signs the message with its private key and then encrypt it with the recipient's public key.

To this end the JSON Web Encryption is used. JWE encapsulates the JWS signed message: JWE( JWS(payload, privateKey), publicKey ) and uses compact serialization before publishing. 


# Secured IoT Domains

## Introduction

In a secured domain, publications are made by publishers whose identity can be verified. Protection of 
secured domains consists of rings. Each ring is an independent layer of security. The implementation of
one ring MUST NOT assume the implementation of another ring.

Ring 5 is the environment outside the message bus, eg the internet. This must be treated as hostile. 
Think of the badlands with predetors roaming freely. Connections to the message bus through this environment
MUST be made with TLS and certificate verification enabled to protect against DNS spoofing and man in 
the middle attacks.

Ring 4 is the LAN environment where the message bus resides. This should be considered just as hostile
as the internet as any computer on the LAN that is compromised can mount an attack. A message bus that
runs on a LAN and is accessible via the internet needs proper firewall configuration and should run in
a DMZ separate from the rest of the LAN. If available it runs on its own VLAN to prevent unintended access
to the rest of the LAN. 

Ring 3 protects the message bus server connection. The server must require TLS connections. Clients are 
required to have proper credentials. Security can be further increased with client side certificates, 
support for certificate revocation and frequent credential rotation. To detect suspicious connections, 
connections from clients are logged; Geolocation restrictions of IP addresses are applied; IP block lists 
are applied; Connection frequency restrictions are in place. Monitoring and alerting of suspicious connections
are in place. Basically best practices for any server exposed to the internet.

Ring 2 protects the message bus publish and subscription environment. This ring protects against clients
subscribing or publishing to topics they are not allowed to. The minimum requirement is to differentiate
between clients that subscribe vs clients that can publish. These permissions are granted separately. The
default approach used MUST be of deny access first and grant access as needed. 

Further enhancements are to control for each client which topics they are are allowed to publish to and
subscribe to. For example, the DSS service is the only service allowed to publish on the  DSS address. 
Similarly, publishers should be the only clients allowed to publish discovery and outputs on their address. 

Ring 1 protects the message publications themselves. Each publication MUST be signed by its publisher 
and each publisher identity MUST be verified as signed by a trusted third party. 

This standard defines the 'DSS', domain security service, to act as a trusted party and control the 
message bus configuration. 


## Joining A Secured IoT Domain - DSS

The DSS - Domain Security Service - signs the publisher identities and issues a trusted public/private key pair for signing and encryption. In order to join a secure domain a publisher must be registered with the DSS as a trusted client. 

The process of joining a secured domain:
1. A publisher generates temporary public/private key pair for signing and encryption on first use
2. The publisher publishes its publisher identity message with the temporary public keys
3. The DSS receives the identity message, adds it to the list of unverified publishers and notifies the administrator
4. The administator verifies and optionally updates the publisher identity information and marks the publisher as trusted
5. The DSS publishes the new signed identity information to the publisher along with signing keys to be used in further messaging. This message is encrypted with the publisher's temporary encryption key using JWE. The address is \{domain}/\{publisherid}/\$set
6. The publisher receives the update to its identity and keys, verifies that they came from the DSS, and persists the information securely
7. The publisher publishes its updated identity for consumers using the new signing keys
8. The DSS periodically re-issues new identity and signing keys before they expire

**1: Generating a temporary signing keyset**

Initially the publisher creates their own private and public keyset. The public key is included in the public publisher's identity. By default the keyset is a 256 bit elliptic curve key pair for use with JWS. (JSON Web Signing)

**2: Publish the publisher identity**

**3: Administrator Marks The Publisher As Trusted**

The **DSS** is the Domain Security Service. Its purpose is to issue keys and identity signatures to publishers that have joined the secured domain.

The DSS needs to be told that the publisher with public key X is indeed the publisher with the ID it claims to have. Optionally additional identity can be required such as location, contact email, phone, administrator name and address. 

The method to establish trust can vary based on the situation. The following method is used in the DSS reference implementation. 

1. On activation of a new publisher, the administrator notes the publisher ID and the public key generated by that publisher. The publisher publishes its discovery using the temporary key.

2. Next, the administrator logs in to the DSS service. The service shows a list of untrusted publishers. The administrator verifies if the publisher public key matches his notes. If there is a match, the administrator informs the DSS that the publisher can be trusted. After this step 4 kicks in.

**4: DSS issues new keys and signs identity**

When a publisher status changes from untrusted to trusted, the DSS starts the cycle of key and identity signature renewal as described below.
   
## Renewing Publisher Identity - DSS

Note this only applies for domains that are secured using a Domain Security Service - DSS, that renews publisher's identity.

A publisher that has joined the secured domain is issued a new full identity record that includes a new public/private key pair, and an identity signature signed by the DSS. The publisher will publish its updated public identity containing a new public key.

Receivers of a publisher identity update must verify that the new public identity is legit before accepting it. This is done by ECDSA verification of the public identity, its signature and the DSS public signing key. If it all matches the new publisher can be trusted.

The identity information has a limited lifespan and is updated periodically by the DSS before the expiry date is reached. By default this is half the lifespan of 48 hours. In low bandwidth situations this might be increased to a week or a month. The expiry check is performed by the DSS when a publisher publishing its own node discovery or periodically by the DSS itself. The publisher must persist the newly issued identity information before using the new keys. 

If the DSS has no record of a new publisher its identity is stored for review by the administrator. The administrator must mark the publisher as trusted before it is invited to join the secured domain. 

If a publisher's identity has expired but the dss has not issued an updated identity, then its messages will be discarded by consumers until the DSS has renewed the identity keys. This should be nearly immediate after the publisher publishes its expired identity. This allows for publishers to be offline for a longer period of time without having to reregister with the  secured domain. However, once the new identity key is issued the old one is no longer valid. 


The update message of a full identity record consists of the publisher identity message with the following additions:
* a sender field that must contain the DSS identity address
* a privateKey field that contains a PEM encoded private key for this publisher


~~~json
{
  "address":   "my.domain.org/openzwave/$set",
  "domain":    "my.domain.org",
  "expires":   "2020-01-22T2:33:44.000PST",
  "issuerId":    "$dss",
  "location":  "my location in BC, Canada",
  "organization": "my organization",
  "publicKey": "PEM encoded public key for signature verification and encryption",
  "timestamp": "2020-01-20T23:33:44.999PST",
  "sender":    "mydomain.org/$dss",
  "signature": "base64encoded ECDSA signature of the sender with signature field blank",
  "privateKey": "PEM encoded private key",
  }


~~~   

**Requirements For Updating A Publisher Identity**

The requirements for a publisher to allow its identity to be updated: 

1. The message must be encrypted with JWE using the publisher's current public key (all configuration updates must be encrypted)
2. The message must originate from the publisher's DSS. Eg the sender address is \{domain\}/$dss and the message must be signed by the DSS.
3. The message must not be sent with the retained flag
4. The identity must contain a correct domain and publisherId. (you cannot assign a publisher an identity of another publisher). 
5. The identity timestamp must be newer than the current identity timestamp
6. The message timestamp must be more recent than the previous received message of the DSS publisher (messages that are a replay of an old identity are discarded)
7. The identity signature must be verified with the DSS public signing key.
8. Remote identity updates must be enabled in the publisher.


## Expiring Identity Keys

By default, identity and crypto keys expire after 30 days. The DSS issues new sets of keys and identity signature when 15 days are remaining. These durations can be changed depending on what policy settings.
Once the identity has expired, the administrator must again go through the procecss of joining the publisher to the domain. 


## Verifying Publisher Identity

When a consumer receives a message from the publisher, it needs to verify that the publisher is indeed who it claims to be using the identity signed by the DSS service. Once this verification succeeds the consumer can assume the identity is valid until it expires or a new identity is received. 

The DSS has the reserved publisher ID of '\$dss'. It publishes its own identity just like any other publisher on address '{domain}/\$dss/\$identity'.

## Verifying The DSS Identity

Just like publishers, the DSS has an identity with a signature. There are two methods for ensuring that the DSS identity is valid:

1. Message bus permissions. Only the DSS has the credentials to publish on the dss publisher address. This is the default in local domains. Restricting access to the DSS publisher address using message bus ACLs is highly recommended.

2. Global Certificate. The DSS is published with a certificate signed by a global CA like Lets Encrypt. Subscribers can verify this certificate with the global CA before trusting the DSS. To facilitate the use of global domains, the domain '\{name}.iotd.zone' is available, where \{name} is a globally unique domain.

The domain 'local' is reserved for local-only domains. In this case message bus permissions must secure the DSS publications and no certificate is used. 

# Sharing Information With Other IoT Domains - Domain Bridge Manager (DBM)

While it is useful to share information within an IoT domain, it can be even more useful if some of this information is shared with other domains.

This is the task of a Domain Bridge. Domain Bridges are managed by the Domain Bridge Manager (DBM) publisher. This publisher is responsible for creating and deleting bridge nodes. Note that the local domain cannot be bridged. If multiple IoT domains share the same message bus then no bridge is needed as consumers can already subscribe to the domains. This is a simple way to aggregate information from multiple domains.

To create a bridge the DBM is given the address and login information of a remote domain message bus. The DBM creates a new bridge node for that domain, which is published in both the local and remote domain. 

When a node output is bridged, the bridge instance listens to publications for that output and republishes the message in the remote domain **under its original address**. The signature and content remain unchanged.

By default the bridge publishes the DSS identity from its home domain into the remote domain to enable consumers in the remote domain to verify the bridge publications. 

Members of a domain can discover remote publishers by subscribing to the  +/+/\$identity address. This discovers the publisher identities of all available domains. 


## Managing Bridges

Bridges are managed through the DBM using its web client if available, or through the message bus. Either way is optional and up to the implementation. If managed through the message bus then the  addresses and messages below must be used.

To create a bridge the DBM service must be active in a domain. Publish the following command to create a new bridge:

>  **\{domain}/\$bridge/\{bridgeId}/\$create**

The payload is a signed message with the new bridge node ID. The new bridge node has address: {domain}/\$bridge/{bridgeId}

Message Content:
| Field       | type     | required     | Description |
|:------------|:-------- |:------------ |:----------- |
| address     | string   | **required** | Address of the publication |
| clientId    | string   | optional     | ID of the client to connect as. Must be unique on a message bus. Default is to generate a temporary ID. |
| host        | string   | **required** | IP address or hostname of the remote bus |
| login       | string   | **required** | Login identifier obtained from the administrator |
| password    | string   | **required** | Password to connect with |
| port        | integer  | optional     | port to connect to. Default is determined by protocol |
| protocol    | enum     | optional     | Protocol to use: "MQTT" (default), "REST" |
| sender      | string   | **required** | domain/publisher of the user that configures the bridge. |
| timestamp   | string   | **required** | Time the record is created |

To delete a bridge:
>  **\{domain}/\$bridge/\{bridgeId}/\$delete**

The payload is a signed message:
| Field       | type     | required     | Description |
|:------------|:-------- |:------------ |:----------- |
| address     | string   | **required** | Address of the publication |
| sender      | string   | **required** | Address of the sender, eg: my.domain.org/mrbob of the user that configures the bridge. |
| timestamp   | string   | **required** | Time the record is created |

A bridge can be deleted from the local or the remote domain.

## Bridge Configuration

Using the standard node configuration mechanism, the bridge node is configured with the domain it is bridging to. 

Bridge configuration can be set on address: {domain}/\$bridge/\{bridgeId}/\$configure following the same
 approach as configuration of other nodes.


Bridges support the following configuration attributes:

| Attribute    | value type   | value        | Description
|:------------ |:---------    |:-----------  |:----------
| clientId     | string       | optional     | ID of the client to connect as. Must be unique on a message bus. Default is to generate a temporary ID.
| host         | string       | **required** | IP address or hostname of the remote bus
| login        | string       | **required** | Login identifier obtained from the administrator
| password     | string       | **required** | Password to connect with
| port         | integer      | optional     | port to connect to. Default is determined by protocol
| protocol     | enum         | optional     | Protocol to use: "MQTT" (default), "REST"

## Forward Nodes, Inputs or Outputs

A bridge node has inputs to manage it forwarding a node, specific input or specific output. 

* To forward a node through the bridge, use the following input set command
> **\{domain}/\$bridge/\{bridgeId}/forward/node/$set**

* To forward an input:  
> **\{domain}/\$bridge/\{bridgeId}/forward/input/$set**

* To forward an output:  
> **\{domain}/\$bridge/\{bridgeId}/forward/output/$set**
 
Message structure:

| Field       | type     | required     | Description |
|:------------|:-------- |:------------ |:----------- |
| address     | string   | **required** | Address of the publication |
| forward     | string   | **required** | Address to forward, eg domain/publisherId/nodeId/$node, input or output
| scope       | object   | **required** | The scope of the information to forward |
| - discovery | boolean  | optional     | Forward node/in/output discovery publications, default=true |
| - batch     | boolean  | optional     | Forward output \$batch publication(s), default=true |
| - event     | boolean  | optional     | Forward output \$event publication(s), default=true |
| - forecast  | boolean  | optional     | Forward output \$forecast publication(s), default=true |
| - history   | boolean  | optional     | Forward output \$history publication(s), default=true |
| - latest    | boolean  | optional     | Forward output \$latest publication(s), default=true |
| - raw       | boolean  | optional     | Forward output \$raw publication(s), default=true |
| sender      | string   | **required** | domain/publisher sending the request |
| timestamp   | string   | **required** | Time the record is created |


## Remove Forwarded Nodes, Inputs or Outputs 

To remove a forward, use the following command:

* **\{domain}/\$bridge/\{bridgeId}/remove/node/\$set**
* **\{domain}/\$bridge/\{bridgeId}/remove/input/\$set**
* **\{domain}/\$bridge/\{bridgeId}/remove/output/\$set**


Message structure:

| Field       | type     | required     | Description |
|:----------- |:-------- |:-----------  |:----------- |
| address     | string   | **required** | Address of the publication |
| remove      | string   | **required** | The address to remove |
| sender      | string   | **required** | domain/publisher sending the request |
| timestamp   | string   | **required** | Time the record is created |


# Appendix A: Value DataTypes

The dataType attribute in input and output discovery messages describe what value is expected in publications. The possbible values are:

| DataType         | Description |
|:-------------    |:------------      |
| bool             | value is boolean converted to text: "true", "false" (case insensitive)  |
| json             | value is a json type with multiple fields, converted to JSON (*)|
| enum             | value is one of the strings provided in the enum attributes of the discovery message|
| float            | value is a floating point number converted to text |
| int              | value is a large, 64bit integer number convert to text |
| string           | value is a string |
| jpeg             | value is a base64 encoded jpeg image |
| png              | value is a base64 encoded png image |
| raw              | value is pure raw content |

**Lists** Lists are supported by starting and ending the value with '[' and ']' and separating each value with a comma.

For dataType int, when the value starts with '\[' it should be considered a list of integers instead of a single integer.  If an application expects a non-list value and receives a list, the first item in the list should be used. If an application expects a list and receives a non-list value it should be treated as a list of 1 item.

**json values** Values with the json dataType are a catch-all for storing multiple fields as json payload. It should be avoided when possible as discovery provides no description of the structure. If possible rather use the $event publication that publishes all output values in a single event.

# Appendix B: Node Types

Nodes represent hardware or software services. The node types standardizes on the names of predefined devices or services.

| Key              | Value Description |
|:-------------    |:------------      |
| alarm            | Node is an alarm emitter |
| avControl        | Audio/video controller, eg remote control |
| avReceiver       | Audio/video smart radio/receiver/amp (eg, denon) |
| beacon           | Location beacon|
| button           | Device with one or more buttons |
| camera           | Web or traffic camera |
| clock            | Time clock |
| computer         | General purpose computer |
| dimmer           | Light dimmer |
| ecu              | Engine control unit with loads of sensors |
| gateway          | Gateway for other nodes (onewire, zwave, etc) |
| gps              | GPS location receiver |
| keypad           | Entry key pad |
| lightswitch      | Light switch |
| lock             | Electronic door lock |
| multisensor      | NodDevicee with multiple sensors |
| netRepeater      | Zwave or other network repeater |
| netRouter        | Network router |
| netSwitch        | Network switch |
| netWifiAP        | Network wifi access point |
| onOffSwitch      | General purpose on/off switch |
| powerMeter       | Power or KW meter |
| sensor           | Device with one sensor. See also multisensor. |
| smartlight       | Smart light, eg philips hue |
| tv               | A (not so) smart TV |
| unknown          | Unknown device or service |
| wallpaper        | Wallpaper montage of multiple images |
| waterValve       | Water valve control unit |
| weatherService   | Service providing current and forecasted weather |
| weatherStation   | Weather station with multiple sensors and controls |
| weighScale       | Electronic weigh scale |

# Appendix C: Predefined Node Attributes

Node attributes provide a description of the device or service including its status. If these 
attributes are configurable they are included in the Node Config section.

| Key              | Value Description |
|:---------------  |:----------------- |
| address          | Node internal address if applicable. Can be used as the node ID |
| color            | color in hex notation |
| description      | device description |
| disabled         | device or sensor is disabled |
| filename         | filename to write images or other values |
| gatewayAddress   | the node gateway address |
| hostname         | network device hostname |
| iotcVersion      | Publishers include the version of the IoTDomain standard. Eg v1.0 |
| localIP          | IP address of the node, for nodes that are publishers themselves |
| latlon           | String with "{latitude}, {longitude}" of device location  |
| locationName     | Name of a location |
| mac              | Node MAC address for nodes that have an IP interface |
| manufacturer     | Device manufacturer |
| max              | maximum value of sensor or config |
| min              | minimum value of sensor or config |
| model            | Device model |
| name             | name of device, sensor |
| netmask          | IP network mask |
| password         | password to connect. Value is not published |
| pollInterval     | polling interval in seconds |
| powerSource      | battery, usb, mains |
| product          | device product or model name |
| softwareVersion  | Software/Firmware identifier or version |
| subnet           | IP subnets configuration |

Node status attributes. These convey the current state of the node and are read-only

| Key           | Value Description |
|:------------  |:------------      |
| errorCount    | nr of errors reported on this device |
| health        | health status of the device 0-100% |
| lastError     | most recent error message |
| lastSeen      | ISO time the device was last seen |
| latencymsec   | duration connect to sensor in milliseconds |
| neighborCount | mesh network nr of neighbors |
| neighborIDs   | mesh network device neighbors ID list [id,id,...] |
| rxCount       | Nr of messages received from device |
| txCount       | Nr of messages send to device |
| runState      | Node runtime state. See the runstate attributes below |

Node 'runstate' attribute values
| Key           | Value Description |
|:------------  |:------------      |
| error         | Node needs servicing |
| disconnected  | Node has cleanly disconnected |
| failed        | Node failed to start |
| initializing  | Node is initializing |
| ready         | Node is ready for use |
| sleeping      | Node has gone into sleep mode, often a battery powered device |


# Appendix D: Predefined Configuration Names

Standard configuration attribute names

| Name          | Value Description |
|:------------- |:------------      |
| ip4           | Device static IP-4 address |
| ip6           | Device static IP-6 address |
| locationName  | Device location name |
| loginName     | login name |
| name          | Device friendly name |
| netmask       | Network netmask |


# Appendix E: Input and Output Types

When available, units used in publication follow the SI standard 
The value content is converted to text before publication.

| input/output type| Units  | Value DataType   | description |
|:--------------  |:--------|:-----------------|:------------|
| acceleration    | m/s2    | List of floats   | coordinates: x,y,z
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
| current         | A       | float     | Electrical current in Amps
| dewpoint        | C, F    | float     | Dewpoint in degrees Celcius
| distance        | m, yd, ft, in | float | distance in meters
| dimmer          | %       | integer   |
| doorwindowsensor|         | boolean   |
| duration        | sec     | float     |
| electricfield   | V/m     | float     | Static electric field in volt per meter
| elevation       | m, ft   | float     | elevation in meters or feet
| energy          | KWh     | float     |
| errors          |         | integer   |
| freememory      | %       | integer   | Relative memory available in percent
| fuel            | L, G    | float     | Amount of fuel in Liters or Gallons
| heatindex       | C, F    | float     | Apparent temperature (humiture) based on air temperature and  relative humidity. Typically used when higher than the air temperature. At 20% relative humidity the heatindex is equal to the temperature.
| heading         | Degrees | float     | Compass heading in 0-360 degrees 
| hue             |         |           |
| humidex         | C       | float     | Humidity temperature index (feels like temperature) derived from    dewpoint, in degrees Celcius 
| humidity        | %       | float     | Relative humidity in %
| image           | jpeg, png | bytes   | Image in jpeg or png format
| latency         | sec     | float     | 
| location        |         | List of 3 floats | \[latitude, longitude, elevation]
| locked          |         | boolean   | Door lock status 
| luminance       | cd/m2, lux | float  | Amount of light in candela/m2 or in lux
| magneticfield   | T, mT, uT, nT, G(auss) | float | Static magnetic field in Tesla or Gauss
| motion          |         | boolean   | Motion detection status |
| power           | W       | float     | Power consumption in watts |
| pressure        | kpa, Psi| float     | gas or liquid pressure
| pushbutton      |         | boolean   | Momentary pushbutton |
| signalstrength  | dBm     | float     | Wireless signal strength
| speed           | m/s, Km/h, mph | float |
| switch          |         | boolean   |
| temperature     | C, F    | float     | Celcius or Fahrenheit. The default is Celcius if available.
| ultraviolet     | UV      | float     | Radiation index with wavelength from 10 nm to 400 nm, range 0-11+
| voltage         | V       | float     | Volts
| volume          | L, G    | float     | Volume in liters or gallons
| waterlevel      | m(eters), ft, in | float |
| wavelength      | m       | float     |
| weight          | Kg, Lbs | float     | 
| windchill       | C, F    | float     | Apparent temperature based on the air temperature and wind speed,    when lower than the air temperature.
| windspeed       | m/s, km/h, mph | float | 

