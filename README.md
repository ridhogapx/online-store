## Synapsis Backend Challenge
As a BackEnd Engineer you are required to create an online store application, you don't need to create a FrontEnd but focus on the BackEnd (RESTful API) only. The programming language you must use is Go-lang or Java spring boot.
You can develop your app by starting with prioritized features first. The following are the priority features to meet the MVP (minimum viable product) criteria:

- Customer can view product list by product category (Check)
- Customer can add product to shopping cart (Check)
- Customers can see a list of products that have been added to the shopping cart (Check)
- Customer can delete product list in shopping cart (Check)
- Customers can checkout and make payment transactions (Check)
- Login and register customers (Check)

## Tech Stack 
This Web Services built with below listed technologies:

- Go (Programming Language)
- Postgres
- Go Fiber (Framework)
- Docker (Deployment)

## Getting Started
In order to run this app in your own machine, you need to clone this repository first. 
And then run below command to build Docker Compose:

```sh
$ docker compose build
```

After that, you can start Docker Compose using below command:

```sh
$ docker compose up
```

## Network Configuration 
This app using custom network adapter configuration static IP Address named <b>Synapsis Net</b> inside Docker Container and using segment IP Address 10.10.10.x/24. Postgres as the database is configure with IP Address 10.10.10.2 and the Web Services is configure with IP Address 10.10.10.3
