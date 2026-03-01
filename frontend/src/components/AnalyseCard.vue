<script setup>
import { marked } from 'marked';
import { onMounted, ref } from 'vue';
import { AnalyseWord } from '../../wailsjs/go/aianalyse/Agent';
const props = defineProps({
    word: String,
})

const parsedContent = ref("");

onMounted(async () => {
    parsedContent.value = marked.parse(await AnalyseWord(props.word));
})

</script>

<template>
    <div class="analyse-content" v-html="parsedContent"></div>
</template>

<style>
.analyse-content {
    padding: 16px;
    margin-top: 16px;
    color: #7393a7;
    border-radius: 8px;
}

</style>