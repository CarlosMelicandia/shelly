CREATE TABLE `events_tracker` (
	`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	`event_id` text NOT NULL,
	`timestamp` integer NOT NULL,
	`user_id` text NOT NULL
);
--> statement-breakpoint
CREATE UNIQUE INDEX `events_tracker_user_id_event_id_unique` ON `events_tracker` (`user_id`,`event_id`);--> statement-breakpoint
CREATE TABLE `hacker_applications` (
	`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	`user_id` text NOT NULL,
	`first_name` text NOT NULL,
	`last_name` text NOT NULL,
	`age` integer NOT NULL,
	`school` text NOT NULL,
	`major` text NOT NULL,
	`grad_year` integer NOT NULL,
	`level_of_study` text NOT NULL,
	`country` text NOT NULL,
	`email` text NOT NULL,
	`phone_number` text NOT NULL,
	`resume_path` text NOT NULL,
	`github` text,
	`linkedin` text,
	`is_international` integer NOT NULL,
	`gender` text NOT NULL,
	`pronouns` text NOT NULL,
	`ethnicity` text NOT NULL,
	`avatar` integer DEFAULT 0 NOT NULL,
	`agreed_mlh_news` integer DEFAULT false NOT NULL,
	`application_status` text DEFAULT 'registered' NOT NULL,
	`created_at` integer NOT NULL
);
--> statement-breakpoint
CREATE UNIQUE INDEX `hacker_applications_user_id_unique` ON `hacker_applications` (`user_id`);--> statement-breakpoint
CREATE UNIQUE INDEX `hacker_applications_email_unique` ON `hacker_applications` (`email`);--> statement-breakpoint
CREATE TABLE `user` (
	`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL,
	`user_id` text NOT NULL,
	`name` text NOT NULL,
	`email` text NOT NULL,
	`discord_username` text,
	`isAdmin` integer DEFAULT false,
	`isVolunteer` integer DEFAULT false,
	`isSponsor` integer DEFAULT false,
	`isMentor` integer DEFAULT false,
	`hacker_id` integer
);
--> statement-breakpoint
CREATE UNIQUE INDEX `user_email_unique` ON `user` (`email`);