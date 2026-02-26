<script setup>
import { ref, onMounted } from 'vue';
import { GetWordsArray } from '../wailsjs/go/words/WordList';
import WordCard from './components/WordCard.vue';

// 定义一个响应式变量来存储单词列表
// 初学Vue

const words = ref([]);

onMounted(async () => {
  try {
    words.value = await GetWordsArray();
  } catch (error) {
    console.error('Failed to load words:', error);
  }
});
</script>

<template>
  <h1>Word Remember</h1>
  <div id="card_container">
    <WordCard v-for="wod in words" :key="wod.Id" :word="wod.Word" :discription="wod.Discription"/>
  </div>
</template>

<style>
h1{
  text-align: center;
} 

#card_container{
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

</style>