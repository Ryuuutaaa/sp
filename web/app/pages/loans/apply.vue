<template>
  <div class="max-w-lg">
    <h1 class="text-2xl font-bold mb-4">Ajukan Pinjaman</h1>
    <form class="bg-white p-6 rounded shadow flex flex-col gap-3" @submit.prevent="submit">
      <input v-model="form.loanNumber" placeholder="Nomor pinjaman" class="border px-3 py-2 rounded" required />
      <input v-model="form.amount" type="number" min="1" placeholder="Jumlah pinjaman" class="border px-3 py-2 rounded" required />
      <input v-model="form.interestRate" type="number" step="0.01" placeholder="Bunga (%)" class="border px-3 py-2 rounded" required />
      <select v-model="form.interestType" class="border px-3 py-2 rounded">
        <option value="flat">Flat</option>
        <option value="anuitas">Anuitas</option>
      </select>
      <input v-model.number="form.tenorMonths" type="number" min="1" placeholder="Tenor (bulan)" class="border px-3 py-2 rounded" required />
      <input v-model="form.monthlyInstallment" type="number" min="1" placeholder="Cicilan per bulan" class="border px-3 py-2 rounded" required />
      <button class="bg-blue-600 text-white px-4 py-2 rounded" :disabled="loading">{{ loading ? "Mengirim..." : "Ajukan" }}</button>
      <p v-if="error" class="text-red-600 text-sm">{{ error }}</p>
      <p v-if="success" class="text-green-600 text-sm">Pengajuan terkirim! <NuxtLink to="/loans" class="underline">Lihat pinjaman saya</NuxtLink></p>
    </form>
  </div>
</template>

<script setup lang="ts">
import { z } from "zod"
import { authClient } from "~/utils/auth-client"

definePageMeta({ middleware: ["auth", "role"], allowedRoles: ["anggota"] })

const loading = ref(false)
const error = ref("")
const success = ref(false)
const form = reactive({ loanNumber: "", amount: "", interestRate: "", interestType: "flat", tenorMonths: 12, monthlyInstallment: "" })

const schema = z.object({
  loanNumber: z.string().min(1), amount: z.string().min(1),
  interestRate: z.string().min(1), tenorMonths: z.number().min(1),
  monthlyInstallment: z.string().min(1),
})

async function submit() {
  error.value = ""
  success.value = false
  if (!schema.safeParse(form).success) {
    error.value = "Lengkapi semua field dengan benar"
    return
  }
  const { data: session } = await authClient.getSession()
  const memberId = (session?.user as any)?.memberId
  if (!memberId) {
    error.value = "Profil anggota belum tertaut — daftar ulang data anggota dulu"
    return
  }
  loading.value = true
  try {
    await applyLoan({ ...form, memberId })
    success.value = true
  } catch (e: any) {
    error.value = e?.data?.message || e?.message || "Pengajuan gagal"
  } finally {
    loading.value = false
  }
}
</script>
