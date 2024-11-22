ALTER TABLE `hacker_applications` RENAME TO `hacker_application`;--> statement-breakpoint
DROP INDEX IF EXISTS `hacker_applications_user_id_unique`;--> statement-breakpoint
DROP INDEX IF EXISTS `hacker_applications_email_unique`;--> statement-breakpoint
CREATE UNIQUE INDEX `hacker_application_user_id_unique` ON `hacker_application` (`user_id`);--> statement-breakpoint
CREATE UNIQUE INDEX `hacker_application_email_unique` ON `hacker_application` (`email`);