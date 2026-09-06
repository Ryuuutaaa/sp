<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">SHU</h1>
    <div v-if="isAdmin" class="bg-white p-4 rounded shadow mb-4 flex gap-2 max-w-lg">
      <input v-model.number="year" type="number" placeholder="Tahun" class="border px-3 py-2 rounded" />
      <button class="px-4 py-2 border rounded" @click="load()">Tampilkan</button>
      <button class="px-4 py-2 bg-blue-600 text-white rounded" @click="calculate()">Hitung & Distribusi</button>
    </div>
    <table class="w-full bg-white rounded shadow text-sm">
      <thead><tr class="border-b text-left"><th class="p-2">Tahun</th><th class="p-2">Anggota</th><th class="p-2">Porsi Simpanan</th><th class="p-2">Porsi Pinjaman</th><th class="p-2">Total SHU</th></tr></thead>
      <tbody>
        <tr v-for="s in rows" :key="s.id" class="border-b">
          <td class="p-2">{{ s.periodYear }}</td>
          <td class="p-2">{{ s.memberId }}</td>
          <td class="p-2">{{ Number(s.savingsProportion).toLocaleString("id-ID") }}</td>
          <td class="p-2">{{ Number(s.loanProportion).toLocaleString("id-ID") }}</td>
          <td class="p-2 font-medium">{{ Number(s.totalShu).toLocaleString("id-ID") }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { authClient } from "~/utils/auth-client"

definePageMeta({ middleware: ["auth"] })

const { data: session } = await authClient.getSession()
const role = (session?.user as any)?.role
const isAdmin = role === "super_admin" || role === "admin"

const year = ref(new Date().getFullYear())
const rows = ref<any[]>([])

async function load() {
  const base = isAdmin ? "/api/admin/shu" : "/api/anggota/shu"
  rows.value = await $fetch<any[]>(`${base}/${year.value}`)
}

async function calculate() {
  await $fetch(`/api/admin/shu/calculate?year=${year.value}`, { method: "POST", body: {} })
  await load()
}

await load()
</script>
