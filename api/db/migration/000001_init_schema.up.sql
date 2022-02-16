CREATE TABLE `teachers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `first_name` longtext,
  `first_name_kana` longtext,
  `last_name` longtext,
  `last_name_kana` longtext,
  `phone_number` varchar(191) DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `address` longtext,
  `hashed_password` varchar(191) DEFAULT NULL,
  `image` longtext,
  `password_changed_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone_number` (`phone_number`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `hashed_password` (`hashed_password`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `students` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `first_name` longtext,
  `first_name_kana` longtext,
  `last_name` longtext,
  `last_name_kana` longtext,
  `phone_number` varchar(191) DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `address` longtext,
  `hashed_password` varchar(191) DEFAULT NULL,
  `image` longtext,
  `stripe_id` longtext,
  `password_changed_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone_number` (`phone_number`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `hashed_password` (`hashed_password`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `student_lecture_schedules` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `student_id` bigint unsigned DEFAULT NULL,
  `start_time` datetime(3) DEFAULT NULL,
  `end_time` datetime(3) DEFAULT NULL,
  `status` enum('empty','pending','reserved','finish','absent') DEFAULT 'empty',
  PRIMARY KEY (`id`),
  KEY `fk_students_student_lecture_schedules` (`student_id`),
  CONSTRAINT `fk_students_student_lecture_schedules` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `teacher_lecture_schedules` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `teacher_id` bigint unsigned DEFAULT NULL,
  `start_time` datetime(3) DEFAULT NULL,
  `end_time` datetime(3) DEFAULT NULL,
  `status` enum('empty','pending','reserved','finish','absent') DEFAULT 'empty',
  PRIMARY KEY (`id`),
  KEY `fk_teachers_teacher_lecture_schedules` (`teacher_id`),
  CONSTRAINT `fk_teachers_teacher_lecture_schedules` FOREIGN KEY (`teacher_id`) REFERENCES `teachers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `lectures` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `teacher_id` bigint unsigned DEFAULT '0',
  `student_id` bigint unsigned DEFAULT '0',
  `start_time` datetime(3) DEFAULT NULL,
  `end_time` datetime(3) DEFAULT NULL,
  `status` enum('empty','pending','reserved','finish','absent') DEFAULT 'empty',
  `zoom_link` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

