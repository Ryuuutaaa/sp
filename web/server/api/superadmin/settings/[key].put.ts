export default defineEventHandler((event) => proxyPut(event, `/api/settings/${getRouterParam(event, "key")}`))
