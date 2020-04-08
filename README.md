# Single event-booking system

This is a single event booking system built on Go with Flutterwave (Rave) implemented for payment.
Feel free to use any component of this mini-application ❤ 

## Requirements
     
[Git](https://git-scm.com/)   
[Go](https://golang.org/doc/install) -v1.1 and above  
[Node](https://nodejs.org/en/download/)  
[MySQL](https://mysql.com) - v5.7 and above  
Create a personal account with [Flutterwave](https://dashboard.flutterwave.com) to collect test payments

## Installation

Clone the front-end of this application built on Vue and set it up with:

```bash
git clone https://github.com/Iamt-chadwick/Event-Booking-Vue.git
cd Event-Booking-Vue
yarn install 
yarn serve
```
The application's front-end would be served here.

Clone this repo

```bash
git clone https://github.com/Iamt-chadwick/event-booking-golang.git
```
Change directory

```bash
cd event-booking-golang
```
Change .env template to set environment variables
```bash
cp .env.example .env
```

Update the .env file with your database credentials including 

* PORT: This is the port the application will be served on
* DB_HOST: This is your database hostname/IP address
* DB_NAME: This is the name of the database created for the application
* DB_NAME: This is your database user
* DB_PASS: This is your database password if any, it should be left blank if no password is configured (localhost)  

You would get the credentials below from your account at [Flutterwave](https://dashboard.flutterwave.com)
* RAVE_PUBLIC_KEY : 
* RAVE_SECRET_KEY


## Run Application

```python
go run server.go
```
You can view the application at the port at which the front end is serving.

Note: Update the 
* API_URL in /src/http.js in the front end directory.
* API_publicKey = <RAVE_PUBLIC_KEY>

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
