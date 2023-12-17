-- Create "authors" table
CREATE TABLE `authors` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `email` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "books" table
CREATE TABLE `books` (`id` bigint NOT NULL AUTO_INCREMENT, `title` varchar(255) NOT NULL, `genre` varchar(255) NOT NULL, `publication_date` varchar(255) NOT NULL, `isbn` varchar(255) NOT NULL, `created_at` timestamp NULL, `author_books` bigint NULL, PRIMARY KEY (`id`), INDEX `books_authors_books` (`author_books`), CONSTRAINT `books_authors_books` FOREIGN KEY (`author_books`) REFERENCES `authors` (`id`) ON UPDATE RESTRICT ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
