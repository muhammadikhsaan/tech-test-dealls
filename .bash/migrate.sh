#!/bin/bash

if [[ $ENVIRONTMENT == "PRERELEASE" ]];
  then
    export $(echo $(cat .env.prerelease | xargs)) && go run dealls.test/migrations
elif [[ $ENVIRONTMENT == "PRODUCTION" ]];
  then
    export $(echo $(cat .env.production | xargs)) && go run dealls.test/migrations
elif [[ $ENVIRONTMENT == "DEVELOPMENT" ]];
  then
    export $(echo $(cat .env.development | xargs)) && go run dealls.test/migrations
elif [[ $ENVIRONTMENT == "TEST" ]];
  then
    export $(echo $(cat .env.test | xargs)) && go run dealls.test/migrations
else
    export $(echo $(cat .env.local | xargs)) && go run dealls.test/migrations
fi