CREATE KEYSPACE etl WITH replication = { 'class' : 'SimpleStrategy', 
'replication_factor' : 1};

USE etl;

CREATE TABLE login_events ( 
	id TIMEUUID PRIMARY KEY,
	event_type varchar, 
	user_id varchar,
	epoch timestamp, 

	device varchar
);

CREATE TABLE purchase_events ( 
	id TIMEUUID PRIMARY KEY,
	event_type varchar, 
	user_id varchar,  
	epoch timestamp, 

	item_id varchar,
	amount float
);


CREATE TABLE levelup_events ( 
	id TIMEUUID PRIMARY KEY,
	event_type varchar, 
	user_id varchar,  
	epoch timestamp, 

	level tinyint
);


