<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Kas Koperasi</h1>
    <div class="bg-white p-4 rounded shadow mb-4 max-w-2xl">
      <h2 class="font-semibold mb-2">Catat Kas</h2>
      <form class="grid md:grid-cols-4 gap-2" @submit.prevent="submit">
        <select v-model="form.type" class="border px-3 py-2 rounded">
          <option value="masuk">Masuk</option>
          <option value="keluar">Keluar</option>
        </select>
        <input v-model="form.category" placeholder="Kategori" class="border px-3 py-2 rounded" required />
        <input v-model="form.amount" type="number" min="1" placeholder="Nominal" class="border px-3 py-2 rounded" required />
        <input v-model="form.description" placeholder="Keterangan" class="border px-3 py-2 rounded" required />
        <button class="md:col-span-4 bg-blue-600 text-white px-4 py-2 rounded">Simpan</button>
      </form>
      <p v-if="msg" class="text-sm mt-2 text-green-600">{{ msg }}</p>
    </div>
    <table class="w-full bg-white rounded shadow text-sm">
      <thead><tr class="border-b text-left"><th class="p-2">Tanggal</th><th class="p-2">Tipe</th><th class="p-2">Kategori</th><th class="p-2">Nominal</th><th class="p-2">Keterangan</th></tr></thead>
      <tbody>
        <tr v-for="c in cash" :key="c.id" class="border-b">
          <td class="p-2">{{ new Date(c.createdAt).toLocaleString("id-ID") }}</td>
          <td class="p-2">{{ c.type }}</td>
          <td class="p-2">{{ c.category }}</td>
          <td class="p-2">{{ Number(c.amount).toLocaleString("id-ID") }}</td>
          <td class="p-2">{{ c.description }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { authClient } from "~/utils/auth-client"

definePageMeta({ middleware: ["auth", "role"], allowedRoles: ["super_admin", "admin"] })

const { data: session } = await authClient.getSession()
const form = reactive({ type: "masuk", category: "", amount: "", description: "" })
const msg = ref("")
const { data: cash, refresh } = await useFetch("/api/admin/cash")

async function submit() {
  await $fetch("/api/admin/cash", {
    method: "POST",
    body: { ...form, createdBy: (session?.user as any)?.id },
  })
  msg.value = "Kas tercatat"
  Object.assign(form, { type: "masuk", category: "", amount: "", description: "" })
  await refresh()
}
</script>
