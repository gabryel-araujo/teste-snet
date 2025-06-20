<script setup>
const props = defineProps({
  type: {
    type: String,
    required: false,
  },
});
</script>

<template>
  <button
    :type="props.type"
    class="bg-emerald-500 p-2 rounded-lg text-slate-100 font-bold hover:shadow-md hover:cursor-pointer hover:bg-emerald-600 transition text-sm inline-flex items-center justify-center"
  >
    <slot />
  </button>
</template>
