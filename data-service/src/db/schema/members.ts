import { pgTable, uuid, varchar, text, date, timestamp } from 'drizzle-orm/pg-core'

export const members = pgTable('members', {
  id: uuid('id').primaryKey().defaultRandom(),
  memberNumber: varchar('member_number', { length: 20 }).unique().notNull(),
  name: varchar('name', { length: 100 }).notNull(),
  nik: varchar('nik', { length: 16 }).unique().notNull(),
  address: text('address').notNull(),
  phone: varchar('phone', { length: 20 }).notNull(),
  occupation: varchar('occupation', { length: 50 }).notNull(),
  ktpPhotoUrl: text('ktp_photo_url').notNull(),
  selfieKtpPhotoUrl: text('selfie_ktp_photo_url').notNull(),
  verificationStatus: varchar('verification_status', { length: 20 }).default('pending').notNull(),
  verifiedBy: text('verified_by'),
  verifiedAt: timestamp('verified_at'),
  joinDate: date('join_date').defaultNow().notNull(),
  status: varchar('status', { length: 20 }).default('active').notNull(),
  createdAt: timestamp('created_at').defaultNow().notNull(),
  updatedAt: timestamp('updated_at').defaultNow().notNull(),
})
