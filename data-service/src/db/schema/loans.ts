import { pgTable, uuid, varchar, numeric, integer, timestamp, text } from 'drizzle-orm/pg-core'
import { members } from './members.js'

export const loans = pgTable('loans', {
  id: uuid('id').primaryKey().defaultRandom(),
  memberId: uuid('member_id').references(() => members.id).notNull(),
  loanNumber: varchar('loan_number', { length: 20 }).unique().notNull(),
  amount: numeric('amount', { precision: 15, scale: 2 }).notNull(),
  interestRate: numeric('interest_rate', { precision: 5, scale: 2 }).notNull(),
  interestType: varchar('interest_type', { length: 10 }).notNull(),
  tenorMonths: integer('tenor_months').notNull(),
  monthlyInstallment: numeric('monthly_installment', { precision: 15, scale: 2 }).notNull(),
  status: varchar('status', { length: 20 }).default('pending').notNull(),
  approvedBy: text('approved_by'),
  approvedAt: timestamp('approved_at'),
  disbursedBy: text('disbursed_by'),
  disbursedAt: timestamp('disbursed_at'),
  createdAt: timestamp('created_at').defaultNow().notNull(),
  updatedAt: timestamp('updated_at').defaultNow().notNull(),
})
