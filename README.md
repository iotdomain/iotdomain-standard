<i src="iotconnect.png"></i>

# The IotConnect Standard

[[TOC]]

## Status

This standard is currently (Feb 2020) draft version 0. The standard is currently being validated with use-cases and a reference implementation.

## Audience

This standard is aimed at software developers and system implementors with a basic knowledge of operating systems and computing devices such as raspberry pi or industrial automation systems.

## Description

IotConnect is a standard for M2M (machine to machine) discovery and exchange of information from IoT devices and related services. The goal is to provide interoperability between a wide range of IoT devices and their consumers. Benefits of following this standard are:

1. Consumers only have to implement support for this standard, instead of a wide variety of other device specific protocols, like for example zwave, zigbee, 1-wire, [and many more....](https://www.ubuntupit.com/top-15-standard-iot-protocols-that-you-must-know-about/)
2. Developers of IoT devices that implement this standard don't have to invent their own protocol. The devices can directly be used which greatly enlarges their market.
3. IoT Devices that follow the standard start on good security footing. Instead of exposing the IoT devices to the internet, a hardened message bus is used instead. It is much easier to secure only one message bus than all IoT devices. IoT devices and consumers only have one outgoing connection to the message bus and can remain behind an otherwise closed firewall.
4. Legacy third party devices can easily be integrated using 'adapters', services that translate between the standard and the third party protocol. In fact, in the near future the bulk of IoT devices are made compatible this way. The reference implementation has various examples. 
5. This standard does not limit itself to devices. Information *enrichment* services, such as automation, analysis and AI services that support the standard, can be added to the bus and use the information that is published. In turn they can publish their own information to be shown or used by other services. These services no longer need to figure out how to connect to information sources and can readily use the information published on the message bus. It greatly simplifies the flow of information.
6. Information sharing is built into the standard using zoning. The owner of a zone can share select information with users in another zone without concerns of security as the zones remains locked off to anything else. IoT devices don't need to know anything about how to share their information. See the use-cases for examples of how useful this can be.

The use of a message bus is key to this standard to ensure security without the headaches. Secure and hardened message bus implementations are readily available for different environments, like the lightweight [mosquitto](https://mosquitto.org/)' that easily run on a raspberry pi, enterprise level implementations such as [HiveMQ](https://www.hivemq.com/), and massive cloud based implementations like [Amazon IoT](https://aws.amazon.com/iot/) and [Google IoT](https://cloud.google.com/iot-core).

The standard is technology agnostic and can be implemented with any programming language and message bus of choice. MQTT and HTTP based protocols are the most common formats for transporting the information.

A reference implementation of library supporting the standard for MQTT based message busses is available along with several adapters that use it to publish information from zwave, onewire, cameras and other. A dashboard with configuration editor provides an example on how to use the published information.

The standard can be found here:  [IoT Connect Standard](./iotconnect-standard.md)

The golang reference implementation: [IotZone-golang](https://github.com/hspaay/iotzone.golang)

## Supported IoT Devices

Being a draft standard in Feb 2020, you will not find an IoT device that supports this standard directly, yet. This matters not as it is easy to write an adapter for existing devices that translate between the device and this standard.  The reference implementation includes a software library and various adapters that can be used as an example.

Implementing the standard is lightweight and should not pose a problem from small devices. A raspberry pi for example can easily accomodate dozens of adapters without a significate use of CPU used up by the standard implementation.  

In case of extremely constrained devices such as embedded micro controllers where there isn't sufficient memory, an optimized proprietary protocol can be used to connect to a gateway that implements the standard. If the device is able to run a TCP/IP stack there is likely sufficient resources available to run a memory optimized library of the standard.

# IoT Connect and IoT Related Protocols 

## Introduction

There are many protocols defined for communication between devices and destination. Rather than rehash what others have written, this section provides links to summaries and describes how IoT Connect uses or overlaps with the more common protocols.

The following article on IoT protocols provides a good overview of internet related protocols broken down in categories that each solve a specific problem. Note that it does not include non-internet protocols such as ZWave and Zigbee.

* https://www.postscapes.com/internet-of-things-protocols/


## Transport Protocols

Transport protocols are about transporting low level packets between two points. Commonly known transports are Wifi and Bluetooth.

IoTConnect is agnostic and unaware of the transport protocol used.

## Infrastructure Protocols

Infrastructure protocols use a transport protocol to deliver low level messages. The most commonly known are TCP/IP and UDP.

Special mention for [CoAP](https://en.wikipedia.org/wiki/Constrained_Application_Protocol). CoAP is a specialized internet protocol to connect constrainted devices using REST based access.
  
IoTConnect allows the use of any infrastructure protocol to obtain information from devices. To publish information it uses the protocol required by the message bus, usually TCP/IP. 


## Identification Protocols

## Discovery Protocols

Discovery protocols such as mDNS, BLE beacon, Hypercat, UPnP aim to discover devices on the same network or in the vicinity. COAP also support discovery and device metadata.

IoTConnect implements discovery by using the publish subscribe mechanism of a message bus. It differs with the mentioned protocols in that:
A) it is secure as it does not require incoming network access
B) it does not depend on the network topology and is independent of the network infrastructure used to connect devices to the message bus.
C) it uses Zoning which gives the user full control over what devices have access, regardless the network they are on..

IoTConnect is decentralized. Protocols such as COAP use a centralized server model where consumers connect to in order to obtain device information. This differs in that IoTConnect does not utilize any connections other than to the message bus. That said, it is possible to implement a COAP service on top of IoTConnect to provide interoperability if needed.


## Data/Message Protocols

Data/Messaging protocols aim to deliver informational messages using a transport protocol. 

* [MQTT](https://en.wikipedia.org/wiki/MQTT) MQ Telemetry Transport, ISO standard for light weight publish/subscribe messaging between devices. Usually runs over TCP/IP but it can use other lossless bi-directional network. Intended for use in small footprint and resource constrainted devices.

* [AMQP](https://en.wikipedia.org/wiki/Advanced_Message_Queuing_Protocol). Open standard for message oriented middleware for point to point and publish/subscribe connectivity. Supports secure and reliable communicate using TCP.


* [DDS](https://en.wikipedia.org/wiki/Data_Distribution_Service). networking middleware to enable data dependable, high performance real-time information exchange.


IoTConnect uses a data messaging protocol, by default MQTT, to publish and subscribe messages on a message bus. It can in theory work with any protocol that supports publish/subscribe. MQTT is the default because it is light weight and well suited for small devices. It also has its drawbacks but other protocols can be use in case these become a problem.


## Device Management Protocols

Device management protocols aim to facilitate asset management.

IoTConnect does not provide a centralized asset directory. Instead it uses a decentralized approach where each device announces itself. Asset management can be build using IoTConnect through its discovery and configuration messaging.

## Semantic Standards

Semantic standards provide a way to describe information using an encoding such a JSON or XML.

IoTConnect overlaps with some of these standards in that it defines the messages to describe information using JSON.


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
To view zwave Bob would have to buy a Vera controller or equivalent with web server and to control the lights a phone app is needed. Instead Bob uses IoTConnect with IotZone adapters for Zwave, Philips Hue, and dashboard. For hardware he purchases the Zwave sensors, a Raspberry Pi 3, a USB zwave controller and the Philips Hue lightbulbs. The price difference between with a Vera and the simpler USB controller pays for the Raspberry Pi.

Putting it all together he plugs the Zwave USB controller into the raspberry pi, installs mosquitto for the MQTT message bus, and the IotZone adapters simply per instructions. Some minimal configuration is needed to secure the message bus and to point the adapters to the message bus and give them credentials to login. 

To setup the dashboard he points the browser on his computer to the raspberry pi. The adapters already connected to the bus and announce themselves using the discovery feature of the standard. The dashboard shows the zwave and Hue available adapters. Bob subscribes to both and is presented with discovered nodes. Next the Hue bulbs are joined to the wifi as per manual, and the Zwave devices are plugged in and joined with the controller as per instructions of the zwave controller. Shortly after, the auto discovery show the devices are available. A few clicks and Bob adds the device to the dashboard and can view the temperature and control the switches.

Embellished by the success, Bob adds a zwave security lock to his front door and pairs it with the zwave controller. The lock automatically shows up thanks to auto discovery and is ready to be configured with pin codes. The dashboard shows several inputs to control and configure the lock, one is to set a pin, and Bob is ready to go. 

a pre-setup Raspberry IoTZone version (future plan) makes things even easier as the message bus and adapters are already setup.


# Credits

* This project was inspired by the [homie convention](https://github.com/homieiot/convention), a convention for home automation that is based on the MQTT bus. 

# License

This project is licensed under GPL-3. Anyone can copy and use this project. Any modifications and derivatives of the source must be made public for the benefit of the community.

