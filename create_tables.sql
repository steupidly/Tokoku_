show databases;	 -- tampilkan database
use tokoku_project; -- pakai database

-- DROP 
DROP TABLE transactions;
DROP TABLE customers;
DROP TABLE items;
DROP TABLE employees;

CREATE TABLE employees (
	id_employee int auto_increment,
	username varchar(255) NOT NULL,
	name varchar(255) NOT NULL,
	password varchar(255) NOT NULL,
	primary key (id_employee)
);


CREATE TABLE items (
	id_item int auto_increment,
	id_employee int NOT NULL,
	item_name varchar(255) NOT NULL,
	quantity int DEFAULT NULL,
	primary key (id_item),
	CONSTRAINT fk_employees_items
	FOREIGN KEY (id_employee) references employees(id_employee)
);

CREATE TABLE customers (
	no_hp varchar(255) NOT NULL,
	id_employee int NOT NULL,
	customer_name varchar(255),
	PRIMARY KEY (no_hp),
	CONSTRAINT fk_employees_customers
	FOREIGN KEY (id_employee) references employees(id_employee)
);

CREATE TABLE transactions (
	id_transaction int NOT NULL,
	id_employee int NOT NULL,
	id_item int NOT NULL,
	no_hp varchar(255) NOT NULL,
	create_date datetime NOT NULL,
	PRIMARY key (id_transaction),
	CONSTRAINT fk_employee_transactions FOREIGN KEY (id_employee) REFERENCES employees(id_employee),
	CONSTRAINT fk_items_transactions FOREIGN KEY (id_item) REFERENCES items(id_item)
	
);

CREATE TABLE item_transaction(
	id_item int,
	id_transaction int,
	quantity int,
	PRIMARY KEY (id_item, id_transaction),
	CONSTRAINT fk_items_item_trasactions FOREIGN KEY (id_item) REFERENCES items(id_item),
	CONSTRAINT fk_transactions_item_transactions FOREIGN KEY (id_transaction) REFERENCES transactions(id_transaction)
	
);



PRIMARY KEY (id_user, id_activity)


-- INSERT 

insert into users 
values (1,"Thomas","thomasoke");

insert into users(user_name,password)
values	("Gianto","patenkali")



