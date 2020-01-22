MyZone Introduction
===================

Mission statement
-----------------

Assist inhabitants of a zone by providing relevant information and
control. Information from IoT devices and other sources is collected,
processed and presented. Collected information can be shared with other
zones. Information is presented through desktop and mobile and wearable
devices based on the consumer's situation.

Information and Control
-----------------------

Information can come from many sources like sensors as well as services
that derive new information from existing information. This includes any
data captured by IoT devices, camera images, data from the internet,
user input, as well as services that generate new information like
analytics and machine learning.

Control of devices or services is a form of information, one that is
provided by users and feeds back into the devices. If control is shared
with other zones, then inhabitants and services in that zone can also
provide control.

Addressing & Zoning
-------------------

Anything that produces information is called a 'node', including people.
Nodes reside in a zone and each node has a unique address.

A zone is a physical or virtual area where information resides and can
be presented. A zone can be a home, a street, a city, or an area in a
virtual game world. All consumers within a zone have access to
information published in that zone.

Information that is collected within a zone can be shared with other
zones of choice. External zones can subscribe to information that is
made available to them. Sharing takes place through a secure connection
between zones.

For example, a water level sensor provides water levels to a city
monitoring zone. A service within the monitoring zone interprets the
water levels from multiple sensors and determines the risk level for
flooding. This risk level information is shared with the city's
community zone and available to residents and visitors of the city
website.

A virtual game uses zones for its street map that are bridged to street
zones in the real world. The number of people in the real world is
reflected in the zone of the game world; An alarm triggered in the real
world shows up in the game world; A message sent in the game world shows
up in the real world. Once support for zones is available in the game,
the limitation is the imagination.

More examples are presented in the MyZone use-cases document.

Services Derive Information From Data
-------------------------------------

The amount of data collected easily becomes overwhelming. In itself this
raw data is often not immediately useful and can lead to information
overload, or simply ignoring the data, which defeats the purpose of
collecting it.

Services in a zone derive information from collected data and turn it
into something that is useful to the consumer. For example, a service
issues a security alert when a motion sensor triggers when no-one is
home. This derived information is useful information, while the motion
sensor trigger on its own is not useful if the consumer only wants to
know if there is a security breach. These services can be simple rule
based logic, or more advanced like a neural network that uses image
recognition to classify the object seen on camera. Services are also
nodes with inputs and outputs.

Information Identification and Transparency
-------------------------------------------

When consuming information from external sources, trust in the validity
of this information and the ability to identify its source is important.
Faulty sensors or bad actors can generate information that is
unreliable. It should be easy to identify these if needed.

Published information carries metadata like the timestamp, location, and
identification of the producer. Producer identification provides
transparency as to the source of the information. The consumer can
choose to include or exclude information from unverified sources.
Identification and location can be omitted in case of privacy concerns.

For example, an alarm company receives a security alert. The alert
signature is verified to be from a client and not from a bad actor
creating a distraction. The location and timestamp are used to identify
where and when the alert was given.

Presenting information
----------------------

Inhabitants of a zone can be notified of updates to information based on
the information priority and the situation of the inhabitant.

Situational awareness can come from location, time of day, activity and
other information. It can be used to filter collected information before
it is presented, or delay its presentation until the situation has
changed. The location and activity of an inhabitant can be determined
via a portable or wearable device linked to that inhabitant, or derived
from cameras or other sensors.

Once information updates are accepted it is presented to the zone
inhabitant through the available presentation device. This can simply be
shown on all devices associated with the inhabitant, or the device
currently in use.

Presented information has a life span. Stale information that has
expired should be removed from presentation, depending on the type of
information. This can be a personal preference.

Notification and presentation of notifications are provided by services
that combine information and decide when and where to pass the
information to the consumer. MyZone provides the infrastructure and ease
of interoperability to create reusable building blocks for these
features.

Discovery and Configuration
---------------------------

Publishers of information also provide discovery and configuration
metadata of the information.The discovery metadata describes the type of
information, its publisher, and its configuration. Standardization of
discovery and configuration allows data and information to be managed
from a single user interface regardless of the various technologies
involved.

Changing configuration and controlling inputs can be limited to specific
users as identified by the signature contained in the configuration and
control messages.

Technology
----------

MyZone is technology agnostic. It is a convention that describes the
information format and exchange, discovery, configuration and zoning,
irrespective of the technology used to implement it. Use of different
technologies will serve to improve further integration and makes it
easier to expand the information network.

A reference implementation is provided, written in the golang and
typescript languages, using the MQTT service bus for publishing
information, discovery and configuration in the JSON format.

Data Format
-----------

The information exchange rules must be followed by all implementations
of the convention. A JSON based encoding of the data is recommended.
Other encodings such as XML or whatever becomes popular tomorrow can
also be used.

The primary requirement is that the information fields described are
preserved and all information publishers within a zone use the same
format. A zone speaks only one data format language.

Future proofing can be achieved by using different zones for different
data formats and using a bridge service to share between zones. The
bridge maps between old and new data format. This allows for incremental
improvements while maintaining interoperability.

Versioning
----------

Future versions of this convention must remain backwards compatible. New
fields can be added as long as they remain optional. Implementations
must accept and ignore unknown fields.

Publishers include the version of the MyZone convention when publishing
their node. See 'discovery' for more information.

Terminology
-----------

