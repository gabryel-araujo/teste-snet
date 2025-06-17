<template>
  <div
    class="w-full bg-slate-900 flex items-center justify-between h-20 gap-2 text-emerald-400 font-semibold p-4"
  >
    <section class="flex items-center gap-2">
      <img
        src="/public/favicon.ico"
        alt="ícone de dois triangulos verdes sobrepostos"
      />
      <h1 class="font-bold">Snet Test</h1>
    </section>
    <section class="flex items-center gap-6 text-white">
      <p class="text-emerald-400">Home</p>
      <p>Empresas</p>
      <section class="bg-white rounded-full p-2">
        <User class="text-emerald-400" />
      </section>
    </section>
  </div>
</template>

<script setup>
import { User } from "lucide-vue-next";
</script>
