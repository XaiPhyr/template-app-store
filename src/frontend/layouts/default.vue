<script setup lang="ts"></script>

<template>
  <div class="">
    <!-- <div class="hidden md:block">
      <CoreHeader />
    </div> -->
    <slot />
    <!-- <div class="hidden md:block">
      <CoreFooter />
    </div> -->

    <div class="block sm:hidden">
      <MobileSharedFloatingButton />
    </div>
  </div>
</template>
