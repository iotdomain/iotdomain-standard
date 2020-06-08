The IotConnect Standard
=======================

The IotConnect standard defines a simple and easy to use information exchange method between IoT publishers and consumers.


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

Addresses can only contain of alphanumeric, hyphen (-), and underscore (\_) characters. Reserved words start with a dollar ($) character. The separator is the '/' character. All addresses end with a message type indicating the content of the message.

> Address format:
>  **\{domain} / \{publisherId} / \[ \{nodeId} \[ / \{ioType} / \{instance} \] \] / \{messagetype\}**

Where:
* \{domain} is the domain in which the node information is published. This can be "local" or an internet domain.
* \{publisherId} is the ID of the service that publishes the node, input and output information
* \{nodeId} is the ID of the node that is being published
* \{ioType} and \{instance} refers to a particular input or output of the node. 
* \{messageType} indicates the content of the message, be it publishing of discovery or values.
For message bus systems that do not support the '/' character as address separator, the separator character of the message bus implementation can be used. However, the message itself must contain the original address using the '/' character as the separator to allow for interoperability between different message bus implementations.

**Reserved message types:**

The standard predefines the following message types.

| Publisher message type | Purpose |
|:--------     |:--------|
| \$identity   | Publication of a publisher's identity (domain/publisherId/\$identity)


| Node message type | Purpose |
|:--------     |:--------|
| \$batch      | Publication of a batch of events |
| \$configure  | Message to update the node configuration |
| \$create     | Message to create a node. Only usable with publishers that can create/delete nodes |
| \$delete     | Message to delete a node. Only usable with publishers that can create/delete nodes |
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

Node Aliases are intended to support replacing devices while retaining the input and output addresses. When devices are replaced, the node identifier of the replacement can differ from the original. ZWave for example generates a new node ID each time a node is added to the network. This leads to the problem that when replacing a node, all consumers must be updated to use the replacement node ID, which can take quite a bit of effort. 

To address this problem, nodes can be configured with an 'alias' ID. When a node alias is set, all input and output publications use the alias instead of the node ID.

The node alias can be set through the node $configure command. Support for node aliases is optional and implemented in the publisher. If the node configuration does not have an 'alias' configuration option then it is not supported.

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

Publisher discovery contains the publisher's identity. The identity contains public signing key used to verify the signature of the messages they publish. The process of signing uses JWS as described in the security section. Messages that fail signature verification MUST be discarded. 

The identity also contains a public encryption key used to send encrypted messages to the publisher or one of its nodes. The process of encryption uses JWE as described in the security section.

Publisher discovery:

  >  **\{domain}/\{publisher}/\$identity**

Message structure

| Field           | type | Description |
|:--------------- |:-------  | ------ |
| address         | string   | **required** | The address of the publication
| identity          | Identity | Public identity record
| | certificate     | Optional x509 certificate, base64 encoded. Included with the ZSS service to be able to verify its identity with a 3rd party. |
| | domain          | IoT domain name as used in the address. "local" or "test" for local domains |
| | issuerName      | Name of issuer, usually this is "ZSS", The ZSS includes the CA such as LetsEncrypt here. |
| | location        | Optional location of the publisher, city, province/state, country |
| | organization    | Organization the publisher belongs to |
| | publicCryptoKey | Base64 encoded public key for encrypting messages to the publisher | 
| | publicSigningKey| Base64 encoded public key for verifying publisher signatures |
| | publisherId     | ID of this publisher
| | timestamp       | Time the identity was signed |
| | validUntil      | ISO8601 Date this identity is valid until |
| signature        | string | base64 encoded signature of the public identity record
| signer           | string | Name of the signer, either 'DSS' or a CA such as Lets Encrypt
| timestamp        | Time this message was created |

## \$lwt: Publisher Last Will & Testament (MQTT)

This message only applies when using a message bus that supports LWT (last will & testament) .

The LWT option lets the message bus publish a message when the publisher connection is lost unexpectedly. A message with status "connected" and "disconnected" is sent by the publisher when connecting or gracefully disconnecting. The status "lost" is set through last will & testament feature and send by the message bus if the publisher unexpectedly disconnects. 

Address:  **\{domain}/\{publisherId}/\$lwt**

Message structure:
The message structure:

| Field        | Data Type | Required     | Description |
|:----------   |:--------  |:-----------  |:------------ |
| address      | string    | **required** | Address of the publication |
| status       | string    | **required** | LWT status: "connected", "disconnected", "lost"


