<script setup>
import axiosInstance from "~/api/axios";
import { Plus } from "lucide-vue-next";

definePageMeta({
  middleware: "fake-auth",
});

const companies = ref([]);
const selectedCompany = ref(null);

onMounted(async () => {
  const response = await axiosInstance.get("/estabelecimentos");
  companies.value = response.data;
});

async function getEstabs() {
  const response = await axiosInstance.get("/estabelecimentos");
  console.log(response.data);
  companies.value = response.data;
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
    <div v-if="companies.length === 0" class="text-center py-12 text-gray-400">
      <div class="mb-4">
        <Plus class="w-12 h-12 mx-auto opacity-50" />
      </div>
      <p class="text-lg font-medium mb-2">Nenhuma empresa cadastrada</p>
      <p class="text-sm">Clique no botão "Nova organização" para começar</p>
    </div>
  </div>
</template>
