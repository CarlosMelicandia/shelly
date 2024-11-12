// todo. fix this code since its gpt generated with running it
import { hacker_applications, events_tracker, user, application_status_enums } from './your-schema-file'; // Adjust the import to your schema file location
import { db } from './db'; // Assume you have a db instance already set up

async function seedDatabase() {
  try {
    // Insert a user
    await db.insert(user).values({
      id: 'user-123',
      name: 'John Doe',
      email: 'johndoe@example.com',
      emailVerified: new Date().toISOString(),
      image: 'https://example.com/johndoe.jpg',
      discordUsername: 'johnnyD',
      admin: false,
    });

    // Insert a hacker application
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
      created_at: new Date().toISOString(),
    });

    // Insert an event tracker
    await db.insert(events_tracker).values({
      event_id: 'event-001',
      user_id: 'user-123',
      timestamp: new Date().toISOString(),
    });

    console.log('Database seeded successfully.');
  } catch (error) {
    console.error('Error seeding the database:', error);
  }
}

seedDatabase();

