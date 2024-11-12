import { drizzle } from "drizzle-orm/libsql";
import { config } from "dotenv";
import findConfig from 'find-config';

// will recursively try to find the .env file
config({ path: findConfig('.env') });

export const db = drizzle({
  connection: {
    url: process.env.TURSO_CONNECTION_URL!,
    authToken: process.env.TURSO_AUTH_TOKEN!,
  }
});
