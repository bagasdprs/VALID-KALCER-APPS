<template>
  <div class="min-h-screen bg-black text-white font-sans flex">
    <AdminSidebar />

    <main class="flex-1 ml-64 p-8">
      <slot />
    </main>
  </div>
</template>
