<script setup lang="ts">
definePageMeta({ layout: "admin" });

// Data dummy sementara buat visualisasi
const users = [
  { id: 1, name: "Bagas", email: "bagas@example.com", role: "admin", status: "Active" },
  { id: 2, name: "User Gabut", email: "gabut@example.com", role: "user", status: "Active" },
  { id: 3, name: "Spammer", email: "spam@judi.com", role: "user", status: "Banned" },
];
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-8">
      <h2 class="text-3xl font-black text-white">User Management 👥</h2>
      <button class="bg-white/10 hover:bg-white/20 text-white px-4 py-2 rounded-lg text-sm font-bold">+ Add Manual</button>
    </div>

    <div class="border border-white/10 rounded-2xl overflow-hidden">
      <table class="w-full text-left text-sm">
        <thead class="bg-zinc-900 text-zinc-400 font-bold uppercase">
          <tr>
            <th class="p-4">Name</th>
            <th class="p-4">Role</th>
            <th class="p-4">Status</th>
            <th class="p-4 text-right">Actions</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-white/10">
          <tr v-for="user in users" :key="user.id" class="hover:bg-white/5 transition-colors">
            <td class="p-4">
              <div class="font-bold text-white">{{ user.name }}</div>
              <div class="text-xs text-zinc-500">{{ user.email }}</div>
            </td>
            <td class="p-4">
              <span class="px-2 py-1 rounded text-xs font-bold bg-blue-500/20 text-blue-400 border border-blue-500/30">
                {{ user.role }}
              </span>
            </td>
            <td class="p-4">
              <span :class="user.status === 'Active' ? 'text-green-400' : 'text-red-400'"> ● {{ user.status }} </span>
            </td>
            <td class="p-4 text-right">
              <button class="text-zinc-400 hover:text-white underline">Edit</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
