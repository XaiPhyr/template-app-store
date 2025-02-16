<script setup lang="ts"></script>

<template>
  <div class="">
    <div class="hidden md:block">
      <CoreHeader />
    </div>

    <slot />

    <div class="block md:hidden">
      <CoreFooter />
    </div>
  </div>
</template>
