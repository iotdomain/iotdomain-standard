<i src="myzone.png"/>

**Status:**

*This convention is currently v0. It is based on a working proof of concept that was built with an earlier version of this convention. This is not considered final until the official v1 release.*

*The working reference implementation that is based on an earlier version of this convention is currently (Jan 2020) being updated. As publishers are converted they will be added to the reference implemention for this version. This includes publishers for openzwave, onewire, isy99, ipnetwork, webcam, wallpaper and a dashboard*. 


# The MyZone Convention

This is a convention for discovery and exchange of information from IoT devices, services and other sources. The goal is to provide interoperability between a wide range of IoT devices and other information sources to build a true information network.

The convention is technology agnostic and can be implemented with programming languages and message bus of choice. This convention defines the addressing and data fields neccesary for discovery and exchange of information.

A reference implementation in the golang language is provided for publishing zwave, onewire, and camera images, as well as a basic user interface with a dashboard and configuraiton editor.

The convention can be found here:  [CONVENTION](./myzone-convention.md)

## Installation

The reference implementation consists of stand-alone 'publishers', written in golang and typescript. Each of these publishers connect to an MQTT message bus and exchange information in JSON format. 

Publishers can be built from source. Each will be a standalone binary that can operate independently. Binary versions are available for raspberry pi and for linux x64 based systems. Publishers can be launched simply by running it, or through the included systemd launcher for autostart. Publishers can run as any user but some publishers like the 'ipnet' network scanner requires elevated priviliges. See instructions of the publishers for details.

The only configuration neccesary is to edit the file myzone.conf and configure the mqtt bus connection information. This file is first read from /etc/myzone/myzone.conf and if not found, in a local subdirectory 'config'. It is read by all publishers.

Minimal manual installation on Raspberry pi:
* $ git clone ... myzone            -> todo, update location of binaries
* $ edit myzone/config/myzone.conf  -> check the mqtt bus and storage settings
* $ sudo cp -a myzone /opt/
* $ /opt/myzone/runall.sh


## Usage

Most publishers work out of the box and start publishing their discovery information.

A handy tool is 'mqttspy'. It allows you to view the traffic on the mqtt bus.

The included 'dashboard' publisher provides a user interface. Open the dashboard with a browser, eg http://localhost:8080" when running locally, and open the settings menu. The dashboard supports multiple mqtt accounts. Edit the example account with the mqtt broker address(es) of your choice and enable the connection. 

This will immediately presents a list of discovered publishers. Enable the publishers you're interested it. This will autopopulate the list of nodes for each publisher. 

Next, edit the default dashboard or add a new dashboard page. From here add tiles to watch.


## Contributing

Help is most welcome. The 'core' library to help write publishers is available for golang and typescript languages. You can write a publisher using these libraries, or add another core library in another language yourself.

The dashboard is a basic implementation in React and in need of improvement by someone who understands react better than I do. Likewise the UX is basic and can benefit from a UX designer's eye.

Shoot me a message if you are interested.


## Credits

* This project was inspired by the [homie convention](https://github.com/homieiot/convention), a convention for home automation that is based on the MQTT bus. 
* The openzwave publisher would not have been possible without the work of the 'go-openzwave' library.

## License

This project is licensed under GPL-3. Anyone can copy and use this project. Any modifications and derivatives of the source must be made public for the benefit of the community.

