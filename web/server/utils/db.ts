import postgres from "postgres"
import { drizzle } from "drizzle-orm/postgres-js"
import * as authSchema from "../db/auth-schema"

const connectionString = process.env.DATABASE_URL!
const client = postgres(connectionString)
export const db = drizzle(client, { schema: authSchema })
