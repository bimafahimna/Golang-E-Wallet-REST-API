INSERT INTO users
(username,email,password,created_at,updated_at)
VALUES
('jojoYoyo','jojo@gmail.com','$2a$10$Ntf6gLYodD1n38VMkr7mW.lG/xyqb1XlkAaKHzXvlcTUtMHY/nE2e',now()- interval '4 years',now()- interval '4 years'),
('bobiRand','bobi99@gmail.com','$2a$10$HkOD9Dm0pDpbohjI0M6DhOb/b4A3rCGd.eU4Hu3EMG9a5cLiq4WbK',now()- interval '4 years',now()- interval '4 years'),
('creepSlayer','yurnero@gmail.com','$2a$10$qs7z4TIcM31zz2z8hdpSOOOMSg9p0HzMI77Bc5GE9IoiGSDDXERmG',now()- interval '4 years',now()- interval '4 years'),
('XFusion','paperrex@gmail.com','$2a$10$rbGf8sSsMvpw32rY/Sj/1u2iUvSip48NHVMvygjKywX/39EWWM8n6',now()- interval '4 years',now()- interval '4 years'),
('Rody','ironroxxx@gmail.com','$2a$10$keQAQCoVCJByipc3HxpQk.wEaba1xM1Clj1FOvr8bIIjd13s3zhFa',now()- interval '4 years',now()- interval '4 years')
;

INSERT INTO wallets
(user_id,balance,created_at,updated_at)
VALUES
(1,123000,now()- interval '4 years',now()- interval '4 years'),
(2,500000,now()- interval '4 years',now()- interval '4 years'),
(3,120000,now()- interval '4 years',now()- interval '4 years'),
(4,230000,now()- interval '4 years',now()- interval '4 years'),
(5,10000000,now()- interval '4 years',now()- interval '4 years')
;

INSERT INTO reset_password_tokens
(user_id,created_at)
VALUES
(1,now()- interval '4 years'),
(2,now()- interval '4 years'),
(3,now()- interval '4 years'),
(4,now()- interval '4 years'),
(5,now()- interval '4 years')
;

INSERT INTO source_of_funds
(name)
VALUES
('Bank Transfer'),
('Credit Card'),
('Cash'),
('Reward');

INSERT INTO transactions
(source_of_fund_id, from_wallet_id, to_wallet_id, amount, description,created_at,updated_at)
VALUES
(NULL,2,1,100000.999,'Give Away',now()- interval '2 years',now()- interval '2 years'),
(NULL,2,1,5000.129,'Give Away again ',now()- interval '1 years',now()- interval '1 years'),
(NULL,3,4,100000,'Debt',now()- interval '2 months',now()- interval '2 months'),
(NULL,1,4,50000.99,'Debt',now()- interval '2 months',now()- interval '2 months'),
(NULL,5,3,23500,'Toys',now()- interval '4 months',now()- interval '4 months'),
(NULL,1,5,100000.00,'Fried Rice',now()- interval '2 years',now()- interval '2 years'),
(NULL,2,1,3100,'',now()- interval '1 years',now()- interval '1 years'),
(NULL,3,4,29000,'',now()- interval '4 months',now()- interval '4 months'),
(NULL,1,4,12000,'',now()- interval '3 months',now()- interval '3 months'),
(NULL,5,3,50000,'',now()- interval '1 months',now()- interval '1 months'),
(1,NULL,1,900000,'Top Up from Bank Transfer',now()- interval '2 years',now()- interval '2 years'),
(1,NULL,2,310000,'Top Up from Bank Transfer',now()- interval '1 years',now()- interval '1 years'),
(3,NULL,3,2900000,'Top Up from Cash',now()- interval '20 days',now()- interval '20 days'),
(2,NULL,4,1200000,'Top Up from Credit Card',now()- interval '1 months',now()- interval '1 months'),
(3,NULL,5,5000000,'Top Up from Cash',now()- interval '10 months',now()- interval '10 months'),
(4,NULL,5,1,'Top Up from Reward',now()- interval '1 years',now()- interval '1 years'),
(4,NULL,4,100000,'Top Up from Reward',now()- interval '1 months',now()- interval '1 months'),
(2,NULL,2,291000,'Top Up from Credit Card',now()- interval '20 days',now()- interval '20 days'),
(2,NULL,3,122000,'Top Up from Credit Card',now()- interval '1 months',now()- interval '1 months'),
(3,NULL,1,5100000,'Top Up from Cash',now()- interval '9 months',now()- interval '9 months')
;