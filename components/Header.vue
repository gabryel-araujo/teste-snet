<script setup>
import { User } from "lucide-vue-next";

const selected = ref("home");

function changeTab(tab) {
  selected.value = tab;

  if (tab === "home") {
    navigateTo("/app");
  } else {
    navigateTo("/app/sobre");
  }
}
</script>

<template>
  <div
    class="w-full bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 flex items-center justify-between h-20 gap-2 text-emerald-400 font-semibold p-4 border border-transparent border-b-emerald-500"
  >
    <section
      class="flex items-center gap-2 hover:scale-105 transition duration-300 cursor-pointer"
      @click="changeTab('home')"
    >
      <img
        src="/public/favicon.ico"
        alt="ícone de dois triangulos verdes sobrepostos"
      />
      <h1 class="font-bold">Snet Test</h1>
    </section>
    <section class="flex items-center gap-6 text-white">
      <p
        :class="[
          selected === 'home' ? 'text-emerald-400' : 'text-white',
          'hover:text-emerald-500 cursor-pointer transition-all duration-200',
        ]"
        @click="changeTab('home')"
      >
        Home
      </p>
      <p
        :class="[
          selected === 'companies' ? 'text-emerald-400' : 'text-white',
          'hover:text-emerald-500 cursor-pointer transition-all duration-200',
        ]"
        @click="changeTab('companies')"
      >
        Sobre
      </p>
      <section class="bg-white rounded-full p-2">
        <User class="text-emerald-400" />
      </section>
    </section>
  </div>
</template>
