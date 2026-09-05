export default defineEventHandler(async (event) => {
  const year = getQuery(event).year
  return await proxyPost(event, `/api/shu/calculate?year=${year}`)
})
