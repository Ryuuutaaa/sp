<template>
  <div>
    <div class="flex justify-between items-center mb-4">
      <h1 class="text-2xl font-bold">Pinjaman</h1>
      <NuxtLink v-if="role === 'anggota'" to="/loans/apply" class="px-4 py-2 bg-blue-600 text-white rounded">Ajukan Pinjaman</NuxtLink>
    </div>
    <table class="w-full bg-white rounded shadow text-sm">
      <thead><tr class="border-b text-left"><th class="p-2">No</th><th class="p-2">Jumlah</th><th class="p-2">Tenor</th><th class="p-2">Cicilan</th><th class="p-2">Status</th><th v-if="isStaff" class="p-2">Aksi</th></tr></thead>
      <tbody>
        <tr v-for="l in loans" :key="l.id" class="border-b">
          <td class="p-2">{{ l.loanNumber }}</td>
          <td class="p-2">{{ Number(l.amount).toLocaleString("id-ID") }}</td>
          <td class="p-2">{{ l.tenorMonths }} bln</td>
          <td class="p-2">{{ Number(l.monthlyInstallment).toLocaleString("id-ID") }}</td>
          <td class="p-2">{{ l.status }}</td>
          <td v-if="isStaff" class="p-2 flex gap-2">
            <NuxtLink :to="`/installments?loanId=${l.id}`" class="text-blue-600 underline">Jadwal</NuxtLink>
            <template v-if="isAdmin">
              <button v-if="l.status === 'pending'" class="text-green-600 underline" @click="act(l.id, 'approve')">Setujui</button>
              <button v-if="l.status === 'pending'" class="text-red-600 underline" @click="act(l.id, 'reject')">Tolak</button>
              <button v-if="l.status === 'approved'" class="text-blue-600 underline" @click="act(l.id, 'disburse')">Cairkan</button>
            </template>
          </td>
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
const isStaff = role === "super_admin" || role === "admin" || role === "teller"
const isAdmin = role === "super_admin" || role === "admin"

const { data: loans, refresh } = isStaff
  ? await useLoans()
  : await useFetch("/api/anggota/loans")

async function act(id: string, action: "approve" | "reject" | "disburse") {
  const key = action === "approve" ? "approvedBy" : action === "reject" ? "rejectedBy" : "disbursedBy"
  await $fetch(`/api/admin/loans/${id}/${action}`, { method: "PATCH", body: { [key]: (session?.user as any)?.id } })
  await refresh()
}
</script>
