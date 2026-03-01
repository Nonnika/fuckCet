<script setup>
// 单词卡片组件，展示单词和释义
const props = defineProps(["word","discription"])


</script>

<template>
    <div class="rectangle">
        <b>{{ props.word }}</b> 
        <p v-for="d in props.discription" :key="d">{{ d }}</p>
    </div>
</template>

<style>
    .rectangle{
        width: 200px;
        background-color: #7393a7;
        color:#d0e7ef;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        border-radius: 5px;
        padding: 5px;
        margin: 10px;

        transition-duration: 0.4s;
    }
    .rectangle b{
        margin-bottom: 1px;
    }
    .rectangle p{
        margin: 0;
        font-size: 14px;
    }
    .rectangle:hover{
        transform: scale(1.1);
    }
</style>