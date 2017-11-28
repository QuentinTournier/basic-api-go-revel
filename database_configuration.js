use cloud_project;
db.createCollection('users');
db.users.save({
    "firstName": "Raymon",
    "lastName": "Boulier",
    "position": {
        "lat": -28.488321,
        "lon": 37.652013
    },
    "birthDay": "17/02/1954"
});
