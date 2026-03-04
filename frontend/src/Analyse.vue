<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue';
import AnalyseCard from './components/AnalyseCard.vue';

const selectedWord = ref('');

const updateSelectedWord = () => {
	const hash = window.location.hash || '#/analyse';
	const query = hash.split('?')[1] || '';
	const params = new URLSearchParams(query);
	selectedWord.value = params.get('word') || '';
};

onMounted(() => {
	updateSelectedWord();
	window.addEventListener('hashchange', updateSelectedWord);
});

onBeforeUnmount(() => {
	window.removeEventListener('hashchange', updateSelectedWord);
});

const backToWords = () => {
	window.location.hash = '#/words';
};


</script>

<template>
	<div class="analyse-page">
		<button v-on:click="backToWords">返回单词列表</button>
		<h2 v-if="selectedWord">{{ selectedWord }}</h2>
		<AnalyseCard v-if="selectedWord" :key="selectedWord" :word="selectedWord" />
		<p v-else>未选择单词</p>
	</div>
</template>

<style>
.analyse-page {
	display: flex;
	flex-direction: column;
	align-items: center;
}

.analyse-page button {
	margin-top: 20px;
}
</style>