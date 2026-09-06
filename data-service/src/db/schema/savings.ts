import { pgTable, uuid, varchar, numeric, boolean, timestamp, text } from 'drizzle-orm/pg-core'
import { members } from './members.js'

export const savingsTypes = pgTable('savings_types', {
  id: uuid('id').primaryKey().defaultRandom(),
  name: varchar('name', { length: 50 }).notNull(),
  defaultAmount: numeric('default_amount', { precision: 15, scale: 2 }).notNull(),
  isMandatory: boolean('is_mandatory').notNull(),
  createdAt: timestamp('created_at').defaultNow().notNull(),
})

export const savingsTransactions = pgTable('savings_transactions', {
  id: uuid('id').primaryKey().defaultRandom(),
  memberId: uuid('member_id').references(() => members.id).notNull(),
  savingsTypeId: uuid('savings_type_id').references(() => savingsTypes.id).notNull(),
  type: varchar('type', { length: 10 }).notNull(),
  amount: numeric('amount', { precision: 15, scale: 2 }).notNull(),
  balanceAfter: numeric('balance_after', { precision: 15, scale: 2 }).notNull(),
  note: varchar('note', { length: 255 }),
  receiptNumber: varchar('receipt_number', { length: 50 }).unique().notNull(),
  createdBy: text('created_by').notNull(),
  createdAt: timestamp('created_at').defaultNow().notNull(),
})
