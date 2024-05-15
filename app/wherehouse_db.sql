-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 15, 2024 at 06:26 AM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.1.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `wherehouse_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `clients`
--

CREATE TABLE `clients` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `clients`
--

INSERT INTO `clients` (`id`, `name`) VALUES
(201, 'Client A'),
(202, 'Client B'),
(203, 'Client C'),
(204, 'Client D'),
(205, 'Client E');

-- --------------------------------------------------------

--
-- Table structure for table `staffs`
--

CREATE TABLE `staffs` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `user_type` varchar(50) NOT NULL,
  `password` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `staffs`
--

INSERT INTO `staffs` (`id`, `name`, `user_type`, `password`) VALUES
(1, 'Joy', 'Staff', 'joy123'),
(2, 'Peter', 'Staff', 'peter123'),
(3, 'Matthew', 'Staff', 'matt123'),
(4, 'Harry', 'Owner', 'harry123');

-- --------------------------------------------------------

--
-- Table structure for table `units`
--

CREATE TABLE `units` (
  `id` bigint(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text DEFAULT NULL,
  `kg` double NOT NULL,
  `warehouse_id` bigint(20) NOT NULL,
  `client_id` int(11) NOT NULL,
  `order_date` date NOT NULL,
  `status` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `units`
--

INSERT INTO `units` (`id`, `name`, `description`, `kg`, `warehouse_id`, `client_id`, `order_date`, `status`) VALUES
(10, 'Rubicon', 'JEEP', 900, 10003, 201, '2024-03-01', 'pending'),
(11, 'GIGABYTE G24F', 'Monitor', 10, 10002, 202, '2024-02-17', 'completed'),
(12, 'Secret-Lab', 'Kursi Gaming', 30, 10002, 203, '2024-02-29', 'completed'),
(13, 'Sedan X', 'Luxury sedan with advanced safety features', 1800, 10001, 201, '2024-05-01', 'pending'),
(14, 'SUV Z', 'Off-road capable SUV with hybrid technology', 2200, 10002, 202, '2024-04-28', 'completed'),
(15, 'Sports Car Y', 'High-performance sports car with turbocharged engine', 1600, 10003, 203, '2024-05-03', 'pending'),
(16, 'Electric Car A', 'Compact electric car with autonomous driving features', 1500, 10004, 204, '2024-04-29', 'pending'),
(17, 'Truck B', 'Heavy-duty truck with towing capacity of 10 tons', 3500, 10001, 205, '2024-05-02', 'completed'),
(18, 'Hybrid Vehicle C', 'Mid-size hybrid vehicle with fuel-efficient engine', 1900, 10002, 201, '2024-04-30', 'pending'),
(19, 'Luxury SUV D', 'Premium SUV with leather interiors and panoramic sunroof', 2400, 10003, 202, '2024-05-05', 'completed'),
(20, 'Convertible E', 'Convertible sports car with retractable hardtop', 1700, 10004, 203, '2024-05-04', 'pending'),
(21, 'Electric Scooter F', 'Urban electric scooter with removable battery', 100, 10001, 204, '2024-05-01', 'completed'),
(22, 'Compact Car G', 'Affordable compact car with good fuel economy', 1200, 10002, 205, '2024-04-29', 'pending');

-- --------------------------------------------------------

--
-- Table structure for table `warehouses`
--

CREATE TABLE `warehouses` (
  `id` bigint(20) NOT NULL,
  `name` varchar(255) NOT NULL,
  `staff_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `warehouses`
--

INSERT INTO `warehouses` (`id`, `name`, `staff_id`) VALUES
(10001, 'Warehouse Jakarta', 1),
(10002, 'Warehouse Bandung', 2),
(10003, 'Warehouse Surabaya', 3),
(10004, 'Warehouse Palembang', 4);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `clients`
--
ALTER TABLE `clients`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `staffs`
--
ALTER TABLE `staffs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `units`
--
ALTER TABLE `units`
  ADD PRIMARY KEY (`id`),
  ADD KEY `warehouse_id` (`warehouse_id`),
  ADD KEY `client_id` (`client_id`);

--
-- Indexes for table `warehouses`
--
ALTER TABLE `warehouses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `staff_id` (`staff_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `clients`
--
ALTER TABLE `clients`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=206;

--
-- AUTO_INCREMENT for table `staffs`
--
ALTER TABLE `staffs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `units`
--
ALTER TABLE `units`
  MODIFY `id` bigint(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

--
-- AUTO_INCREMENT for table `warehouses`
--
ALTER TABLE `warehouses`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10005;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `units`
--
ALTER TABLE `units`
  ADD CONSTRAINT `units_ibfk_1` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`),
  ADD CONSTRAINT `units_ibfk_2` FOREIGN KEY (`client_id`) REFERENCES `clients` (`id`);

--
-- Constraints for table `warehouses`
--
ALTER TABLE `warehouses`
  ADD CONSTRAINT `warehouses_ibfk_1` FOREIGN KEY (`staff_id`) REFERENCES `staffs` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
