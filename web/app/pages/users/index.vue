<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Kelola User & Role</h1>
    <table class="w-full bg-white rounded shadow text-sm max-w-3xl">
      <thead><tr class="border-b text-left"><th class="p-2">Email</th><th class="p-2">Role</th><th class="p-2">Status</th><th class="p-2">Aksi</th></tr></thead>
      <tbody>
        <tr v-for="u in users" :key="u.id" class="border-b">
          <td class="p-2">{{ u.email }}</td>
          <td class="p-2">
            <select v-model="draft[u.id]" class="border px-2 py-1 rounded">
              <option value="super_admin">super_admin</option>
              <option value="admin">admin</option>
              <option value="teller">teller</option>
              <option value="anggota">anggota</option>
            </select>
          </td>
          <td class="p-2">{{ u.status }}</td>
          <td class="p-2"><button class="text-blue-600 underline" @click="save(u.id)">Simpan</button></td>
        </tr>
      </tbody>
    </table>
    <p class="text-xs text-gray-500 mt-2 max-w-3xl">Catatan: role di sini mengatur akses Go backend. Role login Better Auth diatur terpisah.</p>
    <p v-if="msg" class="text-sm mt-2 text-green-600">{{ msg }}</p>
  </div>
</template>

<script setup lang="ts">
definePageMeta({ middleware: ["auth", "role"], allowedRoles: ["super_admin"] })

const msg = ref("")
const draft = reactive<Record<string, string>>({})
const { data: users, refresh } = await useFetch<any[]>("/api/superadmin/users")

watchEffect(() => {
  for (const u of (users.value as any[]) || []) draft[u.id] = u.role
})

async function save(id: string) {
  await $fetch(`/api/superadmin/users/${id}/role`, { method: "PATCH", body: { role: draft[id] } })
  msg.value = "Role diperbarui"
  await refresh()
}
</script>