## Discover Nodes

Node discovery messages contain a detailed description of the node. It does not contain information on inputs and outputs as these are published separately. This reduces the amount of traffic for simple consumers that only subscribe to certain types of inputs or outputs.

Node discovery address:

  >  **\{domain}/\{publisherId}/\{nodeId}/\$node**

Where:
* {domain} is the IoT domain in which the node lives
* {publisherId} is the ID of the publisher of the information. The publisher Id is unique within its domain
* {nodeId} is the ID of the node that is discovered. This is a device or a service identifier and unique within a publisher. 
* $node message type for node discovery

Node discovery message structure:

| Field        | Data Type | Required     | Description
|:-----------  |:--------- |:----------   |:------------
| address      | string    | **required** | The address of the publication
| attr         | map       | **required** | Key value pairs describing the node. The list of predefined attribute keys are part of the standard. See appendix B: Predefined Node Attributes. |
| config       | List of **Configuration Records** | optional | Node configuration, if any exist. Set of configuration objects that describe the attributes that are configurable. The attribute value can be set with a ‘$configure’ message based on the configuration description.|
| nodeId       | string    | **required** | Immutable ID of this node
| publisherId  | string    | **required** | Publisher managing this node
| status       | map       | optional     | key-value pairs describing node performance status
| timestamp    | string    | **required** | Time the record is created |

**Configuration Record**

The configuration record describes the node attributes that can be configured. It includes datatype, and various constraints of the configuration:

| Field    | Data Type | Required  | Description |
|:-------- |:--------- |:--------- |:----------- |
| name     | string    | **required** | Name of the attribute as used in the attr section. See also Appendix C: Predefined Configuration Names |
| datatype | enum      | optional| Type of value. Used to determine the editor to use for the value. One of: bool, enum, float, int, string. Default is ‘string’ |
| default  | string    | optional| Default value for this configuration in string format |
| description| string  | optional | Description of the configuration for human use |
| enum     | \[strings] | optional* | List of valid enum values as strings. Required when datatype is enum |
| max      | float     | optional | Optional maximum value for numeric data |
| min      | float     | optional | Optional minimum value for numeric data | 
| secret   | bool      | optional | Optional flag that the configuration value is secret and will be left empty. When a secret configuration is set in \$configure, the value is encrypted with the publisher node public key. |

Example payload for node discovery

~~~json
local/openzwave/5/\$node:
{
  "address": "local/openzwave/5/$node",
  
  "attr": {
    "make": "AeoTec",
    "type": "multisensor",
    "name": ""
  },
  "config": {
    "alias": {
      "alias": "deck",
      "datatype": "string",
      "description": "Friendly name",
    }, 
  },
  "timestamp": "2020-01-20T23:33:44.999PST",
}
~~~

## Discover Inputs and Outputs

Inputs and outputs discovery are published separately from the node. The discovery of each output and each input is published separately. This facilitates control over which inputs and outputs are shared with other domains. 

Address of input discovery:

