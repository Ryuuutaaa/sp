import { db } from '../../db/index.js'
import * as schema from '../../db/schema/index.js'
import { eq } from 'drizzle-orm'

export const resolvers = {
  Query: {
    users: async () => await db.select().from(schema.users),
    user: async (_: any, { id }: { id: string }) => {
      const res = await db.select().from(schema.users).where(eq(schema.users.id, id))
      return res[0] || null
    },

    members: async () => await db.select().from(schema.members),
    member: async (_: any, { id }: { id: string }) => {
      const res = await db.select().from(schema.members).where(eq(schema.members.id, id))
      return res[0] || null
    },

    savingsTypes: async () => await db.select().from(schema.savingsTypes),
    savingsTransactions: async (_: any, { memberId }: { memberId?: string }) => {
      let query = db.select().from(schema.savingsTransactions)
      if (memberId) {
        query = query.where(eq(schema.savingsTransactions.memberId, memberId)) as any
      }
      return await query
    },

    loans: async (_: any, { memberId }: { memberId?: string }) => {
      let query = db.select().from(schema.loans)
      if (memberId) {
        query = query.where(eq(schema.loans.memberId, memberId)) as any
      }
      return await query
    },
    loan: async (_: any, { id }: { id: string }) => {
      const res = await db.select().from(schema.loans).where(eq(schema.loans.id, id))
      return res[0] || null
    },

    installments: async (_: any, { loanId }: { loanId: string }) => {
      return await db.select().from(schema.installments).where(eq(schema.installments.loanId, loanId))
    },

    cashTransactions: async () => await db.select().from(schema.cashTransactions),

    shuDistributions: async (_: any, { periodYear }: { periodYear: number }) => {
      return await db.select().from(schema.shuDistributions).where(eq(schema.shuDistributions.periodYear, periodYear))
    },

    settings: async () => await db.select().from(schema.settings),
    setting: async (_: any, { key }: { key: string }) => {
      const res = await db.select().from(schema.settings).where(eq(schema.settings.key, key))
      return res[0] || null
    },
  },

  Mutation: {
    createUser: async (_: any, args: any) => {
      const [row] = await db.insert(schema.users).values(args).returning()
      return row
    },
    updateUserRole: async (_: any, { id, role }: any) => {
      const [row] = await db.update(schema.users).set({ role }).where(eq(schema.users.id, id)).returning()
      return row
    },

    createMember: async (_: any, args: any) => {
      const [row] = await db.insert(schema.members).values(args).returning()
      return row
    },
    updateMember: async (_: any, { id, ...args }: any) => {
      const [row] = await db.update(schema.members).set({ ...args, updatedAt: new Date() }).where(eq(schema.members.id, id)).returning()
      return row
    },
    updateMemberStatus: async (_: any, { id, status }: any) => {
      const [row] = await db.update(schema.members).set({ status, updatedAt: new Date() }).where(eq(schema.members.id, id)).returning()
      return row
    },
    updateMemberVerification: async (_: any, { id, status, verifiedBy }: any) => {
      const [row] = await db
        .update(schema.members)
        .set({ verificationStatus: status, verifiedBy, verifiedAt: new Date(), updatedAt: new Date() })
        .where(eq(schema.members.id, id))
        .returning()
      return row
    },

    createSavingsTransaction: async (_: any, args: any) => {
      const [row] = await db.insert(schema.savingsTransactions).values(args).returning()
      return row
    },

    createLoan: async (_: any, args: any) => {
      const [row] = await db.insert(schema.loans).values(args).returning()
      return row
    },
    updateLoanStatus: async (_: any, { id, status, userId }: any) => {
      const updateData: any = { status, updatedAt: new Date() }
      if (status === 'approved') {
        updateData.approvedBy = userId
        updateData.approvedAt = new Date()
      } else if (status === 'disbursed') {
        updateData.disbursedBy = userId
        updateData.disbursedAt = new Date()
      }
      const [row] = await db.update(schema.loans).set(updateData).where(eq(schema.loans.id, id)).returning()
      return row
    },

    createInstallment: async (_: any, args: any) => {
      const [row] = await db.insert(schema.installments).values(args).returning()
      return row
    },
    payInstallment: async (_: any, { id, paidAmount, penalty, receiptNumber, paidReceivedBy }: any) => {
      const [row] = await db
        .update(schema.installments)
        .set({ 
          paidAmount, 
          penalty, 
          receiptNumber, 
          paidReceivedBy, 
          paidAt: new Date(), 
          status: 'paid' 
        })
        .where(eq(schema.installments.id, id))
        .returning()
      return row
    },

    createCashTransaction: async (_: any, args: any) => {
      const [row] = await db.insert(schema.cashTransactions).values(args).returning()
      return row
    },

    createShuDistribution: async (_: any, args: any) => {
      const [row] = await db.insert(schema.shuDistributions).values(args).returning()
      return row
    },

    upsertSetting: async (_: any, args: any) => {
      const existing = await db.select().from(schema.settings).where(eq(schema.settings.key, args.key))
      if (existing.length > 0) {
        const [row] = await db.update(schema.settings).set({ ...args, updatedAt: new Date() }).where(eq(schema.settings.key, args.key)).returning()
        return row
      }
      const [row] = await db.insert(schema.settings).values(args).returning()
      return row
    },
  },
}
