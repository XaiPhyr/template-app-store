<script setup lang="ts">
  const props = defineProps({
    summary: {
      type: String,
      default: 'Success',
    },
    detail: {
      type: String,
      default: 'Success',
    },
  });

  const toast = useToast();

  const show = () => {
    toast.add({
      severity: 'success',
      summary: props.summary,
      detail: props.detail,
      life: 3000,
    });
  };
</script>

<template>
  <div class="">
    <div class="card flex justify-center">
      <Toast />
      <slot :show="show" />
    </div>
  </div>
</template>

<style></style>
