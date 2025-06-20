<script setup>
import axiosInstance from "~/api/axios";

definePageMeta({
  middleware: "fake-auth",
});

const companies = ref([]);

onMounted(async () => {
  const response = await axiosInstance.get("/estabelecimentos");
  companies.value = response.data;
});

async function getEstabs() {
  const response = await axiosInstance.get("/estabelecimentos");
  companies.value = response.data;
}

const selectedCompany = ref(null);

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
  </div>
</template>
