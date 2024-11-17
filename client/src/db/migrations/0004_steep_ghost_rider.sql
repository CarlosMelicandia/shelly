DROP INDEX IF EXISTS "events_tracker_user_id_event_id_unique";--> statement-breakpoint
DROP INDEX IF EXISTS "hacker_applications_user_id_unique";--> statement-breakpoint
DROP INDEX IF EXISTS "hacker_applications_email_unique";--> statement-breakpoint
DROP INDEX IF EXISTS "user_email_unique";--> statement-breakpoint
ALTER TABLE `hacker_applications` ALTER COLUMN "github" TO "github" text DEFAULT '';--> statement-breakpoint
CREATE UNIQUE INDEX `events_tracker_user_id_event_id_unique` ON `events_tracker` (`user_id`,`event_id`);--> statement-breakpoint
CREATE UNIQUE INDEX `hacker_applications_user_id_unique` ON `hacker_applications` (`user_id`);--> statement-breakpoint
CREATE UNIQUE INDEX `hacker_applications_email_unique` ON `hacker_applications` (`email`);--> statement-breakpoint
CREATE UNIQUE INDEX `user_email_unique` ON `user` (`email`);--> statement-breakpoint
ALTER TABLE `hacker_applications` ALTER COLUMN "linkedin" TO "linkedin" text DEFAULT '';--> statement-breakpoint
ALTER TABLE `user` ALTER COLUMN "discord_username" TO "discord_username" text DEFAULT '';--> statement-breakpoint
ALTER TABLE `user` ALTER COLUMN "hacker_id" TO "hacker_id" integer DEFAULT 0;