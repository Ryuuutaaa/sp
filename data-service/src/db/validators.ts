import { createInsertSchema, createSelectSchema } from 'drizzle-zod'
import * as schema from './schema/index.js'

// Members
export const insertMemberSchema = createInsertSchema(schema.members)
export const selectMemberSchema = createSelectSchema(schema.members)

// Savings
export const insertSavingsTransactionSchema = createInsertSchema(schema.savingsTransactions)
export const selectSavingsTransactionSchema = createSelectSchema(schema.savingsTransactions)

// Loans
export const insertLoanSchema = createInsertSchema(schema.loans)
export const selectLoanSchema = createSelectSchema(schema.loans)

// Installments
export const insertInstallmentSchema = createInsertSchema(schema.installments)
export const selectInstallmentSchema = createSelectSchema(schema.installments)

// Cash
export const insertCashTransactionSchema = createInsertSchema(schema.cashTransactions)
export const selectCashTransactionSchema = createSelectSchema(schema.cashTransactions)
