CREATE TABLE membership(
	member_id SERIAL PRIMARY KEY,
	username VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR(50) NOT NULL,
	email VARCHAR(200) UNIQUE NOT NULL,
	member_since TIMESTAMP NOT NULL,
	last_visit TIMESTAMP
);

CREATE TABLE gold_member(
	gold_id SERIAL PRIMARY KEY,
	level VARCHAR(25) NOT NULL
);

CREATE TABLE member_rewards(
	member_id INTEGER REFERENCES membership(member_id),
	gold_id INTEGER REFERENCES gold_member(gold_id),
	award_date TIMESTAMP 
);
