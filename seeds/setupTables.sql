CREATE DATABASE Twitter;

USE Twitter;

CREATE TABLE tweets(
  id int(4) auto_increment PRIMARY KEY,
  authorId int(4),
  FOREIGN key(authorId) references users(userId),
  dateTweeted VARCHAR(255),
  tweetBody VARCHAR(255),
  likes int(4),
  retweets int (4)
);

CREATE TABLE users(
  userId INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255),
  email VARCHAR(255),
  dateCreated VARCHAR(255),
  displayPic VARCHAR(255)
);

CREATE TABLE followings(
  id int(4) auto_increment PRIMARY KEY,
  followerId int(4),
  foreign key(followerId) references users(userId),
  dateFollowed VARCHAR(255)
);

CREATE TABLE likes(
	id int(4) auto_increment PRIMARY KEY, 
  tweetId int(4),
  likerId int(4),
  foreign key(likerId) references users(userId),
  dateLiked VARCHAR(255)
);

SHOW TABLES;
