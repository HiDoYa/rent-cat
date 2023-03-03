INSERT INTO expense_type (type_name, created_at) VALUES ('Electricity', '2022-03-02');
INSERT INTO expense_type (type_name, created_at) VALUES ('Gas', '2022-03-02');

INSERT INTO expense (expense_type, amount, created_at) VALUES ('Electricity', 10.05, '2022-03-02');
INSERT INTO expense (expense_type, amount, created_at) VALUES ('Gas', 20.61, '2022-03-02');
INSERT INTO expense (expense_type, amount, created_at) VALUES ('Electricity', 123.45, '2022-09-02');

INSERT INTO split (my_percentage, her_percentage, created_at) VALUES (68.5, 31.5, '2022-03-02');
INSERT INTO split (my_percentage, her_percentage, created_at) VALUES (75, 25, '2022-03-03');