<table>
<thead>
<tr class="header">
<th>Account</th>
<th>The account used to connect a publisher to a message bus</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>Authentication</td>
<td>Method used to identify the publisher and subscriber with the message bus.</td>
</tr>
<tr class="even">
<td>Configuration</td>
<td>Configuration of the node configuration</td>
</tr>
<tr class="odd">
<td>Discovery</td>
<td>Description of nodes, their inputs and outputs</td>
</tr>
<tr class="even">
<td>Information</td>
<td>Anything that can be published by a producer. This can be sensor data, images,</td>
</tr>
<tr class="odd">
<td>Message Bus</td>
<td>A transport for publication of information and control. Information is published by a node onto a message bus. Consumers subscribe to information they are interested in.</td>
</tr>
<tr class="even">
<td>Node</td>
<td>A node is a device or service that provides information and accepts control input. Information from this node can be published by the node itself or published by a ‘publisher’ that knows how to access the node.</td>
</tr>
<tr class="odd">
<td>Node Input</td>
<td>Node input used to control the node. This can be used for many things, including to manage additional input and outputs.</td>
</tr>
<tr class="even">
<td>Node Output</td>
<td>Zone information is provided through node outputs. For example, the current temperature.</td>
</tr>
<tr class="odd">
<td>Publisher</td>
<td><p>A node that is responsible for publishing information (node outputs) and discovery, and handle configuration updates and control inputs..</p>
<p>Publishers can include identification with published information to provide transparency of the data source.</p></td>
</tr>
<tr class="even">
<td>Subscriber</td>
<td>Consumer of data or information</td>
</tr>
<tr class="odd">
<td>Zone</td>
<td>An area in which information is shared between inhabitants</td>
</tr>
</tbody>
</table>

Zone Overview
-------------

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

1.  Publishing output information. This is mandatory. Every publisher
    > must as a minimum publish their information on the value address
    > (see in/output addresses).

2.  Handling requests to update inputs. This is only for nodes that have
    > inputs. Publishers can implement constraints that only trusted
    > users can update the inputs.

3.  Publishing discovery information for available nodes, node
    > configuration, node inputs, node outputs. This is optional and
    > intended for environments where the computing power is available.

4.  Update node configuration. This is optional and intended for
    > environments where the computing power is available. Publishers
    > can implement constraints that only trusted users can update the
    > inputs.

All publishers in a zone must use the same data interchange format of
the published records. The recommended format is JSON. Other formats
like BSON or XML can be used as long as all publishers of the zone use
the same format. Zone Bridges exchange information in JSON.

Inputs and Outputs
==================

Information flows between nodes to consumers via publishers. Node
information is published as outputs and control is handled via publisher
inputs.

In/Output Addresses
-------------------

The addresses used to publish outputs and control inputs consist of
segments. The address segments include the zone, the publisher of the
information, the node whose information is published or controlled, the
type of information, and the instance of the in- or output.

Segment names consist of alphanumeric, hyphen (-), and underscore (\_)
characters. Other characters are not recommended to allow for various
publish/subscribe technologies like MQTT, REST or other.

MQTT/REST format example. Depending on the implementation environment,
other methods of separating address segments can be used, as long as
they are consistent within a zone:

**{zoneId} / {publisherId} / {nodeId} / {type} / {instance} / value**

**{zoneId} / {publisherId} / {nodeId} / {type} / {instance} / latest**

**{zoneId} / {publisherId} / {nodeId} / {type} / {instance} / 24hours**

**{zoneId} / {publisherId} / {nodeId} / {type} / {instance} / set**

<table>
<thead>
<tr class="header">
<th>Address segment</th>
<th>Description</th>
<th></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>{zoneId}</td>
<td><p>The zone in which publishing takes place. The local and default zoneId is ‘myzone’. This is only valid within a zone and akin to localhost on a computer.</p>
<p>See bridging for more information on connecting multiple zones.</p></td>
<td></td>
</tr>
<tr class="even">
<td>{publisherId}</td>
<td><p>The service that is publishing the information.</p>
<p>A publisher provides its identity when publishing its discovery. The publisher Id is unique within its zone.</p></td>
<td></td>
</tr>
<tr class="odd">
<td>{nodeId}</td>
<td>The node that owns the input or output. This is a device identifier or a service identifier and unique within a publisher.</td>
<td></td>
</tr>
<tr class="even">
<td>{type}</td>
<td>The type of input or output, eg temperature. This convention includes a list of output types.</td>
<td></td>
</tr>
<tr class="odd">
<td>{instance}</td>
<td><p>The instance of the type on the node. For example, a node can have two temperature sensors. The combination type + instance is unique for the node.</p>
<p>The instance can be a name or number. If only a single instance exists the instance can be shortened to “_”</p></td>
<td></td>
</tr>
<tr class="even">
<td></td>
<td>In/Output commands</td>
<td></td>
</tr>
<tr class="odd">
<td></td>
<td>value</td>
<td>The keyword “value” provides the latest known value of the output. The payload is the raw data.</td>
</tr>
<tr class="even">
<td></td>
<td>latest</td>
<td>The keyword “latest” provides a record with the timestamp of the value, a value attribute that contains the output data, the address information and the publisher signature. The value is converted to a string using base64 for binary data.</td>
</tr>
<tr class="odd">
<td></td>
<td>24hours</td>
<td><p>The keyword “24hours” provides a record containing a history attribute with a list of timestamp-value pairs of the last 24 hours. This is intended to be able to determine a trend without having to store these values. The value is provided in its string format.</p>
<p>The content is not required to persist between publisher restarts.</p></td>
</tr>
<tr class="even">
<td></td>
<td>set</td>
<td><p>The keyword “set” provides the string representation of the value to update the input with.</p>
<p>It is published by a consumer that has permission to control it. The publisher subscribes to updates published to the address and updates the node input accordingly.</p></td>
</tr>
</tbody>
</table>

**MQTT Examples:**

1.  The value of the first temperature sensor of node 5 published by a
    > service named 'openzwave' is published on an MQTT bus on topic:
    > **myzone/openzwave/5/temperature/1/latest**

2.  To activate a switch on device node with ID "3", a message is
    > published on topic: **myzone/openzwave/3/switch/1/set**

'Value' Output
--------------

The payload used with the 'value' output is the straight raw data
without metadata.

