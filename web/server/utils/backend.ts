import jwt from "jsonwebtoken"
import { auth } from "./auth"

export function backendUrl(path: string) {
  const config = useRuntimeConfig()
  return `${config.backendUrl}${path}`
}

// Mint short-lived Go JWT from Better Auth session.
// Claims must match Go middleware.Claims (userId, email, role, memberId).
export async function goAuthHeaders(event: any) {
  const session = await auth.api.getSession({ headers: event.headers })
  const user = session?.user as any
  if (!user) {
    throw createError({ statusCode: 401, message: "Login terlebih dahulu" })
  }
  const token = jwt.sign(
    {
      userId: user.id,
      email: user.email,
      role: user.role || "anggota",
      memberId: user.memberId || "",
    },
    process.env.JWT_SECRET!,
    { expiresIn: "15m" },
  )
  return { Authorization: `Bearer ${token}` }
}

export async function proxyGet(event: any, path: string) {
  return await $fetch(backendUrl(path), { headers: await goAuthHeaders(event) })
}

export async function proxyPost(event: any, path: string) {
  const body = await readBody(event).catch(() => ({}))
  return await $fetch(backendUrl(path), {
    method: "POST",
    body,
    headers: await goAuthHeaders(event),
  })
}

export async function proxyPatch(event: any, path: string) {
  const body = await readBody(event).catch(() => ({}))
  return await $fetch(backendUrl(path), {
    method: "PATCH",
    body,
    headers: await goAuthHeaders(event),
  })
}

export async function proxyPut(event: any, path: string) {
  const body = await readBody(event).catch(() => ({}))
  return await $fetch(backendUrl(path), {
    method: "PUT",
    body,
    headers: await goAuthHeaders(event),
  })
}
