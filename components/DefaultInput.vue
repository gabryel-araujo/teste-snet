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
  data: {
    type: String,
    required: true,
  },
});

const emit = defineEmits(["update:data"]);

function handleChange(event) {
  emit("update:data", event.target.value);
}
</script>

<template>
  <input
    :type="props.type"
    :placeholder="props.placeholder"
    class="border border-gray-600 rounded-md p-2 focus:ring-emerald-400 focus:ring-1 outline-none text-sm w-1/4 text-slate-300 focus:text-white"
    @input="handleChange"
    :value="props.data"
  />
</template>
