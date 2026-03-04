<script setup>
import { onMounted, ref } from 'vue';
import { GetWordsArray } from '../wailsjs/go/words/WordList';
import  WordCard  from './components/WordCard.vue';

const words = ref([]);

// Get all words from the backend
onMounted(async () => {
    try {
        words.value = await GetWordsArray();
    } catch (error) {
        console.error('Failed to load words', error);
    }
});

const onHeadClick = () => {
    window.location.hash = '#/home';
};

</script>

<template>
  <div>
    <h1 v-on:click="onHeadClick">Word List</h1>
    <div id="card_container">
      <WordCard 
        v-for="word in words" 
        :key="word.Id" 
        :word="word.Word" 
        :discription="word.Discription"/>
    </div>
  </div>
</template>

<style>

#card_container{
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

h1{
    text-align: center;
    color: #7393a7;
}

</style>