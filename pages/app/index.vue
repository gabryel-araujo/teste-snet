<script setup>
import axiosInstance from "~/api/axios";
import { Plus } from "lucide-vue-next";
import Swal from "sweetalert2";

definePageMeta({
  middleware: "fake-auth",
});

const toast = Swal.mixin({
  toast: true,
  position: "top-end",
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: (toast) => {
    toast.onmouseenter = Swal.stopTimer;
    toast.onmouseleave = Swal.resumeTimer;
  },
});

const companies = ref([]);
const selectedCompany = ref(null);

onMounted(async () => {
  const response = await axiosInstance.get("/estabelecimentos");
  companies.value = response.data;
  toast.fire({
    icon: "success",
    title: "Listando empresas...",
  });
});

async function getEstabs() {
  const response = await axiosInstance.get("/estabelecimentos");
  const ordered = response.data.sort((a, b) => a.id - b.id);
  console.log(ordered);
  companies.value = ordered;
}

function changeTab(company) {
  selectedCompany.value = company;
}

function closeTab() {
  selectedCompany.value = null;
}
</script>

<template>
  <div>
    <Home
      v-if="selectedCompany"
      :company="selectedCompany"
      @close="closeTab"
      @submitted="getEstabs"
    >
      <CompanyTab :company="selectedCompany" @close="closeTab" />
    </Home>

    <Home v-else @submitted="getEstabs">
      <CompanyCard
        v-if="selectedCompany == null"
        v-for="(company, index) in companies"
        :key="index"
        :company="company"
        @click="changeTab(company)"
      />
    </Home>
    <div v-if="!companies" class="text-center py-12 text-gray-400">
      <div class="mb-4">
        <Plus class="w-12 h-12 mx-auto opacity-50" />
      </div>
      <p class="text-lg font-medium mb-2">Nenhuma empresa cadastrada</p>
      <p class="text-sm">Clique no botão "Nova organização" para começar</p>
    </div>
  </div>
</template>
