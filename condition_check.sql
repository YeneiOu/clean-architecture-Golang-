CREATE TABLE tiki_bar(
	bar_id SERIAL PRIMARY KEY,
	birthdate DATE NOT NULL,
	age SMALLINT CHECK(age >= 21),
	payment INTEGER CHECK(payment >= 5)
);
