<script setup>
definePageMeta({
  middleware: "fake-auth",
});

const companies = ref([
  {
    id: 1,
    name: "Truesoft",
    number: 1234,
  },
  {
    id: 2,
    name: "PicoMais",
    number: 4321,
  },
  {
    id: 3,
    name: "Snet",
    number: 1423,
  },
]);

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
    <Home v-if="selectedCompany" :company="selectedCompany">
      <CompanyTab :company="selectedCompany" @close="closeTab" />
    </Home>

    <Home v-else>
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
