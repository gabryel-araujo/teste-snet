<script setup>
import { useForm, useField } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import { estabSchema } from "~/schemas/estabSchema";
import axios from "axios";
import axiosInstance from "~/api/axios";

const props = defineProps({
  isOpen: {
    type: Boolean,
  },
  company: {
    type: Object,
  },
});

const { handleSubmit, errors } = useForm({
  validationSchema: toTypedSchema(estabSchema),
});

const { value: nome } = useField("nome");
const { value: numero_estabelecimento } = useField("numero_estabelecimento");
const { value: razao_social } = useField("razao_social");
const { value: cep } = useField("cep");
const { value: cidade } = useField("cidade");
const { value: estado } = useField("estado");
const { value: endereco } = useField("endereco");
const { value: numero } = useField("numero");

const emit = defineEmits(["submitted"]);

const onSubmit = handleSubmit(async (values) => {
  console.log("Empresa cadastrada:", values);
  let response;

  if (props.company) {
    console.log("Editando empresa:", values);
    response = await axiosInstance.put(
      `/estabelecimentos/${props.company.id}`,
      {
        ...values,
      }
    );
  } else {
    console.log("Cadastrando empresa:", values);
    response = await axiosInstance.post("/estabelecimentos", {
      ...values,
    });
  }

  console.log("resposta", response.data);
  // navigateTo("/app");
  emit("submitted", values);
});

watch(cep, async (newCep) => {
  if (newCep.length === 8) {
    try {
      const { data } = await axios.get(
        `https://viacep.com.br/ws/${newCep}/json/`
      );
      endereco.value = data.logradouro;
      cidade.value = data.localidade;
      estado.value = data.estado;
    } catch (error) {
      console.error("Erro ao buscar CEP", error);
    }
  }
});

if (props.company) {
  onMounted(() => {
    nome.value = props.company.nome;
    numero_estabelecimento.value = props.company.numero_estabelecimento;
    razao_social.value = props.company.razao_social;
    cep.value = props.company.cep;
    cidade.value = props.company.cidade;
    estado.value = props.company.estado;
    endereco.value = props.company.endereco;
    numero.value = props.company.numero;
  });
}
</script>

<template>
  <form @submit.prevent="onSubmit">
    <div class="flex flex-col gap-4">
      <div class="w-full flex gap-2">
        <section class="w-3/5">
          <label for="nome">Nome:</label>
          <DefaultInput
            id="nome"
            :type="'text'"
            class="w-full"
            v-model:data="nome"
          />
          <p class="text-red-500 text-sm transition">{{ errors.nome }}</p>
        </section>
        <section class="w-2/5">
          <label for="numero_estabelecimento" class="text-sm"
            >N° Estabelecimento</label
          >
          <DefaultInput
            id="numero_estabelecimento"
            :type="'text'"
            :placeholder="'Ex.: 49586'"
            class="w-full"
            v-model:data="numero_estabelecimento"
          />
          <p class="text-red-500 text-sm transition">
            {{ errors.numero_estabelecimento }}
          </p>
        </section>
      </div>
      <section>
        <label for="razao_social">Razão Social:</label>
        <DefaultInput
          id="razao_social"
          :type="'text'"
          class="w-full"
          v-model:data="razao_social"
        />
        <p class="text-red-500 text-sm transition">{{ errors.razao_social }}</p>
      </section>

      <div class="w-full flex gap-2">
        <section class="w-1/3">
          <label for="cep">CEP:</label>
          <DefaultInput
            id="cep"
            :type="'text'"
            class="w-full"
            :placeholder="'Apenas números'"
            v-model:data="cep"
          />
          <p class="text-red-500 text-sm transition">
            {{ errors.cep }}
          </p>
        </section>
        <section class="w-1/3">
          <label for="cidade">Cidade:</label>
          <DefaultInput
            id="cidade"
            :type="'text'"
            class="w-full"
            v-model:data="cidade"
          />
          <p class="text-red-500 text-sm transition">
            {{ errors.cidade }}
          </p>
        </section>
        <section class="w-1/3">
          <label for="Estado">Estado:</label>
          <DefaultInput
            id="Estado"
            :type="'text'"
            class="w-full"
            v-model:data="estado"
          />
          <p class="text-red-500 text-sm transition">
            {{ errors.estado }}
          </p>
        </section>
      </div>

      <div class="w-full flex gap-2">
        <section class="w-4/5">
          <label for="endereco">Endereço</label>
          <DefaultInput
            id="endereco"
            :type="'text'"
            class="w-full"
            v-model:data="endereco"
          />
          <p class="text-red-500 text-sm transition">
            {{ errors.endereco }}
          </p>
        </section>
        <section class="w-1/5">
          <label for="numero">Número:</label>
          <DefaultInput
            id="numero"
            :type="'text'"
            class="w-full"
            v-model:data="numero"
          />
          <p class="text-red-500 text-sm transition">
            {{ errors.numero }}
          </p>
        </section>
      </div>
      <DefaultButton> Salvar </DefaultButton>
    </div>
  </form>
</template>
