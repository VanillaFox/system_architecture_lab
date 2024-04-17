#!/bin/sh

docker exec -it mongo-db sh -c 'mongosh <<EOF
use $MONGO_INITDB_DATABASE
var conferenceObjectID = new ObjectId()
var date =  new ISODate("2016-05-18T16:00:00Z");
db.conferences.insertOne({_id: conferenceObjectID, "title":"Тестовая конференция", "date":date})
db.reports.insertOne({"conference_object_id": conferenceObjectID, "title":"Тестовый доклад", "description":"тестовое описание", "date":date})
EOF'

