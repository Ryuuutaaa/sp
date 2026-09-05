export function backendUrl(path: string) {
  const config = useRuntimeConfig()
  return `${config.backendUrl}${path}`
}

export function authHeaders(event: any) {
  const auth = getHeader(event, "authorization")
  return auth ? { Authorization: auth } : {}
}

export async function proxyGet(event: any, path: string) {
  return await $fetch(backendUrl(path), { headers: authHeaders(event) })
}

export async function proxyPost(event: any, path: string) {
  const body = await readBody(event).catch(() => ({}))
  return await $fetch(backendUrl(path), {
    method: "POST",
    body,
    headers: authHeaders(event),
  })
}

export async function proxyPatch(event: any, path: string) {
  const body = await readBody(event).catch(() => ({}))
  return await $fetch(backendUrl(path), {
    method: "PATCH",
    body,
    headers: authHeaders(event),
  })
}
