<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Laporan</h1>
    <div class="grid md:grid-cols-3 gap-4">
      <div class="bg-white p-4 rounded shadow">
        <h2 class="font-semibold mb-2">Simpanan</h2>
        <p class="text-sm text-gray-500 mb-3">{{ txCount }} transaksi</p>
        <button class="px-4 py-2 border rounded" @click="exportCsv('simpanan')">Export CSV</button>
      </div>
      <div class="bg-white p-4 rounded shadow">
        <h2 class="font-semibold mb-2">Pinjaman</h2>
        <p class="text-sm text-gray-500 mb-3">{{ loans?.length ?? 0 }} pinjaman</p>
        <button class="px-4 py-2 border rounded" @click="exportCsv('pinjaman')">Export CSV</button>
      </div>
      <div class="bg-white p-4 rounded shadow">
        <h2 class="font-semibold mb-2">Anggota</h2>
        <p class="text-sm text-gray-500 mb-3">{{ members?.length ?? 0 }} anggota</p>
        <button class="px-4 py-2 border rounded" @click="exportCsv('anggota')">Export CSV</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({ middleware: ["auth", "role"], allowedRoles: ["super_admin", "admin"] })

const { data: members } = await useMembers()
const { data: loans } = await useLoans()
const txns = await $fetch<any[]>("/api/teller/savings/transactions")
const txCount = computed(() => txns.length)

function toCsv(rows: any[]) {
  if (!rows.length) return ""
  const head = Object.keys(rows[0]).join(",")
  const body = rows.map((r) => Object.values(r).map((v) => `"${String(v ?? "")}"`).join(",")).join("\n")
  return `${head}\n${body}`
}

function download(name: string, text: string) {
  const blob = new Blob([text], { type: "text/csv" })
  const a = document.createElement("a")
  a.href = URL.createObjectURL(blob)
  a.download = `${name}.csv`
  a.click()
}

function exportCsv(kind: "simpanan" | "pinjaman" | "anggota") {
  if (kind === "simpanan") return download("laporan-simpanan", toCsv(txns))
  if (kind === "pinjaman") return download("laporan-pinjaman", toCsv((loans.value as any[]) || []))
  return download("laporan-anggota", toCsv((members.value as any[]) || []))
}
</script>
