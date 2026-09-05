import { createAuthClient } from "better-auth/vue"

function getBaseURL() {
  if (typeof window !== "undefined") return window.location.origin
  if (typeof process !== "undefined" && process.env?.BETTER_AUTH_URL) {
    return process.env.BETTER_AUTH_URL
  }
  return "http://localhost:3000"
}

export const authClient = createAuthClient({
  baseURL: getBaseURL(),
})
