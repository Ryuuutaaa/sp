import { relations } from 'drizzle-orm'
import { users } from './users.js'
import { members } from './members.js'
import { savingsTypes, savingsTransactions } from './savings.js'
import { loans } from './loans.js'
import { installments } from './installments.js'
import { cashTransactions } from './cash.js'
import { shuDistributions } from './shu.js'

export const usersRelations = relations(users, ({ one }) => ({
  member: one(members, {
    fields: [users.memberId],
    references: [members.id],
  }),
}))

export const membersRelations = relations(members, ({ one, many }) => ({
  verifier: one(users, {
    fields: [members.verifiedBy],
    references: [users.id],
  }),
  savingsTransactions: many(savingsTransactions),
  loans: many(loans),
  shuDistributions: many(shuDistributions),
}))

export const savingsTypesRelations = relations(savingsTypes, ({ many }) => ({
  transactions: many(savingsTransactions),
}))

export const savingsTransactionsRelations = relations(savingsTransactions, ({ one }) => ({
  member: one(members, {
    fields: [savingsTransactions.memberId],
    references: [members.id],
  }),
  savingsType: one(savingsTypes, {
    fields: [savingsTransactions.savingsTypeId],
    references: [savingsTypes.id],
  }),
  creator: one(users, {
    fields: [savingsTransactions.createdBy],
    references: [users.id],
  }),
}))

export const loansRelations = relations(loans, ({ one, many }) => ({
  member: one(members, {
    fields: [loans.memberId],
    references: [members.id],
  }),
  approver: one(users, {
    fields: [loans.approvedBy],
    references: [users.id],
  }),
  disburser: one(users, {
    fields: [loans.disbursedBy],
    references: [users.id],
  }),
  installments: many(installments),
}))

export const installmentsRelations = relations(installments, ({ one }) => ({
  loan: one(loans, {
    fields: [installments.loanId],
    references: [loans.id],
  }),
  receiver: one(users, {
    fields: [installments.paidReceivedBy],
    references: [users.id],
  }),
}))

export const cashTransactionsRelations = relations(cashTransactions, ({ one }) => ({
  creator: one(users, {
    fields: [cashTransactions.createdBy],
    references: [users.id],
  }),
}))

export const shuDistributionsRelations = relations(shuDistributions, ({ one }) => ({
  member: one(members, {
    fields: [shuDistributions.memberId],
    references: [members.id],
  }),
}))
