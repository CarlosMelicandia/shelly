import { sqliteTable, integer, text, unique } from 'drizzle-orm/sqlite-core';

export const applicationStatusEnums = ['registered', 'in_wave', 'accepted', 'confirmed', 'withdrawn', 'waitlisted', 'checked_in'] as const;

export const hackerApplications = sqliteTable('hacker_applications', {
  id: integer('id').primaryKey({ autoIncrement: true }),
  userId: text('user_id').unique().notNull(),
  firstName: text('first_name').notNull(),
  lastName: text('last_name').notNull(),
  age: integer('age').notNull(),
  school: text('school').notNull(),
  major: text('major').notNull(),
  gradYear: integer('grad_year').notNull(),
  levelOfStudy: text('level_of_study').notNull(),
  country: text('country').notNull(),
  // discuss to see if we want to keep email
  email: text('email').unique().notNull(),
  phoneNumber: text('phone_number').notNull(),
  resumePath: text('resume_path').notNull(),
  github: text('github'),
  linkedin: text('linkedin'),
  isInternational: integer('is_international', { mode: 'boolean' }).notNull(),
  gender: text('gender').notNull(),
  pronouns: text('pronouns').notNull(),
  ethnicity: text('ethnicity').notNull(),
  avatar: integer('avatar').default(0).notNull(),
  agreedMLHNnews: integer('agreed_mlh_news', { mode: 'boolean' }).default(false).notNull(),
  applicationStatus: text('application_status').default('registered').notNull(),
  createdAt: integer('created_at', { mode: 'timestamp' }).notNull(),
});

export const eventsTracker = sqliteTable('events_tracker', {
  id: integer('id').primaryKey({ autoIncrement: true }),
  eventId: text('event_id').notNull(),
  timestamp: integer({ mode: 'timestamp' }).notNull(),
  userId: text('user_id').notNull(),
},
  (t) => ({
    userEventId: unique().on(t.userId, t.eventId)
  }));

export const user = sqliteTable('user', {
  id: text('id').primaryKey(),
  name: text('name').notNull(),
  email: text('email').unique().notNull(),
  discordUsername: text('discord_username'),
  admin: integer({ mode: 'boolean' }).default(false),
  hackerId: integer('hacker_id'),
});

