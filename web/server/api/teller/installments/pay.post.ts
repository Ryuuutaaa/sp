export default defineEventHandler(async (event) => {
  const { id } = await readBody(event).catch(() => ({ id: "" }))
  return await proxyPost(event, `/api/installments/${id}/pay`)
})