Publishing information on this address is required. It is primarily
intended for compatibility with 3rd party systems or for use in
environments with limited bandwidth or computing power.

'Latest' Output
---------------

The payload used with the 'latest' output includes the address and
timeStamp of the information and optionally the publisher signature to
verify the content. Publishing information on this address is
recommended for environments that are not too limited in bandwidth and
computing power.

The payload structure is as follows:

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Data Type</th>
<th>Required</th>
<th>Description</th>
<th></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>address</td>
<td>Address</td>
<td>required</td>
<td>Address of the node</td>
<td></td>
</tr>
<tr class="even">
<td></td>
<td>zoneId</td>
<td>string</td>
<td>required</td>
<td></td>
</tr>
<tr class="odd">
<td></td>
<td>publisherId</td>
<td>string</td>
<td>required</td>
<td></td>
</tr>
<tr class="even">
<td>signature</td>
<td>nodeId</td>
<td>string</td>
<td>required</td>
<td></td>
</tr>
<tr class="odd">
<td>instance</td>
<td>string</td>
<td>required</td>
<td>The output instance of this node</td>
<td></td>
</tr>
<tr class="even">
<td>signature</td>
<td>string</td>
<td>optional</td>
<td><p>Signature signed by the publisher.</p>
<p>The address is required when signature is used.</p></td>
<td></td>
</tr>
<tr class="odd">
<td>timeStamp</td>
<td>string</td>
<td>required</td>
<td>Time the output value was obtained, in ISO8601 format, eg YYYY-MM-DDTHH:MM:SS.sssTZ. The timezone is the local timezone where the value was published.</td>
<td></td>
</tr>
<tr class="even">
<td>value</td>
<td>string</td>
<td>required</td>
<td>The output value</td>
<td></td>
</tr>
</tbody>
</table>

A JSON example:\
{\
"publisherId": "openzwave",

"nodeId": "5",

"Instance": "\_",\
"timeStamp": "2020-01-16T15:00:01.000PST",\
"value" : "20.6",\
"signature": "\..."

}

'24hours' Output
----------------

The payload for the '24hours' output contains a history of the values of
the last 24 hours along with address information and signature. It is
updated each time a value changes. This publication is optional and
intended for environments with plenty of bandwidth and computing power.
Consumers can use this history to display a recent trend, like
temperature rising or falling, or presenting a graph of the last 24
hours..

The payload structure is as follows:

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Data Type</th>
<th>Required</th>
<th>Description</th>
<th></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>address</td>
<td>Address</td>
<td>optional</td>
<td>Address of the node</td>
<td></td>
</tr>
<tr class="even">
<td></td>
<td>zoneId</td>
<td>string</td>
<td>required</td>
<td></td>
</tr>
<tr class="odd">
<td></td>
<td>publisherId</td>
<td>string</td>
<td>required</td>
<td></td>
</tr>
<tr class="even">
<td></td>
<td>nodeId</td>
<td>string</td>
<td>required</td>
<td></td>
</tr>
<tr class="odd">
<td>history</td>
<td>list</td>
<td>required</td>
<td>List of timestamps and values in reverse time order</td>
<td></td>
</tr>
<tr class="even">
<td></td>
<td>History records</td>
<td></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td></td>
<td>timeStamp</td>
<td>string</td>
<td>required</td>
<td>Time the output was obtained, in ISO8601 format, eg YYYY-MM-DDTHH:MM:SS.sssTZ. The timezone is the local timezone where the value was published.</td>
</tr>
<tr class="even">
<td></td>
<td>value</td>
<td>string</td>
<td>required</td>
<td>The output value</td>
</tr>
<tr class="odd">
<td>instance</td>
<td>string</td>
<td>require</td>
<td>The output instance of this node</td>
<td></td>
</tr>
<tr class="even">
<td>signature</td>
<td>string</td>
<td>optional</td>
<td><p>Signature signed by the publisher.</p>
<p>The address is required when signature is used.</p></td>
<td></td>
</tr>
</tbody>
</table>

> A JSON example:\
> {\
> "publisherId: "openzwave",
>
> "nodeId": "5",
>
> "Instance": "\_",\
> "signature": "\...",
>
> "History": \[
>
> {"timeStamp": "2020-01-16T15:20:01.000PST", "value" : "20.4" },
>
> {"timeStamp": "2020-01-16T15:00:01.000PST", "value" : "20.6" },
>
> \...
>
> \]\
> }

'set' Input 
------------

Publishers subscribe to the 'set' input address to receive requests to
update the input of a node.

Subscribing to the set address is only for nodes that have inputs.

The payload structure is as follows:

| Field     | Data Type   | Required | Description                                                                                                                                       |     |
|-----------|-------------|----------|---------------------------------------------------------------------------------------------------------------------------------------------------|-----|
| address   | Address     | optional | Address of the node                                                                                                                               |     |
|           | zoneId      | string   | required                                                                                                                                          |     |
|           | publisherId | string   | required                                                                                                                                          |     |
|           | nodeId      | string   | required                                                                                                                                          |     |
| instance  | string      | require  | The input instance of this node                                                                                                                   |     |
| value     | string      | required | The input value                                                                                                                                   |     |
| senderId  | string      | optional | Id of the sender requesting to update the input                                                                                                   |     |
| signature | string      | optional | Signature of this record, signed by the consumer that wants to set the input.                                                                     |     |
| timeStamp | string      | required | Time this request was created, in ISO8601 format, eg YYYY-MM-DDTHH:MM:SS.sssTZ. The timezone is the local timezone where the value was published. |     |

### In/output Types

The type of data being published is obviously widely varied. To
facilitate interoperability between publishers and consumers the
following data types are defined as part of this convention:

