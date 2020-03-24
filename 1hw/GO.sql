drop database if exists Site;
create database Site;
use Site;

drop table if exists Blogs;
CREATE TABLE Blogs (
  id int NOT NULL primary key AUTO_INCREMENT,
  name TEXT NOT NULL,
  description TEXT NULL, 
   UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE
);

drop table if exists Post;
CREATE TABLE Post (
  id INT NOT NULL primary key AUTO_INCREMENT,
  header TEXT NOT NULL,
  text TEXT NOT NULL,
  UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE
);

insert into Blogs
  (name, description)
values
  ("My Blog", "description");
insert into Post
  (header, text)
values
  ("1 post", "1 text"),
  ("2 post", "2 text"),
  ("3 post", "3 text"),
  ("4 post", "4 text"),
  ("5 post", "5 text");
  -- sudo chmod 666 /var/run/docker.sock
  -- sudo docker run --name go -e MYSQL_ROOT_PASSWORD=12345 -p 3306:3306 -d mysql
  -- sudo docker exec -it go bash
  -- mysql -u root -p