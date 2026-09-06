export default defineEventHandler((event) => proxyPatch(event, `/api/loans/${getRouterParam(event, "id")}/reject`))
