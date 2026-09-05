export function useLoans() {
  return useFetch("/api/admin/loans")
}

export async function applyLoan(body: Record<string, unknown>) {
  return await $fetch("/api/anggota/loans/apply", { method: "POST", body })
}