> **\{domain}/\{publisherId}/\{nodeId|alias}/\{inputType}/\{instance}/\$input/**

Address of output discovery:

> **\{domain}/\{publisherId}/\{nodeId|alias}/\{outputType}/\{instance}/\$output**

| Address segment | Description |
| :-------------- | :---------- |
| {domain}        | The IoT domain in which the node lives, or "local" for local domains |
| {publisherId}   | The service that is publishing the information |
| {nodeId|alias}  | ID or alias of the node that owns the input or output |
| {inputType}     | Type identifier of the input. For a list of predefined types see Appendix D |
| {outputType}    | Type identifier of the output. For a list of predefined types see Appendix D |
| {instance}      | The instance of the input or output on the node. If only a single instance exists the standard is to use 0 unless a name is used to provide more meaning|
| \$input         | Message type for input discovery, or |
| \$output        | Message type for output discovery |

For example, the discovery of a temperature sensor on node '5', published by a service named 'openzwave', is published on address:

  > **local/openzwave/5/temperature/0/\$output**

The message structure:

| Field       | Data Type | Required     | Description |
|:----------- |:--------- |:---------    |:----------- |
| address     | string    | **required** | Address of the publication |
| attr        | map       | optional     | attributes describing the output |
| config      | List of **Configuration Records**|optional|List of Configuration Records of attributes that can be configured. Only available when an input or output has defined their own configuration |
| datatype    | string    | optional     | Value datatype. See appending for datatypes, default is string
| description | string    | optional     | Description of the in/output for humans |
| enumValues  | list      | optional*    | List of possible values. Required when datatype is enum |
| instance    | string    | **required** | Output instance for for multi-I/O nodes |
| max         | number    | optional     | Maximum possible in/output value |
| min         | number    | optional     | Minimum possible in/output value |
| timestamp   | string    | **required** | Time the record is created |
| outputType  | string    | **required** | Type of output. See the output type list for standardized type names |
| unit        | string    | optional     | The unit of the data type |


Example payload for output discovery:

~~~json
local/openzwave/5/\$output/temperature/0:
{
  "address": "local/openzwave/5/temperature/0/$output",
  "datatype": "float",
  "instance": "0",
  "timestamp": "2020-01-20T23:33:44.999PST",
  "outputType": "temperature",
  "unit": "C",
}
~~~   

# Publishing Output Values

Publishers monitor the outputs of their nodes and publish updates to node output values when there is a change. Output values are published using various commands depending on the content, as described in the following paragraphs.

>The general output value address is:
>  **\{domain}/\{publisherId}/\{nodeId|alias}/\{type}/\{instance}/\{$messageType}**

| Address segment | Description|
|:--------------- |:-----------|
| {domain}        | The global IoT domain in which publishing takes place, or "local" |
| {publisherId}   | ID of the publisher of the information |
| {nodeId|alias}  | ID or alias of the node that owns the input or output |
| {type}          | The type of output, for example "temperature". This standard includes a list of output types |
| {instance}      | The instance of the type on the node |
| {$messageType}  | Type of output value publication as described in the following paragraphs: $raw, $latest, ...|

With exception of the \$raw command, all publications contain a payload consisting of a JSON object with the value and additional metadata.


## \$raw: Publish Single 'no frills' Raw Output Value

The payload used with the '\$raw' message type is the pure information as text, without any signature or other metadata. It is the only message without the json message and signature as described above.

The \$raw publication is the fallback that every publisher *MUST* publish. It is intended for interoperability with highly constrained devices or 3rd party software that do not support JSON parsing. The payload is therefore the straight value.

Address:  **\{domain}/\{publisherId}/\{nodeId|alias}/\{type}/\{instance}/\$raw**

Payload: Output value, converted to string. There is no message JSON and no signature.

Example:
~~~
local/openzwave/6/temperature/0/\$raw: "20.6"
~~~

## \$latest: Publish Latest Output With Metadata

The \$latest publication contains the latest known value of the output including metadata such as the unit and timestamp. The value is represented as a string. Binary data is converted to base64. 

This is the recommended publication publishing updates to single value sensors. See also the \$event publication for multiple values that are related. 

Address:  **\{domain}/\{publisherId}/\{nodeId|alias}/\{type}/\{instance}/\$latest**

The message structure is as follows:

| Field        | Data Type | Required     | Description |
|:-------------|:----------|:------------ |:----------- |
| address      | string    | **required** | Address of the publication |
| timestamp    | string    | **required** | timestamp of the value ISO8601 "YYYY-MM-DDTHH:MM:SS.sssTZ" |
| unit         | string    | optional     | unit of value type, if applicable |
| value        | string    | **required** | value in string format |

Example of a publication on local/openzwave/6/\$latest/temperature/0:

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

Address:  **\{domain}/\{publisherId}/\{nodeId|alias}/\{type}/\{instance}/\$forecast**

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
local/openzwave/6/$forecast/temperature/0:
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

The payload for the '\$history' command contains an ordered list of the recent values. The history is published each time a value changes. The history publication is optional and intended for users that like to view a 24 hour trend. It can also be used to check for missing values in case transport reliability is untrusted. The content is not required to persist between publisher restarts.

Address:  **\{domain}/\{publisherId}/\{nodeId|alias}/\{type}/\{instance}/\$history**

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
local/openzwave/6/temperature/0/$history:
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

The optional \$event publication indicates the publisher provides multiple output values with the same timestamp as a single event. This can be used in lieu of publishing output values separately and thus reduce bandwidth. It can also be useful to publish multiple values that are highly correlated. 

The event value can include one, multiple or all node outputs. 

Address:  **\{domain}/\{publisherId}/\{nodeId|alias}/\$event**

The message structure:

| Field        | Data Type | Required     | Description |
|:----------   |:--------  |:-----------  |:------------ |
| address      | string    | **required** | Address of the publication |
| event        | map       | **required** | Map with one or more {output type/instance : value} 
| timestamp    | string    | **required** | timestamp of the event in ISO8601 format |

For Example:

~~~json
local/vehicle-1/\{nodeId}/\$event:
{
  "address" : "local/vehicle-1/\{nodeId}/$event",
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

The optional \$batch publication indicates the publisher provides multiple events. This is intended to reduce bandwidth in case for high frequency sampling of multiple values. Consumers must process the events in the provided order, as if they were sent one at a time.

Address:  **\{domain}/\{publisherId}/\{nodeId}/\$batch**

The message structure:

| Field        | Data Type | Required     | Description |
|:----------   |:--------  |:-----------  |:------------ |
| address      | string    | **required** | Address of the publication |
| batch        | list      | **required** | Time ordered list of events with their timestamp, oldest first and newest last.|
| | timestamp   | string    | timestamp of the event in ISO8601 format "YYYY-MM-DDTHH:MM:SS.sssTZ" |
| | event       | map       | Map with {output type/instance : value} |
| timestamp    | string    | **required** | ISO8601 timestamp this message was created |


# Input Commands

Input commands are send by other publishers to provide input to a node. The messages of all input commands contain the address of the sender. 

In secured domains, only publishers that have joined the secure domain and provide a valid signature are allowed to send input commands. Receivers can verify the message signature with the sender's public key, provided with its discovery message. If this verification fails then the input command must be ignored.

Additional restrictions can be imposed by limiting updates to specific publishers.

## \$set: Set Input Value
Publishers subscribe to receive commands to update the inputs of the node they manage.

Address:  **\{domain}/\{publisher}/\{node}/\{type}/\{instance}/\$set**

The message structure:

| Field        | Data Type | Required      | Description
|:------------ |:--------- |:----------    |:-----------
| address      | string    | **required** | Address of the publication |
| timestamp    | string    | **required**  | Time this request was created, in ISO8601 format, eg: YYYY-MM-DDTHH:MM:SS.sssTZ. The timezone is the local timezone where the value was published. If a request was received with a newer timestamp, up to the current time, then this request is ignored. |
| sender       | string    | **required** | Address of the publisher of the message (domain/publisherId) |
| value        | string    | **required** | The control input value to set |

For Example:

~~~json
local/openzwave/6/switch/0/\$set:
{
  "address" : "local/openzwave/6/switch/0/\$set",
  "sender": "local/mrbob",
  "timestamp": "2020-01-02T22:03:03.000PST",
  "value": "true",
}
~~~

## \$create: Create Node

Publishers where users can create and delete nodes subscribe to this message type. For example to add a new ip camera, the ip camera publisher can be told to create a new node for a new camera where nodeId is the new camera ID.

Address:  **\{domain}/\{publisherId}/\{nodeId}/\$create**

The message structure:

| Field        | Data Type | Required      | Description
|:------------ |:--------- |:----------    |:-----------
| address      | string    | **required**  | Address of the publication |
| configure    | map       | **required**  | key-value pairs for configuration of the node. This is the same content as the config field in the $configure command
| sender       | string    | **required** | Address of the publisher of the message (domain/publisherId) |
| timestamp    | string    | **required**  | Time this request was created, in ISO8601 format, eg: YYYY-MM-DDTHH:MM:SS.sssTZ. The timezone is the local timezone where the value was published. If a request was received with a newer timestamp, up to the current time, then this request is ignored. | nodeID       | string    | **required** | ID of the node to create. Must be unique within the publisher. For example the camera name |

For Example, To create a new camera node with the ipcam publisher:

~~~json
local/ipcam/Bennet-Bridge/\$create:
{
  "address" : "local/ipcam/Bennet-Bridge/\$create",
  "configure" : { 
    "url":"https://images.drivebc.ca/bchighwaycam/pub/cameras/149.jpg",
    "name": "Kelowna Bennett bridge",
    },
  "sender": "local/mrbob",
  "timestamp": "2020-01-02T22:03:03.000PST",
}
~~~


## \$delete: Delete Node

Publishers that support creation and deletion of nodes, subscribe to this message type. For example to delete an ip camera, the ip camera publisher can be told to delete the camera node.

Address:  **\{domain}/\{publisherId}/\{nodeId}/\$delete**

The message structure:

| Field        | Data Type | Required      | Description
|:------------ |:--------- |:----------    |:-----------
| address      | string    | **required**  | Address of the publication |
| sender       | string    | **required**  | Address of the publisher |
| timestamp    | string    | **required**  | Time this request was created, in ISO8601 format, eg: YYYY-MM-DDTHH:MM:SS.sssTZ. The timezone is the local timezone where the value was published. If a request was received with a newer timestamp, up to the current time, then this request is ignored. | 

For Example, To delete a previously created camera node:

~~~json
local/ipcam/Kelowna-Bennet/$delete:
{
  "address" : "local/ipcam/Kelowna-Bennet/$delete",
  "sender": "local/mrbob",
  "timestamp": "2020-01-02T22:03:03.000PST",
}
~~~

# Configuring A Node  

Support for remote configuration of node attributes lets administrators manage devices and services over the message bus. Publishers of node discovery information include the available configurations for the published nodes. These publishers handle the configuration update messages for the nodes they publish. 

In secured domains, the following requirements apply before configuration updates are accepted:
1. The message must be correctly signed by the sender (this goes for all messages in secured domains)
2. The message must be encrypted with the receiver's public crypto key (see section on encryption using JWE)
3. The sender must be a publisher that has joined the secure domain. Eg it must have a valid identity signature issued by the ZCAS.
4. The sender must be allowed to update node configuration. 

If one of the above verification steps fail then the message is discarded and the request is logged.

Address:  **\{domain}/\{publisherId}/\{nodeId}/\$configure**

Configuration Message structure:

| Field        | type     | required     | Description
|:------------ |:-------- |:------------ |:-----------
| address      | string   | **required** | Address of the publication |
| attr         | map      | **required** | key-value pairs for attributes to configure { key: value, …}. **Only fields that require change should be included**. Existing fields remain unchanged.
| sender       | string   | **required** | Address of the sender node of the message |
| timestamp    | string   | **required** | Time this request was created, in ISO8601 format


Example payload for node configuration:

~~~json
local/openzwave/5/\$configure:
{
  "address": "local/openzwave/5/$configure",
  "attr": {
    "name": "My new name"
  },
  "sender": "local/mrbob",
  "timestamp": "2020-01-20T23:33:44.999PST",
}
~~~

In this example, the publisher mrbob must first have published its node discovery containing the identity attribute signed by the ZCAS before the message is accepted.

# Signing Messages Using JWS

Messages can be sent unsigned using plain text JSON and signed using flattened or compact JWS JSON serialization. It is still possible to publish plain text JSON messages but these messages MUST be discarded by all publishers that have joined the secured domain.

## Plain text JSON - unsigned messages

In plain text JSON mode the messages are published as described in this standard and are unsigned. Messages are always in plain text (UTF 8) JSON except for the $raw publication.

This mode allows inspection of the data directly on the message bus and is interoperable with consumers that understand JSON but don't support signatures. It should however only be used in trusted environments.

## Flattened JWS JSON Serialization Signing

In flattened JWS mode, messages are digitally signed and published using flattened JSON serialization. The message signature guarantees that it hasn't been tampered with. See [RFC7515](https://www.rfc-editor.org/rfc/rfc7515.txt) for details.

The flattened JWS JSON serialization syntax is a JSON message with three fields. The protected header is used to set claims for the issuer and the subject address:

```json
{
  "protected": "<integrity protected header contents>",
  "payload": "<base64url encoded payload contents>",
  "signature": "<base64url encoded digital signature of the content>"
}
```

Where "protected" is the base64 encoded protected header, containing the alg claim: {"alg":"ES256"}

for example: 
```json
{
  "protected": "eyJhbGciOiJIUzI1NiJ9",
  "payload": "SXTigJlzIGEgZGFuZ2Vyb3VzIGJ1c2luZXNzLCBGcm9kbywg
      Z29pbmcgb3V0IHlvdXIgZG9vci4gWW91IHN0ZXAgb250byB0aGUgcm9h
      ZCwgYW5kIGlmIHlvdSBkb24ndCBrZWVwIHlvdXIgZmVldCwgdGhlcmXi
      gJlzIG5vIGtub3dpbmcgd2hlcmUgeW91IG1pZ2h0IGJlIHN3ZXB0IG9m
      ZiB0by4",
  "signature": "bWUSVaxorn7bEF1djytBd0kHv70Ly5pvbomzMWSOr20"
}
```

In the above example, the protected header, payload and signature are all base64url encoded before publication of the message. 

1. The protect header contains the algorithm claim as described in JWS JSON header specification:
2. The payload is the base64url encoded message content in JSON or raw format.
3. The signature is the base64url encoded encrypted hash of: \<base64url protected header> . \<base64url encoded payload>. 


## Compact JWS JSON Serialization Signing

Compact JWS JSON serializations is similar to flattened JSON serialization except that instead of using a JSON object to describe the parts, the parts are simply concatenated and separated by a dot '.'. Eg:

> Base64URL(UTF8(protected header)) . Base64URL(payload) . Base64URL(JWS Signature)

For Example: 
  "eyJhbGciOiJIUzI1NiJ9.
   SXTigJlzIGEgZGFuZ2Vyb3VzIGJ1c2luZXNzLCBGcm9kbywgZ29pbmcgb3V0IHlvdXIgZG9vci4gWW
   91IHN0ZXAgb250byB0aGUgcm9hZCwgYW5kIGlmIHlvdSBkb24ndCBrZWVwIHlvdXIgZmVldCwgdGhl
   cmXigJlzIG5vIGtub3dpbmcgd2hlcmUgeW91IG1pZ2h0IGJlIHN3ZXB0IG9mZiB0by4.
   bWUSVaxorn7bEF1djytBd0kHv70Ly5pvbomzMWSOr20"

When combined with the $raw publications it is the smallest publication possible within this spec as the payload is the base64url encoded raw value.

## Creating A Signature

The JWS header describes the algorithm used to generate the signature. Options are ECDSA Elliptic Curve, HMAC or RSA encryption, using SHA-256, 384 or 512 hash algorithms. See https://openid.net/specs/draft-jones-json-web-signature-04.html#RFC3275 for details.

The preferred method of signing messages is [ECDSA Elliptic Curve Cryptography](https://blog.cloudflare.com/ecdsa-the-digital-signature-algorithm-of-a-better-internet). Its keys are shorter than RSA, it has not (yet - May 2020) been broken and it is claimed to be [more secure than RSA](https://blog.cloudflare.com/a-relatively-easy-to-understand-primer-on-elliptic-curve-cryptography/).

The steps to create a signature are:
1. Base64url encode the protected header
2. Base64url encode the payload
3. Combine 1) and 2) separated with a dot
4. Create a hash of 3) using SHA256. 
5. Create the signature by encrypting the hash using ECDSA with the publisher's private signing key 
6. base64url encode the resulting signature and provide it in the 'signature' field.

See the example code:
* golang: https://github.com/hspaay/iotconnect.standard/tree/master/examples/edcsa_text.go
* python: https://github.com/hspaay/iotconnect.standard/tree/master/examples/example.py
* javascript: https://github.com/hspaay/iotconnect.standard/tree/master/examples/example.js


## Verifying A Message Signature

Consumers verify a signature by:
1. Base64url decode the protected header, payload and the signature
2. Determine the publisher of the message from the publication address. 
3. Determine the public signing key of the publisher from its discovery information.
4. Determine the hash of the encoded header and payload concatenated with a dot
5. ECDSA verify the signature using the hash and public key

See the examples for more detail.

If the verification fails, the message content is discarded.

When a publisher has not yet been discovered, its signature cannot be verified as they are from an unknown publisher. When a publisher discovery message is received it contains a public key that will be used for future signature verification. 

In non-secured domains a publisher discovery is accepted when its signature can be verified against the identity in the discovery message. This is akin to believing someone on his word and of limited value unless additional measures are in place.

The purpose of secured domains is to provide a publisher identity verification method, through topic level security or identity verification with a third party. See the ZCAS for more information.


# Encrypted Messaging - JWE

Messages with potentially sensitive content, eg input commands and configuration updates, must be sent encrypted using the public crypto key of the intended recipient. To this end the JSON Web Encryption is used.

Just like signing, JWE supports compact serialization and JSON serialization

.. todo ..

# Secured IoT Domains


## Introduction

In a secured domain publications are made by publishers whose identity can be verified. Protection of secured domains consists of rings. Each ring is an independent layer of security. The implementation of one ring MUST NOT assume the implementation of another ring.

Ring 5 is the environment outside the message bus, eg the internet. This must be treated as hostile. Think of the badlands with predetors roaming freely. Connections to the message bus through this environment MUST be made with TLS and certificate verification enabled to protect against DNS spoofing and man in the middle attacks.

Ring 4 is the LAN environment where the message bus resides. This should be considered just as hostile as the internet as any computer on the LAN that is compromised can mount an attack. A message bus that runs on a LAN and is accessible via the internet needs proper firewall configuration and should run in a DMZ separate from the rest of the LAN. If available it runs on its own VLAN to prevent unintended access to the rest of the LAN. 

Ring 3 protects the message bus server connection. The server must require TLS connections. Clients are required to have proper credentials. Security can be further increased with client side certificates, support for certificate revocation and frequent credential rotation. To detect suspicious connections, connections from clients are logged; Geolocation restrictions of IP addresses are applied; IP block lists are applied; Connection frequency restrictions are in place. Monitoring and alerting of suspicious connections are in place. Basically best practices for any server exposed to the internet.

Ring 2 protects the message bus publish and subscription environment. This ring protects against clients subscribing or publishing to topics they are not allowed to. The minimum requirement is to differentiate between clients that subscribe vs clients that can publish. These permissions are granted separately. The default approach used MUST be of deny access first and grant access as needed. 

Further enhancements are to control for each client which topics they are are allowed to publish to and subscribe to. For example, the DSS service is the only service allowed to publish on the  DSS address. Similarly, publishers should be the only clients allowed to publish discovery and outputs on their address. 

Ring 1 protects the message publications themselves. Each publication MUST be signed by its publisher and each publisher identity MUST be verified as signed by a trusted third party. 

This standard defines the 'DSS', domain security service, to act as a trusted party and control the message bus configuration. 


## Joining A Secured IoT Domain - DSS

The DSS - Domain Security Service - is a ring 1 service that signs publisher identities and issues signing and encryption keys. In order to join a secure domain a publisher must be registered with the DSS as a trusted client.

The process of joining a secured domain:
1. A publisher generates temporary public/private keys for signing and encryption on first use
2. The publisher publishes its publisher identity message with the temporary public keys
3. The DSS receives the identity message, adds it to the list of unverified publishers and notifies the administrator
4. The administator verifies and optionally updates the publisher identity information and marks the publisher as trusted
5. The DSS publishes the new signed identity information to the publisher along with signing keys to be used in further messaging. This message is encrypted with the publisher's temporary encryption key using JWE. The address is \{domain}/\{publisherId}/\$set
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

A publisher that has joined the secured domain is issued a new identity record that includes signing and encryption keys, and an identity signature signed by the DSS. Consumers of messages from this publisher can verify that the publisher is legit by verifying the identity signature with the DSS public signing key. This check is done by consumers each time a publisher publishes an updated identity.

The identity information has a limited lifespan and is updated periodically by the DSS before the expiry date is reached. By default this is half the lifespan of 48 hours. In low bandwidth situations this might be increased to a week or a month. The expiry check is performed by the DSS when a publisher publishing its own node discovery or periodically by the DSS itself. The publisher must persist the newly issued identity information before using the new keys. 

If the DSS has no record of a new publisher its identity is stored for review by the administrator. The administrator must mark the publisher as trusted before it is invited to join the secured domain. 

If a publisher's identity has expired but the dss has not issued an updated identity, then its messages will be discarded by consumers until the DSS has renewed the identity keys. This should be nearly immediate after the publisher publishes its expired identity. This allows for publishers to be offline for a longer period of time without having to reregister with the  secured domain. However, once the new identity key is issued the old one is no longer valid. 


Example Identity Update Message:


~~~json
my.domain.org/openzwave/\$set:
{
  "address": "my.domain.org/openzwave/\$set",
  "public" : {
    "domain": "my.domain.org",
    "expires":  "2020-01-22T2:33:44.000PST",
    "issuerName": "DSS",
    "location":   "my location in BC, Canada",
    "organization": "my organization",
    "publicCryptoKey": "Base64 encoded public key for encrypting messages to the publisher",
    "publicSigningKey": "Base64 encoded public key for verifying publisher signatures",
    "publisherId": "openzwave",
    "timestamp": "2020-01-20T23:33:44.999PST",
  },
  "signature":  "base64encoded ECDSA signature of the DSS",
  "timestamp": "2020-01-20T23:34:00.000PST",
  }


~~~   

**Requirements For Updating A Publisher Identity**

The requirements for a publisher to allow its identity to be updated: 

1. The message must be encrypted with JWE using the publisher's currently public crypto key (all configuration updates must be encrypted)
2. The message must originate from the DSS publisher. Eg the sender address is \{domain\}/dss and the message must be signed by the DSS.
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

2. Global Certificate. The DSS is published with a certificate signed by a global CA like Lets Encrypt. Subscribers can verify this certificate with the global CA before trusting the DSS. To facilitate the use of global domains, the domain '\{name}.iotc.zone' is available, where \{name} is a globally unique domain.

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
| clientId    | string   | optional     | ID of the client to connect as. Must be unique on a message bus. Default is to generate a temporary ID.
| host        | string   | **required** | IP address or hostname of the remote bus
| login       | string   | **required** | Login identifier obtained from the administrator
| nodeId      | string   | **required** | NodeId of the bridge to create |
| password    | string   | **required** | Password to connect with
| port        | integer  | optional     | port to connect to. Default is determined by protocol
| protocol    | enum     | optional     | Protocol to use: "MQTT" (default), "REST"
| sender      | string   | **required** | Address of the sender, eg: my.domain.org/mrbob of the user that configures the bridge. |
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

Bridge configuration can be set on address: {domain}/\$bridge/\{bridgeId}/\$configure:

Bridges support the following configuration settings:

| Field        | value type   | value        | Description
|:------------ |:---------    |:-----------  |:----------
| clientId     | string       | optional | ID of the client to connect as. Must be unique on a message bus. Default is to generate a temporary ID.
| host         | string       | **required** | IP address or hostname of the remote bus
| login        | string       | **required** | Login identifier obtained from the administrator
| password     | string       | **required** | Password to connect with
| port         | integer      | optional     | port to connect to. Default is determined by protocol
| protocol     | enum         | optional     | Protocol to use: "MQTT" (default), "REST"
| sender      | string   | **required** | Address of the sender, eg: my.domain.org/mrbob of the user that configures the bridge. |
| timestamp   | string   | **required** | Time the record is created |

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
| forward     | string   | **required** | The node, input or output address to forward |
| discovery   | boolean  | optional     | Forward the node/output discovery publications, default=true |
| batch       | boolean  | optional     | Forward the output \$batch publication(s), default=true |
| event       | boolean  | optional     | Forward the output \$event publication(s), default=true |
| forecast    | boolean  | optional     | Forward the output \$forecast publication(s), default=true |
| history     | boolean  | optional     | Forward the output \$history publication(s), default=true |
| latest      | boolean  | optional     | Forward the output \$latest publication(s), default=true |
| value       | boolean  | optional     | Forward the output \$raw publication(s), default=true |
| sender      | string   | **required** | Address of the sender, eg: my.domain.org/mrbob of the user that configures the bridge. |
| timestamp   | string    | **required** | Time the record is created |


## Remove Bridged Nodes, Inputs or Outputs 

To remove a forward, use the following command:

* **\{domain}/\$bridge/\{bridgeId}/remove/node/\$set**
* **\{domain}/\$bridge/\{bridgeId}/remove/input/\$set**
* **\{domain}/\$bridge/\{bridgeId}/remove/output/\$set**


Message structure:

| Field       | type     | required     | Description |
|:----------- |:-------- |:-----------  |:----------- |
| address     | string   | **required** | Address of the publication |
| remove      | string   | **required** | The node, input or output address to remove. |
| sender      | string   | **required** | Address of the publisher |
| timestamp   | string   | **required** | Time the record is created |


# Appendix A: Value Datatypes

The datatype attribute in input and output discovery messages describe what value is expected in publications. The possbible values are:

| Datatype         | Description |
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

For datatype int, when the value starts with '\[' it should be considered a list of integers instead of a single integer.  If an application expects a non-list value and receives a list, the first item in the list should be used. If an application expects a list and receives a non-list value it should be treated as a list of 1 item.

**json values** Values with the json datatype are a catch-all for storing multiple fields as json payload. It should be avoided when possible as discovery provides no description of the structure. If possible rather use the $event publication that publishes all output values in a single event.

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
| alias            | Configured node alias. Used as nodeID of all inputs and outputs |
| color            | color in hex notation |
| description      | device description |
| disabled         | device or sensor is disabled |
| filename         | filename to write images or other values |
| gatewayAddress   | the node gateway address |
| hostname         | network device hostname |
| iotcVersion      | Publishers include the version of the IotConnect standard. Eg v1.0 |
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
| publicKey        | public key for encrypting sensitive configuration settings |
| softwareVersion  | Software/Firmware identifier or version |
| subnet           | IP subnets configuration |

Node status attributes. These convey the current state of the node and are read-only

| Key           | Value Description |
|:------------  |:------------      |
| errorCount    | nr of errors reported on this device |
| health        | health status of the device 0-100% |
| lastError     | most recent error message |
| lastSeen      | ISO time the device was last seen |
| latencyMSec   | duration connect to sensor in milliseconds |
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

| input/output type| Units  | Value Datatype   | description |
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

