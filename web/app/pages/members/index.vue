<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Daftar Anggota</h1>
    <table class="w-full bg-white rounded shadow text-sm">
      <thead><tr class="border-b"><th class="p-2 text-left">No</th><th class="p-2 text-left">Nama</th><th class="p-2 text-left">NIK</th><th class="p-2 text-left">Status</th></tr></thead>
      <tbody>
        <tr v-for="m in members" :key="m.id" class="border-b">
          <td class="p-2">{{ m.memberNumber }}</td>
          <td class="p-2">{{ m.name }}</td>
          <td class="p-2">{{ m.nik }}</td>
          <td class="p-2">{{ m.verificationStatus }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
definePageMeta({ middleware: ["auth"] })
const { data: members } = await useMembers()
</script>
