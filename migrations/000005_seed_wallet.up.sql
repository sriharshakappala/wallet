INSERT INTO wallet (user_id, balance) VALUES ((select id from users where username = 'Alice'), 100);
INSERT INTO wallet (user_id, balance) VALUES ((select id from users where username = 'Bob'), 150);
INSERT INTO wallet (user_id, balance) VALUES ((select id from users where username = 'Candy'), 120);
INSERT INTO wallet (user_id, balance) VALUES ((select id from users where username = 'David'), 200);