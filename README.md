# Setup Project

## Initialize a postgres database

Create a postgres database*

Place postgres URL in the environment file

``
DATABASE_URL=user:password@host:port/database
``

*Project made using Postgres 12.3

## Create Google Oauth credentials

Go to https://console.developers.google.com/apis/dashboard and create a new project. Create a new Oauth 2.0 Client ID
and secret in the ``credentials`` tab Populate the projectvars.env file with your Oauth Client ID and secret

``
GOOGLE_CLIENT_ID=<CLIENTID>
GOOGLE_CLIENT_SECRET=<SECRET>
``

## Add Migrations

``migrate create -ext sql -dir database/migrations -seq migration_name``
