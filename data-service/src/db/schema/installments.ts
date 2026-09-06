import { pgTable, uuid, integer, date, numeric, varchar, timestamp, text } from 'drizzle-orm/pg-core'
import { loans } from './loans.js'

export const installments = pgTable('installments', {
  id: uuid('id').primaryKey().defaultRandom(),
  loanId: uuid('loan_id').references(() => loans.id).notNull(),
  installmentNumber: integer('installment_number').notNull(),
  dueDate: date('due_date').notNull(),
  principalAmount: numeric('principal_amount', { precision: 15, scale: 2 }).notNull(),
  interestAmount: numeric('interest_amount', { precision: 15, scale: 2 }).notNull(),
  penalty: numeric('penalty', { precision: 15, scale: 2 }).default('0').notNull(),
  totalAmount: numeric('total_amount', { precision: 15, scale: 2 }).notNull(),
  paidAmount: numeric('paid_amount', { precision: 15, scale: 2 }).default('0').notNull(),
  paidAt: timestamp('paid_at'),
  paidReceivedBy: text('paid_received_by'),
  receiptNumber: varchar('receipt_number', { length: 50 }),
  status: varchar('status', { length: 20 }).default('unpaid').notNull(),
  createdAt: timestamp('created_at').defaultNow().notNull(),
})
