<script setup>
import { useRoute } from 'vue-router';
import UserInfo from '@/components/UserInfo.vue';

const route = useRoute();
const id = route.params.id;

</script>

<template>
  <UserInfo :id="id"/>
</template>
