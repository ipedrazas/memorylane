CREATE TABLE events (
 ID serial NOT NULL PRIMARY KEY,
 info json NOT NULL
);


INSERT INTO events (info)
VALUES
 (
 '{ "location": "London", "event": "Dentist", "timestamp": "Sat 13 May 2017 11:36:07 BST", "_meta": {"user": 1, "client_id": 1}}'
 );