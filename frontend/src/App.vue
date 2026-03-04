<script setup>
import { ref, onMounted, computed } from 'vue';
import  Home  from './Home.vue'
import  WordList  from './WordList.vue'
import Analyse from './Analyse.vue'

const routes = {
  '/': Home,
  '/words': WordList,
  '/analyse': Analyse
}

const currentPath = ref(window.location.hash);

const getRoutePath = (hash) => {
  const route = (hash || '#/').slice(1);
  return route.split('?')[0] || '/';
};

window.addEventListener('hashchange', () => {
  currentPath.value = window.location.hash;
  console.log('Hash changed to:', currentPath.value);
});

const currentView = computed(() => {
  return routes[getRoutePath(currentPath.value)] || Home;
})

</script>

<template>
  <component :is="currentView" :key="currentPath"></component>
</template>