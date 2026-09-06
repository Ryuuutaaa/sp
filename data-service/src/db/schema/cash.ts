import { pgTable, uuid, varchar, numeric, text, timestamp } from 'drizzle-orm/pg-core'

export const cashTransactions = pgTable('cash_transactions', {
  id: uuid('id').primaryKey().defaultRandom(),
  type: varchar('type', { length: 10 }).notNull(),
  category: varchar('category', { length: 50 }).notNull(),
  amount: numeric('amount', { precision: 15, scale: 2 }).notNull(),
  description: text('description').notNull(),
  referenceType: varchar('reference_type', { length: 50 }),
  referenceId: uuid('reference_id'),
  createdBy: text('created_by').notNull(),
  createdAt: timestamp('created_at').defaultNow().notNull(),
})
