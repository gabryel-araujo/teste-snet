<script setup>
import { onMounted, watch } from "vue";
import { useForm, useField } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import { storeSchema } from "~/schemas/storeSchema";
import axios from "axios";
import axiosInstance from "~/api/axios";

const props = defineProps({
  isOpen: Boolean,
  company: Object,
  selectedStore: Object,
});

const emit = defineEmits(["submitted", "close"]);

const { handleSubmit, errors } = useForm({
  validationSchema: toTypedSchema(storeSchema),
});

const { value: nome } = useField("nome");
const { value: numero_loja } = useField("numero_loja");
const { value: razao_social } = useField("razao_social");
const { value: cep } = useField("cep");
const { value: cidade } = useField("cidade");
const { value: estado } = useField("estado");
const { value: endereco } = useField("endereco");
const { value: numero } = useField("numero");

const onSubmit = handleSubmit(async (values) => {
  let response;
  if (props.selectedStore) {
    response = await axiosInstance.put(`/lojas/${props.selectedStore.id}`, {
      ...values,
      estabelecimento_id: props.company.id,
    });
  } else {
    response = await axiosInstance.post("/lojas", {
      ...values,
      estabelecimento_id: props.company.id,
    });
  }

  emit("submitted", response.data);
});

watch(cep, async (newCep) => {
  if (newCep.length === 8) {
    try {
      const { data } = await axios.get(
        `https://viacep.com.br/ws/${newCep}/json/`
      );
      endereco.value = data.logradouro;
      cidade.value = data.localidade;
      estado.value = data.uf;
    } catch (error) {
      console.error("Erro ao buscar CEP:", error);
    }
  }
});

onMounted(() => {
  if (props.selectedStore) {
    nome.value = props.selectedStore.nome || "";
    numero_loja.value = props.selectedStore.numero_loja || "";
    razao_social.value = props.selectedStore.razao_social || "";
    cep.value = props.selectedStore.cep || "";
    cidade.value = props.selectedStore.cidade || "";
    estado.value = props.selectedStore.estado || "";
    endereco.value = props.selectedStore.endereco || "";
    numero.value = props.selectedStore.numero || "";
  }
});
</script>

<template>
  <form @submit.prevent="onSubmit">
    <div class="flex flex-col gap-4">
      <div class="w-full flex gap-2">
        <section class="w-3/5">
          <label for="nome">Nome:</label>
          <DefaultInput
            id="nome"
            type="text"
            class="w-full"
            v-model:data="nome"
          />
          <p class="text-red-500 text-sm">{{ errors.nome }}</p>
        </section>
        <section class="w-2/5">
          <label for="numero_loja" class="text-sm">N° loja</label>
          <DefaultInput
            id="numero_loja"
            type="text"
            placeholder="Ex.: 49586"
            class="w-full"
            v-model:data="numero_loja"
          />
          <p class="text-red-500 text-sm">{{ errors.numero_loja }}</p>
        </section>
      </div>

      <section>
        <label for="razao_social">Razão Social:</label>
        <DefaultInput
          id="razao_social"
          type="text"
          class="w-full"
          v-model:data="razao_social"
        />
        <p class="text-red-500 text-sm">{{ errors.razao_social }}</p>
      </section>

      <div class="w-full flex gap-2">
        <section class="w-1/3">
          <label for="cep">CEP:</label>
          <DefaultInput
            id="cep"
            type="text"
            placeholder="Apenas números"
            class="w-full"
            v-model:data="cep"
          />
          <p class="text-red-500 text-sm">{{ errors.cep }}</p>
        </section>
        <section class="w-1/3">
          <label for="cidade">Cidade:</label>
          <DefaultInput
            id="cidade"
            type="text"
            class="w-full"
            v-model:data="cidade"
          />
          <p class="text-red-500 text-sm">{{ errors.cidade }}</p>
        </section>
        <section class="w-1/3">
          <label for="estado">Estado:</label>
          <DefaultInput
            id="estado"
            type="text"
            class="w-full"
            v-model:data="estado"
          />
          <p class="text-red-500 text-sm">{{ errors.estado }}</p>
        </section>
      </div>

      <div class="w-full flex gap-2">
        <section class="w-4/5">
          <label for="endereco">Endereço</label>
          <DefaultInput
            id="endereco"
            type="text"
            class="w-full"
            v-model:data="endereco"
          />
          <p class="text-red-500 text-sm">{{ errors.endereco }}</p>
        </section>
        <section class="w-1/5">
          <label for="numero">Número:</label>
          <DefaultInput
            id="numero"
            type="text"
            class="w-full"
            v-model:data="numero"
          />
          <p class="text-red-500 text-sm">{{ errors.numero }}</p>
        </section>
      </div>

      <DefaultButton>Salvar</DefaultButton>
    </div>
  </form>
</template>
