<!-- todo: aqui vai ser a tela de login -->
<script setup>
const companies = ref([
  {
    name: "Vivara",
    number: 1234,
  },
  {
    name: "PicoMais",
    number: 4321,
  },
  {
    name: "Snet",
    number: 1423,
  },
]);
</script>

<template>
  <div>
    <Home>
      <CompanyCard
        v-for="(company, index) in companies"
        :key="index"
        :companyName="company.name"
        :companyNumber="company.number"
      />
    </Home>
  </div>
</template>
