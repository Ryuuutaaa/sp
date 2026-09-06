export const typeDefs = /* GraphQL */ `
  type User {
    id: ID!
    memberId: ID
    email: String!
    role: String!
    status: String!
    createdAt: String!
    updatedAt: String!
  }

  type Member {
    id: ID!
    memberNumber: String!
    name: String!
    nik: String!
    address: String!
    phone: String!
    occupation: String!
    ktpPhotoUrl: String!
    selfieKtpPhotoUrl: String!
    verificationStatus: String!
    verifiedBy: ID
    verifiedAt: String
    joinDate: String!
    status: String!
    createdAt: String!
    updatedAt: String!
  }

  type SavingsType {
    id: ID!
    name: String!
    defaultAmount: String!
    isMandatory: Boolean!
    createdAt: String!
  }

  type SavingsTransaction {
    id: ID!
    memberId: ID!
    savingsTypeId: ID!
    type: String!
    amount: String!
    balanceAfter: String!
    note: String
    receiptNumber: String!
    createdBy: ID!
    createdAt: String!
  }

  type Loan {
    id: ID!
    memberId: ID!
    loanNumber: String!
    amount: String!
    interestRate: String!
    interestType: String!
    tenorMonths: Int!
    monthlyInstallment: String!
    status: String!
    approvedBy: ID
    approvedAt: String
    disbursedBy: ID
    disbursedAt: String
    createdAt: String!
    updatedAt: String!
  }

  type Installment {
    id: ID!
    loanId: ID!
    installmentNumber: Int!
    dueDate: String!
    principalAmount: String!
    interestAmount: String!
    penalty: String!
    totalAmount: String!
    paidAmount: String!
    paidAt: String
    paidReceivedBy: ID
    receiptNumber: String
    status: String!
    createdAt: String!
  }

  type CashTransaction {
    id: ID!
    type: String!
    category: String!
    amount: String!
    description: String!
    referenceType: String
    referenceId: ID
    createdBy: ID!
    createdAt: String!
  }

  type ShuDistribution {
    id: ID!
    periodYear: Int!
    memberId: ID!
    savingsProportion: String!
    loanProportion: String!
    totalShu: String!
    createdAt: String!
  }

  type Setting {
    id: ID!
    key: String!
    value: String!
    description: String
    updatedAt: String!
  }

  type Query {
    users: [User!]!
    user(id: ID!): User

    members: [Member!]!
    member(id: ID!): Member

    savingsTypes: [SavingsType!]!
    savingsTransactions(memberId: ID): [SavingsTransaction!]!

    loans(memberId: ID): [Loan!]!
    loan(id: ID!): Loan

    installments(loanId: ID!): [Installment!]!

    cashTransactions: [CashTransaction!]!

    shuDistributions(periodYear: Int!): [ShuDistribution!]!

    settings: [Setting!]!
    setting(key: String!): Setting
  }

  type Mutation {
    createUser(
      email: String!
      role: String!
      memberId: ID
    ): User!

    updateUserRole(id: ID!, role: String!): User!

    createMember(
      memberNumber: String!
      name: String!
      nik: String!
      address: String!
      phone: String!
      occupation: String!
      ktpPhotoUrl: String!
      selfieKtpPhotoUrl: String!
    ): Member!

    updateMember(
      id: ID!
      name: String
      address: String
      phone: String
      occupation: String
    ): Member!

    updateMemberStatus(id: ID!, status: String!): Member!

    updateMemberVerification(
      id: ID!
      status: String!
      verifiedBy: ID!
    ): Member!

    createSavingsTransaction(
      memberId: ID!
      savingsTypeId: ID!
      type: String!
      amount: String!
      balanceAfter: String!
      note: String
      receiptNumber: String!
      createdBy: ID!
    ): SavingsTransaction!

    createLoan(
      memberId: ID!
      loanNumber: String!
      amount: String!
      interestRate: String!
      interestType: String!
      tenorMonths: Int!
      monthlyInstallment: String!
    ): Loan!

    updateLoanStatus(
      id: ID!
      status: String!
      userId: ID!
    ): Loan!

    createInstallment(
      loanId: ID!
      installmentNumber: Int!
      dueDate: String!
      principalAmount: String!
      interestAmount: String!
      totalAmount: String!
    ): Installment!

    payInstallment(
      id: ID!
      paidAmount: String!
      penalty: String!
      receiptNumber: String!
      paidReceivedBy: ID!
    ): Installment!

    createCashTransaction(
      type: String!
      category: String!
      amount: String!
      description: String!
      referenceType: String
      referenceId: ID
      createdBy: ID!
    ): CashTransaction!

    createShuDistribution(
      periodYear: Int!
      memberId: ID!
      savingsProportion: String!
      loanProportion: String!
      totalShu: String!
    ): ShuDistribution!

    upsertSetting(
      key: String!
      value: String!
      description: String
    ): Setting!
  }
`
