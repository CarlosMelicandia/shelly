import { hacker_applications, events_tracker, user, application_status_enums } from './schema';
import { db } from './index';

async function seedDatabase() {

  // reseed db
  await Promise.all([
    db.delete(user),
    db.delete(hacker_applications),
    db.delete(events_tracker)
  ]);

  try {
    await db.insert(user).values({
      id: 'user-123',
      name: 'John Doe',
      email: 'johndoe@example.com',
      discordUsername: 'johnnyD',
      admin: false,
    });

    await db.insert(hacker_applications).values({
      userId: 'user-123',
      first_name: 'John',
      last_name: 'Doe',
      age: 22,
      school: 'Sample University',
      major: 'Computer Science',
      grad_year: 2024,
      level_of_study: 'Undergraduate',
      country: 'USA',
      email: 'johndoe@example.com',
      phone_number: '+1234567890',
      resume_path: '/resumes/johndoe.pdf',
      github: 'https://github.com/johndoe',
      linkedin: 'https://linkedin.com/in/johndoe',
      is_international: false,
      gender: 'Male',
      pronouns: 'He/Him',
      ethnicity: 'Non-Hispanic White',
      dinosaur_avatar: 1,
      agreed_mlh_news: true,
      application_status: application_status_enums[0], // 'registered'
      check_in_status: false,
      created_at: new Date()
    });

    await db.insert(events_tracker).values({
      event_id: 'event-001',
      user_id: 'user-123',
      timestamp: new Date()
    });

    console.log('Database seeded successfully.');
  } catch (error) {
    console.error('Error seeding the database:', error);
  }
}

seedDatabase();