<table>
<thead>
<tr class="header">
<th>input/output type</th>
<th>Units</th>
<th>Payload datatype</th>
<th>description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>acceleration</td>
<td>m/s2</td>
<td>List of floats</td>
<td>[x,y,z]</td>
</tr>
<tr class="even">
<td>airquality</td>
<td></td>
<td>integer</td>
<td>Number representing the air quality</td>
</tr>
<tr class="odd">
<td>alarm</td>
<td></td>
<td>boolean</td>
<td>Indicator of alarm status. True is alarm, False is no alarm.</td>
</tr>
<tr class="even">
<td>atmosphericpressure</td>
<td><p>kpa</p>
<p>mbar</p>
<p>Psi</p>
<p>hg</p></td>
<td>float</td>
<td></td>
</tr>
<tr class="odd">
<td>avchannel</td>
<td></td>
<td>integer</td>
<td></td>
</tr>
<tr class="even">
<td>avmute</td>
<td></td>
<td>boolean</td>
<td></td>
</tr>
<tr class="odd">
<td>avpause</td>
<td></td>
<td>boolean</td>
<td></td>
</tr>
<tr class="even">
<td>avplay</td>
<td></td>
<td>boolean</td>
<td></td>
</tr>
<tr class="odd">
<td>avvolume</td>
<td>%</td>
<td>integer</td>
<td></td>
</tr>
<tr class="even">
<td>battery</td>
<td>%</td>
<td>integer</td>
<td></td>
</tr>
<tr class="odd">
<td>co2level</td>
<td>ppm</td>
<td>float</td>
<td></td>
</tr>
<tr class="even">
<td>colevel</td>
<td>ppm</td>
<td>float</td>
<td></td>
</tr>
<tr class="odd">
<td>color</td>
<td>rgb</td>
<td>string</td>
<td></td>
</tr>
<tr class="even">
<td>colortemperature</td>
<td>K</td>
<td>float</td>
<td></td>
</tr>
<tr class="odd">
<td>compass</td>
<td>degrees</td>
<td>float</td>
<td>0-359 degree compass reading</td>
</tr>
<tr class="even">
<td>contact</td>
<td></td>
<td>boolean</td>
<td></td>
</tr>
<tr class="odd">
<td>cpulevel</td>
<td>%</td>
<td>integer</td>
<td></td>
</tr>
<tr class="even">
<td>current</td>
<td>A</td>
<td>float</td>
<td></td>
</tr>
<tr class="odd">
<td>dewpoint</td>
<td><p>C</p>
<p>F</p></td>
<td>float</td>
<td></td>
</tr>
<tr class="even">
<td>dimmer</td>
<td>%</td>
<td>integer</td>
<td></td>
</tr>
<tr class="odd">
<td>doorwindowsensor</td>
<td></td>
<td>boolean</td>
<td></td>
</tr>
<tr class="even">
<td>duration</td>
<td>sec</td>
<td>float</td>
<td></td>
</tr>
<tr class="odd">
<td>electricfield</td>
<td>V/m</td>
<td>float</td>
<td>Static electric field in volt per meter</td>
</tr>
<tr class="even">
<td>energy</td>
<td>KWh</td>
<td>float</td>
<td></td>
</tr>
<tr class="odd">
<td>errors</td>
<td></td>
<td>integer</td>
<td></td>
</tr>
<tr class="even">
<td>heatindex</td>
<td><p>C</p>
<p>F</p></td>
<td>float</td>
<td>Apparent temperature (humiture) based on air temperature and relative humidity. Typically used when higher than the air temperature. At 20% relative humidity the heatindex is equal to the temperature.</td>
</tr>
<tr class="odd">
<td>heading</td>
<td><p>Degrees</p>
<p>Radians</p></td>
<td>float</td>
<td></td>
</tr>
<tr class="even">
<td>hue</td>
<td></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td>humidex</td>
<td>C</td>
<td>float</td>
<td>Humidity temperature index (feels like temperature) derived from dewpoint, in degrees Celcius</td>
</tr>
<tr class="even">
<td>humidity</td>
<td>%</td>
<td>float</td>
<td>Relative humidity in %</td>
</tr>
<tr class="odd">
<td>image</td>
<td><p>jpeg</p>
<p>png</p></td>
<td>bytes</td>
<td>Image in jpeg or png format</td>
</tr>
<tr class="even">
<td>latency</td>
<td><p>sec</p>
<p>msec</p></td>
<td>float</td>
<td></td>
</tr>
<tr class="odd">
<td>location</td>
<td></td>
<td>List of floats</td>
<td>[latitude, longitude, elevation]</td>
</tr>
<tr class="even">
<td>locked</td>
<td></td>
<td>boolean</td>
<td>Door lock status</td>
</tr>
<tr class="odd">
<td>luminance</td>
<td><p>cd/m2</p>
<p>lux</p></td>
<td>float</td>
<td>Amount of light in candela/m2 or in lux</td>
</tr>
<tr class="even">
<td>magneticfield</td>
<td><p>T</p>
<p>mT</p>
<p>uT</p>
<p>nT</p>
<p>G(auss)</p>
<p>mG</p></td>
<td>float</td>
<td>Static magnetic field in Tesla or (milli) Gauss</td>
</tr>
<tr class="odd">
<td>motion</td>
<td></td>
<td>boolean</td>
<td>Motion detection status</td>
</tr>
<tr class="even">
<td>power</td>
<td>W</td>
<td>float</td>
<td>Power consumption in watts</td>
</tr>
<tr class="odd">
<td>pushbutton</td>
<td></td>
<td>boolean</td>
<td>Momentary pushbutton</td>
</tr>
<tr class="even">
<td>signalstrength</td>
<td>dBm</td>
<td>float</td>
<td></td>
</tr>
<tr class="odd">
<td>speed</td>
<td><p>Kph</p>
<p>Mps</p>
<p>mph</p></td>
<td>float</td>
<td></td>
</tr>
<tr class="even">
<td>switch</td>
<td></td>
<td>boolean</td>
<td></td>
</tr>
<tr class="odd">
<td>temperature</td>
<td><p>C</p>
<p>F</p></td>
<td>float</td>
<td></td>
</tr>
<tr class="even">
<td>ultraviolet</td>
<td>UV</td>
<td>float</td>
<td>Radiation with <a href="https://en.wikipedia.org/wiki/Wavelength"><span class="underline">wavelength</span></a> from 10 nm to 400 nm</td>
</tr>
<tr class="odd">
<td>voltage</td>
<td>V(olt)</td>
<td>float</td>
<td></td>
</tr>
<tr class="even">
<td>waterlevel</td>
<td><p>M(eters)</p>
<p>Foot</p>
<p>Inch</p></td>
<td>float</td>
<td></td>
</tr>
<tr class="odd">
<td>wavelength</td>
<td>M</td>
<td>float</td>
<td></td>
</tr>
<tr class="even">
<td>weight</td>
<td><p>Kg</p>
<p>Lbs</p></td>
<td>float</td>
<td></td>
</tr>
<tr class="odd">
<td>windchill</td>
<td><p>C</p>
<p>F</p></td>
<td>float</td>
<td>Apparent temperature based on the air temperature and wind speed, when lower than the air temperature.</td>
</tr>
<tr class="even">
<td>windspeed</td>
<td><p>kph</p>
<p>mps</p>
<p>mph</p></td>
<td>float</td>
<td></td>
</tr>
<tr class="odd">
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
</tbody>
</table>

