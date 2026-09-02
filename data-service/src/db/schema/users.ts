import { pgTable, uuid, varchar, timestamp } from 'drizzle-orm/pg-core'
import { members } from './members.js'

export const users = pgTable('users', {
  id: uuid('id').primaryKey().defaultRandom(),
  memberId: uuid('member_id'), // can't reference directly if not initialized first, using type only for circular dependency
  email: varchar('email', { length: 100 }).unique().notNull(),
  role: varchar('role', { length: 20 }).notNull(),
  status: varchar('status', { length: 20 }).default('active').notNull(),
  createdAt: timestamp('created_at').defaultNow().notNull(),
  updatedAt: timestamp('updated_at').defaultNow().notNull(),
})
