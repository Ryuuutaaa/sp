<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Dashboard</h1>
    <div class="grid grid-cols-3 gap-4">
      <div class="bg-white p-4 rounded shadow">Anggota: {{ members?.length ?? "-" }}</div>
      <div class="bg-white p-4 rounded shadow">Pinjaman: {{ loans?.length ?? "-" }}</div>
      <div class="bg-white p-4 rounded shadow">Tipe Simpanan: {{ types?.length ?? "-" }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({ middleware: ["auth"] })
const { data: members } = await useMembers()
const { data: loans } = await useLoans()
const { data: types } = await useSavingsTypes()
</script>