Node Discovery 
===============

Node discovery describes the node, its attributes, configuration, inputs
and outputs. It is typically published by the same publisher that is
responsible for publishing the output values and subscribing to input
control values. Discovery messages can be signed by its publisher to
verify its authenticity.

Publishers that publish discovery must also publish a node that
represents themselves. The publisher's node id must be '\$publisher'.
Publishers that have their own sensors can choose to publish the inputs
and outputs under the \$publisher node ID, or publish two records, one
for the \$publisher and one for the node with the inputs and outputs.

Publishing of node discovery is optional but highly recommended. It
enables auto discovery, configuration management and information
verification. For very resource restricted devices it can be omitted
however.

Discovery Address
-----------------

The addresses used to publish node discovery consists of segments that
describe the zone, publisher of the information, and the node being
discovered.

Each segment consists of alphanumeric, hyphen (-), and underscore (\_)
characters. Other characters are not recommended to allow for various
publish/subscribe mediums (like MQTT, REST)

MQTT/REST format:

**{zoneId} / {publisherId} / {nodeId} / discover**

<table>
<thead>
<tr class="header">
<th>Address segment</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>{zoneId}</td>
<td>The zone in which discovery takes place. The default zoneId is ‘myzone’.</td>
</tr>
<tr class="even">
<td>{publisherId}</td>
<td><p>The service that is publishing the information.</p>
<p>A publisher provides its identity when publishing a node discovery. The publisher Id is unique within its zone. The node can its own publisher. In case of incompatible devices the publisher can be an adapter service that publishes the nodes.</p></td>
</tr>
<tr class="odd">
<td>{nodeId}</td>
<td><p>The node that is discovered. This is a device identifier or a service identifier and unique within a publisher.</p>
<p>Two special nodes are defined:</p>
<p>“<strong>$publisher</strong>” is the service that publishes on behalf of the node.</p>
<p>“$gateway” represents the device that acts as a gateway to one or more nodes. For example a zwave controller.</p></td>
</tr>
<tr class="even">
<td>discover</td>
<td>Keyword for node discovery.</td>
</tr>
</tbody>
</table>

MQTT examples:

1.  The discovery of a node '5' with a temperature sensor, published by
    > a service named 'openzwave' is published on an MQTT bus on topic
    > **myzone/openzwave/5/discover**. The payload describes the
    > temperature sensor output.

2.  The discovery of a node '3' with a switch published by a service
    > named 'openzwave' is published on an MQTT bus on topic
    > **myzone/openzwave/3/discover**.\
    > The payload describes the switch input and output.

Discovery Payload
-----------------

The discovery payload describes in detail the node, its configuration,
and its inputs and outputs. The objective is for the node to be
sufficiently described so consumers can use it without further
information.

**Example payload for node discovery in JSON format:**

{

Address: {

\...

},

attr: {

type: multisensor,

\...

},

config: {

name: {

datatype: string,

default: "",

description: "Name of the node for presentation",

value: "barn multisensor",

},

...

},

id: "5",

}

