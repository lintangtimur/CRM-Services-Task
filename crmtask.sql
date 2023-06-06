-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 06, 2023 at 09:14 PM
-- Server version: 10.4.14-MariaDB
-- PHP Version: 7.4.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `crmtask`
--

-- --------------------------------------------------------

--
-- Table structure for table `actors`
--

CREATE TABLE `actors` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role_id` int(10) UNSIGNED DEFAULT NULL,
  `verified` enum('true','false') DEFAULT 'false',
  `active` enum('true','false') DEFAULT 'false',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `actors`
--

INSERT INTO `actors` (`id`, `username`, `password`, `role_id`, `verified`, `active`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'superadmin', '$2a$10$d0zgM0ZFmeCnLINc/qQhUO.QGW7PBv7L6WF.DPtKWaNspdvj0JWDO', 1, 'true', 'true', NULL, NULL, NULL),
(23, 'admin', '$2a$10$d0zgM0ZFmeCnLINc/qQhUO.QGW7PBv7L6WF.DPtKWaNspdvj0JWDO', 2, 'false', 'false', '2023-06-06 08:21:24', '2023-06-06 10:48:02', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `first_name`, `last_name`, `email`, `avatar`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'lintang', 'qwerty', 'lintangtimur@gmail.com', 'https://google.com', '2023-06-04 06:28:25', '2023-06-04 06:28:25', NULL),
(4, 'lintang', 'qwerty', 'lintangtimur2@gmail.com', 'https://google.com', '2023-06-05 08:19:31', '2023-06-05 08:19:31', NULL),
(5, 'stefanus', 'litang', 'stefanus@gmail.com', 'https://google.com', '2023-06-06 07:54:44', '2023-06-06 07:54:44', NULL),
(8, 'stefanus2', 'litang', 'stefanus2@gmail.com', 'https://google.com', '2023-06-06 07:57:20', '2023-06-06 07:57:20', NULL),
(10, 'stefanus3', 'litang', 'stefanus3@gmail.com', 'https://google.com', '2023-06-06 10:57:13', '2023-06-06 10:57:13', '2023-06-06 10:58:51'),
(11, 'Michael', 'Lawson', 'michael.lawson@reqres.in', 'https://reqres.in/img/faces/7-image.jpg', '2023-06-06 12:11:05', '2023-06-06 12:11:05', NULL),
(12, 'Lindsay', 'Ferguson', 'lindsay.ferguson@reqres.in', 'https://reqres.in/img/faces/8-image.jpg', '2023-06-06 12:11:05', '2023-06-06 12:11:05', NULL),
(13, 'Tobias', 'Funke', 'tobias.funke@reqres.in', 'https://reqres.in/img/faces/9-image.jpg', '2023-06-06 12:11:05', '2023-06-06 12:11:05', NULL),
(14, 'Byron', 'Fields', 'byron.fields@reqres.in', 'https://reqres.in/img/faces/10-image.jpg', '2023-06-06 12:11:05', '2023-06-06 12:11:05', NULL),
(15, 'George', 'Edwards', 'george.edwards@reqres.in', 'https://reqres.in/img/faces/11-image.jpg', '2023-06-06 12:11:05', '2023-06-06 12:11:05', NULL),
(16, 'Rachel', 'Howell', 'rachel.howell@reqres.in', 'https://reqres.in/img/faces/12-image.jpg', '2023-06-06 12:11:05', '2023-06-06 12:11:05', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `register_approvals`
--

CREATE TABLE `register_approvals` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `admin_id` bigint(20) UNSIGNED DEFAULT NULL,
  `super_admin_id` bigint(20) UNSIGNED DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `register_approvals`
--

INSERT INTO `register_approvals` (`id`, `admin_id`, `super_admin_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(2, 23, 1, 'rejected', '2023-06-06 08:21:24', '2023-06-06 10:30:38', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` int(10) UNSIGNED NOT NULL,
  `role_name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `role_name`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'superadmin', NULL, NULL, NULL),
(2, 'admin', NULL, NULL, NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `actors`
--
ALTER TABLE `actors`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD KEY `role_id` (`role_id`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `register_approvals`
--
ALTER TABLE `register_approvals`
  ADD PRIMARY KEY (`id`),
  ADD KEY `admin_id` (`admin_id`),
  ADD KEY `super_admin_id` (`super_admin_id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `actors`
--
ALTER TABLE `actors`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;

--
-- AUTO_INCREMENT for table `register_approvals`
--
ALTER TABLE `register_approvals`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `actors`
--
ALTER TABLE `actors`
  ADD CONSTRAINT `actors_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

--
-- Constraints for table `register_approvals`
--
ALTER TABLE `register_approvals`
  ADD CONSTRAINT `register_approvals_ibfk_1` FOREIGN KEY (`admin_id`) REFERENCES `actors` (`id`),
  ADD CONSTRAINT `register_approvals_ibfk_2` FOREIGN KEY (`super_admin_id`) REFERENCES `actors` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
