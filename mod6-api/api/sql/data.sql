insert into users (name, nick, email, password)
values
('User 1', 'user1', 'user1@mail.com', '$2a$10$JbnArVqv9SHFW0tL/yl8a.KbxoT/mzaIbQATYEDfZw6g5ZH4L4W7C'),
('User 2', 'user2', 'user2@mail.com', '$2a$10$JbnArVqv9SHFW0tL/yl8a.KbxoT/mzaIbQATYEDfZw6g5ZH4L4W7C'),
('User 3', 'user3', 'user3@mail.com', '$2a$10$JbnArVqv9SHFW0tL/yl8a.KbxoT/mzaIbQATYEDfZw6g5ZH4L4W7C'),
('User 4', 'user4', 'user4@mail.com', '$2a$10$JbnArVqv9SHFW0tL/yl8a.KbxoT/mzaIbQATYEDfZw6g5ZH4L4W7C'),
('User 5', 'user5', 'user5@mail.com', '$2a$10$JbnArVqv9SHFW0tL/yl8a.KbxoT/mzaIbQATYEDfZw6g5ZH4L4W7C'),
('User 6', 'user6', 'user6@mail.com', '$2a$10$JbnArVqv9SHFW0tL/yl8a.KbxoT/mzaIbQATYEDfZw6g5ZH4L4W7C'),
('User 7', 'user7', 'user7@mail.com', '$2a$10$JbnArVqv9SHFW0tL/yl8a.KbxoT/mzaIbQATYEDfZw6g5ZH4L4W7C'),
('User 8', 'user8', 'user8@mail.com', '$2a$10$JbnArVqv9SHFW0tL/yl8a.KbxoT/mzaIbQATYEDfZw6g5ZH4L4W7C'),
('User 9', 'user9', 'user9@mail.com', '$2a$10$JbnArVqv9SHFW0tL/yl8a.KbxoT/mzaIbQATYEDfZw6g5ZH4L4W7C'),
('User 10', 'user10', 'user10@mail.com', '$2a$10$JbnArVqv9SHFW0tL/yl8a.KbxoT/mzaIbQATYEDfZw6g5ZH4L4W7C');

insert into followers (user_id, follower_id)
values
(1, 2),
(1, 3),
(1, 5),
(1, 10),
(2, 3),
(2, 5),
(3, 5),
(3, 7),
(4, 5),
(6, 7),
(8, 2),
(8, 5),
(9, 5),
(9, 6),
(10, 3);