**Discovery Fields:**

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Data Type</th>
<th>Required</th>
<th>Description</th>
<th></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>address</td>
<td>Address</td>
<td>required</td>
<td>Address of the node</td>
<td></td>
</tr>
<tr class="even">
<td></td>
<td>zoneId</td>
<td>string</td>
<td>required</td>
<td>Zone where the node lives</td>
</tr>
<tr class="odd">
<td></td>
<td>publisherId</td>
<td>string</td>
<td>required</td>
<td>The publisher of this discovery record.</td>
</tr>
<tr class="even">
<td></td>
<td>nodeId</td>
<td>string</td>
<td>required</td>
<td>ID of the node. This is the same value as nodeId in the address.</td>
</tr>
<tr class="odd">
<td>attr</td>
<td>Dictionary of key-value pairs</td>
<td>required</td>
<td><p>Node attributes provided by the node. Collection of key-value pairs that describe the node.</p>
<p>For interoperability, attribute keys that are part of the convention are described below.</p></td>
<td></td>
</tr>
<tr class="even">
<td>config</td>
<td>List of configuration records</td>
<td>optional</td>
<td>Node configuration. Set of configuration objects that describe the configuration options. These can be modified with a ‘configure’ message.</td>
<td></td>
</tr>
<tr class="odd">
<td></td>
<td>Configuration record definition</td>
<td></td>
<td></td>
<td></td>
</tr>
<tr class="even">
<td></td>
<td>datatype</td>
<td>enum</td>
<td>optional</td>
<td><p>Type of value. Used to determine the editor to use for the value.</p>
<p>One of: bool, enum, float, int, string. Default is ‘string’</p></td>
</tr>
<tr class="odd">
<td></td>
<td>default</td>
<td>string</td>
<td>optional</td>
<td>Default value for this configuration in string format</td>
</tr>
<tr class="even">
<td></td>
<td>description</td>
<td>string</td>
<td>optional</td>
<td>Description of the configuration for human use</td>
</tr>
<tr class="odd">
<td></td>
<td>enum</td>
<td>List of strings</td>
<td>optional*</td>
<td><p>Required when datatype is enum.</p>
<p>List of valid enum values as strings</p></td>
</tr>
<tr class="even">
<td></td>
<td>max</td>
<td>number</td>
<td>optional</td>
<td>Optional maximum value for numeric data</td>
</tr>
<tr class="odd">
<td></td>
<td>min</td>
<td>number</td>
<td>optional</td>
<td>Optional minimum value for numeric data</td>
</tr>
<tr class="even">
<td></td>
<td>name</td>
<td>string</td>
<td>required</td>
<td>Name of configuration. This is the same as the key in the configuration dictionary.</td>
</tr>
<tr class="odd">
<td></td>
<td>secret</td>
<td>bool</td>
<td>optional</td>
<td><p>Optional flag that the value is secret and will not be published.</p>
<p>When a secret configuration is set, the value is encrypted with the node public key.</p></td>
</tr>
<tr class="even">
<td></td>
<td>value</td>
<td>string</td>
<td>required</td>
<td>The current configuration value in string format.</td>
</tr>
<tr class="odd">
<td>inputs</td>
<td>List of records</td>
<td>optional</td>
<td>List of records describing each available input. See input/output definition below.</td>
<td></td>
</tr>
<tr class="even">
<td></td>
<td>See below for the Input/Output record definition</td>
<td></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td>outputs</td>
<td>List of records</td>
<td>optional</td>
<td>List of records describing each available output. See input/output definition below.</td>
<td></td>
</tr>
<tr class="even">
<td></td>
<td>Input/Output record definition</td>
<td></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td></td>
<td>config</td>
<td>List of configuration records</td>
<td>optional</td>
<td>Collection of records that describe in/output configuration, if applicable. See Node configuration for the definition.</td>
</tr>
<tr class="even">
<td></td>
<td>datatype</td>
<td>string</td>
<td>required</td>
<td>Value datatype. One of bool, enum, float, image, int, text, raw</td>
</tr>
<tr class="odd">
<td></td>
<td>default</td>
<td>numbers</td>
<td>optional</td>
<td>Default output value</td>
</tr>
<tr class="even">
<td></td>
<td>description</td>
<td>string</td>
<td>optional</td>
<td>Description of the in/output for humans</td>
</tr>
<tr class="odd">
<td></td>
<td>enum</td>
<td>list</td>
<td>optional*</td>
<td>Required when datatype is enum</td>
</tr>
<tr class="even">
<td></td>
<td>instance</td>
<td>string</td>
<td>required</td>
<td>The output instance when multiple instances of the same type exist. Default is ‘_’ when only a single instance exists.</td>
</tr>
<tr class="odd">
<td></td>
<td>max</td>
<td>number</td>
<td>optional</td>
<td>Maximum possible in/output value</td>
</tr>
<tr class="even">
<td></td>
<td>min</td>
<td>number</td>
<td>optional</td>
<td>Minimum possible in/output value</td>
</tr>
<tr class="odd">
<td></td>
<td>type</td>
<td>string</td>
<td>required</td>
<td>Type of input/output. See list below</td>
</tr>
<tr class="even">
<td></td>
<td>value</td>
<td>string|raw</td>
<td>required</td>
<td>The current in/output value</td>
</tr>
<tr class="odd">
<td>signature</td>
<td>string</td>
<td>optional</td>
<td>Signature of this discovery record, signed by the publisher using its private key. It can be verified with the publisher’s public key.</td>
<td></td>
</tr>
<tr class="even">
<td>timeStamp</td>
<td>string</td>
<td>required</td>
<td>Time this discovery record was last updated or confirmed, in ISO8601 format</td>
<td></td>
</tr>
</tbody>
</table>

**Node Attributes:**

| Key          | Data type | Description                                                                                                                                                                                     |
|--------------|-----------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| certificate  | string    | A certificate from a trusted source, like Lets Encrypt. It is included by publishers to provide consumers a means to verify their identity.                                                     |
| localip      | string    | IP address of the node, for nodes that are publishers themselves                                                                                                                                |
| location     | string    | String with "latitude, longitude" of device location                                                                                                                                            |
| mac          | string    | Node MAC address for nodes that are publishers                                                                                                                                                  |
| manufacturer | string    | Node make or manufacturer                                                                                                                                                                       |
| model        | string    | Node model                                                                                                                                                                                      |
| myzone       | string    | Version of the convention this publisher uses. This attribute must be present when a publisher publishes its own node.                                                                          |
| publicKey    | string    | Publisher's public key used verify the signature provided with publications of information. Only accept public keys from publishers that are verified through their certificate or other means. |
| type         | string    | Type of node. Eg, multisensor, binary switch, See the nodeTypes list for predefined values.                                                                                                     |
| version      | string    | Hardware or firmware version                                                                                                                                                                    |
|              |           |                                                                                                                                                                                                 |
|              |           |                                                                                                                                                                                                 |
|              |           |                                                                                                                                                                                                 |

Node Configuration
==================

Nodes that can be configured contain a list of configuration records
described in the node discovery. The configuration value can be updated
with a configuration command as per below.

The configuration of a node can be updated by a consumer by publishing
on the 'configure' address. The node publisher listens to this request
and processes it after validation.

