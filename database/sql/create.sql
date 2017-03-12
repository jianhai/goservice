DROP DATABASE IF EXISTS goserver;
CREATE DATABASE goserver;
USE goserver;

CREATE TABLE `t_User` ( \
    `userId` int NOT NULL AUTO_INCREMENT, \
    `email` varchar(255) NOT NULL, \
    `password` varchar(255) NOT NULL, \
    `nickName` varchar(255), \
    `headURL` varchar(255), \
    `birthday` varchar(20), \
    `sex` int NOT NULL DEFAULT 0, \
    `isOnline` int, \
    `longitude` double, \
    `latitude` double, \
    `livingId` varchar(255), \
    `livingState` int, \
    `livingDuration` int, \
    `golden` int, \
    `gifts` int, \
    `wantSex` int, \
    `topicID` int, \
    PRIMARY KEY (`userId`)
);
