import { hackerApplications, eventsTracker, user, applicationStatusEnums } from './schema';
import { db } from './index';

async function seedDatabase() {

  // reseed db
  await Promise.all([
    db.delete(user),
    db.delete(hackerApplications),
    db.delete(eventsTracker)
  ]);

  try {
    await db.insert(user).values({
      userId: '123',
      name: 'John Doe',
      email: 'johndoe@example.com',
      discordUsername: 'johnnyD',
    });

    await db.insert(hackerApplications).values({
      userId: 'user-123',
      firstName: 'John',
      lastName: 'Doe',
      age: 22,
      school: 'Sample University',
      major: 'Computer Science',
      gradYear: 2024,
      levelOfStudy: 'Undergraduate',
      country: 'USA',
      email: 'johndoe@example.com',
      phoneNumber: '+1234567890',
      resumePath: '/resumes/johndoe.pdf',
      github: 'https://github.com/johndoe',
      linkedin: 'https://linkedin.com/in/johndoe',
      isInternational: false,
      gender: 'Male',
      pronouns: 'He/Him',
      ethnicity: 'Non-Hispanic White',
      avatar: 1,
      agreedMLHNews: true,
      applicationStatus: applicationStatusEnums[0], // 'registered'
      createdAt: new Date()
    });

    await db.insert(eventsTracker).values({
      eventId: 'event-001',
      userId: 'user-123',
      timestamp: new Date()
    });

    console.log('Database seeded successfully.');
  } catch (error) {
    console.error('Error seeding the database:', error);
  }
}

seedDatabase();
