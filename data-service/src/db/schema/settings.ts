import { pgTable, uuid, varchar, text, timestamp } from 'drizzle-orm/pg-core'

export const settings = pgTable('settings', {
  id: uuid('id').primaryKey().defaultRandom(),
  key: varchar('key', { length: 50 }).unique().notNull(),
  value: text('value').notNull(),
  description: text('description'),
  updatedAt: timestamp('updated_at').defaultNow().notNull(),
})
