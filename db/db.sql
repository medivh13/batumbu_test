-- CREATE TABLE l_account (
-- 	id serial,
-- 	account_number varchar(50),
-- 	customer_number varchar(50),
-- 	balance numeric
-- );

-- CREATE TABLE l_customer (
-- 	id serial,
-- 	customer_number varchar(50),
-- 	customer_name varchar(50)
-- );


-- CREATE TABLE l_transaction (
-- 	id serial,
-- 	request_id varchar(50),
-- 	from_account varchar(50),
-- 	to_account varchar(50),
-- 	amount numeric,
-- 	UNIQUE(request_id)
-- );


-- insert into l_account 
-- (account_number,customer_number,balance)
-- values
-- ('555001','1001',100000),
-- ('555002','1002',150000)
-- ;


-- insert into l_customer
-- (customer_number,customer_name)
-- values 
-- ('1001','Bob martin'),
-- ('1002','Linus Torvalds');