Only authorized users can modify the configuration of a node.

Configure Address
-----------------

**{zoneId} / {publisherId} / {nodeId} / configure**

<table>
<thead>
<tr class="header">
<th>Address segment</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>{zoneId}</td>
<td>The zone in which discovery takes place. The default zoneId is ‘myzone’</td>
</tr>
<tr class="even">
<td>{publisherId}</td>
<td>The publisher of the node discovery which is handling the configuration update for that node.</td>
</tr>
<tr class="odd">
<td>{nodeId}</td>
<td>The node whose configuration is updated.</td>
</tr>
<tr class="even">
<td>configure</td>
<td><p>Keyword for node configuration. The payload is a collection of key-value pairs that contain requested changes to the configuration.</p>
<p>It is published by a user of the node that has permission to change the configuration. The publisher subscribes to updates published to this address.</p></td>
</tr>
</tbody>
</table>

Configure Payload
-----------------

A dictionary with keyword and value of changes to the configuration

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>type</th>
<th>required</th>
<th>Description</th>
<th></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>address</td>
<td>Address</td>
<td>optional</td>
<td>Address of the node</td>
<td></td>
</tr>
<tr class="even">
<td></td>
<td>zoneId</td>
<td>string</td>
<td>required</td>
<td></td>
</tr>
<tr class="odd">
<td></td>
<td>publisherId</td>
<td>string</td>
<td>required</td>
<td></td>
</tr>
<tr class="even">
<td></td>
<td>nodeId</td>
<td>string</td>
<td>required</td>
<td></td>
</tr>
<tr class="odd">
<td>config</td>
<td><p>Dictionary with key-value pairs for configuration to update</p>
<p>{ key: value, …}</p></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr class="even">
<td>senderId</td>
<td>string</td>
<td>optional</td>
<td>Id of the sender requesting to update the input. This is the publisherId of the consumer, if the consumer publishes.</td>
<td></td>
</tr>
<tr class="odd">
<td>signature</td>
<td>string</td>
<td>optional</td>
<td>Signature of this configuration record, signed by the consumer that wants to modify the configuration. The node publisher can verify if the consumer has permission to modify the configuration of the node.</td>
<td></td>
</tr>
<tr class="even">
<td>timeStamp</td>
<td>string</td>
<td>required</td>
<td>Time this request was created, in ISO8601 format</td>
<td></td>
</tr>
</tbody>
</table>

Node Status
===========

The availability status of a node is published by its publisher when the
availability changes or errors are encountered.

Status Address
--------------

**{zoneId} / {publisherId} / {nodeId} / status**

<table>
<thead>
<tr class="header">
<th>Address segment</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>{zoneId}</td>
<td>The zone in which discovery takes place. The default zoneId is ‘myzone’</td>
</tr>
<tr class="even">
<td>{publisherId}</td>
<td>The publisher of the node discovery which is handling the configuration update for that node.</td>
</tr>
<tr class="odd">
<td>{nodeId}</td>
<td>The node whose configuration is updated.</td>
</tr>
<tr class="even">
<td>status</td>
<td><p>Keyword for node status. Published when the availability of a node changes or new errors are reported.</p>
<p>It is published by the publisher of the node.</p></td>
</tr>
</tbody>
</table>

Status Payload
--------------

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>type</th>
<th>required</th>
<th>Description</th>
<th></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>zoneId</td>
<td>string</td>
<td>required</td>
<td>Zone where the publisher lives.</td>
<td></td>
</tr>
<tr class="even">
<td>publisherId</td>
<td>string</td>
<td>required</td>
<td><p>The publisher of this discovery record. This is the same value as publisherId in the address.</p>
<p>See also ‘signature’</p></td>
<td></td>
</tr>
<tr class="odd">
<td>nodeId</td>
<td>string</td>
<td>required</td>
<td>ID of the node whose status is provided. This is the same value as nodeId in the address.</td>
<td></td>
</tr>
<tr class="even">
<td>status</td>
<td>record</td>
<td>required</td>
<td>Status record</td>
<td></td>
</tr>
<tr class="odd">
<td></td>
<td>available</td>
<td>enum (awake, asleep, missing)</td>
<td>required</td>
<td>The node is available.</td>
</tr>
<tr class="even">
<td></td>
<td>errorCount</td>
<td>integer</td>
<td>optional</td>
<td>Nr of errors since startup</td>
</tr>
<tr class="odd">
<td></td>
<td>errorMessage</td>
<td>string</td>
<td>optional</td>
<td>Last error message</td>
</tr>
<tr class="even">
<td></td>
<td>errorTime</td>
<td>string</td>
<td>optional</td>
<td>Timestamp of last error message in ISO8601 format.</td>
</tr>
<tr class="odd">
<td></td>
<td>lastSeen</td>
<td>string</td>
<td>required</td>
<td>Timestamp in ISO8601 format that the publisher received information from the node.</td>
</tr>
<tr class="even">
<td>signature</td>
<td>string</td>
<td>optional</td>
<td>Signature of this configuration record, signed by the consumer that wants to modify the configuration. The node publisher can verify if the consumer has permission to modify the configuration of the node.</td>
<td></td>
</tr>
<tr class="odd">
<td>timeStamp</td>
<td>string</td>
<td>required</td>
<td>Time the status was last updated, in ISO8601 format</td>
<td></td>
</tr>
</tbody>
</table>

Trust & Digital Signatures
==========================

Note: this section needs further review and a reference implementation.

Trust is essential to information exchange between publishers and
consumers, especially when the producer and consumer don't know each
other directly. In this case trust means that the consumer can be sure
that the publisher is who he claims to be.

