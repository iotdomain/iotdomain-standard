<i src="iotdomain.png"></i>

# The IoTDomain Standard

[[TOC]]

## Status

*** Archived - 
Note January 2021: Since starting this project a number of international IoT standard have reached or are about to reach maturity. This standard has been a great learning experience and the work will continue on the [WoST project](https://github.com/wostzone/) that aims to be WoT compliant (where possible) ***

This standard is currently (June 2020) a draft version. The standard is currently being validated with use-cases and a reference implementation.

The standard can be found here:  [IoTDomain Standard](./iotdomain-standard.md)

The golang reference implementation: [go-iotdomain](https://github.com/iotdomain/go-iotdomain)

## Audience

This standard is aimed at software developers and system implementors with a basic knowledge of operating systems and computing devices such as raspberry pi or industrial automation systems.

## Description

IoTDomain is a simple standard for M2M (machine to machine) discovery and exchange of information from IoT devices and related services. The goal is to provide interoperability between a wide range of IoT devices and their consumers. Benefits of following this standard are:

1. Consumers only have to implement support for this standard, instead of a wide variety of other device specific protocols, like for example zwave, zigbee, 1-wire, [and many more....](https://www.ubuntupit.com/top-15-standard-iot-protocols-that-you-must-know-about/)
2. Developers of IoT devices that implement this standard don't have to invent their own protocol. The devices can directly be used which greatly enlarges their market.
3. IoT Devices that follow the standard start on good security footing. A hardened message bus is used instead of exposing the IoT devices to the internet. It is much easier to secure only one message bus than all IoT devices. IoT devices and consumers only have one outgoing connection to the message bus and can remain behind an otherwise closed firewall.
4. Incompatible third party devices can easily be integrated using 'adapters', services that translate between the standard and the third party protocol. In fact, in the near future the bulk of IoT devices are made compatible this way. The reference implementation has various examples. 
5. This standard does not limit itself to devices. Information *enrichment* services, such as automation, analysis and AI services that support the standard, can be added to the bus and use the information that is published. In turn they can publish their own information to be shown or used by other services. These services no longer need to figure out how to connect to information sources and can readily use the information published on the message bus. It greatly simplifies the flow of information.
6. Information sharing is built into the standard using zoning. The owner of a zone can share select information with users in another zone without concerns of security as the zones remains locked off to anything else. IoT devices don't need to know anything about how to share their information. See the use-cases for examples of how useful this can be.

The use of a message bus is key to this standard to ensure security without the headaches. All communication takes place over the message bus. Secure and hardened message bus implementations are readily available for different environments, like the lightweight [mosquitto](https://mosquitto.org/)' that easily run on a raspberry pi, enterprise level implementaticabin2ons such as [HiveMQ](https://www.hivemq.com/), and massive cloud based implementations like [Amazon IoT](https://aws.amazon.com/iot/) and [Google IoT](https://cloud.google.com/iot-core).

The standard is technology agnostic and can be implemented with any programming language and message bus of choice. MQTT and HTTP based protocols are the most common formats for transporting the information.

A reference implementation of library supporting the standard for MQTT based message bus is available along with several adapters that use it to publish information from zwave, onewire, cameras and other. A dashboard with configuration editor provides an example on how to use the published information.

## Supported IoT Devices

Being a draft standard in Feb 2020, you will not find an IoT device that supports this standard directly, yet. This matters not as it is easy to write an adapter for existing devices that translate between the device and this standard.  The reference implementation includes a software library and various adapters that can be used as an example.

Implementing the standard is lightweight and should not pose a problem from small devices. A raspberry pi for example can easily accomodate dozens of adapters without a significate use of CPU used up by the standard implementation.  

In case of extremely constrained devices such as embedded micro controllers where there isn't sufficient memory, an optimized proprietary protocol can be used to connect to a gateway that implements the standard. If the device is able to run a TCP/IP stack there is likely sufficient resources available to run a memory optimized library of the standard.

# IoT Related Protocols 

So how does this standard relate to other IoT protocols out there? Lets start with an overview:

* https://www.postscapes.com/internet-of-things-protocols/

* https://www.ubuntupit.com/top-15-standard-iot-protocols-that-you-must-know-about/

## Introduction

As seen above, there are many protocols related to IoT. Rather than rehash what others have written, this section describes how IoTDomain uses or differs with the more commonly known protocols based on the categories they are in. This is not the same breakdown as the layers of the OSI model but focuses on their purpose in the IoT application space.

## Transport Protocols

Transport protocols are about transporting low level packets between two points. A well known transport is TCP/IP and UDP. They make use of a physical transport such as Ethernet, wifi, bluetooth, LoraWAN, RFID, NFC, and so on.

IoTDomain is agnostic of the transport protocol used. Instead it depends on higher level messaging protocols and device gateways that use transport protocols. Simply put, use it when available. It is all good.

## Discovery Protocols

Discovery protocols such as mDNS, BLE beacon, Hypercat, UPnP, AD aim to discover devices on the same network or in the vicinity. One of the more interesting options is COAP's resource discovery as described in [RFC7252#section-7](https://tools.ietf.org/html/rfc7252#section-7) that provides discovery of constrained devices like those in IoT. 

IoTDomain implements discovery by using the publish subscribe mechanism of a message bus. It differs with the mentioned protocols in that:

A) IoTDomain does not depend on the network topology and is independent of the network infrastructure used to connect devices to the message bus. Many older discovery protocols rely on multicast DNS and a specific network topology and are therefore not compatible. COAP is one of the few exceptions and facilitates discovery of remote devices. In this sense it is similar to IoTDomain.

B) IoTDomain devices (nodes) remain hidden on the network which improves their security. Most other protocols require the device to be accessible and connectable. COAP uses HTTP based URL schema to identify the device endpoint. This requires both the directory service and the device to be accessible via the internet, whereas IoTDomain only exposes the message bus. In Addition IoTDomain can use but does not require HTTP/REST and does not utilize a request/response mechanism.

C) IoTDomain uses 'domains' which gives the user control over what devices have access to a domain, regardless the network topology used. COAP supports groups which could achieve the same goal for discovery.


## Message Protocols

Messaging protocols aim to deliver messages to one or multiple consumers using a transport protocol without understanding their content. 

* [MQTT](https://en.wikipedia.org/wiki/MQTT) MQ Telemetry Transport, ISO standard for light weight publish/subscribe messaging between devices. Usually runs over TCP/IP but it can use other lossless bi-directional network. Intended for use in small footprint and resource constrainted devices.

* [AMQP](https://en.wikipedia.org/wiki/Advanced_Message_Queuing_Protocol). Open standard for message oriented middleware for point to point and publish/subscribe connectivity. Supports secure and reliable communicate using TCP.

* [DDS](https://en.wikipedia.org/wiki/Data_Distribution_Service). networking middleware to enable data dependable, high performance real-time information exchange. Like MQTT it supports publish/subscribe but also includes discovery of publishers and subscribers and exclusive ownership of topics (addresses). [DDS-XRCE](https://objectcomputing.com/resources/publications/sett/october-2019-dds-for-extremely-resource-constrained-environments) is aimed at resource constrained devices. A [Micro XRCE-DDS client](https://github.com/eProsima/Micro-XRCE-DDS-Client) is available for C++ clients.
  
IoTDomain requires the use of a message bus with publish/subscribe capability. Information is shared with one or more subscribers in a zone and each zone can have its own message bus. MQTT, AMQP and DDS are all able to work as a message bus for IoTDomain. MQTT is considered the default as it is lightweight and clients are readily available. That said, DDS looks very interesting as an alternative and could offer reliability and security benefits.

The reliance on a message bus has several pros and cons. IoTDomain values security above all and works well in situations where the downsides are acceptable.
1. Upside: Security. Devices remain hidden to the internet and only have a single outgoing connection to the message bus. This avoids many security risks.
2. Upside: Ease of device configuration. The device only needs to be configured to connect to the message bus and remains unaware of who the consumers are. Changes to the consumers do not affect the device.
3. Downside: Message bus subscribers that are disconnected do not receive their messages unless the message bus supports queuing. Some enterprise message busses, like HiveMQ and Azure Message Bus, support queuing of messages until delivery at the cost of a more complicated configuration. IoTDomain addresses this problem by supporting a history message that contains the recently published values at the cost of extra bandwidth. More importantly though, guaranteed delivery in itself is insufficient as it doesn't guarantee the message is properly processed. 

For delivery of critical messages different approaches can be used depending on the problem to solve. The first approach is to use a message queue on the message bus where the message remains in the queue until the subscriber has successfully processed it. This is the most robust approach as it guarantees no message is missed. The second approach, without a queue, is to repeat the message periodically until the receiver publishes a confirmation on the publishers input once the message is successfully processed. For example, a security alarm sensor repeats its alarm notification until the alarm service responds with a message that help is on the way. Last, when only the most recent status is relevant the use of 'retained' messages supported by MQTT would immediately update the subscriber on reconnect. The latter is the simplest and default approach when using MQTT.

## Device Management Protocols

Device management protocols aim to facilitate asset management. COAP and other protocols support a central directory service that can be used for asset management.

Centralized Asset Management is not included in IoTDomain. However, device management is available using 'retained' capability on message busses that support it as discovery messages are published with the retained flag set.
IoTDomain discovery messages include asset information like device make/model. Subscribers to discovery messages will receive all available publishers and nodes when connecting to the message bus. In many cases that makes a directory service unnecesary. Not all message busses support retained capability however. Mosquitto is a simple lightweight message bus that does support retained messages.

A central directory service can easily be added by storing discovery messages and making them available on request.


## Semantic Standards

Semantic standards provide a way to describe information using an encoding such a JSON or XML.

IoTDomain overlaps with some of these standards in that it defines the messages to describe information using JSON.


# Contributing

Help is always welcome, especially in the following areas:
1. Check the text on semantics and proper English. 
2. Ask questions on use-cases where you think this standard can help but are not sure how. 
3. Provide answers with use-cases.
4. Write adapters for existing IoT devices

This will help solidify the draft until it is ready for publication. Just put a ticket in if you are interested in helping.

# Use Cases

Use-cases help developers and implementors understand how best to use this standard. Use cases are split into two main categories of usage:

1. [Home automation usage](./home-automation-usecases.md)
2. [Industrial automation](./industrial-usescases.md)


Example of a home automation use-case:

Bob would like to view the temperature outside and compare it with several places in his home. Bob has heard of openzwave and found a great little multi-sensor that also captures humidity and motion, great for future expansion of his project. A couple of philips Hue lightbulbs are connected to his wifi. 

Bob wants to have a single dashboard where sensors and lightbulbs can be viewed and controlled.
To view zwave Bob would have to buy a Vera controller or equivalent with web server and to control the lights a phone app is needed. Instead Bob uses IoTDomain with adapters for Zwave, Philips Hue, and dashboard. For hardware he purchases the Zwave sensors, a Raspberry Pi 3, a USB zwave controller and the Philips Hue lightbulbs. The price difference between with a Vera and the simpler USB controller pays for the Raspberry Pi.

Putting it all together he plugs the Zwave USB controller into the raspberry pi, installs mosquitto for the MQTT message bus, and the IoTDomain adapters simply per instructions. Some minimal configuration is needed to secure the message bus and to point the adapters to the message bus and give them credentials to login. 

To setup the dashboard he points the browser on his computer to the raspberry pi. The adapters already connected to the bus and announce themselves using the discovery feature of the standard. The dashboard shows the zwave and Hue available adapters. Bob subscribes to both and is presented with discovered nodes. Next the Hue bulbs are joined to the wifi as per manual, and the Zwave devices are plugged in and joined with the controller as per instructions of the zwave controller. Shortly after, the auto discovery show the devices are available. A few clicks and Bob adds the device to the dashboard and can view the temperature and control the switches.

Embellished by the success, Bob adds a zwave security lock to his front door and pairs it with the zwave controller. The lock automatically shows up thanks to auto discovery and is ready to be configured with pin codes. The dashboard shows several inputs to control and configure the lock, one is to set a pin, and Bob is ready to go. 

a pre-setup Raspberry IoTDomain version (future plan) makes things even easier as the message bus and adapters are already setup.


# Credits

* This project was inspired by the [homie convention](https://github.com/homieiot/convention), a convention for home automation that is based on the MQTT bus. 

# License

This project is licensed under GPL-3. Anyone can copy and use this project. Any modifications and derivatives of the source must be made public for the benefit of the community.

