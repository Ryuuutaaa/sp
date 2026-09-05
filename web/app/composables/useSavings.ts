export function useSavingsTypes() {
  return useFetch("/api/teller/savings/types")
}

export async function depositSaving(body: Record<string, unknown>) {
  return await $fetch("/api/teller/savings/deposit", { method: "POST", body })
}

export async function withdrawSaving(body: Record<string, unknown>) {
  return await $fetch("/api/teller/savings/withdraw", { method: "POST", body })
}