To this purpose, publishers include a [[digital
signature]{.underline}](https://en.wikipedia.org/wiki/Digital_signature)
in their publications that lets the consumer verify the records
originate from the publisher. [[This
tutorial]{.underline}](https://www.tutorialspoint.com/cryptography/cryptography_digital_signatures.htm)
explains it with a picture. Common digital signatures are:

-   [[RSA-PSS]{.underline}](https://en.wikipedia.org/wiki/Probabilistic_signature_scheme)
    > part of [[PKCS\#1
    > v2.1]{.underline}](https://en.wikipedia.org/wiki/PKCS_1) and used
    > in OpenSSL

-   DSA and
    > [[ECDSA]{.underline}](https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm)
    > from NIST. NIST produced schemes are considered suspect as there
    > are concerns the NSA has inserted backdoors.

Therefore RSA-PSS is used in this convention. As security is evolving
different future schemes allowed if the signature can identify the
scheme to use.

Verification Process
--------------------

Publishers generate a new key pair on first use and save them in a
secure place. The public key is published with the publisher's node
information. Consumers can decide to trust a publisher by saving their
public key in a trust store. When information is received, the signature
is verified with the public key of that publisher in the trust store.
The implementation examples in the following paragraphs point to tools
for each of these steps.

Consumers subscribe to publishers before using their information. During
the subscription process, the user can be asked whether to trust this
publisher and accept its public key or whether to consider the
information unverified. If accepted, the public key is trusted and the
information published by the publisher is verified against this key. If
verification fails then the consumer is alerted and the information is
ignored.

Before trusting a publisher's public key, the consumer must perform a
verification that this key is indeed from the publisher. There are
several options:

1.  Manually trust the publisher when subscribing to it. This should
    > only be done when there is another way to verify the publisher's
    > identity. For example, the publisher can include a certificate
    > attribute that contains a certificate from Lets Encrypt. On
    > presenting the certificate the consumer can decide to trust the
    > publisher.

2.  Add the publisher's public key using a USB key obtained straight
    > from the publisher.

3.  Add the publisher's public key from a trusted 3rd party.

Zone Bridge
===========

The purpose of a Zone Bridge is to share information between two zones.
To share information, the bridge is configured with the node addresses
to share, the zones they are shared with, the type of information
shared, eg input, output, discovery and status, and whether a valid
publisher signature is required.

A bridge is a service that runs in a zone. It connects to a bridge in
another zone. Both services are nodes in their respective zones. It has
similarities to a firewall that can be configured for what traffic is
allowed in and out of the zone.

Like any other node, bridges have inputs, outputs and configuration. The
inputs define imports from external nodes that are forwarded to the
current zone. The outputs define exports of nodes in the current zone
that are forwarded to external zones.

For example, to export a camera image from 'cam-5' from the the 'ipcam'
publisher in zone-1 with zone-2, a bridge defines an output that exports
the image address: myzone/ipcam/cam5/image/1. It appears as
zone-1/ipcam/cam5/image within the zone-2 network. Zone-ID 'myzone'
identifies the local zone, similar to localhost on a computer.

The bridge node has control inputs for the following commands:

-   Pushbutton to add an import, parameter is its address

-   Pushbutton to add an export, parameter is its address

-   Pushbutton to remove an import, parameter is its address

-   Pushbutton to remove an export, parameter is its address

**A bridge output has configuration options for:**

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>type</th>
<th>required</th>
<th>Description</th>
<th></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td>connection</td>
<td>Connection object for export</td>
<td></td>
<td></td>
<td></td>
</tr>
<tr class="even">
<td></td>
<td>address</td>
<td>string</td>
<td>required</td>
<td><p>URL of the bus to connect to.</p>
<p>For example: mqtt://bus.domain.name:1883 for a connection to another MQTT instance</p></td>
</tr>
<tr class="odd">
<td></td>
<td>clientId</td>
<td>string</td>
<td>required</td>
<td>ID of the client to connect as</td>
</tr>
<tr class="even">
<td></td>
<td>loginId</td>
<td>string</td>
<td>optional</td>
<td>Login identifier</td>
</tr>
<tr class="odd">
<td></td>
<td>credentials</td>
<td>string</td>
<td>optional</td>
<td>Password to connect with</td>
</tr>
<tr class="even">
<td>forward</td>
<td>Forwarding configuration</td>
<td></td>
<td></td>
<td></td>
</tr>
<tr class="odd">
<td></td>
<td>address</td>
<td>string</td>
<td>required</td>
<td>Address to forward: zoneId/publisherId/nodeId[/type[/instance]]</td>
</tr>
<tr class="even">
<td></td>
<td>discovery</td>
<td>boolean</td>
<td>required</td>
<td>Forward discovery: true/false</td>
</tr>
<tr class="odd">
<td></td>
<td>configuration</td>
<td>boolean</td>
<td>required</td>
<td>Forward configuration requests</td>
</tr>
<tr class="even">
<td></td>
<td>status</td>
<td>boolean</td>
<td>required</td>
<td>For node status updates</td>
</tr>
</tbody>
</table>

**A bridge import has configuration options for:**

| Field      | type                         | required | Description |                                                                     |
|------------|------------------------------|----------|-------------|---------------------------------------------------------------------|
| connection | Connection object for import |          |             |                                                                     |
|            | zoneId                       | string   | required    | ID of the zone that is imported from                                |
|            | clientId                     | string   | required    | ID of the client that is connecting                                 |
|            | loginId                      | string   | optional    | Login identifier                                                    |
|            | credentials                  | string   | optional    | Password to connect                                                 |
| forward    | Forwarding configuration     |          |             |                                                                     |
|            | address                      | string   | required    | Address to forward: zoneId/publisherId/nodeId\[/type\[/instance\]\] |
|            | discovery                    | boolean  | required    | Forward discovery: true/false                                       |
|            | configuration                | boolean  | required    | Forward configuration requests                                      |
|            | status                       | boolean  | required    | For node status updates                                             |

Implementation Libraries
========================

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
