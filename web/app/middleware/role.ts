import { authClient } from "~/utils/auth-client"

export default defineNuxtRouteMiddleware(async (to) => {
  const { data } = await authClient.getSession()
  const role = (data?.user as any)?.role
  const allowed = to.meta.allowedRoles as string[] | undefined
  if (allowed && !allowed.includes(role)) return navigateTo("/dashboard")
})
