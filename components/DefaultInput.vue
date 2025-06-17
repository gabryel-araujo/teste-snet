<script setup>
const props = defineProps({
  type: {
    type: String,
    required: true,
  },
  placeholder: {
    type: String,
    required: false,
    default: "Digite aqui...",
  },
});
</script>

<template>
  <input
    :type="type"
    :placeholder="placeholder"
    class="border border-gray-600 rounded-md p-2 focus:ring-emerald-400 focus:ring-1 outline-none text-sm w-1/4"
  />
</template>
