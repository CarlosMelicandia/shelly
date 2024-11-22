import { hackerApplications, eventsTracker, user, applicationStatusEnums } from './schema';
import { db } from './index';

async function seedDatabase() {
  const firstNames = ["Gabriel", "John", "Jack", "Kyle", "Bruna", "jacob", "david", "Jose", "osmany", "Morgan", "Allen", "Nathan", "chris", "Christian", "Dennis"]
  const lastNames = ["Pedroza", "Smith", "Saint", "Jones", "gomez", "Doe", "yi", "Yang", "Dennis", "Alexander", "Alexandra", "Gentil"]

  // reseed db
  await Promise.all([
    db.delete(user),
    db.delete(hackerApplications),
    db.delete(eventsTracker)
  ]);

  // keep in mind that this affects storage quota for the database. keep the total iterations at a moderate number
for (let i = 1; i <= 2; i++) {
  const firstName = firstNames[i % firstNames.length];
  const lastName = lastNames[i % lastNames.length];
  const date = new Date();

  try {
    await db.insert(user).values({
      userId: i,
      firstName,
      lastName,
      email: `${firstName}${lastName}${i}@gmail.com`,
      discordId: `${firstName}${lastName}${i}`,
    });

    await db.insert(hackerApplications).values({
      userId: i,
      firstName,
      lastName,
      age: ~~(Math.random() * 50) + 1,
      school: 'Sample University',
      major: 'Computer Science',
      gradYear: 2024,
      levelOfStudy: 'Undergraduate',
      country: 'USA',
      email: `${firstName}${lastName}${i}@gmail.com`,
      phoneNumber: '1234567890',
      resumePath: 'some_s3_bucket_path',
      github: 'https://github.com/johndoe',
      linkedin: 'https://linkedin.com/in/johndoe',
      isInternational: false,
      gender: 'Male',
      pronouns: 'He/Him',
      ethnicity: 'Non-Hispanic White',
      avatar: 1,
      agreedMLHNews: true,
      applicationStatus: applicationStatusEnums[~~(Math.random() * 6)],
      createdAt: date,
      updatedAt: date,
    });

    await db.insert(eventsTracker).values({
      eventId: i,
      userId: 'user-123',
      timestamp: date,
    });

  } catch (error) {
    console.error('Error seeding the database:', error);
  }
}
    console.log("Successfully seeded database! 🌱")
}

seedDatabase();
