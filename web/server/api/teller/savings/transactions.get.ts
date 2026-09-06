export default defineEventHandler(async (event) => {
  const memberId = getQuery(event).memberId || ""
  return await proxyGet(event, `/api/savings/transactions?memberId=${memberId}`)
})
