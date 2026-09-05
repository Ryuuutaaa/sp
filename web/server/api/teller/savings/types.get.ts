export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const res = await $fetch(`${config.backendUrl}/api/savings/types`)
  return res
})
