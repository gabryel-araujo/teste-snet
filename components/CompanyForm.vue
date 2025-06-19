<template>
  <div class="flex flex-col gap-4">
    <div class="w-full flex gap-2">
      <section class="w-3/5">
        <label for="nome">Nome:</label>
        <DefaultInput id="nome" :type="'text'" class="w-full" />
      </section>
      <section class="w-2/5">
        <label for="codigo" class="text-sm">N° Estabelecimento</label>
        <DefaultInput
          id="codigo"
          :type="'text'"
          :placeholder="'Ex.: 49586'"
          class="w-full"
        />
      </section>
    </div>
    <section>
      <label for="razao">Razão Social:</label>
      <DefaultInput id="razao" :type="'text'" class="w-full" />
    </section>

    <div class="w-full flex gap-2">
      <section class="w-1/3">
        <label for="cep">CEP:</label>
        <DefaultInput id="cep" :type="'text'" class="w-full" />
      </section>
      <section class="w-1/3">
        <label for="cidade">Cidade:</label>
        <DefaultInput id="cidade" :type="'text'" class="w-full" />
      </section>
      <section class="w-1/3">
        <label for="Estado">Estado:</label>
        <DefaultInput id="Estado" :type="'text'" class="w-full" />
      </section>
    </div>

    <div class="w-full flex gap-2">
      <section class="w-4/5">
        <label for="endereco">Endereço</label>
        <DefaultInput id="endereco" :type="'text'" class="w-full" />
      </section>
      <section class="w-1/5">
        <label for="numero">Número:</label>
        <DefaultInput id="numero" :type="'text'" class="w-full" />
      </section>
    </div>
  </div>
</template>
