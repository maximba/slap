# Slap-webservice: Slap Reservation Room Backend

## Table of Contents
- [Description](#description)
- [Requirements and Installation](#requirements-and-installation)

## Description

Slap is a REST server backend. It handle basic operation for creating and
manage room and attendee data structures

## Requirements and Installation

You should provide a data directory named `datadir` where database will
store persistent data.

Initial data are dumped from file `initdb/init.sql`

Project is designed to run into containers. Installation is so easy as run
docker-compose:

	```sh
	docker-compose up -d 
	```

You should provide a data directory named `datadir`



