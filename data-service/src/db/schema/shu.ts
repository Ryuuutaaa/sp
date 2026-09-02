import { pgTable, uuid, integer, numeric, timestamp } from 'drizzle-orm/pg-core'
import { members } from './members.js'

export const shuDistributions = pgTable('shu_distributions', {
  id: uuid('id').primaryKey().defaultRandom(),
  periodYear: integer('period_year').notNull(),
  memberId: uuid('member_id').references(() => members.id).notNull(),
  savingsProportion: numeric('savings_proportion', { precision: 15, scale: 2 }).notNull(),
  loanProportion: numeric('loan_proportion', { precision: 15, scale: 2 }).notNull(),
  totalShu: numeric('total_shu', { precision: 15, scale: 2 }).notNull(),
  createdAt: timestamp('created_at').defaultNow().notNull(),
})
