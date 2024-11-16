CREATE TABLE transactions (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `userId` BIGINT UNSIGNED NOT NULL,
    `createdAt` DATETIME NOT NULL,
    `updatedAt` DATETIME NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `category` VARCHAR(100) NOT NULL,
    `amount` DECIMAL(10,2) NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_transaction`
        FOREIGN KEY (`userId`) REFERENCES users(`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);