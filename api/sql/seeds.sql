INSERT INTO users (name, nickname, email, password)
VALUES -- senha: 123456
("usuario 1", "usuario1", "usuario1@gmail.com", "$2a$10$qIlOi3yehi4hHu.jsLAimehlOphlbATjerGINnmsbh3tjtewT6Mi6"),
("usuario 2", "usuario2", "usuario2@gmail.com", "$2a$10$qIlOi3yehi4hHu.jsLAimehlOphlbATjerGINnmsbh3tjtewT6Mi6"),
("usuario 3", "usuario3", "usuario3@gmail.com", "$2a$10$qIlOi3yehi4hHu.jsLAimehlOphlbATjerGINnmsbh3tjtewT6Mi6");

INSERT INTO followers(user_id, follower_id)
VALUES
(1, 2),
(1, 3),
(2, 1),
(2, 3);