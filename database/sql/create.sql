DROP DATABASE IF EXISTS goserver;
CREATE DATABASE goserver DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE goserver;

CREATE TABLE `t_User` ( \
    `userId` int unsigned NOT NULL AUTO_INCREMENT, \
    `email` varchar(255) NOT NULL, \
    `password` varchar(255) NOT NULL, \
    `nickName` varchar(255), \
    `headURL` varchar(255), \
    `birthday` DOUBLE, \
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
) AUTO_INCREMENT=10000;
