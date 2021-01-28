# Odds API challenge

## How to build it & run it

> make build && make run  

This will start the DB (mongo) and a golang API which will call the ODDS API.  
It will first save all sports into a "sports collection", then save odds into a "odds collection".  
After that a new call will be executed every hour (by default, but configurable) to look up for new fixtures filtered with parameters as requested in the challenge description.

## How to clean the workspace for a fresh start

> make clean  

## Env config
The following configs are taken from the `.env` file.  
> API_KEY  
> DB_CON_STRING  
> MONGO_INITDB_ROOT_USERNAME  
> MONGO_INITDB_ROOT_PASSWORD  
> MINUTES  

## How to debug mongo
Issue the following command to get into the mongo docker instance and open a mongodb cli:  
> make mongodebug  

Then select the proper database issuing:  
> show dbs  
> use oddsdb  

Query the collections with (examples):  
> show collections  
> db.sports.count()   
> db.odds.count()   
> db.sports.find()   
> db.odds.find()   

## Observations

I've used Docker to define my app and docker-compose to describe the deployment.  

I've used a Makefile to define aliases in order to ease the workflow (build, run, clean).  

The logic for the api and the DB was implemented in a separate module each.  

I've used Logrus to have logs of different levels of information.

It's my first time using Mongo DB. I usually use MySQL but I picked Mongo this time because I figured it's a suitable DB for this exercise given the scope of it and the structure of the data.  

I'm aware I should't use `root:example` for db credentials in a production environment.  
I using them here to facilitate the development process and because secret's management it's out of scope of the challenge.  

I'm aware a proper production-ready project should include automated testing (either unit-tests or integration tests).  

Counting the hours spent through the days I was able to focus on it I would say that consumed me approximately, 12 hours in total.