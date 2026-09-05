<template>
  <div class="min-h-screen flex flex-col font-sans bg-gray-50">
    <!-- NAVBAR -->
    <header class="bg-white shadow-sm">
      <nav class="max-w-6xl mx-auto flex items-center justify-between px-6 py-4">
        <span class="text-lg font-bold text-blue-700">Koperasi SP</span>
        <div class="hidden md:flex gap-6 text-sm text-gray-600">
          <a href="#fitur" class="hover:text-blue-700">Fitur</a>
          <a href="#simpanan" class="hover:text-blue-700">Simpanan</a>
          <a href="#pinjaman" class="hover:text-blue-700">Pinjaman</a>
        </div>
        <div class="flex gap-2">
          <NuxtLink to="/auth/login" class="px-4 py-2 text-sm border border-blue-600 text-blue-600 rounded hover:bg-blue-50">
            Login
          </NuxtLink>
          <NuxtLink to="/auth/register" class="px-4 py-2 text-sm bg-blue-600 text-white rounded hover:bg-blue-700">
            Daftar Anggota
          </NuxtLink>
        </div>
      </nav>
    </header>

    <!-- BODY / HERO -->
    <main class="flex-1">
      <section class="max-w-6xl mx-auto px-6 py-16 text-center">
        <h1 class="text-4xl font-bold mb-4">Koperasi Simpan Pinjam</h1>
        <p class="text-gray-600 max-w-xl mx-auto mb-8">
          Kelola simpanan, pinjaman, dan angsuran dalam satu aplikasi.
          Daftar sebagai anggota, verifikasi KTP, dan pantau saldo kapan saja.
        </p>
        <div class="flex gap-3 justify-center">
          <NuxtLink to="/auth/register" class="px-6 py-3 bg-blue-600 text-white rounded hover:bg-blue-700">
            Daftar Jadi Anggota
          </NuxtLink>
          <NuxtLink to="/auth/login" class="px-6 py-3 border border-blue-600 text-blue-600 rounded hover:bg-blue-50">
            Login
          </NuxtLink>
        </div>
      </section>

      <section id="fitur" class="max-w-6xl mx-auto px-6 pb-16 grid md:grid-cols-3 gap-4">
        <div class="bg-white p-6 rounded shadow">
          <h3 class="font-semibold mb-2">Simpanan</h3>
          <p class="text-sm text-gray-600">Pokok, wajib, dan sukarela dengan bukti transaksi.</p>
        </div>
        <div id="pinjaman" class="bg-white p-6 rounded shadow">
          <h3 class="font-semibold mb-2">Pinjaman</h3>
          <p class="text-sm text-gray-600">Pengajuan online, approval pengurus, jadwal angsuran otomatis.</p>
        </div>
        <div id="simpanan" class="bg-white p-6 rounded shadow">
          <h3 class="font-semibold mb-2">SHU & Laporan</h3>
          <p class="text-sm text-gray-600">Distribusi SHU tahunan dan laporan siap export.</p>
        </div>
      </section>
    </main>

    <!-- FOOTER -->
    <footer class="bg-gray-900 text-gray-300">
      <div class="max-w-6xl mx-auto px-6 py-6 flex flex-col md:flex-row justify-between gap-2 text-sm">
        <span>© 2026 Koperasi SP. Hak cipta dilindungi.</span>
        <span>Simpan • Pinjam • Sejahtera</span>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { authClient } from "~/utils/auth-client"

const { data: session } = await authClient.getSession()
if (session?.user) await navigateTo("/dashboard")
</script>
