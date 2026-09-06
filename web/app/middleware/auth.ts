import { authClient } from "~/utils/auth-client"

export default defineNuxtRouteMiddleware(async () => {
  const headers = import.meta.server ? useRequestHeaders(["cookie"]) : undefined
  const { data } = await authClient.getSession({
    fetchOptions: headers ? { headers } : {},
  })
  if (!data?.user) return navigateTo("/auth/login")
})
