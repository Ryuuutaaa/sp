<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Simpanan</h1>

    <!-- STAFF: form setor/tarik -->
    <div v-if="isStaff" class="bg-white p-4 rounded shadow mb-4">
      <h2 class="font-semibold mb-2">Input Transaksi (Teller)</h2>
      <form class="grid md:grid-cols-5 gap-2" @submit.prevent="submitTx">
        <input v-model="tx.memberId" placeholder="ID Anggota" class="border px-3 py-2 rounded" required />
        <select v-model="tx.savingsTypeId" class="border px-3 py-2 rounded" required>
          <option value="">Jenis...</option>
          <option v-for="t in types" :key="t.id" :value="t.id">{{ t.name }}</option>
        </select>
        <input v-model="tx.amount" type="number" min="1" placeholder="Nominal" class="border px-3 py-2 rounded" required />
        <input v-model="tx.receiptNumber" placeholder="No. Bukti" class="border px-3 py-2 rounded" required />
        <div class="flex gap-2">
          <button type="button" class="flex-1 bg-green-600 text-white px-3 py-2 rounded" @click="submitTx('setor')">Setor</button>
          <button type="button" class="flex-1 bg-orange-600 text-white px-3 py-2 rounded" @click="submitTx('tarik')">Tarik</button>
        </div>
      </form>
      <p v-if="txMsg" class="text-sm mt-2" :class="txErr ? 'text-red-600' : 'text-green-600'">{{ txMsg }}</p>
    </div>

    <!-- Riwayat -->
    <div class="bg-white p-4 rounded shadow">
      <h2 class="font-semibold mb-2">Riwayat Transaksi {{ isStaff ? "(filter ID anggota)" : "Saya" }}</h2>
      <div v-if="isStaff" class="flex gap-2 mb-3">
        <input v-model="filterMember" placeholder="ID Anggota (kosong = semua)" class="border px-3 py-2 rounded" />
        <button class="px-4 py-2 border rounded" @click="loadHistory()">Tampilkan</button>
      </div>
      <table class="w-full text-sm">
        <thead><tr class="border-b text-left"><th class="p-2">Tanggal</th><th class="p-2">Jenis</th><th class="p-2">Nominal</th><th class="p-2">Saldo Akhir</th><th class="p-2">Bukti</th></tr></thead>
        <tbody>
          <tr v-for="t in history" :key="t.id" class="border-b">
            <td class="p-2">{{ new Date(t.createdAt).toLocaleString("id-ID") }}</td>
            <td class="p-2">{{ t.type }}</td>
            <td class="p-2">{{ Number(t.amount).toLocaleString("id-ID") }}</td>
            <td class="p-2">{{ Number(t.balanceAfter).toLocaleString("id-ID") }}</td>
            <td class="p-2">{{ t.receiptNumber }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { authClient } from "~/utils/auth-client"

definePageMeta({ middleware: ["auth"] })

const { data: session } = await authClient.getSession()
const role = (session?.user as any)?.role
const memberId = (session?.user as any)?.memberId
const isStaff = role === "super_admin" || role === "admin" || role === "teller"

const { data: types } = await useSavingsTypes()
const filterMember = ref("")
const history = ref<any[]>([])
const tx = reactive({ memberId: "", savingsTypeId: "", amount: "", receiptNumber: "" })
const txMsg = ref("")
const txErr = ref(false)

async function loadHistory() {
  const mid = isStaff ? filterMember.value : memberId
  const q = mid ? `?memberId=${mid}` : ""
  history.value = await $fetch<any[]>(`/api/${isStaff ? "teller/savings/transactions" : "anggota/savings"}${q}`)
}

async function submitTx(kind: "setor" | "tarik") {
  txMsg.value = ""
  txErr.value = false
  try {
    await loadHistoryFor(tx.memberId)
    const balance = history.value.reduce((s: number, t: any) =>
      s + (t.type === "setor" ? Number(t.amount) : -Number(t.amount)), 0)
    const amount = Number(tx.amount)
    const balanceAfter = kind === "setor" ? balance + amount : balance - amount
    if (kind === "tarik" && balanceAfter < 0) throw new Error("Saldo tidak cukup")
    const body = {
      memberId: tx.memberId, savingsTypeId: tx.savingsTypeId,
      amount: String(amount), balanceAfter: String(balanceAfter),
      receiptNumber: tx.receiptNumber, createdBy: (session?.user as any)?.id,
    }
    if (kind === "setor") await depositSaving(body)
    else await withdrawSaving(body)
    txMsg.value = "Transaksi berhasil"
    filterMember.value = tx.memberId
    await loadHistory()
  } catch (e: any) {
    txErr.value = true
    txMsg.value = e?.data?.message || e?.message || "Transaksi gagal"
  }
}

async function loadHistoryFor(mid: string) {
  history.value = await $fetch<any[]>(`/api/teller/savings/transactions?memberId=${mid}`)
}

await loadHistory()
</script>
