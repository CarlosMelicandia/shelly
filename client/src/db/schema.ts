import { sqliteTable, integer, text, unique } from 'drizzle-orm/sqlite-core';

// Enum equivalent for application_status_enums
export const application_status_enums = ['registered', 'in_wave', 'accepted', 'confirmed', 'withdrawn', 'waitlisted', 'checked_in'] as const;

export const hacker_applications = sqliteTable('hacker_applications', {
  id: integer('id').primaryKey({ autoIncrement: true }),
  userId: text('userId').unique().notNull(),
  first_name: text('first_name').notNull(),
  last_name: text('last_name').notNull(),
  age: integer('age').notNull(),
  school: text('school').notNull(),
  major: text('major').notNull(),
  grad_year: integer('grad_year').notNull(),
  level_of_study: text('level_of_study').notNull(),
  country: text('country').notNull(),
  // discuss to see if we want to keep email
  email: text('email').unique().notNull(),
  phone_number: text('phone_number').notNull(),
  resume_path: text('resume_path').notNull(),
  github: text('github'),
  linkedin: text('linkedin'),
  is_international: integer({ mode: 'boolean' }).notNull(),
  gender: text('gender').notNull(),
  pronouns: text('pronouns').notNull(),
  ethnicity: text('ethnicity').notNull(),
  dinosaur_avatar: integer('dinosaur_avatar').default(0).notNull(),
  agreed_mlh_news: integer({ mode: 'boolean' }).default(false).notNull(),
  application_status: text('application_status').default('registered').notNull(),
  created_at: integer({ mode: 'timestamp' }).notNull(),
});

export const events_tracker = sqliteTable('events_tracker', {
  id: integer('id').primaryKey({ autoIncrement: true }),
  event_id: text('event_id').notNull(),
  timestamp: integer({ mode: 'timestamp' }).notNull(),
  user_id: text('user_id').notNull(),
},
  (t) => ({
    user_event_id: unique().on(t.user_id, t.event_id)
  }));

export const user = sqliteTable('user', {
  id: text('id').primaryKey(),
  name: text('name').notNull(),
  email: text('email').unique().notNull(),
  discordUsername: text('discordUsername'),
  admin: integer({ mode: 'boolean' }).default(false),
  hackerId: integer('hackerId'),
});

