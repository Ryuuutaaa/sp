<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Jadwal Angsuran</h1>
    <div class="flex gap-2 mb-4">
      <input v-model="loanId" placeholder="ID Pinjaman" class="border px-3 py-2 rounded" />
      <button class="px-4 py-2 border rounded" @click="load()">Tampilkan</button>
    </div>
    <table v-if="items.length" class="w-full bg-white rounded shadow text-sm">
      <thead><tr class="border-b text-left"><th class="p-2">Ke</th><th class="p-2">Jatuh Tempo</th><th class="p-2">Total</th><th class="p-2">Dibayar</th><th class="p-2">Status</th><th v-if="isStaff" class="p-2">Aksi</th></tr></thead>
      <tbody>
        <tr v-for="i in items" :key="i.id" class="border-b">
          <td class="p-2">{{ i.installmentNumber }}</td>
          <td class="p-2">{{ i.dueDate }}</td>
          <td class="p-2">{{ Number(i.totalAmount).toLocaleString("id-ID") }}</td>
          <td class="p-2">{{ Number(i.paidAmount).toLocaleString("id-ID") }}</td>
          <td class="p-2">{{ i.status }}</td>
          <td v-if="isStaff" class="p-2">
            <button v-if="i.status !== 'paid'" class="text-green-600 underline" @click="pay(i)">Bayar</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-if="isStaff" class="bg-white p-4 rounded shadow mt-4 max-w-lg">
      <h2 class="font-semibold mb-2">Form Pembayaran</h2>
      <form class="grid grid-cols-2 gap-2" @submit.prevent="submitPay">
        <input v-model="payForm.receiptNumber" placeholder="No. Bukti" class="border px-3 py-2 rounded" required />
        <input v-model="payForm.penalty" type="number" min="0" placeholder="Denda (0 jika tidak ada)" class="border px-3 py-2 rounded" required />
        <button class="col-span-2 bg-green-600 text-white px-4 py-2 rounded">Simpan Pembayaran</button>
      </form>
      <p v-if="payMsg" class="text-sm mt-2 text-green-600">{{ payMsg }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { authClient } from "~/utils/auth-client"

definePageMeta({ middleware: ["auth"] })

const route = useRoute()
const { data: session } = await authClient.getSession()
const role = (session?.user as any)?.role
const isStaff = role === "super_admin" || role === "admin" || role === "teller"

const loanId = ref((route.query.loanId as string) || "")
const items = ref<any[]>([])
const selected = ref<any>(null)
const payForm = reactive({ receiptNumber: "", penalty: "0" })
const payMsg = ref("")

async function load() {
  if (!loanId.value) return
  items.value = await $fetch<any[]>(`/api/installments/by-loan?loanId=${loanId.value}`)
}

function pay(i: any) {
  selected.value = i
  payMsg.value = `Bayar angsuran ke-${i.installmentNumber} sebesar ${Number(i.totalAmount).toLocaleString("id-ID")}`
}

async function submitPay() {
  if (!selected.value) return
  await $fetch("/api/teller/installments/pay", {
    method: "POST",
    body: {
      id: selected.value.id,
      paidAmount: selected.value.totalAmount,
      penalty: payForm.penalty,
      receiptNumber: payForm.receiptNumber,
      paidReceivedBy: (session?.user as any)?.id,
    },
  })
  payMsg.value = "Pembayaran tersimpan"
  await load()
}

if (loanId.value) await load()
</script>
