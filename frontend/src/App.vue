<script setup>
import { ref, onMounted } from 'vue';
import { GetWordsArray } from '../wailsjs/go/words/WordList';
import WordCard from './components/WordCard.vue';
import AnalyseCard from './components/AnalyseCard.vue';

// 定义一个响应式变量来存储单词列表
// 初学Vue

const words = ref([]);
const testExpain = ref("这是一个测试释义，展示单词的意思和用法。");

onMounted(async () => {
  try {
    words.value = await GetWordsArray();
  } catch (error) {
    console.error('Failed to load words', error);
  }
});
</script>

<template>
  <div>
    <h1>Word Remember</h1>
    <div id="card_container">
      <WordCard 
        v-for="word in words" 
        :key="word.Id" 
        :word="word.Word" 
        :discription="word.Discription"/>
    </div>
    <!-- 简单测试一下 -->
    <AnalyseCard :word="words[0]?.Word || ''"/>
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