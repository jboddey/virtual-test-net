# Golang virtualised network

This project is for prototyping the use of Golang to automate the provisioning and management of an IoT testing environment.

### Requirements
The interface used in network.go should be a physical ethernet adapter which is connected to a managed switch. The port that is used for this should be tagged on all VLANs on the switch. VLANs should be configured with IDs between 101 and 199, and each port should be untagged on a single VLAN.

For example:

Port 1 = untagged VLAN 101, 

Port 2 = untagged VLAN 102

### Building
Root permissions may be required to run the executable as packet capturing takes place.

Build using:` go build`

Run using:` sudo ./framework`

### Disclaimers:
 - This is my first use of Golang
 - Many variables are hardcoded as this is purely a POC (and will be removed)