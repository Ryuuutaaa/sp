import { authClient } from "~/utils/auth-client"

export default defineNuxtRouteMiddleware(async () => {
  const { data } = await authClient.getSession()
  if (!data?.user) return navigateTo("/auth/login")
})
