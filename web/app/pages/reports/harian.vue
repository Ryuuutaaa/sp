<template>
  <div>
    <h1 class="text-2xl font-bold mb-1">Transaksi Hari Ini</h1>
    <p class="text-sm text-gray-500 mb-4">{{ today }} — {{ todayTxns.length }} transaksi</p>
    <table class="w-full bg-white rounded shadow text-sm">
      <thead><tr class="border-b text-left"><th class="p-2">Waktu</th><th class="p-2">Anggota</th><th class="p-2">Jenis</th><th class="p-2">Nominal</th><th class="p-2">Bukti</th></tr></thead>
      <tbody>
        <tr v-for="t in todayTxns" :key="t.id" class="border-b">
          <td class="p-2">{{ new Date(t.createdAt).toLocaleTimeString("id-ID") }}</td>
          <td class="p-2">{{ t.memberId }}</td>
          <td class="p-2">{{ t.type }}</td>
          <td class="p-2">{{ Number(t.amount).toLocaleString("id-ID") }}</td>
          <td class="p-2">{{ t.receiptNumber }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
definePageMeta({ middleware: ["auth", "role"], allowedRoles: ["super_admin", "admin", "teller"] })

const today = new Date().toISOString().slice(0, 10)
const txns = await $fetch<any[]>("/api/teller/savings/transactions")
const todayTxns = computed(() => txns.filter((t: any) => String(t.createdAt).startsWith(today)))
</script>
