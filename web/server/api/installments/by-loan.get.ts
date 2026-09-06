export default defineEventHandler(async (event) => {
  const loanId = getQuery(event).loanId
  return await proxyGet(event, `/api/installments/loan/${loanId}`)
})
