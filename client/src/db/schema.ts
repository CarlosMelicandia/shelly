import { sqliteTable, integer, text, boolean, primaryKey, unique, date } from 'drizzle-orm/sqlite-core';

// Enum equivalent for application_status_enums
export const application_status_enums = ['registered', 'in_wave', 'accepted', 'confirmed', 'withdrawn', 'waitlisted', 'checked_in'] as const;

export const hacker_applications = sqliteTable('hacker_applications', {
  id: integer('id').primaryKey().autoIncrement(),
  userId: text('userId').unique().notNull(),
  first_name: text('first_name').notNull(),
  last_name: text('last_name').notNull(),
  age: integer('age').notNull(),
  school: text('school').notNull(),
  major: text('major').notNull(),
  grad_year: integer('grad_year').notNull(),
  level_of_study: text('level_of_study').notNull(),
  country: text('country').notNull(),
  email: text('email').unique().notNull(),
  phone_number: text('phone_number').notNull(),
  resume_path: text('resume_path').notNull(),
  github: text('github'),
  linkedin: text('linkedin'),
  is_international: boolean('is_international').notNull(),
  gender: text('gender').notNull(),
  pronouns: text('pronouns').notNull(),
  ethnicity: text('ethnicity').notNull(),
  dinosaur_avatar: integer('dinosaur_avatar').default(0).notNull(),
  agreed_mlh_news: boolean('agreed_mlh_news').notNull(),
  application_status: text('application_status').default('registered').notNull(),
  check_in_status: boolean('check_in_status').default(false).notNull(),
  created_at: date('created_at').default('CURRENT_TIMESTAMP').notNull(),
});

export const events_tracker = sqliteTable('events_tracker', {
  id: integer('id').primaryKey().autoIncrement(),
  event_id: text('event_id').notNull(),
  timestamp: date('timestamp').default('CURRENT_TIMESTAMP').notNull(),
  user_id: text('user_id').notNull(),
}, (table) => ({
  uniqueEventUser: unique('event_id', 'user_id')
}));

export const user = sqliteTable('user', {
  id: text('id').primaryKey(),
  name: text('name'),
  email: text('email').unique(),
  emailVerified: date('emailVerified'),
  image: text('image'),
  discordUsername: text('discordUsername'),
  admin: boolean('admin').default(false),
  hackerId: integer('hackerId'),
});

