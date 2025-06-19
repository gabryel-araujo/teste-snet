<script setup>
import { Plus } from "lucide-vue-next";
import { ref } from "vue";
const isOpen = ref(false);

const props = defineProps({
  company: {
    type: Object,
    required: true,
  },
});
</script>

<template>
  <div class="h-full p-4 flex flex-col gap-4">
    <div class="flex flex-col gap-4">
      <p class="text-2xl font-bold" v-if="company">
        {{ company.name }} - {{ company.number }}
      </p>
      <p class="text-2xl font-bold" v-else>Suas Organizações</p>
      <div class="flex gap-2" v-if="company == null">
        <DefaultButton @click="isOpen = true"
          ><Plus /> Nova organização
        </DefaultButton>
        <Modal
          :isOpen="isOpen"
          @close="isOpen = false"
          :title="'Cadastrar Empresa'"
          :subtitle="'Preencha o formulário abaixo com os dados da nova empresa'"
        >
          <CompanyForm />
        </Modal>
        <DefaultInput :type="'search'" :placeholder="'Digite aqui'" />
      </div>
    </div>
    <div
      :class="[
        company == null
          ? 'grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4'
          : 'flex w-full',
      ]"
    >
      <slot />
    </div>
  </div>
</template>
