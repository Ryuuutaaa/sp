export default defineEventHandler((event) => proxyPatch(event, `/api/members/${getRouterParam(event, "id")}`))
